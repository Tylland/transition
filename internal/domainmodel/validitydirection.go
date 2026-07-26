package domainmodel

// ValidityDirection represents validity direction.
type ValidityDirection int

const (
	ValidityDirectionNone ValidityDirection = iota
	ValidityDirectionWith
	ValidityDirectionAgainst // Omvänd
	ValidityDirectionBoth    // Båda
)

// IsValid checks if the validity direction applies for the given link direction.
func (g ValidityDirection) IsValid(lr LinkDirection) bool {
	switch lr {
	case LinkDirectionWith:
		return g == ValidityDirectionBoth || g == ValidityDirectionWith
	case LinkDirectionAgainst:
		return g == ValidityDirectionBoth || g == ValidityDirectionAgainst
	default:
		panic("LinkDirection unknown")
	}
}

// ToSign returns a directional sign string.
func (g ValidityDirection) ToSign() string {
	switch g {
	case ValidityDirectionWith:
		return ">"
	case ValidityDirectionAgainst:
		return "<"
	case ValidityDirectionBoth:
		return "<>"
	default:
		return "|"
	}
}
