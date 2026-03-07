package domainmodel

import "fmt"

// Lank represents a link in the network.
type Lank struct {
	Id    LankId
	Fran  *Nod
	Till  *Nod
	Langd float64

	Signaler           Entiteter[*Signal]
	Tavlor             Entiteter[*Tavla]
	Platsmittar        Entiteter[*Platsmitt]
	Vagskydd           Entiteter[*Vagskydd]
	Plattformar        Entiteter[*Plattform]
	Sparnummer         Entiteter[*Sparnummer]
	Linjespar          Entiteter[*Linjespar]
	Hastighetsprofiler Entiteter[*Hastighetsprofil]
	Lutningar          Entiteter[*Lutning]
}

// NewLank creates a new Lank.
func NewLank(id LankId, fran, till *Nod, langd float64) *Lank {
	return &Lank{
		Id:    id,
		Fran:  fran,
		Till:  till,
		Langd: langd,
	}
}

// AddHastighetsstracka adds a speed section to the appropriate speed profile,
// creating the profile if it doesn't already exist.
func (l *Lank) AddHastighetsstracka(hs *Hastighetsstracka) {
	var profil *Hastighetsprofil
	for _, hp := range l.Hastighetsprofiler.Items() {
		if hp.Kategori == hs.Kategori {
			profil = hp
			break
		}
	}

	if profil == nil {
		profil = NewHastighetsprofil(hs.Kategori)
		l.Hastighetsprofiler.Add(profil)
	}

	profil.Strackor.Add(hs)
}

// GetNod returns the node at the end of the link in the given direction.
func (l *Lank) GetNod(lankriktning Lankriktning) *Nod {
	switch lankriktning {
	case LankriktningMed:
		return l.Till
	case LankriktningMot:
		return l.Fran
	default:
		panic("Länkriktning okänd")
	}
}

// GetOtherNod returns the other node of the link (not the given one).
func (l *Lank) GetOtherNod(nod *Nod) *Nod {
	if nod == l.Fran {
		return l.Till
	}
	if nod == l.Till {
		return l.Fran
	}
	panic("Lank ansluter inte till Nod!")
}

// GetLankriktning returns the direction of the link relative to the given node.
func (l *Lank) GetLankriktning(nod *Nod) Lankriktning {
	if nod == l.Fran {
		return LankriktningMed
	}
	if nod == l.Till {
		return LankriktningMot
	}
	panic("Lank ansluter inte till Nod!")
}

// String returns a string representation.
func (l *Lank) String() string {
	return fmt.Sprintf("%v->%v %v m", l.Fran, l.Till, l.Langd)
}
