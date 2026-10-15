package domain

import (
	dm "github.com/tylland/transition/internal/core/domain/network"
	"github.com/tylland/transition/internal/core/domain/network/networkvisitor"
)

// Generate traverses the network from every stop signal in both directions
// and returns both the schema signal routes and the domain Signalstracka objects.
func BuildSignalRoutes(network *dm.Network) ([]SignalRoute, []*dm.SignalPath) {
	var result []SignalRoute
	var paths []*dm.SignalPath

	for _, link := range network.Links() {
		for _, signal := range link.Signals.Items() {
			if !signal.IsStopSignal() {
				continue
			}

			for _, direction := range []dm.LinkDirection{dm.LinkDirectionWith, dm.LinkDirectionAgainst} {
				if !signal.LinkPosition.IsValid(direction) {
					continue
				}

				visitor := networkvisitor.NewSignalPathCreator(signal, direction)
				dm.Traverse(visitor)

				if visitor.SignalPath == nil {
					continue
				}

				signalPath := visitor.SignalPath

				signalRoute := SignalRoute{
					StartSignal: signalPath.Start.Id.String(),
					EndSignal:   signalPath.End.Id.String(),
					StartPlace:  signalPath.Start.Id.PlaceSignature,
					EndPlace:    signalPath.End.Id.PlaceSignature,
				}

				for _, path := range signalPath.LinkPaths {
					signalRoute.SegmentProfileRefs = append(signalRoute.SegmentProfileRefs,
						SegmentProfileReference{
							SPID:      createProfileId(path.Link.Id, path.Direction),
							StartPos:  path.Start,
							EndPos:    path.End,
							Direction: string(path.Direction),
						})
				}

				result = append(result, signalRoute)
				paths = append(paths, signalPath)
			}
		}
	}

	return result, paths
}
