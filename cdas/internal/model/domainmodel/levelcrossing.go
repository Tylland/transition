package domainmodel

// Levelcrossing represents road protection with a spread across a link.
type Levelcrossing struct {
	*EntityWithExtent
	Name          string
	HasProtection bool
}

func NewLevelcrossing(name string, hasProtection bool, lankutbredning *LinkExtent) *Levelcrossing {
	return &Levelcrossing{
		EntityWithExtent: NewEntityWithExtent(lankutbredning),
		Name:             name,
		HasProtection:    hasProtection,
	}
}
