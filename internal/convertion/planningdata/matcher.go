// Package planningdata provides conversion logic for matching DagligFärdplan
// (daily timetable) data to TrackRoutes and their underlying SegmentProfiles.
package planningdata

import (
	"fmt"

	pd "github.com/tylland/transition/internal/schema/planningdata"
	"github.com/tylland/transition/internal/schema/routes"
)

// JourneySegment represents one matched step in a train's journey,
// linking timetable data to the corresponding TrackRoute.
type JourneySegment struct {
	// Type is the kind of segment: Arrival, Departure, or Line.
	Type routes.TrackRouteType
	// StartPlace is the location signature of the departure point.
	StartPlace string
	// EndPlace is the location signature of the arrival point.
	EndPlace string
	// StartTrack is the track identifier at the start of the segment.
	StartTrack string
	// EndTrack is the track identifier at the end of the segment.
	EndTrack string
	// TrackRoute is the matched TrackRoute, or nil if no match was found.
	TrackRoute *routes.TrackRoute
}

// MatchResult holds the result of matching a DagligFärdplan to TrackRoutes.
type MatchResult struct {
	// Tågläge is the train number.
	Tågläge int
	// Utgångsdag is the departure date.
	Utgångsdag string
	// Segments is the ordered list of journey segments.
	Segments []JourneySegment
}

// trackRouteKey is used to look up track routes by their identifying properties.
type trackRouteKey struct {
	Type       routes.TrackRouteType
	StartPlace string
	EndPlace   string
	StartTrack string
	EndTrack   string
}

// trackRouteIndex maps trackRouteKey to the matching TrackRoute for fast lookup.
type trackRouteIndex map[trackRouteKey]*routes.TrackRoute

// buildIndex creates a lookup index from a slice of TrackRoutes.
func buildIndex(trackRoutes []routes.TrackRoute) trackRouteIndex {
	idx := make(trackRouteIndex, len(trackRoutes))
	for i := range trackRoutes {
		tr := &trackRoutes[i]
		key := trackRouteKey{
			Type:       tr.Type,
			StartPlace: tr.StartPlace,
			EndPlace:   tr.EndPlace,
			StartTrack: tr.StartTrack,
			EndTrack:   tr.EndTrack,
		}
		idx[key] = tr
	}
	return idx
}

// lookup finds a TrackRoute matching the given properties.
func (idx trackRouteIndex) lookup(typ routes.TrackRouteType, startPlace, endPlace, startTrack, endTrack string) *routes.TrackRoute {
	key := trackRouteKey{
		Type:       typ,
		StartPlace: startPlace,
		EndPlace:   endPlace,
		StartTrack: startTrack,
		EndTrack:   endTrack,
	}
	return idx[key]
}

// MatchFärdplan matches a DagligFärdplan to TrackRoutes, producing an ordered
// sequence of JourneySegments that represent the train's complete journey.
//
// For a journey visiting stops [A, B, C], the segments are:
//
//	Departure(A) → Line(A→B) → Arrival(B) → Departure(B) → Line(B→C) → Arrival(C)
//
// Each segment is matched to a TrackRoute by Type + Place + Track.
// If no matching TrackRoute is found, the segment's TrackRoute field is nil.
func MatchFärdplan(färdplan *pd.DagligFärdplan, trackRoutes []routes.TrackRoute) *MatchResult {
	idx := buildIndex(trackRoutes)

	result := &MatchResult{
		Tågläge:    färdplan.ID.Tågläge,
		Utgångsdag: färdplan.ID.Utgångsdag,
	}

	stops := färdplan.Färdplats
	if len(stops) == 0 {
		return result
	}

	for i, fp := range stops {
		isFirst := i == 0
		isLast := i == len(stops)-1

		platsspår := spårNamn(fp.Platsspår)
		linjespårAnkomst := linjespårNamn(fp.LinjespårAnkomst)
		linjespårAvgång := linjespårNamn(fp.LinjespårAvgång)

		// Arrival at this stop (not for the first stop)
		if !isFirst {
			prevFP := stops[i-1]
			prevLinjespårAvgång := linjespårNamn(prevFP.LinjespårAvgång)

			// Line segment: from previous stop to this stop
			lineSegment := JourneySegment{
				Type:       routes.TrackRouteLine,
				StartPlace: prevFP.Signatur,
				EndPlace:   fp.Signatur,
				StartTrack: prevLinjespårAvgång,
				EndTrack:   linjespårAnkomst,
			}
			lineSegment.TrackRoute = idx.lookup(routes.TrackRouteLine,
				prevFP.Signatur, fp.Signatur, prevLinjespårAvgång, linjespårAnkomst)
			result.Segments = append(result.Segments, lineSegment)

			// Arrival segment at this stop
			arrivalSegment := JourneySegment{
				Type:       routes.TrackRouteArrival,
				StartPlace: fp.Signatur,
				EndPlace:   fp.Signatur,
				StartTrack: linjespårAnkomst,
				EndTrack:   platsspår,
			}
			arrivalSegment.TrackRoute = idx.lookup(routes.TrackRouteArrival,
				fp.Signatur, fp.Signatur, linjespårAnkomst, platsspår)
			result.Segments = append(result.Segments, arrivalSegment)
		}

		// Departure from this stop (not for the last stop)
		if !isLast {
			departureSegment := JourneySegment{
				Type:       routes.TrackRouteDeparture,
				StartPlace: fp.Signatur,
				EndPlace:   fp.Signatur,
				StartTrack: platsspår,
				EndTrack:   linjespårAvgång,
			}
			departureSegment.TrackRoute = idx.lookup(routes.TrackRouteDeparture,
				fp.Signatur, fp.Signatur, platsspår, linjespårAvgång)
			result.Segments = append(result.Segments, departureSegment)
		}
	}

	return result
}

// PrintMatchResult prints a summary of the match result to stdout.
func PrintMatchResult(result *MatchResult) {
	fmt.Printf("Tåg %d (%s): %d segments\n", result.Tågläge, result.Utgångsdag, len(result.Segments))

	matched := 0
	unmatched := 0
	for _, seg := range result.Segments {
		status := "OK"
		if seg.TrackRoute == nil {
			status = "MISSING"
			unmatched++
		} else {
			matched++
		}

		fmt.Printf("  %-9s %-5s → %-5s  track %-4s → %-4s  [%s]\n",
			seg.Type, seg.StartPlace, seg.EndPlace,
			seg.StartTrack, seg.EndTrack, status)
	}

	fmt.Printf("  Matched: %d, Missing: %d\n", matched, unmatched)
}

// spårNamn safely extracts the track name, returning empty string if nil.
func spårNamn(s *pd.Spår) string {
	if s == nil {
		return ""
	}
	return s.Namn
}

// linjespårNamn safely extracts the line track name, returning empty string if nil.
func linjespårNamn(l *pd.Linjespår) string {
	if l == nil {
		return ""
	}
	return l.Namn
}
