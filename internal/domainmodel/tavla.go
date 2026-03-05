package domainmodel

import (
	"fmt"
	"strings"
)

// TavlaTyp constants.
const (
	TavlaTypKilometertavla = "Kmtavla"
)

// TavlaId uniquely identifies a tavla (sign/board).
type TavlaId struct {
	TrafikplatsSignatur string
	Information         string
	Typ                 string
}

// NewTavlaId creates a new TavlaId.
func NewTavlaId(trafikplatsSignatur, information, typ string) *TavlaId {
	return &TavlaId{
		TrafikplatsSignatur: trafikplatsSignatur,
		Information:         information,
		Typ:                 typ,
	}
}

// ParseTavlaId parses a TavlaId from a string in the format "signatur.information.typ".
func ParseTavlaId(id string) (*TavlaId, error) {
	parts := strings.Split(id, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("tavla id felaktigt format: %s", id)
	}
	return NewTavlaId(parts[0], parts[1], parts[2]), nil
}

// String returns a string representation.
func (t *TavlaId) String() string {
	return fmt.Sprintf("%s.%s.%s", t.TrafikplatsSignatur, t.Information, t.Typ)
}

// Tavla represents a sign/board entity with a position on a link.
type Tavla struct {
	*EntitetMedPosition
	Id        *TavlaId
	Kilometer *int // Set for Kilometertavla entries, nil otherwise.
}

// NewTavla creates a new Tavla.
func NewTavla(id *TavlaId, lankposition *Lankposition) *Tavla {
	return &Tavla{
		EntitetMedPosition: NewEntitetMedPosition(lankposition),
		Id:                 id,
	}
}

// Kilometertavla is a type alias kept for readability; in Go it is just a *Tavla with Kilometer set.
type Kilometertavla = Tavla

// NewKilometertavla creates a new Tavla representing a kilometer marker sign.
func NewKilometertavla(identifier string, kilometer int, lankposition *Lankposition) *Tavla {
	t := NewTavla(NewTavlaId("-", identifier, TavlaTypKilometertavla), lankposition)
	t.Kilometer = &kilometer
	return t
}

// IsKilometertavla reports whether the Tavla is a Kilometertavla.
func (t *Tavla) IsKilometertavla() bool {
	return t.Kilometer != nil
}
