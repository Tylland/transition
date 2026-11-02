package service

import (
	"slices"
	"testing"
	"time"

	"github.com/tylland/transition/cdas"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
	"github.com/tylland/transition/specs"
)

var onboardStory = specs.NewStory(specs.UserStory{
	AsA:     "onboard Device",
	IWantTo: "communicate with server side",
	SoThat:  "accurate advices can be given",
})

var (
	messageHeader *sfera.MessageHeader
	trainID       *sfera.TrainIdentificationComplexType
	svc           *SFERAService
	message       any
)

type TestPublisher struct {
	callback func(msg any)
}

func NewTestPublisher(callback func(msg any)) *TestPublisher {
	return &TestPublisher{callback: callback}
}

func (p *TestPublisher) Publish(msg any) error {
	p.callback(msg)

	return nil
}

func TrackSide() *specs.StepResult {
	sr := specs.NewStepResult("track side service")

	settings := &Settings{}
	publisher := NewTestPublisher(func(msg any) { message = msg })

	svc = NewSFERAService(settings, publisher, nil)

	return sr.Success()
}

func Onboard(version, companyCode, sourceDevice string) *specs.StepResult {
	sr := specs.NewStepResult("onboard device with verion %q and run by %q", version, companyCode)

	messageHeader = &sfera.MessageHeader{
		SFERAVersion: version,
		Sender:       &sfera.Sender{TeltsiCompanyCode: new(sfera.TeltsiCompanyCode(companyCode))},
		SourceDevice: sourceDevice,
	}

	return sr.Success()
}

func Train(company, core, variant string) *specs.StepResult {
	sr := specs.NewStepResult("train %q %q", company, core)

	trainID = &sfera.TrainIdentificationComplexType{TAFTAPID: &sfera.TeltsiCompositIdentifierOperationalType{
		TeltsiObjectType:    sfera.ObjectTypeTrain,
		TeltsiCompany:       (*sfera.TeltsiString44)(&company),
		TeltsiCore:          core,
		TeltsiVariant:       variant,
		TeltsiTimetableYear: time.Now().Format(cdas.YearFormat),
		TeltsiStartDate:     time.Now().Format(cdas.DateFormat),
	}}

	return sr.Success()
}

func SessionStarted(drivingMode, architecture, connectivity string) *specs.StepResult {
	sr := specs.NewStepResult("session started with driver mode %q, architecture %q and connectivity %q", drivingMode, architecture, connectivity)

	svc.startSession(trainID, string(*messageHeader.Sender.TeltsiCompanyCode), messageHeader.SourceDevice, drivingMode)

	return sr.Success()
}

func Handshaking(drivingMode, architecture, connectivity string) *specs.StepResult {
	sr := specs.NewStepResult("making handshake request with driver mode %q, architecture %q and connectivity %q", drivingMode, architecture, connectivity)

	hr := &sfera.HandshakeRequest{TrainIdentification: trainID}

	hr.DASOperatingModesSupported = append(hr.DASOperatingModesSupported,
		&sfera.DASModesComplexType{
			DASDrivingMode:  drivingMode,
			DASArchitecture: architecture,
			DASConnectivity: connectivity})

	req := &sfera.SFERAB2GRequestMessage{
		MessageHeader:    messageHeader,
		HandshakeRequest: hr,
	}

	err := svc.HandleB2GRequestMessage(req)

	if err != nil {
		return sr.WithError(err)
	}

	return sr.Success()
}

func TerminatingSession(reason string) *specs.StepResult {
	sr := specs.NewStepResult("terminating session with reason %q", reason)

	eventPayload := &sfera.B2GEventPayload{
		SessionTermination: &sfera.SessionTermination{
			TrainIdentification:      trainID,
			SessionTerminationReason: new(reason),
		},
	}

	evt := &sfera.SFERAB2GEventMessage{
		MessageHeader:   messageHeader,
		B2GEventPayload: eventPayload,
	}

	err := svc.HandleB2GEventMessage(evt)

	if err != nil {
		return sr.WithError(err)
	}

	return sr.Success()
}

func RejectedWithReason(reason string) *specs.StepResult {
	sr := specs.NewStepResult("rejected with reason: %q ", reason)

	if req, ok := message.(*sfera.SFERAG2BReplyMessage); ok {
		if req.HandshakeReject != nil {
			if !slices.Contains(req.HandshakeReject.HandshakeRejectReason, reason) {
				return sr.WithErrorf("Expected %q, got %q", reason, req.HandshakeReject.HandshakeRejectReason)
			}

		} else {
			return sr.WithErrorMessage("Expected HandshakeReject")
		}
	} else {
		return sr.WithErrorMessage("Expected SFERAG2BReplyMessage")
	}

	return sr.Success()
}

func ActiveSessionsIs(n int) *specs.StepResult {
	sr := specs.NewStepResult("active sessions is %d ", n)

	if len(svc.activeTrains) != n {
		return sr.WithErrorf("Expected %d active sessions, got %d", n, len(svc.activeTrains))
	}

	return sr.Success()
}

func TestDX1HandshakeInvalidVersionRejected(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("1.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		When(Handshaking("GoA1", "BoardAdviceCalculation", "Connected")).
		Then(RejectedWithReason(RejectReasonSFERAVersionIncompatible)).
		Execute()
}

func TestDX1HandshakeInvalidConnectivityRejected(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("3.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		When(Handshaking("GoA1", "BoardAdviceCalculation", "Standalone")).
		Then(RejectedWithReason(RejectReasonConnectivityNotSupported)).
		Execute()
}
func TestDX1HandshakeValidAcknowledged(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("3.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		When(Handshaking("GoA1", "BoardAdviceCalculation", "Connected")).
		Then(ActiveSessionsIs(1)).
		Execute()
}

func TestDX2TerminatingSession(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("3.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		Given(SessionStarted("GoA1", "BoardAdviceCalculation", "Connected")).
		When(TerminatingSession("Finished")).
		Then(ActiveSessionsIs(0)).
		Execute()
}
