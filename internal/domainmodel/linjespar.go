package domainmodel

// Linjespar represents a line track (linjespår) with a spread across a link.
// Line tracks identify which track on the line a link belongs to (e.g. "E" for single,
// "U" for up-track, "N" for down-track).
type Linjespar struct {
	*EntitetMedUtbredning
	Nummer string
	Namn   string
}

// NewLinjespar creates a new Linjespar.
func NewLinjespar(nummer, namn string, lankutbredning *Lankutbredning) *Linjespar {
	return &Linjespar{
		EntitetMedUtbredning: NewEntitetMedUtbredning(lankutbredning),
		Nummer:               nummer,
		Namn:                 namn,
	}
}
