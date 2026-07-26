package routes

import (
	"fmt"

	dm "github.com/tylland/transition/internal/domainmodel"
	"github.com/tylland/transition/internal/domainmodel/networkvisitor"
	"github.com/tylland/transition/internal/schema/routes"
)

// createProfileId creates a SFERA segment profile ID from a LankId and direction.
func createProfileId(linkId dm.LinkId, riktning dm.LinkDirection) string {
	return fmt.Sprintf("%v:%d", linkId, int(riktning))
}

// GenerateSignalRoutes traverses the network from every stop signal in both directions
// and returns the resulting signal routes.
func GenerateSignalRoutes(network *dm.Network) []routes.SignalRoute {
	signalRoutes, _ := GenerateSignalRoutesWithDomain(network)
	return signalRoutes
}

// GenerateSignalRoutesWithDomain traverses the network from every stop signal in both directions
// and returns both the schema signal routes and the domain Signalstracka objects.
func GenerateSignalRoutesWithDomain(network *dm.Network) ([]routes.SignalRoute, []*dm.SignalPath) {
	var result []routes.SignalRoute
	var strackor []*dm.SignalPath

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
					StartPlace:  signalPath.Start.Id.TrafikplatsSignatur,
					EndPlace:    signalPath.End.Id.TrafikplatsSignatur,
				}

				for _, path := range signalPath.LinkPaths {
					signalRoute.SegmentProfileRefs = append(signalRoute.SegmentProfileRefs,
						routes.SegmentProfileReference{
							SPID:      createProfileId(path.Link.Id, path.Direction),
							StartPos:  path.Start,
							EndPos:    path.End,
							Direction: int(path.Direction),
						})
				}

				result = append(result, signalRoute)
				strackor = append(strackor, signalPath)
			}
		}
	}

	return result, strackor
}

// findTrackAtPosition returns the track number (spårnummer) that covers the given
// position on the given link. Returns empty string if none found.
func findTrackAtPosition(link *dm.Link, position float64) string {
	for _, sn := range link.StationTracks.Items() {
		start := sn.Extent.Start
		slut := sn.Extent.End
		if start > slut {
			start, slut = slut, start
		}
		if position >= start && position <= slut {
			return sn.Number
		}
	}
	return ""
}

// findLineTrackAtPosition returns the line track (linjespår) that covers the given
// position on the given link. Returns empty string if none found.
func findLineTrackAtPosition(link *dm.Link, position float64) string {
	for _, ls := range link.LineTracks.Items() {
		start := ls.Extent.Start
		slut := ls.Extent.End
		if start > slut {
			start, slut = slut, start
		}
		if position >= start && position <= slut {
			return ls.Number
		}
	}
	return ""
}

// isLineSignal returns true if the signal type is a line signal (INF, UBL, MBL)
// that sits on the line between stations and should use linjespår for track lookup.
func isLineSignal(signalType string) bool {
	return signalType == dm.SignalTypInfart ||
		signalType == dm.SignalTypUtfart ||
		signalType == dm.SignalTypMellanBlock
}

// findTrackForSignal returns the track number for a signal position.
// Line signals (INF, UBL, MBL) use linjespår, station signals (M) use spårnummer.
func findTrackForSignal(link *dm.Link, position float64, signalType string) string {
	if isLineSignal(signalType) {
		return findLineTrackAtPosition(link, position)
	}
	return findTrackAtPosition(link, position)
}

// findFirstLineTrackOnLink returns the first linjespår nummer on the link, or empty string.
func findFirstLineTrackOnLink(link *dm.Link) string {
	if items := link.LineTracks.Items(); len(items) > 0 {
		return items[0].Number
	}
	return ""
}

// findFirstTrackOnLink returns the first spårnummer on the link, or empty string.
func findFirstTrackOnLink(link *dm.Link) string {
	if items := link.StationTracks.Items(); len(items) > 0 {
		return items[0].Number
	}
	return ""
}

// findStartTrack finds the track at the start of a signal route.
// Uses linjespår for line signals (INF, UBL, MBL), spårnummer for station signals (M).
// If not found at the signal position, searches forward through the route's links.
func findStartTrack(st *dm.SignalPath) string {
	startLank := st.Start.LinkPosition.Link
	track := findTrackForSignal(startLank, st.Start.LinkPosition.Position, st.Start.Type)
	if track != "" {
		return track
	}
	line := isLineSignal(st.Start.Type)
	for _, la := range st.LinkPaths {
		if line {
			if t := findFirstLineTrackOnLink(la.Link); t != "" {
				return t
			}
		} else {
			if t := findFirstTrackOnLink(la.Link); t != "" {
				return t
			}
		}
	}
	return ""
}

