package domainmodel

// EntitetMedUtbredning is a base type for entities with a spread across a link.
type EntitetMedUtbredning struct {
	Lankutbredning *Lankutbredning
}

// NewEntitetMedUtbredning creates a new EntitetMedUtbredning.
func NewEntitetMedUtbredning(lankutbredning *Lankutbredning) *EntitetMedUtbredning {
	return &EntitetMedUtbredning{Lankutbredning: lankutbredning}
}

// Gilltig checks if the entity is valid for the given link direction.
func (e *EntitetMedUtbredning) Gilltig(lankriktning Lankriktning) bool {
	return e.Lankutbredning.Gilltig(lankriktning)
}

// GetStartAvstand returns the start distance for the given link direction.
func (e *EntitetMedUtbredning) GetStartAvstand(lankriktning Lankriktning) float64 {
	return e.Lankutbredning.GetStartAvstand(lankriktning)
}

// GetSlutAvstand returns the end distance for the given link direction.
func (e *EntitetMedUtbredning) GetSlutAvstand(lankriktning Lankriktning) float64 {
	return e.Lankutbredning.GetSlutAvstand(lankriktning)
}
