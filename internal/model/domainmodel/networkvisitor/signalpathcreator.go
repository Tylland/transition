package networkvisitor

import (
	"slices"

	dm "github.com/tylland/transition/internal/model/domainmodel"
)

// SignalPathCreator traverses the network to find signal sections.
// It starts from a given signal and traverses in a direction until it finds
// the next stop signal, creating a Signalstracka from start to end.
type SignalPathCreator struct {
	startSignal    *dm.Signal
	startDirection dm.LinkDirection
	visitedLinks   []*dm.Link
	linkPaths      []*dm.LinkPath
	hasCompleted   bool

	// SignalPath is the result after traversal completes.
	SignalPath *dm.SignalPath
}

// NewSignalPathCreator creates a new SignalstrackaVisitor.
func NewSignalPathCreator(startSignal *dm.Signal, startDirection dm.LinkDirection) *SignalPathCreator {
	return &SignalPathCreator{
		startSignal:    startSignal,
		startDirection: startDirection,
		visitedLinks:   make([]*dm.Link, 0),
		linkPaths:      make([]*dm.LinkPath, 0),
	}
}

// Startposition returns the starting link position.
func (v *SignalPathCreator) StartPosition() *dm.LinkPosition {
	return v.startSignal.LinkPosition
}

// Startriktning returns the starting direction.
func (v *SignalPathCreator) StartDirection() dm.LinkDirection {
	return v.startDirection
}

// CanVisitLank checks if the link has not already been visited.
func (v *SignalPathCreator) CanVisitLink(lank *dm.Link) bool {
	return !slices.Contains(v.visitedLinks, lank)
}

// CanVisitNod always returns true.
func (v *SignalPathCreator) CanVisitNode(node *dm.Node) bool {
	return true
}

// LeaveNode is a no-op for this visitor.
func (v *SignalPathCreator) LeaveNode(node *dm.Node) {
}

// LeaveLinkPath removes the LinkPath and link from the tracked lists.
func (v *SignalPathCreator) LeaveLinkPath(path *dm.LinkPath) {
	for i, l := range v.visitedLinks {
		if l == path.Link {
			v.visitedLinks = append(v.visitedLinks[:i], v.visitedLinks[i+1:]...)
			break
		}
	}
	for i, la := range v.linkPaths {
		if la == path {
			v.linkPaths = append(v.linkPaths[:i], v.linkPaths[i+1:]...)
			return
		}
	}
}

// VisitNod always returns false (does not terminate traversal).
func (v *SignalPathCreator) VisitNode(node *dm.Node) bool {
	return false
}

// VisitLinkPath processes a link usage, looking for stop signals.
// Returns true if a stop signal was found (completing the signal section).
func (v *SignalPathCreator) VisitLinkPath(path *dm.LinkPath) bool {
	v.visitedLinks = append(v.visitedLinks, path.Link)

	signals := dm.EntitiesValidInPath(path.Link.Signals.Items(), path)

	var stopSignal *dm.Signal
	for _, s := range signals {
		if s.IsStopSignal() && s != v.startSignal {
			stopSignal = s
			break
		}
	}

	if stopSignal != nil {
		v.linkPaths = append(v.linkPaths, path.GetPath(stopSignal.LinkPosition))

		// Copy the list for the Signalstracka
		paths := make([]*dm.LinkPath, len(v.linkPaths))
		copy(paths, v.linkPaths)

		v.SignalPath = dm.NewSignalPath(v.startSignal, stopSignal, paths)
		v.hasCompleted = true
		return true
	}

	v.linkPaths = append(v.linkPaths, path)
	return false
}

// HasCompleted returns true if the visitor has found a complete signal section.
func (v *SignalPathCreator) HasCompleted() bool {
	return v.hasCompleted
}
