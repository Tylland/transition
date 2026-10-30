package jetstream

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

const (
	SFERASubjectB2GRequest = "sfera.b2g.request"
	SFERASubjectB2GEvent   = "sfera.b2g.event"
	SFERASubjectB2GReply   = "sfera.b2g.reply"
	SFERASubjectG2BRequest = "sfera.g2b.request"
	SFERASubjectG2BEvent   = "sfera.g2b.event"
	SFERASubjectG2BReply   = "sfera.g2b.reply"
	SFERASubjectInvalid    = "sfera.invalid"
)

var B2GSubjects []string = []string{SFERASubjectB2GRequest, SFERASubjectB2GEvent, SFERASubjectB2GReply}
var G2BSubjects []string = []string{SFERASubjectG2BRequest, SFERASubjectG2BEvent, SFERASubjectG2BReply}

type SFERAMessagePublisher struct {
	js  jetstream.JetStream
	b2g jetstream.Stream
	g2b jetstream.Stream
}

func stream(ctx context.Context, js jetstream.JetStream, name string, subjects []string) (jetstream.Stream, error) {
	stream, err := js.Stream(ctx, name)
	if err != nil {
		stream, err = js.CreateStream(ctx, jetstream.StreamConfig{
			Name:     name,
			Subjects: subjects,
		})
		if err != nil {
			return nil, err
		}
	}

	return stream, nil
}

func NewSFERAMessagePublisher(js jetstream.JetStream) (*SFERAMessagePublisher, error) {
	ctx := context.Background()

	b2g, _ := stream(ctx, js, "B2G", B2GSubjects)
	g2b, _ := stream(ctx, js, "G2B", G2BSubjects)

	return &SFERAMessagePublisher{js: js, b2g: b2g, g2b: g2b}, nil
}

func (p *SFERAMessagePublisher) subject(msg any) string {
	switch msg.(type) {
	case sfera.SFERAB2GRequestMessage:
		return SFERASubjectB2GRequest
	case sfera.SFERAB2GEventMessage:
		return SFERASubjectB2GRequest
	case sfera.SFERAB2GReplyMessage:
		return SFERASubjectB2GRequest
	case sfera.SFERAG2BRequestMessage:
		return SFERASubjectG2BRequest
	case sfera.SFERAG2BEventMessage:
		return SFERASubjectG2BEvent
	case sfera.SFERAG2BReplyMessage:
		return SFERASubjectG2BReply
	default:
		return SFERASubjectInvalid
	}
}

func (p *SFERAMessagePublisher) Publish(ctx context.Context, msg sfera.SFERAB2GRequestMessage) error {
	subject := p.subject(msg)

	payload, err := xml.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message %s: %w", subject, err)
	}

	if _, err := p.js.Publish(ctx, subject, payload); err != nil {
		return fmt.Errorf("publish to subject %q: %w", subject, err)
	}

	return nil
}
