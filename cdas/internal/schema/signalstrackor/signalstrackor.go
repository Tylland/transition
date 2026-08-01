// Package signalstrackor provides Go types for the Signalstrackor v0.10 XML schema.
package signalstrackor

import "encoding/xml"

// Signalstrackor is the root element containing signal sections.
type Signalstrackor struct {
	XMLName       xml.Name           `xml:"Signalstrackor"`
	Signalstracka []SignalstrackaTyp `xml:"Signalstracka"`
}

// SignalstrackaTyp represents a signal section.
type SignalstrackaTyp struct {
	Start string `xml:"Start,attr,omitempty"`
	Slut  string `xml:"Slut,attr,omitempty"`
}
