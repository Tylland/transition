package streams

import (
	"context"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/tylland/transition/cdas"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

func Create(ctx context.Context, js jetstream.JetStream, name string, subjects []string) (jetstream.Stream, error) {
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

func Subject(msg any) string {
	switch msg.(type) {
	case sfera.SFERAB2GRequestMessage:
		return cdas.SFERASubjectB2GRequest
	case sfera.SFERAB2GEventMessage:
		return cdas.SFERASubjectB2GRequest
	case sfera.SFERAB2GReplyMessage:
		return cdas.SFERASubjectB2GRequest
	case sfera.SFERAG2BRequestMessage:
		return cdas.SFERASubjectG2BRequest
	case sfera.SFERAG2BEventMessage:
		return cdas.SFERASubjectG2BEvent
	case sfera.SFERAG2BReplyMessage:
		return cdas.SFERASubjectG2BReply
	default:
		return cdas.SFERASubjectInvalid
	}
}

func Message(subject string) any {
	switch subject {
	case cdas.SFERASubjectB2GRequest:
		return &sfera.SFERAB2GRequestMessage{}
	// case sfera.SFERAB2GEventMessage:
	// 	return SFERASubjectB2GRequest
	// case sfera.SFERAB2GReplyMessage:
	// 	return SFERASubjectB2GRequest
	// case sfera.SFERAG2BRequestMessage:
	// 	return SFERASubjectG2BRequest
	// case sfera.SFERAG2BEventMessage:
	// 	return SFERASubjectG2BEvent
	// case sfera.SFERAG2BReplyMessage:
	// 	return SFERASubjectG2BReply
	default:
		return cdas.SFERASubjectInvalid
	}
}
