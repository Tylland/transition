package service

import (
	"context"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/tylland/transition/cdas"
	"github.com/tylland/transition/cdas/internal/adapter/provider"
	"github.com/tylland/transition/cdas/internal/core/domain"
	"github.com/tylland/transition/cdas/internal/core/port"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
	"github.com/tylland/transition/specs"
)

var onboardStory = specs.NewStory(specs.UserStory{
	AsA:     "onboard Device",
	IWantTo: "communicate with server side",
	SoThat:  "accurate advices can be given",
})

var (
	jpKey    port.JPKeyProvider        = &provider.JPKeyProvider{}
	spKey    port.SPKeyProvider        = &provider.SPKeyProvider{}
	jpReader *TestJourneyProfileReader = NewTestJourneyProfileReader()
	spReader *TestSegmentProfileReader = NewTestSegmentProfileReader()
)

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

func (p *TestPublisher) Publish(ctx context.Context, msg any) error {
	p.callback(msg)

	return nil
}

type TestReader[T any] struct {
	Items map[string]*T
}

func NewTestReader[T any]() *TestReader[T] {
	return &TestReader[T]{Items: make(map[string]*T)}
}

func (r *TestReader[T]) Read(ctx context.Context, key string) (*T, error) {
	if item, ok := r.Items[key]; ok {
		return item, nil
	}

	return nil, fmt.Errorf("item with key %q not found", key)
}

type TestJourneyProfileReader struct {
	JourneyProfiles map[string]*sfera.JourneyProfile
}

func NewTestJourneyProfileReader() *TestJourneyProfileReader {
	return &TestJourneyProfileReader{JourneyProfiles: make(map[string]*sfera.JourneyProfile)}
}

func (r *TestJourneyProfileReader) Read(ctx context.Context, key string) (*sfera.JourneyProfile, error) {
	if jp, ok := r.JourneyProfiles[key]; ok {
		return jp, nil
	}

	return nil, fmt.Errorf("JourneyProfile with key %q not found", key)
}

type TestSegmentProfileReader struct {
	*TestReader[sfera.SegmentProfile]
}

func NewTestSegmentProfileReader() *TestSegmentProfileReader {
	return &TestSegmentProfileReader{TestReader: NewTestReader[sfera.SegmentProfile]()}
}

func TrackSide() *specs.StepResult {
	sr := specs.NewStepResult("track side service")

	settings := &Settings{}
	publisher := NewTestPublisher(func(msg any) { message = msg })

	svc = NewSFERAService(settings, publisher, spKey, spReader, jpKey, jpReader)

	return sr.Success()
}

func Onboard(version, companyCode, sourceDevice string) *specs.StepResult {
	sr := specs.NewStepResult("onboard device %s (%s) version %q", sourceDevice, companyCode, version)

	messageHeader = &sfera.MessageHeader{
		SFERAVersion: version,
		Sender:       &sfera.Sender{TeltsiCompanyCode: cdas.CompanyCode(companyCode)},
		SourceDevice: sourceDevice,
	}

	return sr.Success()
}

func Train(company, core, variant string) *specs.StepResult {
	core = cdas.Core(core)
	companyCode := cdas.CompanyCode(company)

	sr := specs.NewStepResult("train %q %q", *companyCode, core)

	trainID = &sfera.TrainIdentificationComplexType{TAFTAPID: &sfera.TeltsiCompositIdentifierOperationalType{
		TeltsiObjectType:    sfera.ObjectTypeTrain,
		TeltsiCompany:       companyCode,
		TeltsiCore:          core,
		TeltsiVariant:       variant,
		TeltsiTimetableYear: time.Now().Format(cdas.YearFormat),
		TeltsiStartDate:     time.Now().Format(cdas.DateFormat),
	}}

	return sr.Success()
}

func SessionStarted(drivingMode, architecture, connectivity string) *specs.StepResult {
	sr := specs.NewStepResult("session started for %q %q in mode %q", connectivity, architecture, drivingMode)

	svc.startSession(trainID, string(*messageHeader.Sender.TeltsiCompanyCode), messageHeader.SourceDevice, drivingMode)

	return sr.Success()
}

func JourneyProfileExists() *specs.StepResult {
	sr := specs.NewStepResult("journey profile exists")

	key := jpKey.Create(trainID)

	jpReader.JourneyProfiles[key] = &sfera.JourneyProfile{TrainIdentification: trainID}

	return sr.Success()
}

func SegmentProfileExists(spid, major, minor string) *specs.StepResult {
	sr := specs.NewStepResult("segment %q profile exists", spid)

	key := spKey.Data(spid)

	spReader.Items[key] = &sfera.SegmentProfile{SPVersionMajor: major, SPVersionMinor: new(minor),
		SegmentProfileIdentificationComplexType: &sfera.SegmentProfileIdentificationComplexType{
			SPID: spid,
		},
	}

	return sr.Success()
}

