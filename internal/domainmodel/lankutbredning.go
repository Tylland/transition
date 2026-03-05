package domainmodel

import "fmt"

// Lankutbredning represents a spread (extent) across a link.
type Lankutbredning struct {
	Lank     *Lank
	Start    float64
	Slut     float64
	Riktning Gilltighetsriktning
}

// NewLankutbredning creates a new Lankutbredning.
func NewLankutbredning(lank *Lank, start, slut float64, riktning Gilltighetsriktning) *Lankutbredning {
	return &Lankutbredning{
		Lank:     lank,
		Start:    start,
		Slut:     slut,
		Riktning: riktning,
	}
}

// GetStart returns the start position adjusted for the given direction.
func (lu *Lankutbredning) GetStart(lankriktning Lankriktning) float64 {
	switch lankriktning {
	case LankriktningMed:
		return lu.Start
	case LankriktningMot:
		return 1.0 - lu.Slut
	default:
		panic("Länkriktning okänd")
	}
}

// GetSlut returns the end position adjusted for the given direction.
func (lu *Lankutbredning) GetSlut(lankriktning Lankriktning) float64 {
	switch lankriktning {
	case LankriktningMed:
		return lu.Slut
	case LankriktningMot:
		return 1.0 - lu.Start
	default:
		panic("Länkriktning okänd")
	}
}

// GetStartAvstand returns the start distance for the given direction.
func (lu *Lankutbredning) GetStartAvstand(lankriktning Lankriktning) float64 {
	return lu.Lank.Langd * lu.GetStart(lankriktning)
}

// GetSlutAvstand returns the end distance for the given direction.
func (lu *Lankutbredning) GetSlutAvstand(lankriktning Lankriktning) float64 {
	return lu.Lank.Langd * lu.GetSlut(lankriktning)
}

// Gilltig checks if the spread is valid for the given direction.
func (lu *Lankutbredning) Gilltig(lankriktning Lankriktning) bool {
	return lu.Riktning.Gilltig(lankriktning)
}

// String returns a string representation.
func (lu *Lankutbredning) String() string {
	return fmt.Sprintf("%v->%v %s", lu.Start, lu.Slut, lu.Riktning.ToSign())
}
