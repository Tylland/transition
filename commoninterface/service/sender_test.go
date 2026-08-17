package service

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/tylland/transition/commoninterface/message"
	"github.com/tylland/transition/commoninterface/repository"
	"github.com/tylland/transition/specs"
)

// ── Story declarations ────────────────────────────────────────────────────────

var senderStory = specs.NewStory(specs.UserStory{
	AsA:     "CI operator",
	IWantTo: "validate outbound messages before sending",
	SoThat:  "schema-violating messages are never posted to a partner",
})

type SenderTestContext struct {
	t *testing.T
	*serviceTestContext
	outboundMessage *CIMessage
	sndr            *Sender
	called          bool
	err             error
}

func NewSenderTestContext(t *testing.T) *SenderTestContext {
	baseDir := t.TempDir()

	schemasDir := filepath.Join(baseDir, "schemas")

	os.MkdirAll(schemasDir, 0x644)

	return &SenderTestContext{
		t:                  t,
		serviceTestContext: NewServiceTestContext(t),
	}
}

// ── Step functions — Given (input state) ────────────────────────────────────

func (c *SenderTestContext) createSender() (*Sender, error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		c.called = true
		ack := message.AckEnvelope{
			NsSOAP: message.NsSOAP,
			Body:   message.AckBody{Response: message.AckResponse{ResponseStatus: "ACK", AckIdentifier: "ACKID-test"}},
		}
		data, _ := xml.Marshal(ack)
		w.Header().Set("Content-Type", "text/xml")
		w.Write(append([]byte(xml.Header), data...))
	}))

	c.t.Cleanup(server.Close)

	sndr, err := NewSender(SenderConfig{
		LocalIP: "127.0.0.1",
		Partners: map[string]PartnerConfig{
			"0000099": {CompanyCode: "0000099", EndpointURL: server.URL},
		},
		Validator: c.repo,
	})

	if err != nil {
		return nil, fmt.Errorf("creating sender: %w", err)
	}

	return sndr, nil
}

func (c *SenderTestContext) OutboundMessage(companyCode, msgID, body string) *specs.StepResult {
	res := specs.NewStepResult("message %s/%s", companyCode, msgID)

	c.outboundMessage = &CIMessage{companyCode: companyCode, msgID: msgID, body: body}

	return res.Success()
}

func (c *SenderTestContext) ValidOutboundMessage() *specs.StepResult {
	res := specs.NewStepResult("valid outbound message")

	c.outboundMessage = &CIMessage{companyCode: defaultCompanyCode, msgID: defaultMsgID, body: validPayload}

	return res.Success()
}

// ── Step functions — When (actions) ─────────────────────────────────────────

func (c *SenderTestContext) Sending() *specs.StepResult {
	if c.outboundMessage == nil {
		panic("No outbound message")
	}

	companyCode := c.outboundMessage.companyCode
	msgID := c.outboundMessage.msgID
	body := c.outboundMessage.body

	sr := specs.NewStepResult("sending %s/%s", companyCode, msgID)

	sender := specs.Must(c.createSender())
	_, c.err = sender.Send(companyCode, msgID, body)

	return sr.Success()
}

// ── Step functions — Then (assertions) ──────────────────────────────────────

func (c *SenderTestContext) RequestWasMade() *specs.StepResult {
	sr := specs.NewStepResult("the message is posted to the partner endpoint")

	if !c.called {
		return sr.WithErrorMessage("no HTTP request was made to the partner")
	}

	return sr.Success()
}

func (c *SenderTestContext) NoRequestWasMade() *specs.StepResult {
	sr := specs.NewStepResult("no HTTP request is made to the partner endpoint")

	if c.called {
		return sr.WithErrorMessage("HTTP request was made despite validation failure")
	}

	return sr.Success()
}

func (c *SenderTestContext) PartnerNotRegistredError() *specs.StepResult {
	sr := specs.NewStepResult("partner not registred error")

	if _, ok := errors.AsType[ErrPartnerNotRegistred](c.err); !ok {
		return sr.WithErrorMessage("No or wrong error received")
	}

	return sr.Success()
}

func (c *SenderTestContext) SchemaNotRegistredError() *specs.StepResult {
	sr := specs.NewStepResult("schema not registred error")

	if _, ok := errors.AsType[repository.ErrSchemaNotRegistred](c.err); !ok {
		return sr.WithErrorMessage("No or wrong error received")
	}

	return sr.Success()
}

// ── Outbound validation (Sender) tests ───────────────────────────────────────

func TestValidOutboundMessageIsSent(t *testing.T) {
	c := NewSenderTestContext(t)

	senderStory.Scenario(t).
		Given(c.RepositoryWithSchemaFor("TrainRunningInformationMessage")).
		And(c.ValidOutboundMessage()).
		When(c.Sending()).
		Then(c.RequestWasMade()).
		Execute()
}

func TestOutboundMessageFailingValidationIsRejected(t *testing.T) {
	c := NewSenderTestContext(t)

	senderStory.Scenario(t).
		Given(c.RepositoryWithSchemaFor("TrainRunningInformationMessage")).
		And(c.OutboundMessage("0000099", "test-msg-id", wrongRootPayload)).
		When(c.Sending()).
		Then(c.NoRequestWasMade()).
		Execute()
}

func TestSendingMessageWithNoRegisteredSchemaReturnsError(t *testing.T) {
	c := NewSenderTestContext(t)

	senderStory.Scenario(t).
		Given(c.RepositoryLacksSchemaFor("PathRequestMessage")).
		And(c.ValidOutboundMessage()).
		When(c.Sending()).
		Then(c.NoRequestWasMade()).
		And(c.SchemaNotRegistredError()).
		Execute()
}
