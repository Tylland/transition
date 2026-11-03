package service

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
	"uuid"

	"github.com/tylland/transition/cdas"
	"github.com/tylland/transition/cdas/internal/core/domain"
	"github.com/tylland/transition/cdas/internal/core/port"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

var NoRequestHandledErr = errors.New("No request was handled")

type SFERASettings interface {
	SchemaVersion() string
	Zone() *sfera.SPZoneComplexType
	CompanyCode() *sfera.TeltsiCompanyCode
	Sender() *sfera.Sender
}

type Settings struct{}

func (s *Settings) SchemaVersion() string { return "3.0.0" }
func (s *Settings) Zone() *sfera.SPZoneComplexType {
	return &sfera.SPZoneComplexType{IMID: cdas.CompanyCode("0099")}
}

func (s *Settings) CompanyCode() *sfera.TeltsiCompanyCode {
	return cdas.CompanyCode("0099")
}

func (s *Settings) Sender() *sfera.Sender {
	return &sfera.Sender{TeltsiCompanyCode: s.CompanyCode()}
}

const (
	SupportedDrivingMode  = "GoA1"                   // Inactive, Read-only, Timetable, DAS not connected to ATP, GoA1, GoA2, GoA3, GoA4
	SupportedArchitecture = "BoardAdviceCalculation" // DAS-O or DAS-C
	SupportedConnectivity = "Connected"              // Connected or Standalone
)

var SupportedDrivingModes = []string{"Inactive", "Read-only", "GoA1"} // Inactive, Read-only, Timetable, DAS not connected to ATP, GoA1, GoA2, GoA3, GoA4
var SupportedArchitectures = []string{"BoardAdviceCalculation"}       // DAS-O or DAS-C
var SupportedConnectivities = []string{"Connected"}                   // Connected or Standalone

const (
	RejectReasonSFERAVersionIncompatible     = "SFERA version incompatible"
	RejectReasonATOSystemVersionIncompatible = "ATO system version incompatible"
	RejectReasonAnotherDASATOInCharge        = "Another DAS-TS/ATO-TS in charge"
	RejectReasonUnknownDASATOInCharge        = "DAS-TS/ATO-TS in charge unknown"
	RejectReasonArchitectureNotSupported     = "Architecture not supported"
	RejectReasonConnectivityNotSupported     = "Connectivity not supported"
	RejectReasonDrivingModesNotSupported     = "Driving Modes not supported"
)

type SFERAService struct {
	settings     SFERASettings
	publisher    port.SFERAMessagePublisher
	spKey        port.SPKeyProvider
	spReader     port.SegmentProfileReader
	jpKey        port.JPKeyProvider
	jpReader     port.JourneyProfileReader
	activeTrains map[string]*domain.ActiveTrainSession
	mutex        *sync.RWMutex
}

func NewSFERAService(settings SFERASettings,
	publisher port.SFERAMessagePublisher,
	spKey port.SPKeyProvider,
	spReader port.SegmentProfileReader,
	jpKey port.JPKeyProvider,
	jpReader port.JourneyProfileReader) *SFERAService {

	return &SFERAService{settings: settings,
		publisher:    publisher,
		spKey:        spKey,
		spReader:     spReader,
		jpKey:        jpKey,
		jpReader:     jpReader,
		activeTrains: make(map[string]*domain.ActiveTrainSession)}
}

func timestamp() string {
	return time.Now().UTC().Format(cdas.TimestampFormat)
}

func (s *SFERAService) replyHeader(incoming *sfera.MessageHeader) *sfera.MessageHeader {
	return &sfera.MessageHeader{
		SFERAVersion:      s.settings.SchemaVersion(),
		MessageID:         uuid.New().String(),
		Timestamp:         timestamp(),
		MessageReference:  incoming.MessageReference,
		SourceDevice:      "TMS",
		DestinationDevice: &incoming.SourceDevice,
		Sender:            s.settings.Sender(),
		Recipient:         &sfera.Recipient{TeltsiCompanyCode: incoming.Sender.TeltsiCompanyCode},
	}
}

