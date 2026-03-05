package domainmodel

import (
	"fmt"
	"strconv"
	"strings"
)

// NodId uniquely identifies a node.
type NodId struct {
	TrafikplatsId int
	Nummer        int
}

// NewNodId creates a new NodId.
func NewNodId(trafikplatsId, nummer int) NodId {
	return NodId{TrafikplatsId: trafikplatsId, Nummer: nummer}
}

// String returns a string representation of the NodId.
func (n NodId) String() string {
	return fmt.Sprintf("%d:%d", n.TrafikplatsId, n.Nummer)
}

// ParseNodId parses a NodId from a string in the format "trafikplatsId:nummer".
func ParseNodId(id string) (NodId, error) {
	if id == "" {
		return NodId{}, fmt.Errorf("nod id inte angivet")
	}

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return NodId{}, fmt.Errorf("nod id felaktigt format: %s", id)
	}

	trafikplatsId, err := strconv.Atoi(parts[0])
	if err != nil {
		return NodId{}, fmt.Errorf("nod id felaktigt format: %w", err)
	}

	nummer, err := strconv.Atoi(parts[1])
	if err != nil {
		return NodId{}, fmt.Errorf("nod id felaktigt format: %w", err)
	}

	return NodId{TrafikplatsId: trafikplatsId, Nummer: nummer}, nil
}
