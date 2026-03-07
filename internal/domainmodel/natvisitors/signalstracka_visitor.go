package natvisitors

import dm "github.com/tylland/transition/internal/domainmodel"

// SignalstrackaVisitor traverses the network to find signal sections.
// It starts from a given signal and traverses in a direction until it finds
// the next stop signal, creating a Signalstracka from start to end.
type SignalstrackaVisitor struct {
	startSignal      *dm.Signal
	startriktning    dm.Lankriktning
	besöktaLankar    []*dm.Lank
	lankanvandningar []*dm.Lankanvandning
	hasCompleted     bool

	// Signalstracka is the result after traversal completes.
	Signalstracka *dm.Signalstracka
}

// NewSignalstrackaVisitor creates a new SignalstrackaVisitor.
func NewSignalstrackaVisitor(startSignal *dm.Signal, lankriktning dm.Lankriktning) *SignalstrackaVisitor {
	return &SignalstrackaVisitor{
		startSignal:      startSignal,
		startriktning:    lankriktning,
		besöktaLankar:    make([]*dm.Lank, 0),
		lankanvandningar: make([]*dm.Lankanvandning, 0),
	}
}

// Startposition returns the starting link position.
func (v *SignalstrackaVisitor) Startposition() *dm.Lankposition {
	return v.startSignal.Lankposition
}

// Startriktning returns the starting direction.
func (v *SignalstrackaVisitor) Startriktning() dm.Lankriktning {
	return v.startriktning
}

// CanVisitLank checks if the link has not already been visited.
func (v *SignalstrackaVisitor) CanVisitLank(lank *dm.Lank) bool {
	for _, l := range v.besöktaLankar {
		if l == lank {
			return false
		}
	}
	return true
}

// CanVisitNod always returns true.
func (v *SignalstrackaVisitor) CanVisitNod(nod *dm.Nod) bool {
	return true
}

// LeaveNod is a no-op for this visitor.
func (v *SignalstrackaVisitor) LeaveNod(nod *dm.Nod) {
}

// LeaveLankanvandning removes the Lankanvandning and link from the tracked lists.
func (v *SignalstrackaVisitor) LeaveLankanvandning(anvandning *dm.Lankanvandning) {
	for i, l := range v.besöktaLankar {
		if l == anvandning.Lank {
			v.besöktaLankar = append(v.besöktaLankar[:i], v.besöktaLankar[i+1:]...)
			break
		}
	}
	for i, la := range v.lankanvandningar {
		if la == anvandning {
			v.lankanvandningar = append(v.lankanvandningar[:i], v.lankanvandningar[i+1:]...)
			return
		}
	}
}

// VisitNod always returns false (does not terminate traversal).
func (v *SignalstrackaVisitor) VisitNod(nod *dm.Nod) bool {
	return false
}

// VisitLankanvandning processes a link usage, looking for stop signals.
// Returns true if a stop signal was found (completing the signal section).
func (v *SignalstrackaVisitor) VisitLankanvandning(anvandning *dm.Lankanvandning) bool {
	v.besöktaLankar = append(v.besöktaLankar, anvandning.Lank)

	signals := dm.GilltigaInom(anvandning.Lank.Signaler.Items(), anvandning)

	var stopSignal *dm.Signal
	for _, s := range signals {
		if s.IsStopSignal() && s != v.startSignal {
			stopSignal = s
			break
		}
	}

	if stopSignal != nil {
		v.lankanvandningar = append(v.lankanvandningar, anvandning.GetLankanvandning(stopSignal.Lankposition))

		// Copy the list for the Signalstracka
		anvandningar := make([]*dm.Lankanvandning, len(v.lankanvandningar))
		copy(anvandningar, v.lankanvandningar)

		v.Signalstracka = dm.NewSignalstracka(v.startSignal, stopSignal, anvandningar)
		v.hasCompleted = true
		return true
	}

	v.lankanvandningar = append(v.lankanvandningar, anvandning)
	return false
}

// HasCompleted returns true if the visitor has found a complete signal section.
func (v *SignalstrackaVisitor) HasCompleted() bool {
	return v.hasCompleted
}
