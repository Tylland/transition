package domain

import (
	"fmt"
	"strings"

	dm "github.com/tylland/transition/cdas/internal/core/domain/network"
	"github.com/tylland/transition/cdas/internal/core/domain/network/visitor"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

// Generate traverses the network from every stop signal in both directions
// and returns both the schema signal routes and the domain Signalstracka objects.
func BuildSignalRoutes(network *dm.Network, segmentLookup map[string]*sfera.SegmentProfile) ([]*SignalRoute, []*dm.SignalPath) {
	var result []*SignalRoute
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

				visitor := visitor.NewSignalPathFinder(signal, direction)
				dm.Traverse(visitor)

				if visitor.SignalPath == nil {
					continue
				}

				signalPath := visitor.SignalPath

				signalRoute := &SignalRoute{
					Id:          createID(signalPath),
					StartSignal: signalPath.Start.Id.String(),
					EndSignal:   signalPath.End.Id.String(),
					StartPlace:  signalPath.Start.Id.PlaceSignature,
					EndPlace:    signalPath.End.Id.PlaceSignature,
				}

				for _, path := range signalPath.LinkPaths {
					spid := createProfileId(path.Link.Id, path.Direction)

					if sp, ok := segmentLookup[spid]; ok {

						signalRoute.SegmentProfileRefs = append(signalRoute.SegmentProfileRefs,
							&SegmentProfileReference{
								SPID:           sp.SPID,
								SPVersionMajor: sp.SPVersionMajor,
								SPVersionMinor: *sp.SPVersionMinor,
								StartPos:       path.Start,
								EndPos:         path.End,
								Direction:      string(path.Direction),
							})
					}
				}

				result = append(result, signalRoute)
				paths = append(paths, signalPath)
			}
		}
	}

	return result, paths
}

func signalName(name string) string {
	parts := strings.Split(name, "(")

	if len(parts) > 0 {
		name = parts[0]
	}

	return strings.ToLower(strings.TrimSpace(name))
}

func createID(signalPath *dm.SignalPath) string {

	return fmt.Sprintf("%d.%d.%s.%s",
		signalPath.Start.Link.From.Id.LocationNumber,
		signalPath.End.Link.To.Id.LocationNumber,
		signalName(signalPath.Start.Id.Name),
		signalName(signalPath.End.Id.Name))
}
