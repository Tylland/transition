package domainmodel

// Vagskydd represents road protection with a spread across a link.
type Vagskydd struct {
	*EntitetMedUtbredning
	Namn     string
	HarSkydd bool
}

// NewVagskydd creates a new Vagskydd.
func NewVagskydd(namn string, harSkydd bool, lankutbredning *Lankutbredning) *Vagskydd {
	return &Vagskydd{
		EntitetMedUtbredning: NewEntitetMedUtbredning(lankutbredning),
		Namn:                 namn,
		HarSkydd:             harSkydd,
	}
}
