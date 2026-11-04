package jetstream

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/tylland/transition/cdas"
	streams "github.com/tylland/transition/cdas/internal/adapter/sfera"
)

type SFERAMessagePublisher struct {
	js  jetstream.JetStream
	b2g jetstream.Stream
	g2b jetstream.Stream
}

func NewSFERAMessagePublisher(js jetstream.JetStream) (*SFERAMessagePublisher, error) {
	ctx := context.Background()

	b2g, _ := streams.Create(ctx, js, cdas.B2GStreamName, cdas.B2GSubjects)
	g2b, _ := streams.Create(ctx, js, cdas.G2BStreamName, cdas.G2BSubjects)

	return &SFERAMessagePublisher{js: js, b2g: b2g, g2b: g2b}, nil
}

func (p *SFERAMessagePublisher) Publish(ctx context.Context, msg any) error {
	subject := streams.Subject(msg)

	payload, err := xml.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message %s: %w", subject, err)
	}

	if _, err := p.js.Publish(ctx, subject, payload); err != nil {
		return fmt.Errorf("publish to subject %q: %w", subject, err)
	}

	return nil
}
