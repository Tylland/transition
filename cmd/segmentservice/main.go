package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/tylland/transition/internal/model/sfera/segmentprofile"
	"github.com/tylland/transition/internal/model/signalroute"
	"github.com/tylland/transition/internal/model/trackroute"
)

func main() {
	inputDir := filepath.Join("testdata", "input")
	outputDir := filepath.Join("testdata", "sfera")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("error creating output dir: %v", err)
	}

	spWriter := segmentprofile.NewFileWriter(outputDir)
	srWriter := signalroute.NewFileWriter(filepath.Join("testdata", "signalroutes.xml"))
	trWriter := trackroute.NewFileWriter(filepath.Join("testdata", "trackroutes.xml"))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("watching %s for FacilityData XML files", inputDir)
	watchInputFolder(ctx, inputDir, spWriter, srWriter, trWriter, 5*time.Second)
}
