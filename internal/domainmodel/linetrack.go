package domainmodel

// LineTrack represents a line track (linjespår) with a spread across a link.
// Line tracks identify which track on the line a link belongs to (e.g. "E" for single,
// "U" for up-track, "N" for down-track).
type LineTrack struct {
	*EntityWithExtent
	Number string
	Name   string
}

// NewLineTrack creates a new Linjespar.
func NewLineTrack(number, name string, extent *LinkExtent) *LineTrack {
	return &LineTrack{
		EntityWithExtent: NewEntityWithExtent(extent),
		Number:           number,
		Name:             name,
	}
}
