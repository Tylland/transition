package domainmodel

// LinkDirection represents direction along a link (Med = with, Mot = against).
type LinkDirection int

const (
	LinkDirectionWith LinkDirection = iota
	LinkDirectionAgainst
)

func (ld LinkDirection) String() string {
	switch ld {
	case LinkDirectionWith:
		return "With"
	case LinkDirectionAgainst:
		return "Against"
	default:
		return "Okänd"
	}
}
