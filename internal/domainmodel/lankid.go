package domainmodel

import (
	"fmt"
	"strconv"
	"strings"
)

// LankId uniquely identifies a link.
type LankId struct {
	Fran          NodId
	Till          NodId
	Sekvensnummer int
}

// NewLankId creates a new LankId.
func NewLankId(fran, till NodId, sekvensnummer int) LankId {
	return LankId{Fran: fran, Till: till, Sekvensnummer: sekvensnummer}
}

// GetIntersectingNodId returns the node ID that is shared between two links.
func GetIntersectingNodId(lankId1, lankId2 LankId) (NodId, error) {
	if lankId1.Fran == lankId2.Fran || lankId1.Fran == lankId2.Till {
		return lankId1.Fran, nil
	}
	if lankId1.Till == lankId2.Fran || lankId1.Till == lankId2.Till {
		return lankId1.Till, nil
	}
	return NodId{}, fmt.Errorf("länkarna innehåller ingen gemensam nod")
}

// ParseLankId parses a LankId from a string.
// Expected format: "NN<trafikplatsId>:<nummer>:<trafikplatsId>:<nummer>:<sekvensnummer>"
// where NN is a two-character prefix that is stripped from the first part.
func ParseLankId(id string) (LankId, error) {
	if id == "" {
		return LankId{}, fmt.Errorf("länk id inte angivet")
	}

	parts := strings.Split(id, ":")
	if len(parts) != 5 {
		return LankId{}, fmt.Errorf("länk id felaktigt format: %s", id)
	}

	// The first part has a two-character prefix to remove
	franTrafikplatsStr := parts[0]
	if len(franTrafikplatsStr) > 2 {
		franTrafikplatsStr = franTrafikplatsStr[2:]
	}

	franTrafikplats, err := strconv.Atoi(franTrafikplatsStr)
	if err != nil {
		return LankId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	franNummer, err := strconv.Atoi(parts[1])
	if err != nil {
		return LankId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	tillTrafikplats, err := strconv.Atoi(parts[2])
	if err != nil {
		return LankId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	tillNummer, err := strconv.Atoi(parts[3])
	if err != nil {
		return LankId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	sekvensnummer, err := strconv.Atoi(parts[4])
	if err != nil {
		return LankId{}, fmt.Errorf("länk id felaktigt format: %w", err)
	}

	fran := NodId{TrafikplatsId: franTrafikplats, Nummer: franNummer}
	till := NodId{TrafikplatsId: tillTrafikplats, Nummer: tillNummer}

	return LankId{Fran: fran, Till: till, Sekvensnummer: sekvensnummer}, nil
}

// String returns a string representation of the LankId.
func (l LankId) String() string {
	return fmt.Sprintf("%v:%v:%d", l.Fran, l.Till, l.Sekvensnummer)
}
