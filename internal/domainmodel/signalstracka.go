package domainmodel

// Signalstracka represents a signal section between two signals.
type Signalstracka struct {
	Start            *Signal
	Slut             *Signal
	Lankanvandningar []*Lankanvandning
}

// NewSignalstracka creates a new Signalstracka.
func NewSignalstracka(start, slut *Signal, lankanvandningar []*Lankanvandning) *Signalstracka {
	return &Signalstracka{
		Start:            start,
		Slut:             slut,
		Lankanvandningar: lankanvandningar,
	}
}
