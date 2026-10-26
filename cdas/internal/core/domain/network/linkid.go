package network

import (
	"fmt"
	"strconv"
	"strings"
)

// LinkID uniquely identifies a link.
type LinkID struct {
	From  NodeID
	To    NodeID
	SeqNr int
}

// NewLinkId creates a new LinkId.
func NewLinkId(fran, till NodeID, sekvensnummer int) LinkID {
	return LinkID{From: fran, To: till, SeqNr: sekvensnummer}
}

// GetIntersectingNodId returns the node ID that is shared between two links.
func GetIntersectingNodId(lankId1, lankId2 LinkID) (NodeID, error) {
	if lankId1.From == lankId2.From || lankId1.From == lankId2.To {
		return lankId1.From, nil
	}
	if lankId1.To == lankId2.From || lankId1.To == lankId2.To {
		return lankId1.To, nil
	}
	return NodeID{}, fmt.Errorf("länkarna innehåller ingen gemensam nod")
}

// ParseLinkId parses a LinkId from a string.
// Expected format: "NN<trafikplatsId>:<nummer>:<trafikplatsId>:<nummer>:<sekvensnummer>"
// where NN is a two-character prefix that is stripped from the first part.
func ParseLinkId(id string) (LinkID, error) {
	if id == "" {
		return LinkID{}, fmt.Errorf("länk id inte angivet")
	}

	parts := strings.Split(id, ":")
	if len(parts) != 5 {
		return LinkID{}, fmt.Errorf("länk id felaktigt format: %s", id)
	}

	// The first part has a two-character prefix to remove
	franTrafikplatsStr := parts[0]
	if len(franTrafikplatsStr) > 2 {
		franTrafikplatsStr = franTrafikplatsStr[2:]
	}

	franTrafikplats, err := strconv.Atoi(franTrafikplatsStr)
	if err != nil {
		return LinkID{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	franNummer, err := strconv.Atoi(parts[1])
	if err != nil {
		return LinkID{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	tillTrafikplats, err := strconv.Atoi(parts[2])
	if err != nil {
		return LinkID{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	tillNummer, err := strconv.Atoi(parts[3])
	if err != nil {
		return LinkID{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	sekvensnummer, err := strconv.Atoi(parts[4])
	if err != nil {
		return LinkID{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	fran := NodeID{LocationNumber: franTrafikplats, Number: franNummer}
	till := NodeID{LocationNumber: tillTrafikplats, Number: tillNummer}

	return LinkID{From: fran, To: till, SeqNr: sekvensnummer}, nil
}

// String returns a string representation of the LinkId.
func (l LinkID) String() string {
	return fmt.Sprintf("%v:%v:%d", l.From, l.To, l.SeqNr)
}

func (l LinkID) Key() string {
	return fmt.Sprintf("%d.%d.%d.%d.%d", l.From.LocationNumber, l.To.LocationNumber, l.From.Number, l.To.Number, l.SeqNr)
}
