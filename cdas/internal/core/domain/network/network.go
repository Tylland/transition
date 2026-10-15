package network

import (
	"fmt"
	"sort"
)

// NetworkVisitor is the interface for traversing the network.
type NetworkVisitor interface {
	StartPosition() *LinkPosition
	StartDirection() LinkDirection

	CanVisitLink(*Link) bool
	CanVisitNode(*Node) bool

	VisitLinkPath(*LinkPath) bool
	VisitNode(*Node) bool

	LeaveLinkPath(*LinkPath)
	LeaveNode(*Node)

	HasCompleted() bool
}

// Network represents the network (nät).
type Network struct {
	nodes  map[NodeId]*Node
	links  map[LinkId]*Link
	places map[int]*Place
}

func NewNetwork() *Network {
	return &Network{
		nodes:  make(map[NodeId]*Node),
		links:  make(map[LinkId]*Link),
		places: make(map[int]*Place),
	}
}

// Nodes returns all nodes in the network, sorted by NodId for deterministic iteration.
func (n *Network) Nodes() []*Node {
	result := make([]*Node, 0, len(n.nodes))
	for _, node := range n.nodes {
		result = append(result, node)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Id.LocationNumber != result[j].Id.LocationNumber {
			return result[i].Id.LocationNumber < result[j].Id.LocationNumber
		}
		return result[i].Id.Number < result[j].Id.Number
	})
	return result
}

// GetOrCreateNode gets an existing node or creates a new one.
// An optional name can be provided; defaults to "-".
func (n *Network) GetOrCreateNode(nodeId NodeId, names ...string) *Node {
	name := "-"
	if len(names) > 0 {
		name = names[0]
	}
	if _, ok := n.nodes[nodeId]; !ok {
		n.nodes[nodeId] = NewNode(nodeId, name)
	}
	return n.nodes[nodeId]
}

// SaveNode stores a node if it doesn't already exist.
func (n *Network) SaveNode(nodeId NodeId, name string) {
	if _, ok := n.nodes[nodeId]; !ok {
		n.nodes[nodeId] = NewNode(nodeId, name)
	}
}

// Links returns all links in the network, sorted by LinkId for deterministic iteration.
func (n *Network) Links() []*Link {
	result := make([]*Link, 0, len(n.links))
	for _, link := range n.links {
		result = append(result, link)
	}
	sort.Slice(result, func(i, j int) bool {
		li, lj := result[i].Id, result[j].Id
		if li.From.LocationNumber != lj.From.LocationNumber {
			return li.From.LocationNumber < lj.From.LocationNumber
		}
		if li.From.Number != lj.From.Number {
			return li.From.Number < lj.From.Number
		}
		if li.To.LocationNumber != lj.To.LocationNumber {
			return li.To.LocationNumber < lj.To.LocationNumber
		}
		if li.To.Number != lj.To.Number {
			return li.To.Number < lj.To.Number
		}
		return li.SeqNr < lj.SeqNr
	})
	return result
}

// GetOrCreateLink gets an existing link or creates a new one.
func (n *Network) GetOrCreateLink(linkId LinkId, length float64) *Link {
	if _, ok := n.links[linkId]; !ok {
		from := n.GetOrCreateNode(linkId.From)
		to := n.GetOrCreateNode(linkId.To)
		newLink := NewLink(linkId, from, to, length)
		n.links[linkId] = newLink
	}
	return n.links[linkId]
}

// AddLink adds a link to the network.
func (n *Network) AddLink(link *Link) {
	n.links[link.Id] = link
}

// GetLink retrieves a link by its ID. Panics if not found.
func (n *Network) GetLink(linkId LinkId) *Link {
	link, ok := n.links[linkId]
	if !ok {
		panic(fmt.Sprintf("Link not found: %v", linkId))
	}
	return link
}

// WithLinkStr executes an action on a link identified by its string ID.
func (n *Network) WithLinkStr(id string, action func(*Link)) error {
	linkId, err := ParseLinkId(id)
	if err != nil {
		return err
	}
	n.WithLink(linkId, action)
	return nil
}

// WithLink executes an action on a link identified by its LinkId.
func (n *Network) WithLink(id LinkId, action func(*Link)) {
	action(n.GetLink(id))
}

// Places returns all places in the network, sorted by PLC for deterministic iteration.
func (n *Network) Places() []*Place {
	result := make([]*Place, 0, len(n.places))
	for _, place := range n.places {
		result = append(result, place)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].PLC < result[j].PLC
	})
	return result
}

// AddPlace adds a place to the network.
func (n *Network) AddPlace(place *Place) {
	n.places[place.PLC] = place
}

// Traverse traverses the network using the given visitor starting from its start position.
func Traverse(visitor NetworkVisitor) {
	node := visitor.StartPosition().Link.GetNode(visitor.StartDirection())
	linkPath := visitor.StartPosition().GetLinkPath(visitor.StartDirection())
	traverseRecursive(linkPath, node, visitor)
}

func traverseRecursive(path *LinkPath, node *Node, visitor NetworkVisitor) {
	if path.Link == nil {
		return
	}

	if !visitor.CanVisitLink(path.Link) {
		return
	}

	pathCompleted := visitor.VisitLinkPath(path)

	if !pathCompleted && visitor.CanVisitNode(node) {
		pathCompleted = visitor.VisitNode(node)

		if !pathCompleted && !visitor.HasCompleted() {
			links := node.GetLinks(path.Link)
			for _, link := range links {
				if !visitor.HasCompleted() {
					otherNod := link.GetOtherNode(node)
					traverseRecursive(
						NewLinkPathFromLink(link, link.GetLinkDirection(node)),
						otherNod,
						visitor,
					)
				}
			}
		}

		visitor.LeaveNode(node)
	}

	visitor.LeaveLinkPath(path)
}
