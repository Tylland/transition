package routes

import (
	"fmt"

	dm "github.com/tylland/transition/internal/domainmodel"
	"github.com/tylland/transition/internal/domainmodel/natvisitors"
	"github.com/tylland/transition/internal/schema/routes"
)

// createProfileId creates a SFERA segment profile ID from a LankId and direction.
func createProfileId(lankId dm.LankId, riktning dm.Lankriktning) string {
	return fmt.Sprintf("%v:%d", lankId, int(riktning))
}

// GenerateSignalRoutes traverses the network from every stop signal in both directions
// and returns the resulting signal routes.
func GenerateSignalRoutes(nat *dm.Nat) []routes.SignalRoute {
	signalRoutes, _ := GenerateSignalRoutesWithDomain(nat)
	return signalRoutes
}

// GenerateSignalRoutesWithDomain traverses the network from every stop signal in both directions
// and returns both the schema signal routes and the domain Signalstracka objects.
func GenerateSignalRoutesWithDomain(nat *dm.Nat) ([]routes.SignalRoute, []*dm.Signalstracka) {
	var result []routes.SignalRoute
	var strackor []*dm.Signalstracka

	for _, lank := range nat.Lankar() {
		for _, signal := range lank.Signaler.Items() {
			if !signal.IsStopSignal() {
				continue
			}

			for _, riktning := range []dm.Lankriktning{dm.LankriktningMed, dm.LankriktningMot} {
				if !signal.Lankposition.Gilltig(riktning) {
					continue
				}

				visitor := natvisitors.NewSignalstrackaVisitor(signal, riktning)
				dm.Traverse(visitor)

				if visitor.Signalstracka == nil {
					continue
				}

				st := visitor.Signalstracka

				signalRoute := routes.SignalRoute{
					StartSignal: st.Start.Id.String(),
					EndSignal:   st.Slut.Id.String(),
					StartPlace:  st.Start.Id.TrafikplatsSignatur,
					EndPlace:    st.Slut.Id.TrafikplatsSignatur,
				}

				for _, la := range st.Lankanvandningar {
					signalRoute.SegmentProfileRefs = append(signalRoute.SegmentProfileRefs,
						routes.SegmentProfileReference{
							SPID:      createProfileId(la.Lank.Id, la.Riktning),
							StartPos:  la.Start,
							EndPos:    la.Slut,
							Direction: int(la.Riktning),
						})
				}

				result = append(result, signalRoute)
				strackor = append(strackor, st)
			}
		}
	}

	return result, strackor
}

// findTrackAtPosition returns the track number (spårnummer) that covers the given
// position on the given link. Returns empty string if none found.
func findTrackAtPosition(lank *dm.Lank, position float64) string {
	for _, sn := range lank.Sparnummer.Items() {
		start := sn.Lankutbredning.Start
		slut := sn.Lankutbredning.Slut
		if start > slut {
			start, slut = slut, start
		}
		if position >= start && position <= slut {
			return sn.Nummer
		}
	}
	return ""
}

// findLineTrackAtPosition returns the line track (linjespår) that covers the given
// position on the given link. Returns empty string if none found.
func findLineTrackAtPosition(lank *dm.Lank, position float64) string {
	for _, ls := range lank.Linjespar.Items() {
		start := ls.Lankutbredning.Start
		slut := ls.Lankutbredning.Slut
		if start > slut {
			start, slut = slut, start
		}
		if position >= start && position <= slut {
			return ls.Nummer
		}
	}
	return ""
}

// isLineSignal returns true if the signal type is a line signal (INF, UBL, MBL)
// that sits on the line between stations and should use linjespår for track lookup.
func isLineSignal(signalTyp string) bool {
	return signalTyp == dm.SignalTypInfart ||
		signalTyp == dm.SignalTypUtfart ||
		signalTyp == dm.SignalTypMellanBlock
}

// findTrackForSignal returns the track number for a signal position.
// Line signals (INF, UBL, MBL) use linjespår, station signals (M) use spårnummer.
func findTrackForSignal(lank *dm.Lank, position float64, signalTyp string) string {
	if isLineSignal(signalTyp) {
		return findLineTrackAtPosition(lank, position)
	}
	return findTrackAtPosition(lank, position)
}

// findFirstLineTrackOnLink returns the first linjespår nummer on the link, or empty string.
func findFirstLineTrackOnLink(lank *dm.Lank) string {
	if items := lank.Linjespar.Items(); len(items) > 0 {
		return items[0].Nummer
	}
	return ""
}

// findFirstTrackOnLink returns the first spårnummer on the link, or empty string.
func findFirstTrackOnLink(lank *dm.Lank) string {
	if items := lank.Sparnummer.Items(); len(items) > 0 {
		return items[0].Nummer
	}
	return ""
}

// findStartTrack finds the track at the start of a signal route.
// Uses linjespår for line signals (INF, UBL, MBL), spårnummer for station signals (M).
// If not found at the signal position, searches forward through the route's links.
func findStartTrack(st *dm.Signalstracka) string {
	startLank := st.Start.Lankposition.Lank
	track := findTrackForSignal(startLank, st.Start.Lankposition.Position, st.Start.Typ)
	if track != "" {
		return track
	}
	line := isLineSignal(st.Start.Typ)
	for _, la := range st.Lankanvandningar {
		if line {
			if t := findFirstLineTrackOnLink(la.Lank); t != "" {
				return t
			}
		} else {
			if t := findFirstTrackOnLink(la.Lank); t != "" {
				return t
			}
		}
	}
	return ""
}

// findEndTrack finds the track at the end of a signal route.
// Uses linjespår for line signals (INF, UBL, MBL), spårnummer for station signals (M).
// If not found at the signal position, searches backward through the route's links.
func findEndTrack(st *dm.Signalstracka) string {
	endLank := st.Slut.Lankposition.Lank
	track := findTrackForSignal(endLank, st.Slut.Lankposition.Position, st.Slut.Typ)
	if track != "" {
		return track
	}
	line := isLineSignal(st.Slut.Typ)
	for i := len(st.Lankanvandningar) - 1; i >= 0; i-- {
		la := st.Lankanvandningar[i]
		if line {
			if t := findFirstLineTrackOnLink(la.Lank); t != "" {
				return t
			}
		} else {
			if t := findFirstTrackOnLink(la.Lank); t != "" {
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
func classifySignalRoute(st *dm.Signalstracka) routes.TrackRouteType {
	startIsLine := isLineSignal(st.Start.Typ)
	endIsLine := isLineSignal(st.Slut.Typ)

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
	samePlats := st.Start.Id.TrafikplatsSignatur == st.Slut.Id.TrafikplatsSignatur
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
func GenerateTrackRoutes(nat *dm.Nat, signalRoutes []routes.SignalRoute, signalstrackor []*dm.Signalstracka) []routes.TrackRoute {
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
func GenerateRoutes(nat *dm.Nat) *routes.Routes {
	signalRoutes, signalstrackor := GenerateSignalRoutesWithDomain(nat)
	trackRoutes := GenerateTrackRoutes(nat, signalRoutes, signalstrackor)

	return &routes.Routes{
		SignalRoutes: signalRoutes,
		TrackRoutes:  trackRoutes,
	}
}
