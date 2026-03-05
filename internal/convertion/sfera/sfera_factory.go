package sfera

import (
	"fmt"
	"math"

	dm "github.com/tylland/transition/internal/domainmodel"
	"github.com/tylland/transition/internal/schema/sfera"
)

// GenerateSegmentProfiles generates SFERA segment profiles for both directions
// and writes them using the given writer.
func GenerateSegmentProfiles(nat *dm.Nat, writer SegmentProfileWriter) {
	GenerateSegmentProfilesForDirection(nat, dm.LankriktningMed, writer)
	GenerateSegmentProfilesForDirection(nat, dm.LankriktningMot, writer)
}

// GenerateSegmentProfilesForDirection generates SFERA segment profiles for a
// specific direction and writes them using the given writer.
func GenerateSegmentProfilesForDirection(nat *dm.Nat, lankriktning dm.Lankriktning, writer SegmentProfileWriter) {
	for _, lank := range nat.Lankar() {
		imID := uint16(74)

		segment := &sfera.SegmentProfile{
			SPID:           CreateProfileId(lank.Id, lankriktning),
			SPName:         "namn",
			SPZone:         sfera.SPZone{IMID: &imID},
			SPVersionMajor: 1,
			SPVersionMinor: 0,
			SPLength:       lank.Langd,
			SPStatus:       sfera.SegmentProfileStatusValid,
		}

		// Points
		points := &sfera.SPPoints{}

		// Timing points from platsmittar
		for _, platsmitt := range lank.Platsmittar.Items() {
			var sparnummer *dm.Sparnummer
			for _, sn := range lank.Sparnummer.Items() {
				if sn.Gilltig(lankriktning) {
					sparnummer = sn
					break
				}
			}

			if sparnummer != nil {
				points.TimingPoints = append(points.TimingPoints, CreateTimingPoint(platsmitt, sparnummer, lankriktning))
			}
		}

		// Signals
		for _, signal := range lank.Signaler.Items() {
			if signal.Lankposition.Gilltig(dm.LankriktningMed) {
				signalId := sfera.SignalID{
					SignalIDPhysical: signal.Id.String(),
					SignalIDObject:   signal.Id.String(),
					Location:         signal.Lankposition.GetAvstand(lankriktning),
				}
				points.Signals = append(points.Signals, sfera.SignalComplexType{SignalID: signalId})
			}
		}

		// Balise tavlor
		for _, baliseTavla := range lank.Tavlor.Items() {
			if baliseTavla.Id.Typ == "Bal" && baliseTavla.Lankposition.Gilltig(lankriktning) {
				balise := sfera.Balise{
					Identifier: baliseTavla.Id.String(),
					Location:   baliseTavla.Lankposition.GetAvstand(lankriktning),
				}
				points.Balises = append(points.Balises, balise)
			}
		}

		// Unprotected level crossings
		for _, vagskydd := range lank.Vagskydd.Items() {
			if !vagskydd.HarSkydd {
				points.UnprotectedLevelCrossingStops = append(points.UnprotectedLevelCrossingStops,
					sfera.UnprotectedLevelCrossingStop{
						Identifier: vagskydd.Namn,
						Location:   vagskydd.Lankutbredning.GetStartAvstand(lankriktning),
					})
			}
		}

		segment.SPPoints = points

		// Areas
		areas := &sfera.SPAreas{}
		for _, plattform := range lank.Plattformar.Items() {
			startLoc := plattform.Lankutbredning.GetStartAvstand(lankriktning)
			endLoc := plattform.Lankutbredning.GetSlutAvstand(lankriktning)
			areas.Platforms = append(areas.Platforms, sfera.SPGenericArea{
				Identifier:        plattform.Namn,
				StartLocation:     &startLoc,
				EndLocation:       &endLoc,
				StartEndQualifier: sfera.StartEndQualifierStartsEnds,
			})
		}
		segment.SPAreas = areas

		// Characteristics
		characteristics := &sfera.SPCharacteristics{}

		// Default speed profile (category A)
		var defaultHastighetsprofil *dm.Hastighetsprofil
		for _, hp := range lank.Hastighetsprofiler.Items() {
			if hp.Kategori == "A" {
				defaultHastighetsprofil = hp
				break
			}
		}

		if defaultHastighetsprofil != nil {
			characteristics.StaticSpeedProfiles = append(characteristics.StaticSpeedProfiles,
				CreateStaticSpeedProfile(defaultHastighetsprofil, lankriktning, true))
		}

		for _, hastighetsprofil := range lank.Hastighetsprofiler.Items() {
			characteristics.StaticSpeedProfiles = append(characteristics.StaticSpeedProfiles,
				CreateStaticSpeedProfile(hastighetsprofil, lankriktning, false))
		}

		ga := CreateGradientAverage(lank, lankriktning)
		characteristics.GradientAverage = &ga
		segment.SPCharacteristics = characteristics

		// Context information
		contextInfo := &sfera.SPContextInformation{}

		for _, tavla := range lank.Tavlor.Items() {
			if kmTavla, ok := isKilometertavla(tavla); ok {
				contextInfo.KilometreReferencePoints = append(contextInfo.KilometreReferencePoints,
					sfera.KilometreReferencePoint{
						Identifier: kmTavla.Id.Information,
						Location:   kmTavla.Lankposition.GetAvstand(lankriktning),
						KmRef:      float64(*kmTavla.Kilometer),
					})
			}
		}

		contextInfo.LineAndTrackIdentifier = &sfera.LineAndTrackIdentifier{
			Start: &sfera.LineAndTrackIdentifierData{},
		}
		segment.SPContextInfo = contextInfo

		writer.Write(segment)
	}
}

