package service

import (
	"github.com/tylland/transition/cdas/internal/core/domain"
	dm "github.com/tylland/transition/cdas/internal/core/domain/network"
	"github.com/tylland/transition/cdas/internal/core/port"
)

type FacilityService struct {
	srWriter port.SignalRouteWriter
	trWriter port.TrackRouteWriter
	spWriter port.SegmentProfileWriter
}

func NewFacilityService(srWriter port.SignalRouteWriter, trWriter port.TrackRouteWriter, spWriter port.SegmentProfileWriter) *FacilityService {
	return &FacilityService{srWriter: srWriter, trWriter: trWriter, spWriter: spWriter}
}

func (s *FacilityService) ProcessNetwork(net *dm.Network) error {

	signalRoutes, signalPaths := domain.BuildSignalRoutes(net)

	err := s.srWriter.Write(&domain.SignalRoutes{Routes: signalRoutes})

	if err != nil {
		return err
	}

	trackRoutes := domain.BuildTrackRoutes(net, signalRoutes, signalPaths)

	err = s.trWriter.Write(&domain.TrackRoutes{Routes: trackRoutes})

	if err != nil {
		return err
	}

	segmentProfiles := domain.BuildSegmentProfiles(net)

	for _, sp := range segmentProfiles {
		_, err = s.spWriter.Write(sp)

		if err != nil {
			return err
		}
	}

	return nil
}
