package main

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/tylland/transition/cdas"
	consumer "github.com/tylland/transition/cdas/internal/adapter/consumer/jetstream"
	"github.com/tylland/transition/cdas/internal/adapter/provider"
	publisher "github.com/tylland/transition/cdas/internal/adapter/publisher/jetstream"
	reader "github.com/tylland/transition/cdas/internal/adapter/reader/jetstream"
	"github.com/tylland/transition/cdas/internal/core/service"
)

func main() {
	nc, err := nats.Connect("nuclear:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, _ := jetstream.New(nc)

	ctx := context.Background()

	settings := &service.Settings{}
	msgPublisher := cdas.Must(publisher.NewSFERAMessagePublisher(js))
	spKey := &provider.SPKeyProvider{}
	jpKey := &provider.JPKeyProvider{}

	spReader := &reader.SegmentProfileReader{KeyValueReader: cdas.Must(reader.NewKeyValueReader(js, cdas.SegmentProfileBucket))}
	jpReader := &reader.JourneyProfileReader{KeyValueReader: cdas.Must(reader.NewKeyValueReader(js, cdas.JourneyProfileBucket))}

	sferaService := service.NewSFERAService(settings, msgPublisher, spKey, spReader, jpKey, jpReader)

	msgConsumer := cdas.Must(consumer.NewSFERAMessageConsumer(ctx, nc, js, cdas.B2GStreamName, cdas.B2GDefaultConsumerName, sferaService))
	msgConsumer.Start(ctx)
}
