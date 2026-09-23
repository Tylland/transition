package infrastructure

import (
	"encoding/xml"
	"fmt"

	"github.com/nats-io/nats.go"
	sfera "github.com/tylland/transition/internal/schema/sfera/v3"
)

// NatsReader subscribes to a NATS subject and delivers decoded segment profiles.
type NatsReader struct {
	conn    *nats.Conn
	subject string
}

// NewNatsReader creates a NatsReader that subscribes to the given subject.
func NewNatsReader(conn *nats.Conn, subject string) *NatsReader {
	return &NatsReader{conn: conn, subject: subject}
}

// Subscribe registers a handler that is called for each segment profile received
// on the subject. The returned subscription can be used to unsubscribe.
func (r *NatsReader) Subscribe(handler func(*sfera.SegmentProfile) error) (*nats.Subscription, error) {
	sub, err := r.conn.Subscribe(r.subject, func(msg *nats.Msg) {
		var segment sfera.SegmentProfile
		if err := xml.Unmarshal(msg.Data, &segment); err != nil {
			fmt.Printf("natsreader: failed to unmarshal segment profile from %s: %v\n", r.subject, err)
			return
		}
		if err := handler(&segment); err != nil {
			fmt.Printf("natsreader: handler error for segment profile on %s: %v\n", r.subject, err)
		}
	})
	if err != nil {
		return nil, fmt.Errorf("natsreader: failed to subscribe to %s: %w", r.subject, err)
	}
	return sub, nil
}
