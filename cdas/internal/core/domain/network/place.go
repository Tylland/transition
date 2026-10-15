package network

// Place represents a place/location.
type Place struct {
	Signature string
	Name      string
	Type      string
	PLC       int
}

// NewPlace creates a new Plats.
func NewPlace(signature, name, t string, plc int) *Place {
	return &Place{
		Signature: signature,
		Name:      name,
		Type:      t,
		PLC:       plc,
	}
}

// PlaceCenter represents the center of a place positioned on a link.
type PlaceCenter struct {
	*EntityWithPosition
	*Place
}

// NewPlaceCenter creates a new Platsmitt.
func NewPlaceCenter(place *Place, position *LinkPosition) *PlaceCenter {
	return &PlaceCenter{
		EntityWithPosition: NewEntityWithPosition(position),
		Place:              place,
	}
}
