package segmentprofile

import (
	"fmt"
	"math"

	dm "github.com/tylland/transition/internal/model/domainmodel"
	"github.com/tylland/transition/internal/schema/sfera"
)

// Generate generates SFERA segment profiles for both directions
// and writes them using the given writer.
func Generate(network *dm.Network, writer Writer) {
	GenerateForDirection(network, dm.LinkDirectionWith, writer)
	GenerateForDirection(network, dm.LinkDirectionAgainst, writer)
}

// GenerateForDirection generates SFERA segment profiles for a
// specific direction and writes them using the given writer.
func GenerateForDirection(network *dm.Network, direction dm.LinkDirection, writer Writer) {
	for _, link := range network.Links() {
		imID := uint16(74)

		segment := &sfera.SegmentProfile{
			SPID:           createProfileId(link.Id, direction),
			SPName:         "namn",
			SPZone:         sfera.SPZone{IMID: &imID},
			SPVersionMajor: 1,
			SPVersionMinor: 0,
			SPLength:       link.Length,
			SPStatus:       sfera.SegmentProfileStatusValid,
		}

		// Points
		points := &sfera.SPPoints{}

		// Timing points from place centers
		for _, placeCenter := range link.PlaceCenters.Items() {
			var stationTrack *dm.StationTrack
			for _, sn := range link.StationTracks.Items() {
				if sn.IsValid(direction) {
					stationTrack = sn
					break
				}
			}

			if stationTrack != nil {
				points.TimingPoints = append(points.TimingPoints, CreateTimingPoint(placeCenter, stationTrack, direction))
			}
		}

		// Signals
		for _, signal := range link.Signals.Items() {
			if signal.LinkPosition.IsValid(direction) {
				signalId := sfera.SignalID{
					SignalIDPhysical: signal.Id.String(),
					SignalIDObject:   signal.Id.String(),
					Location:         signal.LinkPosition.GetDistance(direction),
				}
				points.Signals = append(points.Signals, sfera.SignalComplexType{SignalID: signalId})
			}
		}

		// Balise tavlor
		for _, baliseBoard := range link.Boards.Items() {
			if baliseBoard.Id.Type == "Bal" && baliseBoard.LinkPosition.IsValid(direction) {
				balise := sfera.Balise{
					Identifier: baliseBoard.Id.String(),
					Location:   baliseBoard.LinkPosition.GetDistance(direction),
				}
				points.Balises = append(points.Balises, balise)
			}
		}

		// Unprotected level crossings
		for _, vagskydd := range link.Levelcrossings.Items() {
			if !vagskydd.HasProtection {
				points.UnprotectedLevelCrossingStops = append(points.UnprotectedLevelCrossingStops,
					sfera.UnprotectedLevelCrossingStop{
						Identifier: vagskydd.Name,
						Location:   vagskydd.Extent.GetStartDistance(direction),
					})
			}
		}

		segment.SPPoints = points

		// Areas
		areas := &sfera.SPAreas{}
		for _, platform := range link.Platforms.Items() {
			startLoc := platform.Extent.GetStartDistance(direction)
			endLoc := platform.Extent.GetEndDistance(direction)
			areas.Platforms = append(areas.Platforms, sfera.SPGenericArea{
				Identifier:        platform.Name,
				StartLocation:     &startLoc,
				EndLocation:       &endLoc,
				StartEndQualifier: sfera.StartEndQualifierStartsEnds,
			})
		}
		segment.SPAreas = areas

		// Characteristics
		characteristics := &sfera.SPCharacteristics{}

		// Default speed profile (category A)
		var defaultSpeedProfile *dm.SpeedProfile
		for _, speedProfile := range link.SpeedProfiles.Items() {
			if speedProfile.Category == "A" {
				defaultSpeedProfile = speedProfile
				break
			}
		}

		if defaultSpeedProfile != nil {
			characteristics.StaticSpeedProfiles = append(characteristics.StaticSpeedProfiles,
				CreateStaticSpeedProfile(defaultSpeedProfile, direction, true))
		}

		for _, hastighetsprofil := range link.SpeedProfiles.Items() {
			characteristics.StaticSpeedProfiles = append(characteristics.StaticSpeedProfiles,
				CreateStaticSpeedProfile(hastighetsprofil, direction, false))
		}

		ga := CreateGradientAverage(link, direction)
		characteristics.GradientAverage = &ga
		segment.SPCharacteristics = characteristics

		// Context information
		contextInfo := &sfera.SPContextInformation{}

		for _, board := range link.Boards.Items() {
			if kmBoard, ok := isKilometerBoard(board); ok {
				contextInfo.KilometreReferencePoints = append(contextInfo.KilometreReferencePoints,
					sfera.KilometreReferencePoint{
						Identifier: kmBoard.Id.Information,
						Location:   kmBoard.LinkPosition.GetDistance(direction),
						KmRef:      float64(*kmBoard.Kilometer),
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

// isKilometerBoard checks if a Board is a Kilometertavla and returns it if so.
func isKilometerBoard(board *dm.Board) (*dm.Board, bool) {
	return board, board.IsKilometerBoard()
}

// CreateTimingPoint creates a SFERA TimingPoint from domain model entities.
func CreateTimingPoint(placeCenter *dm.PlaceCenter, track *dm.StationTrack, direction dm.LinkDirection) sfera.TimingPoint {
	tp := sfera.TimingPoint{
		TPID:         fmt.Sprintf("%d:%s", placeCenter.Place.PLC, track.Number),
		Location:     placeCenter.LinkPosition.GetDistance(direction),
		StationTrack: track.Number,
	}
	tp.TPNames = append(tp.TPNames, sfera.TPName{
		Name:     fmt.Sprintf("%s, %s", track.Name, placeCenter.Place.Name),
		Language: "se",
	})
	return tp
}

// CreateStaticSpeedProfile creates a SFERA StaticSpeedProfile from domain model entities.
func CreateStaticSpeedProfile(speedProfile *dm.SpeedProfile, direction dm.LinkDirection, isDefaultProfile bool) sfera.StaticSpeedProfile {
	ssp := sfera.StaticSpeedProfile{}

	segments := dm.ExtentEntitiesInDirection(speedProfile.Segments.Items(), direction)

	for i, segment := range segments {
		speed := float64(segment.Speed)
		if i == 0 {
			start := &sfera.SSPComplexType{
				SSPSpeed: &speed,
			}

			if !isDefaultProfile {
				category := sfera.SpecificSSPCategoryOtherReplaces
				ncDiff := sfera.NCTrainTypePassenger
				start.SpecificSSPs = append(start.SpecificSSPs, sfera.SpecificSSP{
					OtherSpecificSSP: &sfera.OtherSpecificSSP{
						TrainCategoryCode: string(segment.Category),
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
				Location: segment.Extent.GetStartDistance(direction),
			}

			if !isDefaultProfile {
				change.SpecificSSPs = append(change.SpecificSSPs, sfera.SpecificSSP{
					OtherSpecificSSP: &sfera.OtherSpecificSSP{
						TrainCategoryCode: string(segment.Category),
					},
				})
			}

			ssp.Changes = append(ssp.Changes, change)
		}
	}

	return ssp
}

// CreateGradientAverage creates a SFERA GradientAverage from a Lank's gradients.
func CreateGradientAverage(link *dm.Link, direction dm.LinkDirection) sfera.GradientAverage {
	ga := sfera.GradientAverage{}

	slopes := dm.ExtentEntitiesInDirection(link.Slopes.Items(), direction)

	for i, slope := range slopes {
		if i == 0 {
			ga.Start = &sfera.GradientComplexType{
				GradientValue:     math.Abs(slope.Gradient),
				GradientDirection: getGradientDirection(slope, direction),
			}
		} else {
			ga.Changes = append(ga.Changes, sfera.GradientAverageChange{
				GradientComplexType: sfera.GradientComplexType{
					GradientValue:     math.Abs(slope.Gradient),
					GradientDirection: getGradientDirection(slope, direction),
				},
				Location: slope.Extent.GetStartDistance(direction),
			})
		}
	}

	return ga
}

// createProfileId creates a SFERA segment profile ID from a LankId and direction.
func createProfileId(linkId dm.LinkId, direction dm.LinkDirection) string {
	return fmt.Sprintf("%v:%s", linkId, direction)
}

// getGradientDirection returns the SFERA gradient direction for a Slope.
func getGradientDirection(slope *dm.Slope, direction dm.LinkDirection) sfera.GradientDirectionType {
	gradient := slope.Gradient
	if direction == dm.LinkDirectionAgainst {
		gradient = -gradient
	}

	if gradient >= 0 {
		return sfera.GradientDirectionUphill
	}
	return sfera.GradientDirectionDownhill
}
