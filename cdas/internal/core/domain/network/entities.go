package network

import "sort"

// PositionedEntity is a constraint for entities that have a position on a link.
type PositionedEntity interface {
	IsValid(LinkDirection) bool
	GetLinkposition() *LinkPosition
}

// EntitiesValidInPath returns positioned entities that are valid within the given LinkPath,
// sorted by distance in the direction of the LinkPath.
func EntitiesValidInPath[T PositionedEntity](items []T, la *LinkPath) []T {
	var result []T
	for _, e := range items {
		if e.IsValid(la.Direction) && la.Contains(e.GetLinkposition()) {
			result = append(result, e)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].GetLinkposition().GetDistance(la.Direction) <
			result[j].GetLinkposition().GetDistance(la.Direction)
	})
	return result
}

// ExtentEntitiesInDirection returns spread entities valid for the given direction, sorted by start distance.
func ExtentEntitiesInDirection[T IEntity](entities []T, direction LinkDirection) []T {
	var result []T
	for _, entity := range entities {
		if entity.IsValid(direction) {
			result = append(result, entity)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].GetStartDistance(direction) < result[j].GetStartDistance(direction)
	})
	return result
}

// PositionEntitiesInDirection returns positioned entities valid for the given direction, sorted by distance.
func PositionEntitiesInDirection[T PositionedEntity](entities []T, direction LinkDirection) []T {
	var result []T
	for _, entity := range entities {
		if entity.IsValid(direction) {
			result = append(result, entity)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].GetLinkposition().GetDistance(direction) <
			result[j].GetLinkposition().GetDistance(direction)
	})
	return result
}
