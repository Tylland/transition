package domainmodel

import (
	"fmt"
	"strings"
)

// Signal type constants.
const (
	SignalTypInfart      = "INF"
	SignalTypUtfart      = "UBL"
	SignalTypMellan      = "M"
	SignalTypMellanBlock = "MBL"
)

var stopSignals = []string{SignalTypInfart, SignalTypUtfart, SignalTypMellan, SignalTypMellanBlock}

// IsStopSignalTyp checks if the given signal type is a stop signal.
func IsStopSignalTyp(typ string) bool {
	for _, s := range stopSignals {
		if s == typ {
			return true
		}
	}
	return false
}

// SignalId uniquely identifies a signal.
type SignalId struct {
	TrafikplatsSignatur string
	Namn                string
	Typ                 string
}

// NewSignalId creates a new SignalId.
func NewSignalId(trafikplatsSignatur, namn, typ string) *SignalId {
	return &SignalId{
		TrafikplatsSignatur: trafikplatsSignatur,
		Namn:                namn,
		Typ:                 typ,
	}
}

// ParseSignalId parses a SignalId from a string in the format "signatur.namn.typ".
func ParseSignalId(id string) (*SignalId, error) {
	parts := strings.Split(id, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("signal id felaktigt format: %s", id)
	}
	return NewSignalId(parts[0], parts[1], parts[2]), nil
}

// String returns a string representation.
func (s *SignalId) String() string {
	return fmt.Sprintf("%s.%s.%s", s.TrafikplatsSignatur, s.Namn, s.Typ)
}

// Equals checks equality with another SignalId.
func (s *SignalId) Equals(other *SignalId) bool {
	if other == nil {
		return false
	}
	return s.TrafikplatsSignatur == other.TrafikplatsSignatur &&
		s.Namn == other.Namn &&
		s.Typ == other.Typ
}

// Signal represents a signal on the network.
type Signal struct {
	*EntitetMedPosition
	Id  *SignalId
	Typ string
}

// NewSignal creates a new Signal.
func NewSignal(id *SignalId, typ string, lankposition *Lankposition) *Signal {
	return &Signal{
		EntitetMedPosition: NewEntitetMedPosition(lankposition),
		Id:                 id,
		Typ:                typ,
	}
}

// IsStopSignal checks if this signal is a stop signal.
func (s *Signal) IsStopSignal() bool {
	return IsStopSignalTyp(s.Typ)
}

// String returns a string representation.
func (s *Signal) String() string {
	return s.Id.String()
}
