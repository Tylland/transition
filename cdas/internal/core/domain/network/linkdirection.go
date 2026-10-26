package network

// LinkDirection represents direction along a link (Med = with, Mot = against).
type LinkDirection string

const (
	LinkDirectionUnknown LinkDirection = "Unknown"
	LinkDirectionWith    LinkDirection = "With"
	LinkDirectionAgainst LinkDirection = "Against"
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

func (ld LinkDirection) Reverse() LinkDirection {
	switch ld {
	case LinkDirectionWith:
		return LinkDirectionAgainst
	case LinkDirectionAgainst:
		return LinkDirectionWith
	default:
		return LinkDirectionUnknown
	}
}
