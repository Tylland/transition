package network

// EntityWithExtent is a base type for entities with a spread across a link.
type EntityWithExtent struct {
	Extent *LinkExtent
}

// NewEntityWithExtent creates a new EntitetMedUtbredning.
func NewEntityWithExtent(direction *LinkExtent) *EntityWithExtent {
	return &EntityWithExtent{Extent: direction}
}

// Gilltig checks if the entity is valid for the given link direction.
func (e *EntityWithExtent) IsValid(direction LinkDirection) bool {
	return e.Extent.IsValid(direction)
}

// GetStartAvstand returns the start distance for the given link direction.
func (e *EntityWithExtent) GetStartDistance(direction LinkDirection) float64 {
	return e.Extent.GetStartDistance(direction)
}

// GetSlutAvstand returns the end distance for the given link direction.
func (e *EntityWithExtent) GetEndDistance(direction LinkDirection) float64 {
	return e.Extent.GetEndDistance(direction)
}
