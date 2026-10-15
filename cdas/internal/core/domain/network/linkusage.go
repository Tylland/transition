package network

import (
	"fmt"
	"math"
)

// LinkPath represents the usage of a link (or part of a link).
type LinkPath struct {
	Link      *Link
	Start     float64
	End       float64
	Direction LinkDirection
}

func NewLinkPath(link *Link, start, end float64, direction LinkDirection) *LinkPath {
	return &LinkPath{
		Link:      link,
		Start:     start,
		End:       end,
		Direction: direction,
	}
}

// NewLinkPathFromLink creates a LinkPath covering the entire link in the given direction.
func NewLinkPathFromLink(lank *Link, direction LinkDirection) *LinkPath {
	start := 0.0
	slut := 1.0
	if direction == LinkDirectionAgainst {
		start = 1.0
		slut = 0.0
	}
	return &LinkPath{
		Link:      lank,
		Start:     start,
		End:       slut,
		Direction: direction,
	}
}

func (lp *LinkPath) GetStartDistance() float64 {
	return lp.Link.Length * lp.Start
}

func (lp *LinkPath) GetEndDistance() float64 {
	return lp.Link.Length * lp.End
}

// Contains checks if the given link position is contained within this path.
func (lp *LinkPath) Contains(position *LinkPosition) bool {
	if lp.Link != position.Link {
		return false
	}
	min := math.Min(lp.Start, lp.End)
	max := math.Max(lp.Start, lp.End)
	return min <= position.Position && position.Position <= max
}

// GetPath returns a truncated usage from this usage's start up to the given position.
func (lp *LinkPath) GetPath(position *LinkPosition) *LinkPath {
	return NewLinkPath(lp.Link, lp.Start, position.Position, lp.Direction)
}

func (lp *LinkPath) String() string {
	return fmt.Sprintf("%v->%v %v", lp.Start, lp.End, lp.Direction)
}
