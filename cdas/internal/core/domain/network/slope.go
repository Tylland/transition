package network

// Slope represents a gradient/slope with a spread across a link.
type Slope struct {
	*EntityWithExtent
	Gradient float64
}

// NewLutning creates a new Lutning.
func NewLutning(gradient float64, extent *LinkExtent) *Slope {
	return &Slope{
		EntityWithExtent: NewEntityWithExtent(extent),
		Gradient:         gradient,
	}
}
