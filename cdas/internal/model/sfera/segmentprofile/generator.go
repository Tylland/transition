package segmentprofile

// import (
// 	"fmt"
// 	"math"

// 	dm "github.com/tylland/transition/cdas/internal/core/domain/network"
// 	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
// )

// func ptr[T any](t T) *T {
// 	return &t
// }

// // Generate generates SFERA segment profiles for both directions
// func Generate(network *dm.Network) []*sfera.SegmentProfile {
// 	var profiles []*sfera.SegmentProfile

// 	profiles = append(profiles, GenerateForDirection(network, dm.LinkDirectionWith)...)
// 	profiles = append(profiles, GenerateForDirection(network, dm.LinkDirectionAgainst)...)

// 	return profiles
// }

// // GenerateForDirection generates SFERA segment profiles for a
// // specific direction and writes them using the given writer.
// func GenerateForDirection(network *dm.Network, direction dm.LinkDirection) []*sfera.SegmentProfile {
// 	var profiles []*sfera.SegmentProfile

// 	for _, link := range network.Links() {
// 		imID := sfera.NewID("0074")
// 		name := "namn"

// 		segment := &sfera.SegmentProfile{
// 			SegmentProfileIdentificationComplexType: &sfera.SegmentProfileIdentificationComplexType{
// 				SPID:   createProfileId(link.Id, direction),
// 				SPZone: &sfera.SPZoneComplexType{IMID: &imID}},

// 			SPName:         &name,
// 			SPVersionMajor: "1",
// 			SPVersionMinor: ptr("0"),
// 			SPLength:       link.Length,
// 			SPStatus:       sfera.SPStatusValid,
// 		}

// 		// Points
// 		points := &sfera.SPPoints{}

// 		// Timing points from place centers
// 		for _, placeCenter := range link.PlaceCenters.Items() {
// 			var stationTrack *dm.StationTrack
// 			for _, sn := range link.StationTracks.Items() {
// 				if sn.IsValid(direction) {
// 					stationTrack = sn
// 					break
// 				}
// 			}

// 			if stationTrack != nil {
// 				points.TimingPoint = append(points.TimingPoint, CreateTimingPoint(placeCenter, stationTrack, direction))
// 			}
// 		}

// 		// Signals
// 		for _, signal := range link.Signals.Items() {
// 			if signal.LinkPosition.IsValid(direction) {
// 				signalId := sfera.SignalID{
// 					SignalIDPhysical: signal.Id.String(),
// 					SignalIDObject:   ptr(signal.Id.String()),
// 					Location:         signal.LinkPosition.GetDistance(direction),
// 				}
// 				points.Signal = append(points.Signal, &sfera.SignalComplexType{SignalID: &signalId})
// 			}
// 		}

// 		// // Balise tavlor
// 		// for _, baliseBoard := range link.Boards.Items() {
// 		// 	if baliseBoard.Id.Type == "Bal" && baliseBoard.LinkPosition.IsValid(direction) {
// 		// 		balise := &sfera.Balise{SPGenericPointComplexType: &sfera.SPGenericPointComplexType{
// 		// 			Identifier: ptr(baliseBoard.Id.String()),
// 		// 			Location:   baliseBoard.LinkPosition.GetDistance(direction),
// 		// 		}}

// 		// 		points.Balise = append(points.Balise, balise)
// 		// 	}
// 		// }

// 		// Unprotected level crossings
// 		for _, vagskydd := range link.Levelcrossings.Items() {
// 			if !vagskydd.HasProtection {
// 				points.UnprotectedLevelCrossingStop = append(points.UnprotectedLevelCrossingStop,
// 					&sfera.UnprotectedLevelCrossingStop{SPGenericPointComplexType: createSPGenericPoint(vagskydd.Name, vagskydd.Extent.GetStartDistance(direction))})
// 			}
// 		}

// 		segment.SPPoints = points

