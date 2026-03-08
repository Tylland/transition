// Package planningdata provides Go types for the DagligFärdplan XML schema (v2.1),
// representing Swedish railway planning data (daily timetables).
package planningdata

import "encoding/xml"

// Response is the root XML element wrapping planning data results.
type Response struct {
	XMLName xml.Name `xml:"RESPONSE"`
	Results []Result `xml:"RESULT"`
}

// Result contains the daily timetables for one response.
type Result struct {
	DagligFärdplan []DagligFärdplan `xml:"DagligFärdplan"`
}

// DagligFärdplan represents a daily timetable for a single train.
type DagligFärdplan struct {
	ID        DagligFärdplanID `xml:"ID"`
	Färdplats []Färdplats      `xml:"Färdplats"`
}

// DagligFärdplanID identifies a daily timetable by train number and date.
type DagligFärdplanID struct {
	Tågläge    int    `xml:"Tågläge"`
	Utgångsdag string `xml:"Utgångsdag"`
}

// Färdplats represents a stop in the timetable with track and timing information.
type Färdplats struct {
	Signatur         string     `xml:"Signatur"`
	Namn             string     `xml:"Namn"`
	Ordningsnummer   int        `xml:"Ordningsnummer"`
	Ankomst          string     `xml:"Ankomst,omitempty"`
	Avgång           string     `xml:"Avgång,omitempty"`
	Platsspår        *Spår      `xml:"Platsspår,omitempty"`
	LinjespårAnkomst *Linjespår `xml:"LinjespårAnkomst,omitempty"`
	LinjespårAvgång  *Linjespår `xml:"LinjespårAvgång,omitempty"`
}

// Spår represents a station track (platsspår) with an ID and name.
type Spår struct {
	ID   string `xml:"ID,omitempty"`
	Namn string `xml:"Namn"`
}

// Linjespår represents a line track with a name.
type Linjespår struct {
	Namn string `xml:"Namn"`
}
