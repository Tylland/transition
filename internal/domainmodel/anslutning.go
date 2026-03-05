package domainmodel

// Anslutning represents a connection between a node and a link.
type Anslutning struct {
	Nod      *Nod
	Lank     *Lank
	Position int
}

// NewAnslutning creates a new Anslutning.
func NewAnslutning(nod *Nod, lank *Lank, position int) *Anslutning {
	return &Anslutning{
		Nod:      nod,
		Lank:     lank,
		Position: position,
	}
}
