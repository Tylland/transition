package facilitydata

import (
	dm "github.com/tylland/transition/internal/domainmodel"
	fd "github.com/tylland/transition/internal/schema/facilitydata"
)

// DomainModelBuilder creates a Nat (network) from FacilityData schema data.
type DomainModelBuilder struct {
	Infrastructure fd.Infrastructure
	Facility       fd.Facility
	Divisions      []fd.Division
}

// NewDomainModelBuilder creates a new FacilityDataNatSource from a FacilityData Root.
func NewDomainModelBuilder(baseRoot fd.Root) *DomainModelBuilder {
	return &DomainModelBuilder{
		Infrastructure: baseRoot.Infrastructure,
		Facility:       baseRoot.Facility,
		Divisions:      baseRoot.Divisions,
	}
}

// Build builds a domain model Network from the FacilityData.
func (s *DomainModelBuilder) Build() *dm.Network {
	network := dm.NewNetwork()

	// Import nodes
	for _, node := range s.Infrastructure.Nodes {
		network.SaveNode(CreateNodeIdForNade(node), node.Identification.Designation)
	}

	// Import links
	for _, link := range s.Infrastructure.Links {
		linkId := CreateLinkIdForLink(link)
		network.GetOrCreateLink(linkId, link.Length)
	}

	// Create node paths (nodvägar) by connecting all links that share a node.
	// For each node, every pair of connected links gets a bidirectional path.
	for _, node := range network.Nodes() {
		// Collect all links connected to this node
		var connectedLinks []*dm.Link
		for _, link := range network.Links() {
			if link.From == node || link.To == node {
				connectedLinks = append(connectedLinks, link)
			}
		}
		// Create bidirectional paths between all pairs
		for i := 0; i < len(connectedLinks); i++ {
			for j := 0; j < len(connectedLinks); j++ {
				if i != j {
					node.AddRoute(connectedLinks[i], connectedLinks[j])
				}
			}
		}
	}

	// Import locations (platsmittar)
	for _, location := range s.Infrastructure.Locations {
		place := dm.NewPlace(
			location.Identification.BusinessId.LocationSignature,
			location.Identification.BusinessId.Name,
			location.Type,
			location.PLCCode,
		)
		network.AddPlace(place)

		for _, position := range location.LinkPositions {
			FDWithLink(network, position, func(lank *dm.Link) {
				lank.PlaceCenters.Add(dm.NewPlaceCenter(place, CreateLinkPosition(lank, position, dm.ValidityDirectionBoth)))
			})
		}
	}

	// Import track numbers (spårnummer)
	for _, division := range s.Divisions {
		if division.DivisionType == fd.DivisionTypesTracknumber {
			for _, subdivision := range division.Subdivisions {
				for _, position := range subdivision.LinkPositions {
					FDWithLink(network, position, func(link *dm.Link) {
						link.StationTracks.Add(dm.NewStationTrack(
							subdivision.DivisionCode,
							subdivision.Description,
							CreateLinkExtent(link, position, dm.ValidityDirectionBoth),
						))
					})
				}
			}
		}
	}

	// Import line tracks (linjespår)
	for _, division := range s.Divisions {
		if division.DivisionType == fd.DivisionTypesTrackUpDownSingle {
			for _, subdivision := range division.Subdivisions {
				for _, linkPosition := range subdivision.LinkPositions {
					FDWithLink(network, linkPosition, func(lank *dm.Link) {
						lank.LineTracks.Add(dm.NewLineTrack(
							subdivision.DivisionCode,
							subdivision.Description,
							CreateLinkExtent(lank, linkPosition, dm.ValidityDirectionBoth),
						))
					})
				}
			}
		}
	}

	// Import signals
	for _, informationPoint := range s.Facility.SignalingFacilities.InformationPoints {
		if informationPoint.Signal != nil {
			signal := informationPoint.Signal
			for _, linkPosition := range informationPoint.LinkPositions {
				linkId := CreateLinkId(linkPosition.LinkIdentification.LinkId)
				network.WithLink(linkId, func(lank *dm.Link) {
					signalId := CreateSignalId(informationPoint.Identification.BusinessId)
					lank.Signals.Add(dm.NewSignal(
						signalId,
						signal.TypeOfMainSignal,
						CreateLinkPosition(lank, linkPosition, CreateValidityDirection(signal.Direction)),
					))
				})
			}
		}
	}

	// Import kilometer boards
	for _, kilometerBoard := range s.Facility.SignalingFacilities.KilometerBoards {
		for _, linkPosition := range kilometerBoard.LinkPositions {
			linkId := CreateLinkId(linkPosition.LinkIdentification.LinkId)
			network.WithLink(linkId, func(lank *dm.Link) {
				lank.Boards.Add(dm.NewKilometerBoard(
					kilometerBoard.Identification.GlobalId.UniversalId,
					kilometerBoard.KilometerPost,
					CreateLinkPosition(lank, linkPosition, dm.ValidityDirectionBoth),
				))
			})
		}
	}

	// Import speed profiles
	for _, speedProfile := range s.Facility.OtherFacilities.SpeedProfiles {
		for _, linkPosition := range speedProfile.LinkPositions {
			linkId := CreateLinkId(linkPosition.LinkIdentification.LinkId)
			network.WithLink(linkId, func(lank *dm.Link) {
				validityDirection := CreateValidityDirection(speedProfile.Direction)
				lu := CreateLinkExtent(lank, linkPosition, validityDirection)

				lank.AddSpeedSegment(dm.NewSpeedSegment(dm.TrainCateoryA, speedProfile.ATrainSpeed, lu))
				lank.AddSpeedSegment(dm.NewSpeedSegment(dm.TrainCateoryB, speedProfile.BTrainSpeed, CreateLinkExtent(lank, linkPosition, validityDirection)))
				lank.AddSpeedSegment(dm.NewSpeedSegment(dm.TrainCateoryC, speedProfile.CTrainSpeed, CreateLinkExtent(lank, linkPosition, validityDirection)))
				lank.AddSpeedSegment(dm.NewSpeedSegment(dm.TrainCateoryS, speedProfile.STrainSpeed, CreateLinkExtent(lank, linkPosition, validityDirection)))
				_ = lu // used for first call
			})
		}
	}

	// Import vertical geometries (gradients)
	for _, verticalGeometry := range s.Facility.Geometries.VerticalGeometries {
		for _, linkPosition := range verticalGeometry.LinkPositions {
			FDWithLink(network, linkPosition, func(lank *dm.Link) {
				lank.Slopes.Add(dm.NewLutning(
					verticalGeometry.Gradient,
					CreateLinkExtent(lank, linkPosition, dm.ValidityDirectionBoth),
				))
			})
		}
	}

	return network
}
