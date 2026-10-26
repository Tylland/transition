package visitor

import (
	"slices"

	dm "github.com/tylland/transition/cdas/internal/core/domain/network"
)

// PlaceFinder traverses the network to find signal sections.
// It starts from a given signal and traverses in a direction until it finds
// the next stop signal, creating a Signalstracka from start to end.
type PlaceFinder struct {
	startSignal    *dm.Signal
	startDirection dm.LinkDirection
	reverse        bool
	visitedLinks   []*dm.Link
	hasCompleted   bool

	// Place is the result after traversal completes.
	placeSignature string
}

// NewPlaceFinder creates a new PlaceFinder.
func NewPlaceFinder(startSignal *dm.Signal, reverse bool) *PlaceFinder {
	startDirection, _ := startSignal.LinkPosition.Direction.ToLinkDirection()

	if reverse {
		startDirection = startDirection.Reverse()
	}

	return &PlaceFinder{
		startSignal:    startSignal,
		startDirection: startDirection,
		reverse:        reverse,
		visitedLinks:   make([]*dm.Link, 0),
	}
}

// Startposition returns the starting link position.
func (v *PlaceFinder) StartPosition() *dm.LinkPosition {
	return v.startSignal.LinkPosition
}

// Startriktning returns the starting direction.
func (v *PlaceFinder) StartDirection() dm.LinkDirection {
	return v.startDirection
}

// CanVisitLank checks if the link has not already been visited.
func (v *PlaceFinder) CanVisitLink(lank *dm.Link) bool {
	return !slices.Contains(v.visitedLinks, lank)
}

// CanVisitNod always returns true.
func (v *PlaceFinder) CanVisitNode(node *dm.Node) bool {
	return true
}

// LeaveNode is a no-op for this visitor.
func (v *PlaceFinder) LeaveNode(node *dm.Node) {
}

// LeaveLinkPath removes the LinkPath and link from the tracked lists.
func (v *PlaceFinder) LeaveLinkPath(path *dm.LinkPath) {
	for i, l := range v.visitedLinks {
		if l == path.Link {
			v.visitedLinks = append(v.visitedLinks[:i], v.visitedLinks[i+1:]...)
			break
		}
	}
}

// VisitNod always returns false (does not terminate traversal).
func (v *PlaceFinder) VisitNode(node *dm.Node) bool {
	return false
}

// VisitLinkPath processes a link usage, looking for stop signals.
// Returns true if a stop signal was found (completing the signal section).
func (v *PlaceFinder) VisitLinkPath(path *dm.LinkPath) bool {
	v.visitedLinks = append(v.visitedLinks, path.Link)

	signals := dm.EntitiesValidInPath(path.Link.Signals.Items(), path.InDirection(v.reverse))

	var stopSignal *dm.Signal
	for _, s := range signals {
		if s.IsStopSignal() && s != v.startSignal {
			stopSignal = s
			break
		}
	}

	if stopSignal != nil && stopSignal.Id.PlaceSignature != v.startSignal.Id.PlaceSignature {
		v.placeSignature = stopSignal.Id.PlaceSignature
		v.hasCompleted = true

		return true
	}

	return false
}

// HasCompleted returns true if the visitor has found a complete signal section.
func (v *PlaceFinder) HasCompleted() bool {
	return v.hasCompleted
}

func (v *PlaceFinder) PlaceSignature() string {
	return v.placeSignature
}
