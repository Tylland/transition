package domainmodel

import (
	"fmt"
	"math"
)

// Lankanvandning represents the usage of a link (or part of a link).
type Lankanvandning struct {
	Lank     *Lank
	Start    float64
	Slut     float64
	Riktning Lankriktning
}

// NewLankanvandning creates a new Lankanvandning with explicit start and end positions.
func NewLankanvandning(lank *Lank, start, slut float64, riktning Lankriktning) *Lankanvandning {
	return &Lankanvandning{
		Lank:     lank,
		Start:    start,
		Slut:     slut,
		Riktning: riktning,
	}
}

// NewLankanvandningFromLank creates a Lankanvandning covering the entire link in the given direction.
func NewLankanvandningFromLank(lank *Lank, riktning Lankriktning) *Lankanvandning {
	start := 0.0
	slut := 1.0
	if riktning == LankriktningMot {
		start = 1.0
		slut = 0.0
	}
	return &Lankanvandning{
		Lank:     lank,
		Start:    start,
		Slut:     slut,
		Riktning: riktning,
	}
}

// GetStartAvstand returns the start distance.
func (la *Lankanvandning) GetStartAvstand() float64 {
	return la.Lank.Langd * la.Start
}

// GetSlutAvstand returns the end distance.
func (la *Lankanvandning) GetSlutAvstand() float64 {
	return la.Lank.Langd * la.Slut
}

// Inehaller checks if the given link position is contained within this usage.
func (la *Lankanvandning) Inehaller(lankposition *Lankposition) bool {
	if la.Lank != lankposition.Lank {
		return false
	}
	min := math.Min(la.Start, la.Slut)
	max := math.Max(la.Start, la.Slut)
	return min <= lankposition.Position && lankposition.Position <= max
}

// GetLankanvandning returns a truncated usage from this usage's start up to the given position.
func (la *Lankanvandning) GetLankanvandning(lankposition *Lankposition) *Lankanvandning {
	return NewLankanvandning(la.Lank, la.Start, lankposition.Position, la.Riktning)
}

// String returns a string representation.
func (la *Lankanvandning) String() string {
	return fmt.Sprintf("%v->%v %v", la.Start, la.Slut, la.Riktning)
}
