package domainmodel

import "sort"

// PositionedEntity is a constraint for entities that have a position on a link.
type PositionedEntity interface {
	Gilltig(Lankriktning) bool
	GetLankposition() *Lankposition
}

// GilltigaInom returns positioned entities that are valid within the given Lankanvandning,
// sorted by distance in the direction of the Lankanvandning.
func GilltigaInom[T PositionedEntity](items []T, la *Lankanvandning) []T {
	var result []T
	for _, e := range items {
		if e.Gilltig(la.Riktning) && la.Inehaller(e.GetLankposition()) {
			result = append(result, e)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].GetLankposition().GetAvstand(la.Riktning) <
			result[j].GetLankposition().GetAvstand(la.Riktning)
	})
	return result
}

// GilltigaUtbredningar returns spread entities valid for the given direction, sorted by start distance.
func GilltigaUtbredningar[T IEntitet](items []T, lr Lankriktning) []T {
	var result []T
	for _, e := range items {
		if e.Gilltig(lr) {
			result = append(result, e)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].GetStartAvstand(lr) < result[j].GetStartAvstand(lr)
	})
	return result
}

// GilltigaPositioner returns positioned entities valid for the given direction, sorted by distance.
func GilltigaPositioner[T PositionedEntity](items []T, lr Lankriktning) []T {
	var result []T
	for _, e := range items {
		if e.Gilltig(lr) {
			result = append(result, e)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].GetLankposition().GetAvstand(lr) <
			result[j].GetLankposition().GetAvstand(lr)
	})
	return result
}
