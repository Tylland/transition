package network

// LinkPosition represents a position on a link.
type LinkPosition struct {
	Link      *Link
	Position  float64
	Direction ValidityDirection
}

// NewLinkPosition creates a new Linkposition.
func NewLinkPosition(link *Link, position float64, direction ValidityDirection) *LinkPosition {
	return &LinkPosition{
		Link:      link,
		Position:  position,
		Direction: direction,
	}
}

// IsValid checks if the position is valid for the given link direction.
func (lp *LinkPosition) IsValid(direction LinkDirection) bool {
	return lp.Direction.IsValid(direction)
}

// GetPosition returns the position adjusted for the given direction.
func (lp *LinkPosition) GetPosition(direction LinkDirection) float64 {
	switch direction {
	case LinkDirectionWith:
		return lp.Position
	case LinkDirectionAgainst:
		return 1.0 - lp.Position
	default:
		panic("LinkDirection unknown")
	}
}

// GetDistance returns the distance adjusted for the given direction.
func (lp *LinkPosition) GetDistance(direction LinkDirection) float64 {
	return lp.Link.Length * lp.GetPosition(direction)
}

// GetLinkPath returns a link usage from this position in the given direction.
func (lp *LinkPosition) GetLinkPath(direction LinkDirection) *LinkPath {
	switch direction {
	case LinkDirectionWith:
		return NewLinkPath(lp.Link, lp.Position, 1.0, direction)
	case LinkDirectionAgainst:
		return NewLinkPath(lp.Link, lp.Position, 0.0, direction)
	default:
		panic("LinkDirection unknown")
	}
}