// findEndTrack finds the track at the end of a signal route.
// Uses linjespår for line signals (INF, UBL, MBL), spårnummer for station signals (M).
// If not found at the signal position, searches backward through the route's links.
func findEndTrack(st *dm.SignalPath) string {
	endLink := st.End.LinkPosition.Link
	track := findTrackForSignal(endLink, st.End.LinkPosition.Position, st.End.Type)
	if track != "" {
		return track
	}
	line := isLineSignal(st.End.Type)
	for i := len(st.LinkPaths) - 1; i >= 0; i-- {
		la := st.LinkPaths[i]
		if line {
			if t := findFirstLineTrackOnLink(la.Link); t != "" {
				return t
			}
		} else {
			if t := findFirstTrackOnLink(la.Link); t != "" {
				return t
			}
		}
	}
	return ""
}

// classifySignalRoute determines the TrackRouteType for a signal route based on
// the start and end signal types:
//   - Arrival: line signal → M (entering a station)
//   - Departure: M → line signal (leaving a station)
//   - Line: line signal → line signal (between stations)
func classifySignalRoute(st *dm.SignalPath) routes.TrackRouteType {
	startIsLine := isLineSignal(st.Start.Type)
	endIsLine := isLineSignal(st.End.Type)

	if startIsLine && endIsLine {
		return routes.TrackRouteLine
	}
	if startIsLine && !endIsLine {
		return routes.TrackRouteArrival
	}
	if !startIsLine && endIsLine {
		return routes.TrackRouteDeparture
	}

	// M → M: use place difference as hint
	samePlats := st.Start.Id.TrafikplatsSignatur == st.End.Id.TrafikplatsSignatur
	if !samePlats {
		return routes.TrackRouteLine
	}

	return routes.TrackRouteArrival
}

// trackRouteKey is the grouping key for track routes.
type trackRouteKey struct {
	Type       routes.TrackRouteType
	StartPlace string
	EndPlace   string
	StartTrack string
	EndTrack   string
}

// GenerateTrackRoutes creates track routes by classifying each signal route as
// Arrival, Departure, or Line based on signal types, then grouping by type and
// place/track.
func GenerateTrackRoutes(network *dm.Network, signalRoutes []routes.SignalRoute, signalstrackor []*dm.SignalPath) []routes.TrackRoute {
	grouped := make(map[trackRouteKey][]routes.SignalRouteRef)

	for i, sr := range signalRoutes {
		st := signalstrackor[i]

		routeType := classifySignalRoute(st)

		// Determine track based on signal type:
		// Line signals (INF, UBL, MBL) → Linjespar
		// Station signals (M) → Sparnummer
		startTrack := findStartTrack(st)
		endTrack := findEndTrack(st)

		key := trackRouteKey{
			Type:       routeType,
			StartPlace: sr.StartPlace,
			EndPlace:   sr.EndPlace,
			StartTrack: startTrack,
			EndTrack:   endTrack,
		}

		ref := routes.SignalRouteRef{
			StartSignal: sr.StartSignal,
			EndSignal:   sr.EndSignal,
		}

		grouped[key] = append(grouped[key], ref)
	}

	var result []routes.TrackRoute
	for key, refs := range grouped {
		tr := routes.TrackRoute{
			Type:            key.Type,
			StartPlace:      key.StartPlace,
			EndPlace:        key.EndPlace,
			StartTrack:      key.StartTrack,
			EndTrack:        key.EndTrack,
			StartSignal:     refs[0].StartSignal,
			EndSignal:       refs[len(refs)-1].EndSignal,
			SignalRouteRefs: refs,
		}
		result = append(result, tr)
	}

	return result
}

// GenerateRoutes generates all routes (signal routes and track routes) for the network.
func GenerateRoutes(network *dm.Network) *routes.Routes {
	signalRoutes, signalstrackor := GenerateSignalRoutesWithDomain(network)
	trackRoutes := GenerateTrackRoutes(network, signalRoutes, signalstrackor)

	return &routes.Routes{
		SignalRoutes: signalRoutes,
		TrackRoutes:  trackRoutes,
	}
}
