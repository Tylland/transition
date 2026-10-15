package network

import (
	"fmt"
	"strconv"
	"strings"
)

// LinkId uniquely identifies a link.
type LinkId struct {
	From  NodeId
	To    NodeId
	SeqNr int
}

// NewLinkId creates a new LinkId.
func NewLinkId(fran, till NodeId, sekvensnummer int) LinkId {
	return LinkId{From: fran, To: till, SeqNr: sekvensnummer}
}

// GetIntersectingNodId returns the node ID that is shared between two links.
func GetIntersectingNodId(lankId1, lankId2 LinkId) (NodeId, error) {
	if lankId1.From == lankId2.From || lankId1.From == lankId2.To {
		return lankId1.From, nil
	}
	if lankId1.To == lankId2.From || lankId1.To == lankId2.To {
		return lankId1.To, nil
	}
	return NodeId{}, fmt.Errorf("länkarna innehåller ingen gemensam nod")
}

// ParseLinkId parses a LinkId from a string.
// Expected format: "NN<trafikplatsId>:<nummer>:<trafikplatsId>:<nummer>:<sekvensnummer>"
// where NN is a two-character prefix that is stripped from the first part.
func ParseLinkId(id string) (LinkId, error) {
	if id == "" {
		return LinkId{}, fmt.Errorf("länk id inte angivet")
	}

	parts := strings.Split(id, ":")
	if len(parts) != 5 {
		return LinkId{}, fmt.Errorf("länk id felaktigt format: %s", id)
	}

	// The first part has a two-character prefix to remove
	franTrafikplatsStr := parts[0]
	if len(franTrafikplatsStr) > 2 {
		franTrafikplatsStr = franTrafikplatsStr[2:]
	}

	franTrafikplats, err := strconv.Atoi(franTrafikplatsStr)
	if err != nil {
		return LinkId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	franNummer, err := strconv.Atoi(parts[1])
	if err != nil {
		return LinkId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	tillTrafikplats, err := strconv.Atoi(parts[2])
	if err != nil {
		return LinkId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	tillNummer, err := strconv.Atoi(parts[3])
	if err != nil {
		return LinkId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	sekvensnummer, err := strconv.Atoi(parts[4])
	if err != nil {
		return LinkId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	fran := NodeId{LocationNumber: franTrafikplats, Number: franNummer}
	till := NodeId{LocationNumber: tillTrafikplats, Number: tillNummer}

	return LinkId{From: fran, To: till, SeqNr: sekvensnummer}, nil
}

// String returns a string representation of the LinkId.
func (l LinkId) String() string {
	return fmt.Sprintf("%v:%v:%d", l.From, l.To, l.SeqNr)
}
