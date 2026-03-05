package domainmodel

// Lankposition represents a position on a link.
type Lankposition struct {
	Lank     *Lank
	Position float64
	Riktning Gilltighetsriktning
}

// NewLankposition creates a new Lankposition.
func NewLankposition(lank *Lank, position float64, riktning Gilltighetsriktning) *Lankposition {
	return &Lankposition{
		Lank:     lank,
		Position: position,
		Riktning: riktning,
	}
}

// Gilltig checks if the position is valid for the given link direction.
func (lp *Lankposition) Gilltig(lankriktning Lankriktning) bool {
	return lp.Riktning.Gilltig(lankriktning)
}

// GetPosition returns the position adjusted for the given direction.
func (lp *Lankposition) GetPosition(lankriktning Lankriktning) float64 {
	switch lankriktning {
	case LankriktningMed:
		return lp.Position
	case LankriktningMot:
		return 1.0 - lp.Position
	default:
		panic("Länkriktning okänd")
	}
}

// GetAvstand returns the distance adjusted for the given direction.
func (lp *Lankposition) GetAvstand(lankriktning Lankriktning) float64 {
	return lp.Lank.Langd * lp.GetPosition(lankriktning)
}

// GetAnvandning returns a link usage from this position in the given direction.
func (lp *Lankposition) GetAnvandning(lankriktning Lankriktning) *Lankanvandning {
	switch lankriktning {
	case LankriktningMed:
		return NewLankanvandning(lp.Lank, lp.Position, 1.0, lankriktning)
	case LankriktningMot:
		return NewLankanvandning(lp.Lank, lp.Position, 0.0, lankriktning)
	default:
		panic("Länkriktning okänd")
	}
}
