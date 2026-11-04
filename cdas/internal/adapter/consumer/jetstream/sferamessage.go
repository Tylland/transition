package jetstream

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/tylland/transition/cdas"
	streams "github.com/tylland/transition/cdas/internal/adapter/sfera"
	"github.com/tylland/transition/cdas/internal/core/port"
)

type SFERAMessageConsumer struct {
	streamName   string
	consumerName string
	nc           *nats.Conn
	js           jetstream.JetStream
	consumer     jetstream.Consumer
	service      port.SFERAService
}

func NewSFERAMessageConsumer(ctx context.Context, nc *nats.Conn, js jetstream.JetStream, streamName, consumerName string, service port.SFERAService) (*SFERAMessageConsumer, error) {

	consumer, err := js.CreateOrUpdateConsumer(ctx, streamName, jetstream.ConsumerConfig{
		Name:          consumerName,
		Durable:       consumerName,
		FilterSubject: cdas.B2GSubjectFilter,
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckWait:       30 * time.Second,
		MaxDeliver:    5, // must be > len(BackOff)
		BackOff:       []time.Duration{time.Second, 5 * time.Second, 15 * time.Second},
		MaxAckPending: 100,
	})

	if err != nil {
		return nil, fmt.Errorf("consumer: %w", err)
	}

	return &SFERAMessageConsumer{nc: nc, js: js, streamName: streamName, consumerName: consumerName, consumer: consumer, service: service}, nil
}

func (c *SFERAMessageConsumer) Start(ctx context.Context) {
	cc, err := c.consumer.Consume(
		func(msg jetstream.Msg) { c.handleMessage(ctx, msg) },
		jetstream.ConsumeErrHandler(func(_ jetstream.ConsumeContext, err error) {
			log.Printf("consume error: %v", err)
		}),
	)
	if err != nil {
		log.Fatalf("consume: %v", err)
	}

	log.Printf("consuming %q messages from stream %q as %q", cdas.B2GSubjectFilter, c.streamName, c.consumerName)

	// Wait for shutdown, then stop cleanly.
	<-ctx.Done()
	log.Println("shutting down...")
	cc.Stop()

	// Drain flushes pending acks before closing the connection.
	if err := c.nc.Drain(); err != nil {
		log.Printf("drain: %v", err)
	}
}

func (c *SFERAMessageConsumer) handleMessage(ctx context.Context, msg jetstream.Msg) {
	meta, _ := msg.Metadata()

	b2gMessage := streams.Message(msg.Subject())

	if err := xml.Unmarshal(msg.Data(), b2gMessage); err != nil {
		// Poison message: retrying won't help, so terminate it.
		log.Printf("bad payload, terminating: %v", err)
		_ = msg.Term()
		return
	}

	if err := c.service.HandleB2GMessage(ctx, b2gMessage); err != nil {
		log.Printf("message %s failed (delivery %d): %v", msg.Subject(), meta.NumDelivered, err)
		// Redelivery delay follows the consumer's BackOff schedule when using Nak().
		_ = msg.Nak()
		return
	}

	if err := msg.Ack(); err != nil {
		log.Printf("ack failed: %v", err)
	}
}
