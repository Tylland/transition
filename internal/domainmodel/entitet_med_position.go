package domainmodel

// EntitetMedPosition is a base type for entities with a position on a link.
type EntitetMedPosition struct {
	Lankposition *Lankposition
}

// NewEntitetMedPosition creates a new EntitetMedPosition.
func NewEntitetMedPosition(lankposition *Lankposition) *EntitetMedPosition {
	return &EntitetMedPosition{Lankposition: lankposition}
}

// Gilltig checks if the entity is valid for the given link direction.
func (e *EntitetMedPosition) Gilltig(lankriktning Lankriktning) bool {
	return e.Lankposition.Gilltig(lankriktning)
}

// GetStartAvstand returns the start distance for the given link direction.
func (e *EntitetMedPosition) GetStartAvstand(lankriktning Lankriktning) float64 {
	return e.Lankposition.GetAvstand(lankriktning)
}

// GetLankposition returns the link position.
func (e *EntitetMedPosition) GetLankposition() *Lankposition {
	return e.Lankposition
}
