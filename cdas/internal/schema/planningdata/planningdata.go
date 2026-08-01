// Package planningdata provides Go types for the DagligFärdplan JSON schema (v2.1),
// representing Swedish railway planning data (daily timetables).
package planningdata

// JSONRoot is the top-level JSON wrapper: { "RESPONSE": { ... } }
type JSONRoot struct {
	Response Response `json:"RESPONSE"`
}

// Response wraps the result array.
type Response struct {
	Results []Result `json:"RESULT"`
}

// Result contains the daily timetables for one response.
type Result struct {
	DagligFärdplan []DagligFärdplan `json:"DagligFärdplan"`
}

// DagligFärdplan represents a daily timetable for a single train.
type DagligFärdplan struct {
	ID        DagligFärdplanID `json:"ID"`
	Färdplats []Färdplats      `json:"Färdplats"`
}

// DagligFärdplanID identifies a daily timetable by train number and date.
type DagligFärdplanID struct {
	Tågläge    int    `json:"Tågläge"`
	Utgångsdag string `json:"Utgångsdag"`
}

// Färdplats represents a stop in the timetable with track and timing information.
type Färdplats struct {
	Signatur         string     `json:"Signatur"`
	Namn             string     `json:"Namn"`
	Ordningsnummer   int        `json:"Ordningsnummer"`
	Ankomst          string     `json:"Ankomst,omitempty"`
	Avgång           string     `json:"Avgång,omitempty"`
	Platsspår        *Spår      `json:"Platsspår,omitempty"`
	LinjespårAnkomst *Linjespår `json:"LinjespårAnkomst,omitempty"`
	LinjespårAvgång  *Linjespår `json:"LinjespårAvgång,omitempty"`
}

// Spår represents a station track (platsspår) with an ID and name.
type Spår struct {
	ID   string `json:"ID,omitempty"`
	Namn string `json:"Namn"`
}

// Linjespår represents a line track with a name.
type Linjespår struct {
	Namn string `json:"Namn"`
}
