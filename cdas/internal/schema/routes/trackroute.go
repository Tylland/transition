package routes

import "encoding/xml"

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

// TrackRoutes is the root XML element for track routes.
type TrackRoutes struct {
	XMLName xml.Name     `xml:"TrackRoutes"`
	Routes  []TrackRoute `xml:"TrackRoute"`
}

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