func Handshaking(ctx context.Context, drivingMode, architecture, connectivity string) *specs.StepResult {
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

	err := svc.HandleB2GRequestMessage(ctx, req)

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

	err := svc.HandleB2GMessage(context.Background(), evt)

	if err != nil {
		return sr.WithError(err)
	}

	return sr.Success()
}

func RequestingJourneyProfile() *specs.StepResult {
	sr := specs.NewStepResult("requesting journey profile")

	req := &sfera.SFERAB2GRequestMessage{
		MessageHeader: messageHeader,
		B2GRequest: &sfera.B2GRequest{
			JPRequest: []*sfera.JPRequest{{
				TrainIdentification: trainID,
			}},
		},
	}

	err := svc.HandleB2GMessage(context.Background(), req)

	if err != nil {
		return sr.WithError(err)
	}

	return sr.Success()
}

func RequestingSegmentProfile(spid string) *specs.StepResult {
	sr := specs.NewStepResult("requesting segment profile %q", spid)

	req := &sfera.SFERAB2GRequestMessage{
		MessageHeader: messageHeader,
		B2GRequest: &sfera.B2GRequest{
			SPRequest: []*sfera.SPRequest{{SegmentProfileIdentificationComplexType: &sfera.SegmentProfileIdentificationComplexType{
				SPID: spid,
			}}},
		},
	}

	err := svc.HandleB2GMessage(context.Background(), req)

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

func JourneyProfileReplied() *specs.StepResult {
	sr := specs.NewStepResult("journey profile replied")

	if req, ok := message.(*sfera.SFERAG2BReplyMessage); ok {
		if req.G2BReplyPayload != nil {
			if req.G2BReplyPayload.JourneyProfile != nil {
				for _, jp := range req.G2BReplyPayload.JourneyProfile {
					if jp.TrainIdentification == trainID {
						return sr.Success()
					}
				}
			} else {
				return sr.WithErrorMessage("Expected JourneyProfile")
			}

		} else {
			return sr.WithErrorMessage("Expected G2BReplyPayload")
		}
	} else {
		return sr.WithErrorMessage("Expected SFERAG2BReplyMessage")
	}

	return sr.WithErrorf("Expected journey profile %q not replied", domain.NewActiveTrainSessionKey(trainID))
}

func SegmentProfileReplied(spid string) *specs.StepResult {
	sr := specs.NewStepResult("segment profile %q replied", spid)

	if req, ok := message.(*sfera.SFERAG2BReplyMessage); ok {
		if req.G2BReplyPayload != nil {
			if req.G2BReplyPayload.SegmentProfile != nil {
				for _, sp := range req.G2BReplyPayload.SegmentProfile {
					if sp.SegmentProfileIdentificationComplexType.SPID == spid {
						return sr.Success()
					}
				}
			} else {
				return sr.WithErrorMessage("Expected JourneyProfile")
			}

		} else {
			return sr.WithErrorMessage("Expected G2BReplyPayload")
		}
	} else {
		return sr.WithErrorMessage("Expected SFERAG2BReplyMessage")
	}

	return sr.WithErrorf("Expected journey profile %q not replied", domain.NewActiveTrainSessionKey(trainID))
}

func TestDX1HandshakeInvalidVersionRejected(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("1.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		When(Handshaking(t.Context(), "GoA1", "BoardAdviceCalculation", "Connected")).
		Then(RejectedWithReason(RejectReasonSFERAVersionIncompatible)).
		Execute()
}

func TestDX1HandshakeInvalidConnectivityRejected(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("3.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		When(Handshaking(t.Context(), "GoA1", "BoardAdviceCalculation", "Standalone")).
		Then(RejectedWithReason(RejectReasonConnectivityNotSupported)).
		Execute()
}
func TestDX1HandshakeValidAcknowledged(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("3.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		When(Handshaking(t.Context(), "GoA1", "BoardAdviceCalculation", "Connected")).
		Then(ActiveSessionsIs(1)).
		Execute()
}

func TestDX2TerminatingSession(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("3.0.0", "99", "OB1")).
		Given(Train("1", "123", "01")).
		Given(SessionStarted("GoA1", "BoardAdviceCalculation", "Connected")).
		When(TerminatingSession("Finished")).
		Then(ActiveSessionsIs(0)).
		Execute()
}

func TestJP2RequestJourneyProfile(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("3.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		Given(SessionStarted("GoA1", "BoardAdviceCalculation", "Connected")).
		And(JourneyProfileExists()).
		When(RequestingJourneyProfile()).
		Then(ActiveSessionsIs(1)).
		And(JourneyProfileReplied()).
		Execute()
}

func TestJP4RequestSegmentProfile(t *testing.T) {
	onboardStory.Scenario(t).
		Given(TrackSide()).
		Given(Onboard("3.0.0", "0099", "OB1")).
		Given(Train("0001", "123", "01")).
		Given(SessionStarted("GoA1", "BoardAdviceCalculation", "Connected")).
		And(SegmentProfileExists("sp01", "1", "0")).
		When(RequestingSegmentProfile("sp01")).
		Then(SegmentProfileReplied("sp01")).
		Execute()
}
