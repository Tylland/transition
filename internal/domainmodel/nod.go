package domainmodel

// Nod represents a node in the network.
type Nod struct {
	Id    NodId
	Namn  string
	vagar []*NodVag
}

// NewNod creates a new Nod.
func NewNod(id NodId, namn string) *Nod {
	return &Nod{
		Id:    id,
		Namn:  namn,
		vagar: make([]*NodVag, 0),
	}
}

// AddVag adds a node path connecting two links.
func (n *Nod) AddVag(lankA, lankB *Lank) {
	n.vagar = append(n.vagar, NewNodVag(n, lankA, lankB))
}

// Vagar returns the node paths.
func (n *Nod) Vagar() []*NodVag {
	return n.vagar
}

// GetLankar returns the links reachable from the given link through this node.
func (n *Nod) GetLankar(fran *Lank) []*Lank {
	var result []*Lank
	for _, nv := range n.vagar {
		if nv.LankA == fran {
			result = append(result, nv.LankB)
		}
	}
	return result
}

// String returns a string representation of the Nod.
func (n *Nod) String() string {
	return n.Id.String()
}

// NodVag represents a path through a node connecting two links.
type NodVag struct {
	Nod   *Nod
	LankA *Lank
	LankB *Lank
}

// NewNodVag creates a new NodVag.
func NewNodVag(nod *Nod, lankA, lankB *Lank) *NodVag {
	return &NodVag{Nod: nod, LankA: lankA, LankB: lankB}
}
