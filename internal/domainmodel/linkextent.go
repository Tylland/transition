package domainmodel

import "fmt"

// LinkExtent represents a spread (extent) across a link.
type LinkExtent struct {
	Link      *Link
	Start     float64
	End       float64
	Direction ValidityDirection
}

// NewLinkExtent creates a new Lankutbredning.
func NewLinkExtent(link *Link, start, end float64, riktning ValidityDirection) *LinkExtent {
	return &LinkExtent{
		Link:      link,
		Start:     start,
		End:       end,
		Direction: riktning,
	}
}

// GetStart returns the start position adjusted for the given direction.
func (le *LinkExtent) GetStart(direction LinkDirection) float64 {
	switch direction {
	case LinkDirectionWith:
		return le.Start
	case LinkDirectionAgainst:
		return 1.0 - le.End
	default:
		panic("Länkriktning okänd")
	}
}

// GetEnd returns the end position adjusted for the given direction.
func (le *LinkExtent) GetEnd(direction LinkDirection) float64 {
	switch direction {
	case LinkDirectionWith:
		return le.End
	case LinkDirectionAgainst:
		return 1.0 - le.Start
	default:
		panic("Länkriktning okänd")
	}
}

// GetStartDistance returns the start distance for the given direction.
func (le *LinkExtent) GetStartDistance(direction LinkDirection) float64 {
	return le.Link.Length * le.GetStart(direction)
}

// GetEndDistance returns the end distance for the given direction.
func (le *LinkExtent) GetEndDistance(direction LinkDirection) float64 {
	return le.Link.Length * le.GetEnd(direction)
}

// IsValid checks if the spread is valid for the given direction.
func (le *LinkExtent) IsValid(direction LinkDirection) bool {
	return le.Direction.IsValid(direction)
}

// String returns a string representation.
func (le *LinkExtent) String() string {
	return fmt.Sprintf("%v->%v %s", le.Start, le.End, le.Direction.ToSign())
}