func (s *SFERAService) startSession(trainID *sfera.TrainIdentificationComplexType, companyCode, sourceDevice, drivingMode string) {
	sessionKey := domain.NewActiveTrainSessionKey(trainID)

	var session *domain.ActiveTrainSession

	if existing, ok := s.activeTrains[sessionKey]; ok {
		session = existing
	} else {
		session = &domain.ActiveTrainSession{
			TrainID: trainID,
			Devices: make(map[string]*domain.ActiveDevice),
		}

		s.activeTrains[sessionKey] = session
	}

	deviceKey := domain.NewActiveDeviceKey(companyCode, sourceDevice)

	device := &domain.ActiveDevice{
		CompanyCode:  companyCode,
		SourceDevice: sourceDevice,
		DrivingMode:  drivingMode,
		Started:      time.Now(),
		LastSeen:     time.Now(),
	}

	session.Devices[deviceKey] = device
}

func (s *SFERAService) terminateSession(trainID *sfera.TrainIdentificationComplexType, companyCode string, sourceDevice string) {
	sessionKey := domain.NewActiveTrainSessionKey(trainID)

	if existing, ok := s.activeTrains[sessionKey]; ok {
		deviceKey := domain.NewActiveDeviceKey(companyCode, sourceDevice)

		delete(existing.Devices, deviceKey)

		if len(existing.Devices) == 0 {
			delete(s.activeTrains, sessionKey)
		}
	}

}

func (s *SFERAService) HandleB2GMessage(ctx context.Context, msg any) error {
	switch m := msg.(type) {
	case *sfera.SFERAB2GRequestMessage:
		return s.HandleB2GRequestMessage(ctx, m)
	case *sfera.SFERAB2GEventMessage:
		return s.HandleB2GEventMessage(m)
	case *sfera.SFERAB2GReplyMessage:
		return s.HandleB2GReplyMessage(m)
	default:
		return fmt.Errorf("Unsupported message type: %t", m)
	}
}

func (s *SFERAService) HandleB2GRequestMessage(ctx context.Context, req *sfera.SFERAB2GRequestMessage) error {
	if req.HandshakeRequest != nil {
		ack, reject, err := s.Handshake(req.HandshakeRequest, req.MessageHeader)

		if err != nil {
			return err
		}

		header := s.replyHeader(req.MessageHeader)

		if ack != nil {
			return s.reply(ctx, ack, header)
		}

		if reject != nil {
			return s.reply(ctx, reject, header)
		}
	}

	if req.B2GRequest != nil {
		payload, err := s.HandleB2GRequest(ctx, req.B2GRequest)
		if err != nil {
			return err
		}

		s.reply(ctx, payload, req.MessageHeader)
	}

	return nil
}

func (s *SFERAService) HandleB2GEventMessage(evt *sfera.SFERAB2GEventMessage) error {
	if evt.B2GEventPayload != nil {
		if evt.B2GEventPayload.SessionTermination != nil {
			s.terminateSession(evt.B2GEventPayload.SessionTermination.TrainIdentification,
				string(*evt.MessageHeader.Sender.TeltsiCompanyCode), evt.MessageHeader.SourceDevice)
		}
	}

	return nil
}

func (s *SFERAService) HandleB2GReplyMessage(evt *sfera.SFERAB2GReplyMessage) error {
	return nil
}

func (s *SFERAService) reply(ctx context.Context, msg any, header *sfera.MessageHeader) error {
	reply := &sfera.SFERAG2BReplyMessage{
		MessageHeader: header,
	}

	switch v := msg.(type) {
	case *sfera.HandshakeAcknowledgement:
		reply.HandshakeAcknowledgement = v
	case *sfera.HandshakeReject:
		reply.HandshakeReject = v
	case *sfera.G2BReplyPayload:
		reply.G2BReplyPayload = v
	default:
		return fmt.Errorf("Invalid message type: %v", msg)
	}

	return s.publisher.Publish(ctx, reply)
}

