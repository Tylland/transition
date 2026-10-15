package network

import (
	"fmt"
	"strings"
)

// TavlaTyp constants.
const (
	BoardTypeKilometer = "KmBoard"
)

// BoardId uniquely identifies a tavla (sign/board).
type BoardId struct {
	LocationSignature string
	Information       string
	Type              string
}

// NewBoardId creates a new TavlaId.
func NewBoardId(locationSignature, information, t string) *BoardId {
	return &BoardId{
		LocationSignature: locationSignature,
		Information:       information,
		Type:              t,
	}
}

// ParseBoardId parses a TavlaId from a string in the format "signatur.information.typ".
func ParseBoardId(id string) (*BoardId, error) {
	parts := strings.Split(id, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("tavla id felaktigt format: %s", id)
	}
	return NewBoardId(parts[0], parts[1], parts[2]), nil
}

// String returns a string representation.
func (t *BoardId) String() string {
	return fmt.Sprintf("%s.%s.%s", t.LocationSignature, t.Information, t.Type)
}

// Board represents a sign/board entity with a position on a link.
type Board struct {
	*EntityWithPosition
	Id        *BoardId
	Kilometer *int // Set for Kilometertavla entries, nil otherwise.
}

// NewBoard creates a new Tavla.
func NewBoard(id *BoardId, position *LinkPosition) *Board {
	return &Board{
		EntityWithPosition: NewEntityWithPosition(position),
		Id:                 id,
	}
}

// KilometerBoard is a type alias kept for readability; in Go it is just a *Board with Kilometer set.
type KilometerBoard = Board

// NewKilometerBoard creates a new Board representing a kilometer marker sign.
func NewKilometerBoard(identifier string, kilometer int, position *LinkPosition) *Board {
	t := NewBoard(NewBoardId("-", identifier, BoardTypeKilometer), position)
	t.Kilometer = &kilometer
	return t
}

// IsKilometerBoard reports whether the Tavla is a Kilometertavla.
func (t *Board) IsKilometerBoard() bool {
	return t.Kilometer != nil
}
