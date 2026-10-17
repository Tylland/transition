package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tylland/transition/commoninterface/message"
	"github.com/tylland/transition/specs"
)

// ── Story declarations ────────────────────────────────────────────────────────

var receiverStory = specs.NewStory(specs.UserStory{
	AsA:     "CI operator",
	IWantTo: "validate inbound messages before forwarding them",
	SoThat:  "downstream consumers only receive schema-conformant messages",
})

type ReceiverTestContext struct {
	t *testing.T
	*serviceTestContext
	inboundMessage *CIMessage
	receiver       *Receiver
	forwarded      bool
	responseCode   int
	response       *message.AckResponse
	err            error
}

func NewReceiverTestContext(t *testing.T) *ReceiverTestContext {
	return &ReceiverTestContext{
		t:                  t,
		serviceTestContext: NewServiceTestContext(t),
	}
}

// ── Step functions — Given (input state) ────────────────────────────────────

func (c *ReceiverTestContext) InboundMessage(companyCode, msgID, body string) *specs.StepResult {
	res := specs.NewStepResult("inbound message %s/%s", companyCode, msgID)

	c.inboundMessage = &CIMessage{companyCode: companyCode, msgID: msgID, body: body}

	return res.Success()
}

func (c *ReceiverTestContext) ValidInboundMessage(schemaName, code string) *specs.StepResult {
	res := specs.NewStepResult("valid inbound message")

	payload := fmt.Sprintf(messageTemplate, c.schemaName, code, c.schemaName)

	c.inboundMessage = &CIMessage{companyCode: defaultCompanyCode, msgID: defaultMsgID, body: payload}

	return res.Success()
}

func MessageWithWrongPayload() *specs.StepResult {
	return specs.NewStepResult("inbound message with non-conforming payload").Success()
}

// ── Step functions — When (actions) ─────────────────────────────────────────

func (c *ReceiverTestContext) Receiving() *specs.StepResult {
	res := specs.NewStepResult("receiving message").Success()

	receiver := &Receiver{cfg: ReceiverConfig{
		Validator: c.repo,
		Handler: func(_, _, _ string) error {
			c.forwarded = true
			return nil
		},
	}}

	env := message.NewEnvelope(c.inboundMessage.companyCode, c.inboundMessage.msgID, c.inboundMessage.body)

	reqData := specs.Must(xml.Marshal(env))

	body := bytes.NewReader(reqData)

	req := httptest.NewRequest(http.MethodPost, inboundPath, body)
	rr := httptest.NewRecorder()

	receiver.handleInboundMessage(rr, req)

	c.responseCode = rr.Code

	var resp message.AckEnvelope

	respData := rr.Body.Bytes()

	if err := xml.Unmarshal(respData, &resp); err == nil {
		c.response = &resp.Body.Response
	} else {
		fmt.Print(string(respData))
		c.err = err
	}

	return res.Success()
}

// ── Step functions — Then (assertions) ──────────────────────────────────────

func (c *ReceiverTestContext) MessageIsForwarded() *specs.StepResult {
	sr := specs.NewStepResult("the message is forwarded")
	if !c.forwarded {
		return sr.WithErrorMessage("handler was not called")
	}
	return sr.Success()
}

func (c *ReceiverTestContext) MessageIsNotForwarded() *specs.StepResult {
	sr := specs.NewStepResult("the message is not forwarded")
	if c.forwarded {
		return sr.WithErrorMessage("handler was called despite validation failure")
	}
	return sr.Success()
}

func (c *ReceiverTestContext) MessageIsAccepted() *specs.StepResult {
	sr := specs.NewStepResult("message accepted, ACK is returned")

	if c.response == nil {
		return sr.WithErrorMessage("no response received")
	}

	if c.response.ResponseStatus != "ACK" {
		return sr.WithErrorMessage(fmt.Sprintf("expected %q but got %q", "ACK", c.response.ResponseStatus))
	}

	return sr.Success()
}

func (c *ReceiverTestContext) MessageIsNotAccepted() *specs.StepResult {
	sr := specs.NewStepResult("message not accepted, NACK is returned")

	if c.response == nil {
		return sr.WithErrorMessage("no response received")
	}

	if c.response.ResponseStatus != "NACK" {
		return sr.WithErrorMessage(fmt.Sprintf("expected %q but got %q", "NACK", c.response.ResponseStatus))
	}

	return sr.Success()
}

// ── Inbound validation (Receiver) tests ──────────────────────────────────────

func TestValidInboundMessageIsAcceptedAndForwarded(t *testing.T) {
	c := NewReceiverTestContext(t)

	receiverStory.Scenario(t).
		Given(c.Schema("TrainRunningInformationMessage")).
		And(c.MessageType("0011", "TrainRunningInformationMessage")).
		Given(c.RepositoryExists()).
		Given(c.ValidInboundMessage("TrainRunningInformationMessage", "0011")).
		When(c.Receiving()).
		Then(c.MessageIsAccepted()).
		And(c.MessageIsForwarded()).
		Execute()
}

func TestInboundMessageFailingValidationIsNACKed(t *testing.T) {
	c := NewReceiverTestContext(t)

	receiverStory.Scenario(t).
		Given(c.Schema("TrainRunningInformationMessage")).
		And(c.MessageType("0011", "TrainRunningInformationMessage")).
		Given(c.RepositoryExists()).
		Given(c.InboundMessage("99", "mrg-001", wrongRootPayload)).
		When(c.Receiving()).
		Then(c.MessageIsNotAccepted()).
		And(c.MessageIsNotForwarded()).
		Execute()
}

func TestInboundMessageWithNoRegisteredSchemaIsNACKed(t *testing.T) {
	c := NewReceiverTestContext(t)

	receiverStory.Scenario(t).
		Given(c.Schema("TrainCompositionMessage")).
		And(c.MessageType("3003", "TrainCompositionMessage")).
		Given(c.RepositoryLacksSchemaFor("TrainCompositionMessage")).
		Given(c.ValidInboundMessage("TrainCompositionMessage", "3003")).
		When(c.Receiving()).
		Then(c.MessageIsNotAccepted()).
		And(c.MessageIsNotForwarded()).
		Execute()
}
