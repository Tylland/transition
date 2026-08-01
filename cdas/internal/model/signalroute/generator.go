package signalroute

import (
	"fmt"

	dm "github.com/tylland/transition/internal/model/domainmodel"
	"github.com/tylland/transition/internal/model/domainmodel/networkvisitor"
	"github.com/tylland/transition/internal/schema/routes"
)

// createProfileId creates a SFERA segment profile ID from a LankId and direction.
func createProfileId(linkId dm.LinkId, riktning dm.LinkDirection) string {
	return fmt.Sprintf("%v:%s", linkId, riktning)
}

// Generate traverses the network from every stop signal in both directions
// and returns both the schema signal routes and the domain Signalstracka objects.
func Generate(network *dm.Network) ([]routes.SignalRoute, []*dm.SignalPath) {
	var result []routes.SignalRoute
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

				signalRoute := routes.SignalRoute{
					StartSignal: signalPath.Start.Id.String(),
					EndSignal:   signalPath.End.Id.String(),
					StartPlace:  signalPath.Start.Id.PlaceSignature,
					EndPlace:    signalPath.End.Id.PlaceSignature,
				}

				for _, path := range signalPath.LinkPaths {
					signalRoute.SegmentProfileRefs = append(signalRoute.SegmentProfileRefs,
						routes.SegmentProfileReference{
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
