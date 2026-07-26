package domainmodel

// IEntity defines the interface for entities with validity and distance.
type IEntity interface {
	IsValid(direction LinkDirection) bool
	GetStartDistance(direction LinkDirection) float64
}

// Entities is a generic typed collection.
type Entities[T any] struct {
	items []T
}

// Add appends an item to the collection.
func (e *Entities[T]) Add(item T) {
	e.items = append(e.items, item)
}

// Items returns all items in the collection.
func (e *Entities[T]) Items() []T {
	return e.items
}

// Len returns the number of items in the collection.
func (e *Entities[T]) Len() int {
	return len(e.items)
}
