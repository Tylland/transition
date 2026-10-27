package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/tylland/transition/cdas"
	consumer "github.com/tylland/transition/cdas/internal/adapter/consumer/jetstream"
	"github.com/tylland/transition/cdas/internal/adapter/provider"
	publisher "github.com/tylland/transition/cdas/internal/adapter/publisher/jetstream"
	"github.com/tylland/transition/cdas/internal/core/port"
	"github.com/tylland/transition/cdas/internal/core/service"
)

func main() {
	inputDir := filepath.Join("testdata", "input", "facilitydata")
	outputDir := filepath.Join("testdata", "output", "sfera")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("error creating output dir: %v", err)
	}

	var (
		facilityService port.FacilityService
		srWriter        port.SignalRouteWriter
		trWriter        port.TrackRouteWriter
		spWriter        port.SegmentProfileWriter
	)

	nc, err := nats.Connect("nuclear:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, _ := jetstream.New(nc)

	//srWriter = writer.NewSignalRoutesXMLFileWriter(filepath.Join("testdata", "output", "signalroutes.xml"))
	srWriter = cdas.Must(publisher.NewSignalRouteKVWriter(js))

	//	trWriter = writer.NewTrackRoutesXMLFileWriter(filepath.Join("testdata", "output", "sfera", "trackroutes.xml"))

	trWriter = cdas.Must(publisher.NewTrackRouteKVWriter(js))

	spWriter = cdas.Must(publisher.NewSegmentProfileKVWriter(js))

	facilityService = service.NewFacilityService(srWriter, trWriter, spWriter)

	networkProvider := provider.NewFacilityDataFileNetworkProvider(inputDir, 5*time.Second, facilityService)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger := cdas.Must(consumer.NewLogger(js))
	go logger.Start(ctx)

	networkProvider.Start(ctx)
}
