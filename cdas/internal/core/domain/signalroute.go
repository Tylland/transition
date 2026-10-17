package domain

// Package routes provides Go types for the Routes XML schema,
// containing SignalRoutes and TrackRoutes used by Journey Service
// to map between timetable data and SFERA Segment Profiles.

// SignalRoutes is the root XML element for signal routes.
type SignalRoutes struct {
	Routes  []SignalRoute `xml:"SignalRoute"`
}

// SignalRoute represents a route between two stop signals (signalstrÃ¤cka).
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
	Direction string  `xml:"Direction,attr"`
}