// 		// Areas
// 		areas := &sfera.SPAreas{}
// 		for _, platform := range link.Platforms.Items() {
// 			startLoc := platform.Extent.GetStartDistance(direction)
// 			endLoc := platform.Extent.GetEndDistance(direction)
// 			areas.Platform = append(areas.Platform, &sfera.Platform{SPGenericAreaComplexType: &sfera.SPGenericAreaComplexType{
// 				Identifier:        ptr(platform.Name),
// 				StartLocation:     &startLoc,
// 				EndLocation:       &endLoc,
// 				StartEndQualifier: sfera.StartEndQualifierStartsEnds,
// 			}})
// 		}
// 		segment.SPAreas = areas

// 		// Characteristics
// 		characteristics := &sfera.SPCharacteristics{}

// 		// Default speed profile (category A)
// 		var defaultSpeedProfile *dm.SpeedProfile
// 		for _, speedProfile := range link.SpeedProfiles.Items() {
// 			if speedProfile.Category == "A" {
// 				defaultSpeedProfile = speedProfile
// 				break
// 			}
// 		}

// 		if defaultSpeedProfile != nil {
// 			characteristics.StaticSpeedProfile = append(characteristics.StaticSpeedProfile,
// 				CreateStaticSpeedProfile(defaultSpeedProfile, direction, true))
// 		}

// 		for _, hastighetsprofil := range link.SpeedProfiles.Items() {
// 			characteristics.StaticSpeedProfile = append(characteristics.StaticSpeedProfile,
// 				CreateStaticSpeedProfile(hastighetsprofil, direction, false))
// 		}

// 		ga := CreateGradientAverage(link, direction)
// 		characteristics.Gradient = &ga
// 		segment.SPCharacteristics = characteristics

// 		// Context information
// 		contextInfo := &sfera.SPContextInformation{}

// 		for _, board := range link.Boards.Items() {
// 			if kmBoard, ok := isKilometerBoard(board); ok {
// 				contextInfo.KilometreReferencePoint = append(contextInfo.KilometreReferencePoint,
// 					&sfera.KilometreReferencePoint{
// 						SPGenericPointComplexType: createSPGenericPoint(kmBoard.Id.Information, kmBoard.LinkPosition.GetDistance(direction)),
// 						KMReference:               &sfera.KMReference{KmRef: float64(*kmBoard.Kilometer)}},
// 				)
// 			}
// 		}

// 		contextInfo.LineAndTrackIdentifier = &sfera.LineAndTrackIdentifier{
// 			LineAndTrackIdentifierStart: &sfera.LineAndTrackIdentifierComplexType{},
// 		}

// 		segment.SPContextInformation = contextInfo

// 		profiles = append(profiles, segment)
// 	}

// 	return profiles
// }

// func createSPGenericPoint(id string, loc float64) *sfera.SPGenericPointComplexType {
// 	return &sfera.SPGenericPointComplexType{
// 		Identifier: ptr(id),
// 		Location:   loc,
// 	}
// }

// func createSPGenericArea(id string, start, end float64) *sfera.SPGenericAreaComplexType {
// 	return &sfera.SPGenericAreaComplexType{Identifier: ptr(id), StartLocation: &start, EndLocation: &end}
// }

// // isKilometerBoard checks if a Board is a Kilometertavla and returns it if so.
// func isKilometerBoard(board *dm.Board) (*dm.Board, bool) {
// 	return board, board.IsKilometerBoard()
// }

// // CreateTimingPoint creates a SFERA TimingPoint from domain model entities.
// func CreateTimingPoint(placeCenter *dm.PlaceCenter, track *dm.StationTrack, direction dm.LinkDirection) *sfera.TimingPoint {
// 	tp := &sfera.TimingPoint{
// 		TPID:         fmt.Sprintf("%d:%s", placeCenter.Place.PLC, track.Number),
// 		Location:     placeCenter.LinkPosition.GetDistance(direction),
// 		StationTrack: ptr(track.Number),
// 	}
// 	tp.TPName = append(tp.TPName, &sfera.TPName{
// 		Name:     fmt.Sprintf("%s, %s", track.Name, placeCenter.Place.Name),
// 		Language: ptr("se"),
// 	})
// 	return tp
// }

