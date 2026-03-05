package domainmodel

// Lutning represents a gradient/slope with a spread across a link.
type Lutning struct {
	*EntitetMedUtbredning
	Grader float64
}

// NewLutning creates a new Lutning.
func NewLutning(grader float64, lankutbredning *Lankutbredning) *Lutning {
	return &Lutning{
		EntitetMedUtbredning: NewEntitetMedUtbredning(lankutbredning),
		Grader:               grader,
	}
}
