package network

import (
	"fmt"
	"strconv"
	"strings"
)

// NodeID uniquely identifies a node.
type NodeID struct {
	LocationNumber int
	Number         int
}

func NewNodeId(locationNumber, number int) NodeID {
	return NodeID{LocationNumber: locationNumber, Number: number}
}

func (n NodeID) String() string {
	return fmt.Sprintf("%d:%d", n.LocationNumber, n.Number)
}

// ParseNodeId parses a NodId from a string in the format "trafikplatsId:nummer".
func ParseNodeId(id string) (NodeID, error) {
	if id == "" {
		return NodeID{}, fmt.Errorf("nod id inte angivet")
	}

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return NodeID{}, fmt.Errorf("node id in invalid format: %s", id)
	}

	locationNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		return NodeID{}, fmt.Errorf("location number in invalid format: %w", err)
	}

	number, err := strconv.Atoi(parts[1])
	if err != nil {
		return NodeID{}, fmt.Errorf("number in invalid format: %w", err)
	}

	return NodeID{LocationNumber: locationNumber, Number: number}, nil
}
