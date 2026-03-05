package domainmodel

// Plattform represents a platform with a spread across a link.
type Plattform struct {
	*EntitetMedUtbredning
	Namn  string
	Langd float64
}

// NewPlattform creates a new Plattform.
func NewPlattform(namn string, langd float64, lankutbredning *Lankutbredning) *Plattform {
	return &Plattform{
		EntitetMedUtbredning: NewEntitetMedUtbredning(lankutbredning),
		Namn:                 namn,
		Langd:                langd,
	}
}
