package main

import (
	"encoding/xml"
	"fmt"
	"os"

	dm "github.com/tylland/transition/internal/domainmodel"

	conv "github.com/tylland/transition/internal/convertion/facilitydata"
	fd "github.com/tylland/transition/internal/schema/facilitydata"
)

func main() {
	data, _ := os.ReadFile("testdata/FacilityDataExport_SWE_Svv-Bjf_2022_06_30_09_09_41.xml")
	var facilityData fd.FacilityData
	xml.Unmarshal(data, &facilityData)
	source := conv.NewDomainModelBuilder(facilityData.BaseRoot)
	nat := source.Build()

	// Count totals
	totalSpar := 0
	totalLinje := 0
	lanksWithSpar := 0
	lanksWithLinje := 0
	for _, lank := range nat.Links() {
		sc := len(lank.StationTracks.Items())
		lc := len(lank.LineTracks.Items())
		totalSpar += sc
		totalLinje += lc
		if sc > 0 {
			lanksWithSpar++
		}
		if lc > 0 {
			lanksWithLinje++
		}
	}
	fmt.Printf("Total links: %d\n", len(nat.Links()))
	fmt.Printf("Sparnummer: %d entries on %d links\n", totalSpar, lanksWithSpar)
	fmt.Printf("Linjespar:  %d entries on %d links\n", totalLinje, lanksWithLinje)

	// Check specific stations
	for _, place := range []string{"Bjf", "BJF", "Kv", "Rut", "Kia"} {
		fmt.Printf("\n=== %s ===\n", place)
		for _, lank := range nat.Links() {
			hasSignalAtPlace := false
			for _, sig := range lank.Signals.Items() {
				if sig.Id.TrafikplatsSignatur == place {
					hasSignalAtPlace = true
					break
				}
			}
			if hasSignalAtPlace {
				fmt.Printf("  Link %v (len=%.1f):\n", lank.Id, lank.Length)
				for _, sig := range lank.Signals.Items() {
					if sig.Id.TrafikplatsSignatur == place {
						fmt.Printf("    Signal: %s typ=%s pos=%.4f\n", sig.Id, sig.Type, sig.LinkPosition.Position)
					}
				}
				for _, sn := range lank.StationTracks.Items() {
					fmt.Printf("    Sparnr: %s (%.4f-%.4f)\n", sn.Number, sn.Extent.Start, sn.Extent.End)
				}
				for _, ls := range lank.LineTracks.Items() {
					fmt.Printf("    Linje:  %s (%.4f-%.4f)\n", ls.Number, ls.Extent.Start, ls.Extent.End)
				}
				if len(lank.StationTracks.Items()) == 0 {
					fmt.Printf("    (no sparnummer)\n")
				}
				if len(lank.LineTracks.Items()) == 0 {
					fmt.Printf("    (no linjespar)\n")
				}
			}
		}
	}

	_ = dm.LinkDirectionWith
}
