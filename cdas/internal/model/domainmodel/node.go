package domainmodel

// Node represents a node in the network.
type Node struct {
	Id     NodeId
	Name   string
	routes []*NodeRoute
}

// NewNode creates a new Nod.
func NewNode(id NodeId, name string) *Node {
	return &Node{
		Id:     id,
		Name:   name,
		routes: make([]*NodeRoute, 0),
	}
}

// AddRoute adds a node path connecting two links.
func (n *Node) AddRoute(lankA, lankB *Link) {
	n.routes = append(n.routes, NewNodeRoute(n, lankA, lankB))
}

// Routes returns the node paths.
func (n *Node) Routes() []*NodeRoute {
	return n.routes
}

// GetLinks returns the links reachable from the given link through this node.
func (n *Node) GetLinks(fran *Link) []*Link {
	var result []*Link
	for _, nv := range n.routes {
		if nv.LankA == fran {
			result = append(result, nv.LankB)
		}
	}
	return result
}

// String returns a string representation of the Nod.
func (n *Node) String() string {
	return n.Id.String()
}

// NodeRoute represents a path through a node connecting two links.
type NodeRoute struct {
	Nod   *Node
	LankA *Link
	LankB *Link
}

func NewNodeRoute(node *Node, linkA, linkB *Link) *NodeRoute {
	return &NodeRoute{Nod: node, LankA: linkA, LankB: linkB}
}

type NodeType int

const (
	NodTypOkand NodeType = iota
	NodTypEnkel
	NodTypKorsning
)
