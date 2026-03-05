package domainmodel

// Balis represents a balise.
type Balis struct {
	Id       string
	Position int
}

// NewBalis creates a new Balis.
func NewBalis(id string, position int) *Balis {
	return &Balis{Id: id, Position: position}
}

// BalisGrupp represents a group of balises.
type BalisGrupp struct {
	Id           string
	Namn         string
	Typ          string
	Lankposition *Lankposition
	Baliser      Entiteter[*Balis]
}

// NewBalisGrupp creates a new BalisGrupp.
func NewBalisGrupp(id, namn, typ string, lankposition *Lankposition) *BalisGrupp {
	return &BalisGrupp{
		Id:           id,
		Namn:         namn,
		Typ:          typ,
		Lankposition: lankposition,
	}
}
