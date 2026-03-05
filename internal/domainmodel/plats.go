package domainmodel

// Plats represents a place/location.
type Plats struct {
	Signatur string
	Namn     string
	Typ      string
	PLC      int
}

// NewPlats creates a new Plats.
func NewPlats(signatur, namn, typ string, plc int) *Plats {
	return &Plats{
		Signatur: signatur,
		Namn:     namn,
		Typ:      typ,
		PLC:      plc,
	}
}

// Platsmitt represents the center of a place positioned on a link.
type Platsmitt struct {
	*EntitetMedPosition
	Plats *Plats
}

// NewPlatsmitt creates a new Platsmitt.
func NewPlatsmitt(plats *Plats, lankposition *Lankposition) *Platsmitt {
	return &Platsmitt{
		EntitetMedPosition: NewEntitetMedPosition(lankposition),
		Plats:              plats,
	}
}
