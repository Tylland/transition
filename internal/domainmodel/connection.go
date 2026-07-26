package domainmodel

// Connection represents a connection between a node and a link.
type Connection struct {
	Node     *Node
	Link     *Link
	Position int
}

// NewConnection creates a new Anslutning.
func NewConnection(nod *Node, lank *Link, position int) *Connection {
	return &Connection{
		Node:     nod,
		Link:     lank,
		Position: position,
	}
}
