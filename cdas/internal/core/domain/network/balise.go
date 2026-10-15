package network

// Balise represents a balise.
type Balise struct {
	Id       string
	Position int
}

// NewBalise creates a new Balis.
func NewBalise(id string, position int) *Balise {
	return &Balise{Id: id, Position: position}
}

// BaliseGroup represents a group of balises.
type BaliseGroup struct {
	Id   string
	Name string
	Type string
	*LinkPosition
	Balises Entities[*Balise]
}

// NewBalisGroup creates a new BalisGrupp.
func NewBalisGroup(id, name, t string, position *LinkPosition) *BaliseGroup {
	return &BaliseGroup{
		Id:           id,
		Name:         name,
		Type:         t,
		LinkPosition: position,
	}
}
