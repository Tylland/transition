package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tylland/transition/internal/convertion/facilitydata"
	pdconv "github.com/tylland/transition/internal/convertion/planningdata"
	routeconv "github.com/tylland/transition/internal/convertion/routes"
	sferaconv "github.com/tylland/transition/internal/convertion/sfera"
)

func main() {
	inputPath := filepath.Join("testdata", "FacilityDataExport_SWE_Svv-Bjf_2022_06_30_09_09_41.xml")
	outputPath := filepath.Join("testdata", "sfera")

	// Read FacilityData XML
	facilityData, err := facilitydata.ReadXML(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading facility data: %v\n", err)
		os.Exit(1)
	}

	// Build domain model network from FacilityData
	builder := facilitydata.NewDomainModelBuilder(facilityData.BaseRoot)
	network := builder.Build()

	// Ensure output directory exists
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Write SFERA segment profiles
	writer := sferaconv.NewSegmentProfileFileWriter(outputPath)
	sferaconv.GenerateSegmentProfiles(network, writer)

	fmt.Println("Segment profiles written to", outputPath)

	// Generate and write routes
	routesPath := filepath.Join("testdata", "routes.xml")
	r := routeconv.GenerateRoutes(network)
	if err := routeconv.WriteRoutesFile(r, routesPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing routes: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Routes written to", routesPath)

	// Match planning data to track routes (if planning data file exists)
	planningDataPath := filepath.Join("testdata", "2022-04-19.7st.planningdata.json")
	if _, err := os.Stat(planningDataPath); err == nil {
		response, err := pdconv.ReadPlanningData(planningDataPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading planning data: %v\n", err)
			os.Exit(1)
		}

		fmt.Println()
		for _, result := range response.Results {
			for i := range result.DagligFärdplan {
				matchResult := pdconv.MatchFärdplan(&result.DagligFärdplan[i], r.TrackRoutes)
				pdconv.PrintMatchResult(matchResult)
			}
		}
	}
}
