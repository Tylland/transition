package domainmodel

// Sparnummer represents a track number with a spread across a link.
type Sparnummer struct {
	*EntitetMedUtbredning
	Nummer string
	Namn   string
}

// NewSparnummer creates a new Sparnummer.
func NewSparnummer(nummer, namn string, lankutbredning *Lankutbredning) *Sparnummer {
	return &Sparnummer{
		EntitetMedUtbredning: NewEntitetMedUtbredning(lankutbredning),
		Nummer:               nummer,
		Namn:                 namn,
	}
}
