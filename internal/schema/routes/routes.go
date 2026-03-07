// Package routes provides Go types for the Routes XML schema,
// containing SignalRoutes and TrackRoutes used by Journey Service
// to map between timetable data and SFERA Segment Profiles.
package routes

import "encoding/xml"

// Routes is the root XML element containing signal routes and track routes.
type Routes struct {
	XMLName      xml.Name      `xml:"Routes"`
	SignalRoutes []SignalRoute `xml:"SignalRoutes>SignalRoute"`
	TrackRoutes  []TrackRoute  `xml:"TrackRoutes>TrackRoute"`
}

// SignalRoute represents a route between two stop signals (signalsträcka).
// It contains the list of segment profile references that make up the route.
type SignalRoute struct {
	StartSignal        string                    `xml:"StartSignal,attr"`
	EndSignal          string                    `xml:"EndSignal,attr"`
	StartPlace         string                    `xml:"StartPlace,attr"`
	EndPlace           string                    `xml:"EndPlace,attr"`
	SegmentProfileRefs []SegmentProfileReference `xml:"SegmentProfileRefs>SegmentProfileRef"`
}

// SegmentProfileReference identifies a segment profile that is part of a route.
type SegmentProfileReference struct {
	SPID      string  `xml:"SPID,attr"`
	StartPos  float64 `xml:"StartPos,attr"`
	EndPos    float64 `xml:"EndPos,attr"`
	Direction int     `xml:"Direction,attr"`
}

// TrackRouteType classifies the type of track route.
type TrackRouteType string

const (
	// TrackRouteArrival is a route from an INF signal to an M signal inside a station.
	TrackRouteArrival TrackRouteType = "Arrival"
	// TrackRouteDeparture is a route from an M signal to a UBL signal inside a station.
	TrackRouteDeparture TrackRouteType = "Departure"
	// TrackRouteLine is a route from a UBL signal at one place to an INF signal at the next place.
	TrackRouteLine TrackRouteType = "Line"
)

// TrackRoute maps a track section (linjespår/stationsspår) to a sequence of signal routes,
// enabling lookup from timetable track data to segment profiles.
type TrackRoute struct {
	Type            TrackRouteType   `xml:"Type,attr"`
	StartPlace      string           `xml:"StartPlace,attr"`
	EndPlace        string           `xml:"EndPlace,attr"`
	StartTrack      string           `xml:"StartTrack,attr,omitempty"`
	EndTrack        string           `xml:"EndTrack,attr,omitempty"`
	StartSignal     string           `xml:"StartSignal,attr"`
	EndSignal       string           `xml:"EndSignal,attr"`
	SignalRouteRefs []SignalRouteRef `xml:"SignalRouteRefs>SignalRouteRef"`
}

// SignalRouteRef references a signal route by its start and end signals.
type SignalRouteRef struct {
	StartSignal string `xml:"StartSignal,attr"`
	EndSignal   string `xml:"EndSignal,attr"`
}
