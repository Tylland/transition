package domainmodel

// EntityWithPosition is a base type for entities with a position on a link.
type EntityWithPosition struct {
	*LinkPosition
}

// NewEntityWithPosition creates a new EntitetMedPosition.
func NewEntityWithPosition(lankposition *LinkPosition) *EntityWithPosition {
	return &EntityWithPosition{LinkPosition: lankposition}
}

// IsValid checks if the entity is valid for the given link direction.
func (e *EntityWithPosition) IsValid(direction LinkDirection) bool {
	return e.LinkPosition.IsValid(direction)
}

// GetStartDistance returns the start distance for the given link direction.
func (e *EntityWithPosition) GetStartDistance(lankriktning LinkDirection) float64 {
	return e.LinkPosition.GetDistance(lankriktning)
}

// GetLinkposition returns the link position.
func (e *EntityWithPosition) GetLinkposition() *LinkPosition {
	return e.LinkPosition
}
