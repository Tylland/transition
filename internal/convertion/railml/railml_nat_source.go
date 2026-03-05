package convertion

import (
	"strconv"

	dm "github.com/tylland/transition/domainmodel"
	rml "github.com/tylland/transition/schema/railml"
)

// RailMLNatSource creates a Nat (network) from RailML infrastructure data.
type RailMLNatSource struct {
	Infrastructure rml.Infrastructure
}

// NewRailMLNatSource creates a new RailMLNatSource.
func NewRailMLNatSource(infrastructure rml.Infrastructure) *RailMLNatSource {
	return &RailMLNatSource{Infrastructure: infrastructure}
}

// CreateNat builds a domain model Nat from the RailML infrastructure.
func (s *RailMLNatSource) CreateNat() *dm.Nat {
	infra := NewRailMLTopologyReader(s.Infrastructure)
	nat := dm.NewNat()

	// First pass: save nodes from topology relations
	for _, element := range s.Infrastructure.Topology.NetElements {
		lankId, err := dm.ParseLankId(element.Id)
		if err != nil {
			continue
		}

		for _, relation := range infra.GetRelations(element) {
			relatedElement, err := infra.GetRelatedElement(element, relation)
			if err != nil {
				continue
			}
			relatedLankId, err := dm.ParseLankId(relatedElement.Id)
			if err != nil {
				continue
			}

			nodId, err := dm.GetIntersectingNodId(lankId, relatedLankId)
			if err != nil {
				continue
			}

			name := ""
			if len(relation.Name) > 0 {
				name = relation.Name[0].NameProperty
			}
			nat.SaveNod(nodId, name)
		}
	}

	// Second pass: create links and node paths (vägar)
	for _, element := range s.Infrastructure.Topology.NetElements {
		lankId, err := dm.ParseLankId(element.Id)
		if err != nil {
			continue
		}
		lank := nat.GetOrCreateLank(lankId, element.Length)

		for _, relation := range infra.GetRelations(element) {
			relatedElement, err := infra.GetRelatedElement(element, relation)
			if err != nil {
				continue
			}
			relatedLankId, err := dm.ParseLankId(relatedElement.Id)
			if err != nil {
				continue
			}
			relatedLank := nat.GetOrCreateLank(relatedLankId, relatedElement.Length)

			nodId, err := dm.GetIntersectingNodId(lankId, relatedLankId)
			if err != nil {
				continue
			}

			nod := nat.GetOrCreateNod(nodId)
			nod.AddVag(lank, relatedLank)
		}
	}

	// Import balises (placeholder - matches C# comment-out behavior)
	for _, balise := range s.Infrastructure.FunctionalInfrastructure.Balises {
		for _, location := range balise.SpotLocation {
			lankId, err := dm.ParseLankId(location.NetElementRef)
			if err != nil {
				continue
			}
			_ = nat.GetLank(lankId)
			// Balis creation is commented out in original C# code
			_ = dm.NewBalis("zsfdasf", 0)
		}
	}

	// Import level crossings
	for _, levelCrossing := range s.Infrastructure.FunctionalInfrastructure.LevelCrossingsIS {
		for _, location := range levelCrossing.AreaLocation {
			for _, associatedNetElement := range location.AssociatedNetElement {
				harSkydd := levelCrossing.Protection != nil && levelCrossing.Protection.HasActiveProtection

				lankId, err := dm.ParseLankId(associatedNetElement.NetElementRef)
				if err != nil {
					continue
				}
				lank := nat.GetLank(lankId)

				vagskydd := dm.NewVagskydd(
					RMLGetFirstName(levelCrossing.Name),
					harSkydd,
					dm.NewLankutbredning(lank, associatedNetElement.IntrinsicCoordBegin, associatedNetElement.IntrinsicCoordEnd, dm.GilltighetsriktningBada),
				)
				lank.Vagskydd.Add(vagskydd)
			}
		}
	}

	// Import signals
	for _, signalIs := range s.Infrastructure.FunctionalInfrastructure.SignalsIS {
		if signalIs.IsAnnouncementSignal != nil {
			signalId, err := dm.ParseSignalId(RMLGetFirstName(signalIs.Name))
			if err != nil {
				continue
			}

			for _, location := range signalIs.SpotLocation {
				lankId, err := dm.ParseLankId(location.NetElementRef)
				if err != nil {
					continue
				}
				lank := nat.GetLank(lankId)

				signal := dm.NewSignal(
					signalId,
					signalId.Typ,
					dm.NewLankposition(lank, location.IntrinsicCoord, RMLSpotLocationGetRiktning(location)),
				)
				lank.Signaler.Add(signal)
			}
		}

		if signalIs.IsInformationSignal != nil {
			tavlaId, err := dm.ParseTavlaId(RMLGetFirstName(signalIs.Name))
			if err != nil {
				continue
			}

			for _, location := range signalIs.SpotLocation {
				err := nat.WithLankStr(location.NetElementRef, func(lank *dm.Lank) {
					tavla := dm.NewTavla(
						tavlaId,
						dm.NewLankposition(lank, location.IntrinsicCoord, RMLSpotLocationGetRiktning(location)),
					)
					lank.Tavlor.Add(tavla)
				})
				_ = err
			}
		}

		if signalIs.IsMilepost != nil {
			for _, location := range signalIs.SpotLocation {
				shownValue, err := strconv.Atoi(signalIs.IsMilepost.ShownValue)
				if err != nil {
					continue
				}

				err = nat.WithLankStr(location.NetElementRef, func(lank *dm.Lank) {
					tavla := dm.NewKilometertavla(
						RMLGetFirstName(signalIs.Name),
						shownValue,
						dm.NewLankposition(lank, location.IntrinsicCoord, RMLSpotLocationGetRiktning(location)),
					)
					lank.Tavlor.Add(tavla)
				})
				_ = err
			}
		}
	}

	// Import operational points
	for _, operationalPoint := range s.Infrastructure.FunctionalInfrastructure.OperationalPoints {
		for _, location := range operationalPoint.SpotLocation {
			plats := dm.NewPlats(
				RMLGetFirstName(operationalPoint.Name),
				RMLGetFirstDescription(operationalPoint.Name),
				"-",
				0,
			)

			err := nat.WithLankStr(location.NetElementRef, func(lank *dm.Lank) {
				lank.Platsmittar.Add(dm.NewPlatsmitt(
					plats,
					dm.NewLankposition(lank, location.IntrinsicCoord, RMLSpotLocationGetRiktning(location)),
				))
			})
			_ = err
		}
	}

	// Import platforms
	for _, platform := range s.Infrastructure.FunctionalInfrastructure.Platforms {
		langd := 0.0
		if len(platform.Width) > 0 {
			langd = platform.Width[0].Value
		}

		for _, location := range platform.AreaLocation {
			for _, associatedNetElement := range location.AssociatedNetElement {
				err := nat.WithLankStr(associatedNetElement.NetElementRef, func(lank *dm.Lank) {
					lank.Plattformar.Add(dm.NewPlattform(
						RMLGetFirstName(platform.Name),
						langd,
						dm.NewLankutbredning(lank, associatedNetElement.IntrinsicCoordBegin, associatedNetElement.IntrinsicCoordEnd, dm.GilltighetsriktningBada),
					))
				})
				_ = err
			}
		}
	}

	// Import speed sections
	for _, speedSection := range s.Infrastructure.FunctionalInfrastructure.Speeds {
		for _, location := range speedSection.LinearLocation {
			for _, associatedNetElement := range location.AssociatedNetElement {
				maxSpeed, err := strconv.Atoi(speedSection.MaxSpeed)
				if err != nil {
					continue
				}

				err = nat.WithLankStr(associatedNetElement.NetElementRef, func(lank *dm.Lank) {
					lank.AddHastighetsstracka(dm.NewHastighetsstracka(
						RMLGetFirstName(speedSection.Name),
						maxSpeed,
						dm.NewLankutbredning(lank, associatedNetElement.IntrinsicCoordBegin, associatedNetElement.IntrinsicCoordEnd, RMLApplicationDirectionToGilltighetsriktning(location.ApplicationDirection)),
					))
				})
				_ = err
			}
		}
	}

	return nat
}
