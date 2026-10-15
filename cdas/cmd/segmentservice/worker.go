package main

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"github.com/tylland/transition/internal/model/facilitydata"
// 	"github.com/tylland/transition/internal/model/sfera/segmentprofile"
// 	"github.com/tylland/transition/internal/model/signalroute"
// 	"github.com/tylland/transition/internal/model/trackroute"
// 	schroutes "github.com/tylland/transition/internal/core/domain"
// )

// // watchInputFolder scans inputDir for FacilityData XML files on each tick and
// // processes any file that is new or has been modified since last seen.
// // It returns when ctx is cancelled.
// func watchInputFolder(ctx context.Context, inputDir string, spWriter segmentprofile.Writer, srWriter signalroute.Writer, trWriter trackroute.Writer, interval time.Duration) {
// 	seen := make(map[string]time.Time)

// 	// Process files already present before the first tick.
// 	scanAndProcess(inputDir, seen, spWriter, srWriter, trWriter)

// 	ticker := time.NewTicker(interval)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case <-ticker.C:
// 			scanAndProcess(inputDir, seen, spWriter, srWriter, trWriter)
// 		}
// 	}
// }

// func scanAndProcess(inputDir string, seen map[string]time.Time, spWriter segmentprofile.Writer, srWriter signalroute.Writer, trWriter trackroute.Writer) {
// 	entries, err := os.ReadDir(inputDir)
// 	if err != nil {
// 		log.Printf("error reading input dir %s: %v", inputDir, err)
// 		return
// 	}

// 	for _, entry := range entries {
// 		if entry.IsDir() || filepath.Ext(entry.Name()) != ".xml" {
// 			continue
// 		}
// 		info, err := entry.Info()
// 		if err != nil {
// 			log.Printf("error reading file info for %s: %v", entry.Name(), err)
// 			continue
// 		}
// 		path := filepath.Join(inputDir, entry.Name())
// 		if mod, ok := seen[path]; ok && !info.ModTime().After(mod) {
// 			continue
// 		}
// 		seen[path] = info.ModTime()
// 		processFile(path, spWriter, srWriter, trWriter)
// 	}
// }

// func processFile(path string, spWriter segmentprofile.Writer, srWriter signalroute.Writer, trWriter trackroute.Writer) {
// 	log.Printf("processing %s", path)

// 	fd, err := facilitydata.ReadXML(path)
// 	if err != nil {
// 		log.Printf("error reading %s: %v", path, err)
// 		return
// 	}

// 	network := facilitydata.NewNetworkBuilder(fd.BaseRoot).Build()

// 	segmentprofile.Generate(network, spWriter)

// 	signalRoutes, signalPaths := signalroute.Generate(network)

// 	if err := srWriter.Write(&schroutes.SignalRoutes{Routes: signalRoutes}); err != nil {
// 		log.Printf("error writing signal routes from %s: %v", path, err)
// 	}

// 	trackRoutes := trackroute.Generate(network, signalRoutes, signalPaths)
// 	if err := trWriter.Write(&schroutes.TrackRoutes{Routes: trackRoutes}); err != nil {
// 		log.Printf("error writing track routes from %s: %v", path, err)
// 	}

// 	log.Printf("segment profiles and routes generated from %s", path)
// }