// isKilometertavla checks if a Tavla is a Kilometertavla and returns it if so.
func isKilometertavla(tavla *dm.Tavla) (*dm.Tavla, bool) {
	return tavla, tavla.IsKilometertavla()
}

// CreateTimingPoint creates a SFERA TimingPoint from domain model entities.
func CreateTimingPoint(platsmitt *dm.Platsmitt, sparnummer *dm.Sparnummer, lankriktning dm.Lankriktning) sfera.TimingPoint {
	tp := sfera.TimingPoint{
		TPID:         fmt.Sprintf("%d:%s", platsmitt.Plats.PLC, sparnummer.Nummer),
		Location:     platsmitt.Lankposition.GetAvstand(lankriktning),
		StationTrack: sparnummer.Nummer,
	}
	tp.TPNames = append(tp.TPNames, sfera.TPName{
		Name:     fmt.Sprintf("%s, %s", sparnummer.Namn, platsmitt.Plats.Namn),
		Language: "se",
	})
	return tp
}

// CreateStaticSpeedProfile creates a SFERA StaticSpeedProfile from domain model entities.
func CreateStaticSpeedProfile(hastighetsprofil *dm.Hastighetsprofil, lankriktning dm.Lankriktning, isDefaultProfile bool) sfera.StaticSpeedProfile {
	ssp := sfera.StaticSpeedProfile{}

	gilltiga := dm.GilltigaUtbredningar(hastighetsprofil.Strackor.Items(), lankriktning)

	for i, hs := range gilltiga {
		speed := float64(hs.Hastighet)
		if i == 0 {
			start := &sfera.SSPComplexType{
				SSPSpeed: &speed,
			}

			if !isDefaultProfile {
				category := sfera.SpecificSSPCategoryOtherReplaces
				ncDiff := sfera.NCTrainTypePassenger
				start.SpecificSSPs = append(start.SpecificSSPs, sfera.SpecificSSP{
					OtherSpecificSSP: &sfera.OtherSpecificSSP{
						TrainCategoryCode: hs.Kategori,
						SSPNCDIFF:         &ncDiff,
					},
					Category: &category,
				})
			}

			ssp.Start = start
		} else {
			change := sfera.StaticSpeedProfileChange{
				SSPComplexType: sfera.SSPComplexType{
					SSPSpeed: &speed,
				},
				Location: hs.Lankutbredning.GetStartAvstand(lankriktning),
			}

			if !isDefaultProfile {
				change.SpecificSSPs = append(change.SpecificSSPs, sfera.SpecificSSP{
					OtherSpecificSSP: &sfera.OtherSpecificSSP{
						TrainCategoryCode: hs.Kategori,
					},
				})
			}

			ssp.Changes = append(ssp.Changes, change)
		}
	}

	return ssp
}

// CreateGradientAverage creates a SFERA GradientAverage from a Lank's gradients.
func CreateGradientAverage(lank *dm.Lank, lankriktning dm.Lankriktning) sfera.GradientAverage {
	ga := sfera.GradientAverage{}

	gilltiga := dm.GilltigaUtbredningar(lank.Lutningar.Items(), lankriktning)

	for i, lutning := range gilltiga {
		if i == 0 {
			ga.Start = &sfera.GradientComplexType{
				GradientValue:     math.Abs(lutning.Grader),
				GradientDirection: GetGradientDirection(lutning, lankriktning),
			}
		} else {
			ga.Changes = append(ga.Changes, sfera.GradientAverageChange{
				GradientComplexType: sfera.GradientComplexType{
					GradientValue:     math.Abs(lutning.Grader),
					GradientDirection: GetGradientDirection(lutning, lankriktning),
				},
				Location: lutning.Lankutbredning.GetStartAvstand(lankriktning),
			})
		}
	}

	return ga
}
