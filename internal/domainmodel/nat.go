package domainmodel

import (
	"fmt"
	"sort"
)

// NatVisitor is the interface for traversing the network.
// Method names differ from C# due to Go not supporting method overloading:
//   - CanVisit(Lank) / CanVisit(Nod) → CanVisitLank / CanVisitNod
//   - Visit(Lankanvandning) / Visit(Nod) → VisitLankanvandning / VisitNod
//   - Leave(Lankanvandning) / Leave(Nod) → LeaveLankanvandning / LeaveNod
type NatVisitor interface {
	Startposition() *Lankposition
	Startriktning() Lankriktning

	CanVisitLank(lank *Lank) bool
	CanVisitNod(nod *Nod) bool

	VisitLankanvandning(anvandning *Lankanvandning) bool
	VisitNod(nod *Nod) bool

	LeaveLankanvandning(anvandning *Lankanvandning)
	LeaveNod(nod *Nod)

	HasCompleted() bool
}

// Nat represents the network (nät).
type Nat struct {
	noder   map[NodId]*Nod
	lankar  map[LankId]*Lank
	platser map[int]*Plats
}

// NewNat creates a new empty Nat.
func NewNat() *Nat {
	return &Nat{
		noder:   make(map[NodId]*Nod),
		lankar:  make(map[LankId]*Lank),
		platser: make(map[int]*Plats),
	}
}

// Noder returns all nodes in the network, sorted by NodId for deterministic iteration.
func (n *Nat) Noder() []*Nod {
	result := make([]*Nod, 0, len(n.noder))
	for _, nod := range n.noder {
		result = append(result, nod)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Id.TrafikplatsId != result[j].Id.TrafikplatsId {
			return result[i].Id.TrafikplatsId < result[j].Id.TrafikplatsId
		}
		return result[i].Id.Nummer < result[j].Id.Nummer
	})
	return result
}

// GetOrCreateNod gets an existing node or creates a new one.
// An optional name can be provided; defaults to "-".
func (n *Nat) GetOrCreateNod(nodId NodId, namn ...string) *Nod {
	name := "-"
	if len(namn) > 0 {
		name = namn[0]
	}
	if _, ok := n.noder[nodId]; !ok {
		n.noder[nodId] = NewNod(nodId, name)
	}
	return n.noder[nodId]
}

// SaveNod stores a node if it doesn't already exist.
func (n *Nat) SaveNod(nodId NodId, namn string) {
	if _, ok := n.noder[nodId]; !ok {
		n.noder[nodId] = NewNod(nodId, namn)
	}
}

// Lankar returns all links in the network, sorted by LankId for deterministic iteration.
func (n *Nat) Lankar() []*Lank {
	result := make([]*Lank, 0, len(n.lankar))
	for _, lank := range n.lankar {
		result = append(result, lank)
	}
	sort.Slice(result, func(i, j int) bool {
		li, lj := result[i].Id, result[j].Id
		if li.Fran.TrafikplatsId != lj.Fran.TrafikplatsId {
			return li.Fran.TrafikplatsId < lj.Fran.TrafikplatsId
		}
		if li.Fran.Nummer != lj.Fran.Nummer {
			return li.Fran.Nummer < lj.Fran.Nummer
		}
		if li.Till.TrafikplatsId != lj.Till.TrafikplatsId {
			return li.Till.TrafikplatsId < lj.Till.TrafikplatsId
		}
		if li.Till.Nummer != lj.Till.Nummer {
			return li.Till.Nummer < lj.Till.Nummer
		}
		return li.Sekvensnummer < lj.Sekvensnummer
	})
	return result
}

// GetOrCreateLank gets an existing link or creates a new one.
func (n *Nat) GetOrCreateLank(lankId LankId, langd float64) *Lank {
	if _, ok := n.lankar[lankId]; !ok {
		fran := n.GetOrCreateNod(lankId.Fran)
		till := n.GetOrCreateNod(lankId.Till)
		newLank := NewLank(lankId, fran, till, langd)
		n.lankar[lankId] = newLank
	}
	return n.lankar[lankId]
}

// AddLank adds a link to the network.
func (n *Nat) AddLank(lank *Lank) {
	n.lankar[lank.Id] = lank
}

// GetLank retrieves a link by its ID. Panics if not found.
func (n *Nat) GetLank(lankId LankId) *Lank {
	lank, ok := n.lankar[lankId]
	if !ok {
		panic(fmt.Sprintf("Lank not found: %v", lankId))
	}
	return lank
}

// WithLankStr executes an action on a link identified by its string ID.
func (n *Nat) WithLankStr(id string, action func(*Lank)) error {
	lankId, err := ParseLankId(id)
	if err != nil {
		return err
	}
	n.WithLank(lankId, action)
	return nil
}

// WithLank executes an action on a link identified by its LankId.
func (n *Nat) WithLank(id LankId, action func(*Lank)) {
	action(n.GetLank(id))
}

// Platser returns all places in the network, sorted by PLC for deterministic iteration.
func (n *Nat) Platser() []*Plats {
	result := make([]*Plats, 0, len(n.platser))
	for _, plats := range n.platser {
		result = append(result, plats)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].PLC < result[j].PLC
	})
	return result
}

// AddPlats adds a place to the network.
func (n *Nat) AddPlats(plats *Plats) {
	n.platser[plats.PLC] = plats
}

// Traverse traverses the network using the given visitor starting from its start position.
func Traverse(visitor NatVisitor) {
	nod := visitor.Startposition().Lank.GetNod(visitor.Startriktning())
	anvandning := visitor.Startposition().GetAnvandning(visitor.Startriktning())
	traverseRecursive(anvandning, nod, visitor)
}

func traverseRecursive(anvandning *Lankanvandning, nod *Nod, visitor NatVisitor) {
	if anvandning.Lank == nil {
		return
	}

	if !visitor.CanVisitLank(anvandning.Lank) {
		return
	}

	pathCompleted := visitor.VisitLankanvandning(anvandning)

	if !pathCompleted && visitor.CanVisitNod(nod) {
		pathCompleted = visitor.VisitNod(nod)

		if !pathCompleted && !visitor.HasCompleted() {
			lankar := nod.GetLankar(anvandning.Lank)
			for _, lank := range lankar {
				if !visitor.HasCompleted() {
					otherNod := lank.GetOtherNod(nod)
					traverseRecursive(
						NewLankanvandningFromLank(lank, lank.GetLankriktning(nod)),
						otherNod,
						visitor,
					)
				}
			}
		}

		visitor.LeaveNod(nod)
	}

	visitor.LeaveLankanvandning(anvandning)
}
