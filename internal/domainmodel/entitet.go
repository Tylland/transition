package domainmodel

// IEntitet defines the interface for entities with validity and distance.
type IEntitet interface {
	Gilltig(lankriktning Lankriktning) bool
	GetStartAvstand(lankriktning Lankriktning) float64
}

// Entiteter is a generic typed collection.
type Entiteter[T any] struct {
	items []T
}

// Add appends an item to the collection.
func (e *Entiteter[T]) Add(item T) {
	e.items = append(e.items, item)
}

// Items returns all items in the collection.
func (e *Entiteter[T]) Items() []T {
	return e.items
}

// Len returns the number of items in the collection.
func (e *Entiteter[T]) Len() int {
	return len(e.items)
}
