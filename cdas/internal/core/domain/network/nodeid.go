package network

import (
	"fmt"
	"strconv"
	"strings"
)

// NodeId uniquely identifies a node.
type NodeId struct {
	LocationNumber int
	Number         int
}

func NewNodeId(locationNumber, number int) NodeId {
	return NodeId{LocationNumber: locationNumber, Number: number}
}

func (n NodeId) String() string {
	return fmt.Sprintf("%d:%d", n.LocationNumber, n.Number)
}

// ParseNodeId parses a NodId from a string in the format "trafikplatsId:nummer".
func ParseNodeId(id string) (NodeId, error) {
	if id == "" {
		return NodeId{}, fmt.Errorf("nod id inte angivet")
	}

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return NodeId{}, fmt.Errorf("node id in invalid format: %s", id)
	}

	locationNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		return NodeId{}, fmt.Errorf("location number in invalid format: %w", err)
	}

	number, err := strconv.Atoi(parts[1])
	if err != nil {
		return NodeId{}, fmt.Errorf("number in invalid format: %w", err)
	}

	return NodeId{LocationNumber: locationNumber, Number: number}, nil
}
