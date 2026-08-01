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
	PlaceSignature string
	Name           string
	Type           string
}

// NewSignalId creates a new SignalId.
func NewSignalId(placeSignature, name, t string) *SignalId {
	return &SignalId{
		PlaceSignature: placeSignature,
		Name:           name,
		Type:           t,
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
	return fmt.Sprintf("%s.%s.%s", s.PlaceSignature, s.Name, s.Type)
}

// Equals checks equality with another SignalId.
func (s *SignalId) Equals(other *SignalId) bool {
	if other == nil {
		return false
	}
	return s.PlaceSignature == other.PlaceSignature &&
		s.Name == other.Name &&
		s.Type == other.Type
}

// Signal represents a signal on the network.
type Signal struct {
	*EntityWithPosition
	Id   *SignalId
	Type string
}

// NewSignal creates a new Signal.
func NewSignal(id *SignalId, t string, position *LinkPosition) *Signal {
	return &Signal{
		EntityWithPosition: NewEntityWithPosition(position),
		Id:                 id,
		Type:               t,
	}
}

// IsStopSignal checks if this signal is a stop signal.
func (s *Signal) IsStopSignal() bool {
	return IsStopSignalTyp(s.Type)
}

// String returns a string representation.
func (s *Signal) String() string {
	return s.Id.String()
}