func (s *SFERAService) Handshake(req *sfera.HandshakeRequest, hdr *sfera.MessageHeader) (*sfera.HandshakeAcknowledgement, *sfera.HandshakeReject, error) {
	if hdr.SFERAVersion != s.settings.SchemaVersion() {
		reject := &sfera.HandshakeReject{
			SPZone:                s.settings.Zone(),
			HandshakeRejectReason: []string{RejectReasonSFERAVersionIncompatible},
		}

		return nil, reject, nil
	}

	for _, mode := range req.DASOperatingModesSupported {
		if slices.Contains(SupportedDrivingModes, mode.DASDrivingMode) &&
			slices.Contains(SupportedArchitectures, mode.DASArchitecture) &&
			slices.Contains(SupportedConnectivities, mode.DASConnectivity) {

			s.startSession(req.TrainIdentification, string(*hdr.Sender.TeltsiCompanyCode), hdr.SourceDevice, mode.DASDrivingMode)

			ack := &sfera.HandshakeAcknowledgement{
				DASOperatingModeSelected: &sfera.DASOperatingModeSelected{
					DASArchitecture: mode.DASArchitecture,
					DASConnectivity: mode.DASConnectivity},
			}

			return ack, nil, nil
		}
	}

	for _, mode := range req.DASOperatingModesSupported {
		reason := []string{}

		if !slices.Contains(SupportedDrivingModes, mode.DASDrivingMode) {
			reason = append(reason, RejectReasonDrivingModesNotSupported)
		}

		if !slices.Contains(SupportedArchitectures, mode.DASArchitecture) {
			reason = append(reason, RejectReasonArchitectureNotSupported)
		}

		if !slices.Contains(SupportedConnectivities, mode.DASConnectivity) {
			reason = append(reason, RejectReasonConnectivityNotSupported)
		}

		if len(reason) > 0 {
			reject := &sfera.HandshakeReject{
				SPZone:                s.settings.Zone(),
				HandshakeRejectReason: reason,
			}

			return nil, reject, nil
		}
	}

	return nil, nil, fmt.Errorf("No supported mode without reason")
}

func (s *SFERAService) HandleB2GRequest(ctx context.Context, req *sfera.B2GRequest) (*sfera.G2BReplyPayload, error) {
	switch {
	case req.SPRequest != nil:
		return s.handleSPRequest(ctx, req.SPRequest)
	case req.JPRequest != nil:
		return s.handleJPRequest(ctx, req.JPRequest)
	}

	return nil, NoRequestHandledErr
}

func (s *SFERAService) handleSPRequest(ctx context.Context, reqs []*sfera.SPRequest) (*sfera.G2BReplyPayload, error) {
	var segments []*sfera.SegmentProfile

	for _, req := range reqs {
		key := s.spKey.Data(req.SPID)
		segment, err := s.spReader.Read(ctx, key)

		if err != nil {
			return nil, fmt.Errorf("read %q: %w", key, err)
		}

		segments = append(segments, segment)
	}

	return &sfera.G2BReplyPayload{SegmentProfile: segments}, nil
}

func (s *SFERAService) handleJPRequest(ctx context.Context, reqs []*sfera.JPRequest) (*sfera.G2BReplyPayload, error) {
	var journeys []*sfera.JourneyProfile

	for _, req := range reqs {
		key := s.jpKey.Create(req.TrainIdentification)
		journey, err := s.jpReader.Read(ctx, key)

		if err != nil {
			return nil, fmt.Errorf("read %q: %w", key, err)
		}

		journeys = append(journeys, journey)
	}

	return &sfera.G2BReplyPayload{JourneyProfile: journeys}, nil
}
