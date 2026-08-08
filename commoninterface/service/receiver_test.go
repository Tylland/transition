package service

import (
	"bytes"
	"encoding/json"
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

	data := specs.Must(xml.Marshal(env))

	body := bytes.NewReader(data)

	req := httptest.NewRequest(http.MethodPost, "/hello", body)
	rr := httptest.NewRecorder()

	receiver.handleInboundMessage(rr, req)

	c.responseCode = rr.Code

	// Parse JSON response
	var resp outAckEnvelope

	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err == nil {
		c.response = &resp.Body.Response
	}

	return res.Success()
}

func (c *ReceiverTestContext) InboundMessage(companyCode, msgID, body string) *specs.StepResult {
	res := specs.NewStepResult("inbound message %s/%s", companyCode, msgID)

	c.inboundMessage = &CIMessage{companyCode: companyCode, msgID: msgID, body: body}

	return res.Success()
}

func payload(text string) *specs.StepResult {
	return specs.NewStepResult("inbound payload").Success()
}

// func SchemaExists(messageType string) *specs.StepResult {
// 	sr := specs.NewStepResult("schema exists for %q", messageType)
// 	dir, err := os.MkdirTemp("", "")
// 	if err != nil {
// 		return sr.WithError(err)
// 	}
// 	if err := os.WriteFile(filepath.Join(dir, messageType+".xsd"), minimalXSD, 0644); err != nil {
// 		return sr.WithError(err)
// 	}
// 	repo, err := repository.NewSchemaRepository(dir)
// 	if err != nil {
// 		return sr.WithError(err)
// 	}
// 	f := new(bool)
// 	forwarded = f
// 	r = &Receiver{cfg: ReceiverConfig{
// 		Validator: repo,
// 		Handler:   func(_, _, _ string) error { *f = true; return nil },
// 	}}
// 	return sr.Success()
// }

func MessageWithWrongPayload() *specs.StepResult {
	return specs.NewStepResult("inbound message with non-conforming payload").Success()
}

// ── Step functions — When (actions) ─────────────────────────────────────────

// func validating(r *Receiver, text string) *specs.StepResult {
// 	sr := specs.NewStepResult("validating inbound payload")
// 	if err := r.cfg.Validator.Validate(text); err == nil {
// 		_ = r.cfg.Handler("test-id", "127.0.0.1", text)
// 	}
// 	return sr.Success()
// }

// func Validating() *specs.StepResult {
// 	return validating(r, wrongRootPayload)
// }

// ── Step functions — Then (assertions) ──────────────────────────────────────

func (c *ReceiverTestContext) MessageIsForwarded() *specs.StepResult {
	sr := specs.NewStepResult("the message is forwarded")
	if !c.forwarded {
		return sr.WithErrorMessage("handler was not called")
	}
	return sr.Success()
}

func (c *ReceiverTestContext) MessageIsNotForwarded() *specs.StepResult {
	sr := specs.NewStepResult("the message is not forwarded and a NACK is returned")
	if c.forwarded {
		return sr.WithErrorMessage("handler was called despite validation failure")
	}
	return sr.Success()
}

func (c *ReceiverTestContext) MessageIsNotAccepted() *specs.StepResult {
	sr := specs.NewStepResult("message not accepted")

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
		Given(c.RepositoryWithSchemaFor("TrainRunningInformationMessage")).
		Given(c.InboundMessage("99", "mrg-001", validPayload)).
		When(c.Receiving()).
		Then(c.MessageIsForwarded()).
		Execute()
}

func TestInboundMessageFailingValidationIsNACKed(t *testing.T) {
	c := NewReceiverTestContext(t)

	receiverStory.Scenario(t).
		Given(c.RepositoryWithSchemaFor("TrainRunningInformationMessage")).
		Given(c.InboundMessage("99", "mrg-001", wrongRootPayload)).
		When(c.Receiving()).
		Then(c.MessageIsNotAccepted()).
		And(c.MessageIsNotForwarded()).
		Execute()
}

func TestInboundMessageWithNoRegisteredSchemaIsNACKed(t *testing.T) {
	c := NewReceiverTestContext(t)

	receiverStory.Scenario(t).
		Given(c.RepositoryLacksSchemaFor("PathRequestMessage")).
		Given(c.InboundMessage("99", "mrg-001", wrongRootPayload)).
		When(c.Receiving()).
		Then(c.MessageIsNotAccepted()).
		And(c.MessageIsNotForwarded()).
		Execute()
}
