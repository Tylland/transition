package main

import (
	"fmt"
	"os"
	"path/filepath"

	fdconv "github.com/tylland/transition/internal/convertion/facilitydata"
	routeconv "github.com/tylland/transition/internal/convertion/routes"
	sferaconv "github.com/tylland/transition/internal/convertion/sfera"
)

func main() {
	inputPath := filepath.Join("testdata", "FacilityDataExport_SWE_Svv-Bjf_2022_06_30_09_09_41.xml")
	outputPath := filepath.Join("testdata", "sfera")

	// Read FacilityData XML
	facilityData, err := fdconv.ReadFacilityData(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading facility data: %v\n", err)
		os.Exit(1)
	}

	// Build domain model network from FacilityData
	natSource := fdconv.NewFacilityDataNatSource(facilityData.BaseRoot)
	nat := natSource.CreateNat()

	// Ensure output directory exists
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Write SFERA segment profiles
	writer := sferaconv.NewSegmentProfileFileWriter(outputPath)
	sferaconv.GenerateSegmentProfiles(nat, writer)

	fmt.Println("Segment profiles written to", outputPath)

	// Generate and write routes
	routesPath := filepath.Join("testdata", "routes.xml")
	r := routeconv.GenerateRoutes(nat)
	if err := routeconv.WriteRoutesFile(r, routesPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing routes: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Routes written to", routesPath)
}
