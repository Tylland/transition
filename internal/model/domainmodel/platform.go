package domainmodel

// Platform represents a platform with a spread across a link.
type Platform struct {
	*EntityWithExtent
	Name   string
	Length float64
}

// NewPlattform creates a new Plattform.
func NewPlattform(name string, length float64, extent *LinkExtent) *Platform {
	return &Platform{
		EntityWithExtent: NewEntityWithExtent(extent),
		Name:             name,
		Length:           length,
	}
}