// // CreateStaticSpeedProfile creates a SFERA StaticSpeedProfile from domain model entities.
// func CreateStaticSpeedProfile(speedProfile *dm.SpeedProfile, direction dm.LinkDirection, isDefaultProfile bool) *sfera.StaticSpeedProfile {
// 	ssp := &sfera.StaticSpeedProfile{}

// 	segments := dm.ExtentEntitiesInDirection(speedProfile.Segments.Items(), direction)

// 	for i, segment := range segments {
// 		speed := float64(segment.Speed)
// 		if i == 0 {
// 			start := &sfera.SSPComplexType{
// 				SSPSpeed: &speed,
// 			}

// 			if !isDefaultProfile {
// 				category := sfera.SpecificSSPCategoryOtherReplaces
// 				ncDiff := sfera.NCTrainTypePassenger
// 				start.SpecificSSP = append(start.SpecificSSP, &sfera.SpecificSSPComplexType{
// 					OtherSpecificSSP: &sfera.OtherSpecificSSP{
// 						TrainCategoryCode: ptr(string(segment.Category)),
// 						SSPNCDIFF:         ptr(string(ncDiff)),
// 					},
// 					SpecificSSPCategory: ptr(string(category)),
// 				})
// 			}

// 			ssp.StaticSpeedProfileStart = start
// 		} else {
// 			change := &sfera.StaticSpeedProfileChange{
// 				SSPComplexType: &sfera.SSPComplexType{
// 					SSPSpeed: &speed,
// 				},
// 				Location: segment.Extent.GetStartDistance(direction),
// 			}

// 			if !isDefaultProfile {
// 				change.SpecificSSP = append(change.SpecificSSP, &sfera.SpecificSSPComplexType{
// 					OtherSpecificSSP: &sfera.OtherSpecificSSP{
// 						TrainCategoryCode: ptr(string(segment.Category)),
// 					},
// 				})
// 			}

// 			ssp.StaticSpeedProfileChange = append(ssp.StaticSpeedProfileChange, change)
// 		}
// 	}

// 	return ssp
// }

// // CreateGradientAverage creates a SFERA GradientAverage from a Lank's gradients.
// func CreateGradientAverage(link *dm.Link, direction dm.LinkDirection) sfera.Gradient {
// 	ga := sfera.Gradient{}

// 	slopes := dm.ExtentEntitiesInDirection(link.Slopes.Items(), direction)

// 	for i, slope := range slopes {
// 		if i == 0 {
// 			ga.GradientStart = &sfera.GradientComplexType{
// 				GradientValue:     math.Abs(slope.Gradient),
// 				GradientDirection: getGradientDirection(slope, direction),
// 			}
// 		} else {
// 			ga.GradientChange = append(ga.GradientChange, &sfera.GradientChange{
// 				GradientComplexType: &sfera.GradientComplexType{
// 					GradientValue:     math.Abs(slope.Gradient),
// 					GradientDirection: getGradientDirection(slope, direction),
// 				},
// 				Location: slope.Extent.GetStartDistance(direction),
// 			})
// 		}
// 	}

// 	return ga
// }

// // createProfileId creates a SFERA segment profile ID from a LankId and direction.
// func createProfileId(linkId dm.LinkId, direction dm.LinkDirection) string {
// 	return fmt.Sprintf("%v:%s", linkId, direction)
// }

// // getGradientDirection returns the SFERA gradient direction for a Slope.
// func getGradientDirection(slope *dm.Slope, direction dm.LinkDirection) sfera.GradientDirectionType {
// 	gradient := slope.Gradient
// 	if direction == dm.LinkDirectionAgainst {
// 		gradient = -gradient
// 	}

// 	if gradient >= 0 {
// 		return sfera.GradientDirectionUphill
// 	}
// 	return sfera.GradientDirectionDownhill
// }
