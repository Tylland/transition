package domainmodel

// Lankriktning represents direction along a link (Med = with, Mot = against).
type Lankriktning int

const (
	LankriktningMed Lankriktning = iota
	LankriktningMot
)

func (lr Lankriktning) String() string {
	switch lr {
	case LankriktningMed:
		return "Med"
	case LankriktningMot:
		return "Mot"
	default:
		return "Okänd"
	}
}

// Gilltighetsriktning represents validity direction.
type Gilltighetsriktning int

const (
	GilltighetsriktningIngen Gilltighetsriktning = iota
	GilltighetsriktningNormal
	GilltighetsriktningOmvand // Omvänd
	GilltighetsriktningBada   // Båda
)

// Gilltig checks if the validity direction applies for the given link direction.
func (g Gilltighetsriktning) Gilltig(lr Lankriktning) bool {
	switch lr {
	case LankriktningMed:
		return g == GilltighetsriktningBada || g == GilltighetsriktningNormal
	case LankriktningMot:
		return g == GilltighetsriktningBada || g == GilltighetsriktningOmvand
	default:
		panic("Länkriktning okänd")
	}
}

// ToSign returns a directional sign string.
func (g Gilltighetsriktning) ToSign() string {
	switch g {
	case GilltighetsriktningNormal:
		return ">"
	case GilltighetsriktningOmvand:
		return "<"
	case GilltighetsriktningBada:
		return "<>"
	default:
		return "|"
	}
}

// Profilriktning represents profile direction.
type Profilriktning int

const (
	ProfilriktningMed Profilriktning = iota
	ProfilriktningMot
)

// NodTyp represents node type.
type NodTyp int

const (
	NodTypOkand NodTyp = iota
	NodTypEnkel
	NodTypKorsning
)
