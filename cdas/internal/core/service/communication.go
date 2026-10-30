package service

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"
	"uuid"

	"github.com/tylland/transition/cdas/internal/core/port"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

var NoRequestHandledErr = errors.New("No request was handled")

var zone = &sfera.SPZoneComplexType{IMID: new(sfera.TeltsiString44("0099"))}
var companyCode = new(sfera.TeltsiCompanyCode("0099"))
var sender = &sfera.Sender{TeltsiCompanyCode: companyCode}

type SFERASettings interface {
	SchemaVersion() string
	Zone() *sfera.SPZoneComplexType
	CompanyCode() *sfera.TeltsiCompanyCode
	Sender() *sfera.Sender
}

type Settings struct{}

func (s *Settings) SchemaVersion() string                 { return "3.0.0" }
func (s *Settings) Zone() *sfera.SPZoneComplexType        { return zone }
func (s *Settings) CompanyCode() *sfera.TeltsiCompanyCode { return companyCode }
func (s *Settings) Sender() *sfera.Sender                 { return sender }

const (
	SupportedDrivingMode  = "GoA1"      // Inactive, Read-only, Timetable, DAS not connected to ATP, GoA1, GoA2, GoA3, GoA4
	SupportedArchitecture = "DAS-C"     // DAS-O or DAS-C
	SupportedConnectivity = "Connected" // Connected or Standalone
)

var SupportedDrivingModes = []string{"GoA1"}                    // Inactive, Read-only, Timetable, DAS not connected to ATP, GoA1, GoA2, GoA3, GoA4
var SupportedArchitectures = []string{"BoardAdviceCalculation"} // DAS-O or DAS-C
var SupportedConnectivities = []string{"Connected"}             // Connected or Standalone

const (
	RejectReasonSFERAVersionIncompatible     = "SFERA version incompatible"
	RejectReasonATOSystemVersionIncompatible = "ATO system version incompatible"
	RejectReasonAnotherDASATOInCharge        = "Another DAS-TS/ATO-TS in charge"
	RejectReasonUnknownDASATOInCharge        = "DAS-TS/ATO-TS in charge unknown"
	RejectReasonArchitectureNotSupported     = "Architecture not supported"
	RejectReasonConnectivityNotSupported     = "Connectivity not supported"
	RejectReasonDrivingModesNotSupported     = "Driving Modes not supported"
)

type SessionService struct {
	settings  SFERASettings
	publisher port.SFERAMessagePublisher
	spKey     port.SPKeyProvider
	spReader  port.SegmentProfileReader
}

func NewSessionService(settings SFERASettings,
	publisher port.SFERAMessagePublisher,
	spReader port.SegmentProfileReader) *SessionService {

	return &SessionService{settings: settings, publisher: publisher, spReader: spReader}
}

func timestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z07:00")
}

func (s *SessionService) replyHeader(incoming *sfera.MessageHeader) *sfera.MessageHeader {
	return &sfera.MessageHeader{
		SFERAVersion:      s.settings.SchemaVersion(),
		MessageID:         uuid.New().String(),
		Timestamp:         timestamp(),
		MessageReference:  incoming.MessageReference,
		SourceDevice:      "TMS",
		DestinationDevice: new("DAS"),
		Sender:            s.settings.Sender(),
		Recipient:         &sfera.Recipient{TeltsiCompanyCode: incoming.Sender.TeltsiCompanyCode},
	}
}

func (s *SessionService) Handle(req sfera.SFERAB2GRequestMessage) error {
	if req.HandshakeRequest != nil {
		ack, reject, err := s.Handshake(req.HandshakeRequest, req.MessageHeader)

		if err != nil {
			return err
		}

		header := s.replyHeader(req.MessageHeader)

		if ack != nil {
			return s.reply(ack, header)
		}

		if reject != nil {
			return s.reply(reject, header)
		}
	}

	if req.B2GRequest != nil {
		payload, err := s.HandleB2GRequest(req.B2GRequest)
		if err != nil {
			return err
		}

		s.reply(payload, req.MessageHeader)
	}

	return nil
}

func (s *SessionService) reply(msg any, header *sfera.MessageHeader) error {
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

	return s.publisher.Publish(reply)
}

func (s *SessionService) Handshake(req *sfera.HandshakeRequest, hdr *sfera.MessageHeader) (*sfera.HandshakeAcknowledgement, *sfera.HandshakeReject, error) {
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

			//TODO: req.TrainIdentification

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

func (s *SessionService) HandleB2GRequest(req *sfera.B2GRequest) (*sfera.G2BReplyPayload, error) {

	// switch {
	// case req.SPRequest != nil:
	// 	return s.handleSPRequest(req.SPRequest)
	// }

	if req.SPRequest != nil {
		return s.handleSPRequest(req.SPRequest)
	}

	return nil, NoRequestHandledErr
}

func (s *SessionService) handleSPRequest(reqs []*sfera.SPRequest) (*sfera.G2BReplyPayload, error) {
	var segments []*sfera.SegmentProfile

	for _, req := range reqs {
		key := s.spKey.Data(req.SPID)
		segment, err := s.spReader.Read(context.Background(), key)

		if err != nil {
			return nil, fmt.Errorf("read %q: %w", key, err)
		}

		segments = append(segments, segment)
	}

	return &sfera.G2BReplyPayload{SegmentProfile: segments}, nil
}
