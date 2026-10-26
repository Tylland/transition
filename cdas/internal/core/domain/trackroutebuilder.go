package domain

import (
	"fmt"
	"strings"

	dm "github.com/tylland/transition/cdas/internal/core/domain/network"
	"github.com/tylland/transition/cdas/internal/core/domain/network/visitor"
)

// Generate creates track routes by classifying each signal route as
// Arrival, Departure, or Line based on signal types, then grouping by type and
// place/track.
func BuildTrackRoutes(network *dm.Network, signalRoutes []*SignalRoute, signalPaths []*dm.SignalPath) []*TrackRoute {
	grouped := make(map[trackRouteKey][]SignalRouteRef)

	for i, sr := range signalRoutes {
		sp := signalPaths[i]

		routeType := classifySignalRoute(sp)

		startTrack := findStartTrack(sp)
		endTrack := findEndTrack(sp)

		directionPlace := findDirectionPlace(sp, routeType)

		key := trackRouteKey{
			Type:           routeType,
			StartPlace:     sr.StartPlace,
			EndPlace:       sr.EndPlace,
			DirectionPlace: directionPlace,
			StartTrack:     startTrack,
			EndTrack:       endTrack,
		}

		ref := SignalRouteRef{
			StartSignal: sr.StartSignal,
			EndSignal:   sr.EndSignal,
		}

		grouped[key] = append(grouped[key], ref)
	}

	var result []*TrackRoute
	for key, refs := range grouped {
		tr := &TrackRoute{
			Id:              key.String(),
			Type:            key.Type,
			StartPlace:      key.StartPlace,
			EndPlace:        key.EndPlace,
			DirectionPlace:  key.DirectionPlace,
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

// findDirectionPlace finds the place before an arrival or after a departure.
func findDirectionPlace(signalPath *dm.SignalPath, trt TrackRouteType) string {
	if trt == TrackRouteLine {
		return signalPath.Start.Id.PlaceSignature
	}

	signal := signalPath.End
	var reverse bool

	if trt == TrackRouteArrival {
		signal = signalPath.Start
		reverse = true
	}

	visitor := visitor.NewPlaceFinder(signal, reverse)
	dm.Traverse(visitor)
	if visitor.PlaceSignature() != "" {
		return visitor.PlaceSignature()
	}

	return signal.Id.PlaceSignature
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
func classifySignalRoute(st *dm.SignalPath) TrackRouteType {
	startIsLine := isLineSignal(st.Start.Type)
	endIsLine := isLineSignal(st.End.Type)

	if startIsLine && endIsLine {
		return TrackRouteLine
	}
	if startIsLine && !endIsLine {
		return TrackRouteArrival
	}
	if !startIsLine && endIsLine {
		return TrackRouteDeparture
	}

	// M → M: use place difference as hint
	samePlats := st.Start.Id.PlaceSignature == st.End.Id.PlaceSignature
	if !samePlats {
		return TrackRouteLine
	}

	return TrackRouteArrival
}

// trackRouteKey is the grouping key for track
type trackRouteKey struct {
	Type           TrackRouteType
	StartPlace     string
	EndPlace       string
	DirectionPlace string
	StartTrack     string
	EndTrack       string
}

func (k *trackRouteKey) String() string {
	return strings.ToLower(fmt.Sprintf("%s.%s.%s.%s.%s.%s", k.StartPlace, k.EndPlace, k.DirectionPlace, k.StartTrack, k.EndTrack, k.Type))
}
