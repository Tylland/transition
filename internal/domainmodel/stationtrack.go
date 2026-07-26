package domainmodel

// StationTrack represents a track number on a station with a spread across a link.
type StationTrack struct {
	*EntityWithExtent
	Number string
	Name   string
}

// NewStationTrack creates a new Sparnummer.
func NewStationTrack(number, name string, extent *LinkExtent) *StationTrack {
	return &StationTrack{
		EntityWithExtent: NewEntityWithExtent(extent),
		Number:           number,
		Name:             name,
	}
}
