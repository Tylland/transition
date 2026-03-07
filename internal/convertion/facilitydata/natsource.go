package convertion

import (
	dm "github.com/tylland/transition/internal/domainmodel"
	fd "github.com/tylland/transition/internal/schema/facilitydata"
)

// FacilityDataNatSource creates a Nat (network) from FacilityData schema data.
type FacilityDataNatSource struct {
	Infrastructure fd.Infrastructure
	Facility       fd.Facility
	Divisions      []fd.Division
}

// NewFacilityDataNatSource creates a new FacilityDataNatSource from a FacilityData Root.
func NewFacilityDataNatSource(baseRoot fd.Root) *FacilityDataNatSource {
	return &FacilityDataNatSource{
		Infrastructure: baseRoot.Infrastructure,
		Facility:       baseRoot.Facility,
		Divisions:      baseRoot.Divisions,
	}
}

// CreateNat builds a domain model Nat from the FacilityData.
func (s *FacilityDataNatSource) CreateNat() *dm.Nat {
	nat := dm.NewNat()

	// Import nodes
	for _, node := range s.Infrastructure.Nodes {
		nat.SaveNod(FDNodeCreateNodId(node), node.Identification.Designation)
	}

	// Import links
	for _, link := range s.Infrastructure.Links {
		lankId := FDLinkCreateLankId(link)
		nat.GetOrCreateLank(lankId, link.Length)
	}

	// Create node paths (nodvägar) by connecting all links that share a node.
	// For each node, every pair of connected links gets a bidirectional path.
	for _, nod := range nat.Noder() {
		// Collect all links connected to this node
		var connectedLinks []*dm.Lank
		for _, lank := range nat.Lankar() {
			if lank.Fran == nod || lank.Till == nod {
				connectedLinks = append(connectedLinks, lank)
			}
		}
		// Create bidirectional paths between all pairs
		for i := 0; i < len(connectedLinks); i++ {
			for j := 0; j < len(connectedLinks); j++ {
				if i != j {
					nod.AddVag(connectedLinks[i], connectedLinks[j])
				}
			}
		}
	}

	// Import locations (platsmittar)
	for _, location := range s.Infrastructure.Locations {
		plats := dm.NewPlats(
			location.Identification.BusinessId.LocationSignature,
			location.Identification.BusinessId.Name,
			location.Type,
			location.PLCCode,
		)
		nat.AddPlats(plats)

		for _, linkPosition := range location.LinkPositions {
			FDWithLank(nat, linkPosition, func(lank *dm.Lank) {
				lank.Platsmittar.Add(dm.NewPlatsmitt(plats, FDGetLankposition(lank, linkPosition, dm.GilltighetsriktningBada)))
			})
		}
	}

	// Import track numbers (spårnummer)
	for _, division := range s.Divisions {
		if division.DivisionType == fd.DivisionTypesTracknumber {
			for _, subdivision := range division.Subdivisions {
				for _, linkPosition := range subdivision.LinkPositions {
					FDWithLank(nat, linkPosition, func(lank *dm.Lank) {
						lank.Sparnummer.Add(dm.NewSparnummer(
							subdivision.DivisionCode,
							subdivision.Description,
							FDGetLankutbredning(lank, linkPosition, dm.GilltighetsriktningBada),
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
					FDWithLank(nat, linkPosition, func(lank *dm.Lank) {
						lank.Linjespar.Add(dm.NewLinjespar(
							subdivision.DivisionCode,
							subdivision.Description,
							FDGetLankutbredning(lank, linkPosition, dm.GilltighetsriktningBada),
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
				lankId := FDLinkIdToLankId(linkPosition.LinkIdentification.LinkId)
				nat.WithLank(lankId, func(lank *dm.Lank) {
					signalId := FDBusinessIdToSignalId(informationPoint.Identification.BusinessId)
					lank.Signaler.Add(dm.NewSignal(
						signalId,
						signal.TypeOfMainSignal,
						FDGetLankposition(lank, linkPosition, FDDirectionToGilltighetsriktning(signal.Direction)),
					))
				})
			}
		}
	}

	// Import kilometer boards
	for _, kilometerBoard := range s.Facility.SignalingFacilities.KilometerBoards {
		for _, linkPosition := range kilometerBoard.LinkPositions {
			lankId := FDLinkIdToLankId(linkPosition.LinkIdentification.LinkId)
			nat.WithLank(lankId, func(lank *dm.Lank) {
				lank.Tavlor.Add(dm.NewKilometertavla(
					kilometerBoard.Identification.GlobalId.UniversalId,
					kilometerBoard.KilometerPost,
					FDGetLankposition(lank, linkPosition, dm.GilltighetsriktningBada),
				))
			})
		}
	}

	// Import speed profiles
	for _, speedProfile := range s.Facility.OtherFacilities.SpeedProfiles {
		for _, linkPosition := range speedProfile.LinkPositions {
			lankId := FDLinkIdToLankId(linkPosition.LinkIdentification.LinkId)
			nat.WithLank(lankId, func(lank *dm.Lank) {
				gilltighetsriktning := FDDirectionToGilltighetsriktning(speedProfile.Direction)
				lu := FDGetLankutbredning(lank, linkPosition, gilltighetsriktning)

				lank.AddHastighetsstracka(dm.NewHastighetsstracka(dm.TagkategoriA, speedProfile.ATrainSpeed, lu))
				lank.AddHastighetsstracka(dm.NewHastighetsstracka(dm.TagkategoriB, speedProfile.BTrainSpeed, FDGetLankutbredning(lank, linkPosition, gilltighetsriktning)))
				lank.AddHastighetsstracka(dm.NewHastighetsstracka(dm.TagkategoriC, speedProfile.CTrainSpeed, FDGetLankutbredning(lank, linkPosition, gilltighetsriktning)))
				lank.AddHastighetsstracka(dm.NewHastighetsstracka(dm.TagkategoriS, speedProfile.STrainSpeed, FDGetLankutbredning(lank, linkPosition, gilltighetsriktning)))
				_ = lu // used for first call
			})
		}
	}

	// Import vertical geometries (gradients)
	for _, verticalGeometry := range s.Facility.Geometries.VerticalGeometries {
		for _, linkPosition := range verticalGeometry.LinkPositions {
			FDWithLank(nat, linkPosition, func(lank *dm.Lank) {
				lank.Lutningar.Add(dm.NewLutning(
					verticalGeometry.Gradient,
					FDGetLankutbredning(lank, linkPosition, dm.GilltighetsriktningBada),
				))
			})
		}
	}

	return nat
}
