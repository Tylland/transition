package nats

import (
	"encoding/xml"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tylland/transition/cdas/internal/core/domain"
)

// Writer publishes SFERA segment profiles as XML to NATS.
type SignalRouteNATSWriter struct {
	conn *nats.Conn
}

// NewWriter creates a NatsWriter that publishes segment profiles to NATS.
func NewSignalRouteNATSWriter(conn *nats.Conn) *SignalRouteNATSWriter {
	return &SignalRouteNATSWriter{conn: conn}
}

// subjectFromSignalRoute builds the NATS subject from the signal route.
// Format: signalroute.{from_trafikplats}.{to_trafikplats}.{from_signal}.{to_signal}
func subjectFromSignalRoute(route *domain.SignalRoute) (string, error) {
	return fmt.Sprintf("signalroute.%s.%s.%s.%s",
		route.StartPlace, route.EndPlace, route.StartSignal, route.EndSignal), nil
}

// Write marshals the signal route to XML and publishes it to its derived NATS subject.
// It returns the subject on success.
func (w *SignalRouteNATSWriter) Write(route *domain.SignalRoute) (string, error) {
	subject, err := subjectFromSignalRoute(route)
	if err != nil {
		return "", err
	}

	data, err := xml.Marshal(route)
	if err != nil {
		return subject, fmt.Errorf("failed to marshal signal route %s: %w", subject, err)
	}

	msg := &nats.Msg{
		Subject: subject,
		Data:    data,
	}

	if err := w.conn.PublishMsg(msg); err != nil {
		return subject, fmt.Errorf("failed to publish signal route %s: %w", subject, err)
	}

	return subject, nil
}
