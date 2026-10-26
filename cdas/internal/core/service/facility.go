package service

import (
	"github.com/tylland/transition/cdas/internal/core/domain"
	dm "github.com/tylland/transition/cdas/internal/core/domain/network"
	"github.com/tylland/transition/cdas/internal/core/port"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
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
	segmentLookup := make(map[string]*sfera.SegmentProfile)

	segmentProfiles := domain.BuildSegmentProfiles(net)

	var err error
	for _, sp := range segmentProfiles {
		_, err = s.spWriter.Write(sp)

		if err != nil {
			return err
		}

		segmentLookup[sp.SPID] = sp
	}

	signalRoutes, signalPaths := domain.BuildSignalRoutes(net, segmentLookup)

	err = s.srWriter.Write(&domain.SignalRoutes{Routes: signalRoutes})

	if err != nil {
		return err
	}

	trackRoutes := domain.BuildTrackRoutes(net, signalRoutes, signalPaths)

	err = s.trWriter.Write(&domain.TrackRoutes{Routes: trackRoutes})

	if err != nil {
		return err
	}

	return nil
}
