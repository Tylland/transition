package provider

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/tylland/transition/cdas/internal/core/port"
	"github.com/tylland/transition/cdas/internal/model/facilitydata"
)

type FacilityDataFileNetworkProvider struct {
	inputDir  string
	interval  time.Duration
	service   port.FacilityService
	filesSeen map[string]time.Time
}

func NewFacilityDataFileNetworkProvider(inputDir string, interval time.Duration, service port.FacilityService) *FacilityDataFileNetworkProvider {
	return &FacilityDataFileNetworkProvider{inputDir: inputDir, interval: interval, service: service, filesSeen: make(map[string]time.Time)}
}

// scans inputDir for FacilityData XML files on each tick and
// processes any file that is new or has been modified since last seen.
// It returns when ctx is cancelled.
func (p *FacilityDataFileNetworkProvider) Start(ctx context.Context) {
	log.Printf("Start watching %s for FacilityData XML files", p.inputDir)

	// Process files already present before the first tick.
	p.scanAndProcess()

	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.scanAndProcess()
		}
	}
}

func (p *FacilityDataFileNetworkProvider) scanAndProcess() {
	entries, err := os.ReadDir(p.inputDir)
	if err != nil {
		log.Printf("error reading input dir %s: %v", p.inputDir, err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".xml" {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			log.Printf("error reading file info for %s: %v", entry.Name(), err)
			continue
		}
		path := filepath.Join(p.inputDir, entry.Name())
		if mod, ok := p.filesSeen[path]; ok && !info.ModTime().After(mod) {
			continue
		}
		p.filesSeen[path] = info.ModTime()

		p.processFile(path)
	}
}

func (p *FacilityDataFileNetworkProvider) processFile(path string) {
	log.Printf("processing %s", path)

	fd, err := facilitydata.ReadXML(path)
	if err != nil {
		log.Printf("error reading %s: %v", path, err)
		return
	}

	network := facilitydata.NewNetworkBuilder(fd.BaseRoot).Build()

	if p.service != nil {
		p.service.ProcessNetwork(network)
	}
}
