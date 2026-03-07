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
	source := conv.NewFacilityDataNatSource(facilityData.BaseRoot)
	nat := source.CreateNat()

	// Count totals
	totalSpar := 0
	totalLinje := 0
	lanksWithSpar := 0
	lanksWithLinje := 0
	for _, lank := range nat.Lankar() {
		sc := len(lank.Sparnummer.Items())
		lc := len(lank.Linjespar.Items())
		totalSpar += sc
		totalLinje += lc
		if sc > 0 {
			lanksWithSpar++
		}
		if lc > 0 {
			lanksWithLinje++
		}
	}
	fmt.Printf("Total links: %d\n", len(nat.Lankar()))
	fmt.Printf("Sparnummer: %d entries on %d links\n", totalSpar, lanksWithSpar)
	fmt.Printf("Linjespar:  %d entries on %d links\n", totalLinje, lanksWithLinje)

	// Check specific stations
	for _, place := range []string{"Bjf", "BJF", "Kv", "Rut", "Kia"} {
		fmt.Printf("\n=== %s ===\n", place)
		for _, lank := range nat.Lankar() {
			hasSignalAtPlace := false
			for _, sig := range lank.Signaler.Items() {
				if sig.Id.TrafikplatsSignatur == place {
					hasSignalAtPlace = true
					break
				}
			}
			if hasSignalAtPlace {
				fmt.Printf("  Link %v (len=%.1f):\n", lank.Id, lank.Langd)
				for _, sig := range lank.Signaler.Items() {
					if sig.Id.TrafikplatsSignatur == place {
						fmt.Printf("    Signal: %s typ=%s pos=%.4f\n", sig.Id, sig.Typ, sig.Lankposition.Position)
					}
				}
				for _, sn := range lank.Sparnummer.Items() {
					fmt.Printf("    Sparnr: %s (%.4f-%.4f)\n", sn.Nummer, sn.Lankutbredning.Start, sn.Lankutbredning.Slut)
				}
				for _, ls := range lank.Linjespar.Items() {
					fmt.Printf("    Linje:  %s (%.4f-%.4f)\n", ls.Nummer, ls.Lankutbredning.Start, ls.Lankutbredning.Slut)
				}
				if len(lank.Sparnummer.Items()) == 0 {
					fmt.Printf("    (no sparnummer)\n")
				}
				if len(lank.Linjespar.Items()) == 0 {
					fmt.Printf("    (no linjespar)\n")
				}
			}
		}
	}

	_ = dm.LankriktningMed
}
