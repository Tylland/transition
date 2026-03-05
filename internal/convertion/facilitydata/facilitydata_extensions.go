package convertion

import (
	dm "github.com/tylland/transition/internal/domainmodel"
	fd "github.com/tylland/transition/internal/schema/facilitydata"
)

// FacilityDataExtensions provides conversion utilities from FacilityData schema types
// to domain model types.

// FDNodeIdToNodId converts a FacilityData NodeId to a domain model NodId.
func FDNodeIdToNodId(nodeId fd.NodeId) dm.NodId {
	return dm.NewNodId(nodeId.LocationNumber, nodeId.NodeNumber)
}

// FDNodeCreateNodId creates a domain NodId from a FacilityData Node.
func FDNodeCreateNodId(node fd.Node) dm.NodId {
	return FDNodeIdToNodId(node.Identification.NodeId)
}

// FDLinkIdToLankId converts a FacilityData LinkId to a domain model LankId.
func FDLinkIdToLankId(linkId fd.LinkId) dm.LankId {
	franNodId := FDNodeIdToNodId(linkId.NodeFrom)
	tillNodId := FDNodeIdToNodId(linkId.NodeTo)
	return dm.NewLankId(franNodId, tillNodId, linkId.Linknumber)
}

// FDLinkCreateLankId creates a domain LankId from a FacilityData Link.
func FDLinkCreateLankId(link fd.Link) dm.LankId {
	return FDLinkIdToLankId(link.Identification.LinkId)
}

// FDDirectionToGilltighetsriktning converts a FacilityData Direction to domain Gilltighetsriktning.
func FDDirectionToGilltighetsriktning(direction fd.Direction) dm.Gilltighetsriktning {
	switch direction {
	case fd.DirectionIn:
		return dm.GilltighetsriktningNormal
	case fd.DirectionAgainst:
		return dm.GilltighetsriktningOmvand
	case fd.DirectionAny:
		return dm.GilltighetsriktningBada
	default:
		return dm.GilltighetsriktningIngen
	}
}

// FDBusinessIdToSignalId converts a FacilityData BusinessId to a domain SignalId.
func FDBusinessIdToSignalId(businessId fd.BusinessId) *dm.SignalId {
	return dm.NewSignalId(businessId.LocationSignature, businessId.Name, businessId.ObjectType)
}

// FDGetLankposition creates a Lankposition from a FacilityData LinkPosition on a given Lank.
func FDGetLankposition(lank *dm.Lank, linkPosition fd.LinkPosition, gilltighetsriktning dm.Gilltighetsriktning) *dm.Lankposition {
	return dm.NewLankposition(lank, linkPosition.StartDistanceFromNode/lank.Langd, gilltighetsriktning)
}

// FDGetLankutbredning creates a Lankutbredning from a FacilityData LinkPosition on a given Lank.
func FDGetLankutbredning(lank *dm.Lank, linkPosition fd.LinkPosition, gilltighetsriktning dm.Gilltighetsriktning) *dm.Lankutbredning {
	return dm.NewLankutbredning(
		lank,
		linkPosition.StartDistanceFromNode/lank.Langd,
		linkPosition.EndDistanceFromNode/lank.Langd,
		gilltighetsriktning,
	)
}

// FDWithLank executes an action on a Lank identified by a FacilityData LinkPosition.
func FDWithLank(nat *dm.Nat, linkPosition fd.LinkPosition, action func(*dm.Lank)) {
	lankId := FDLinkIdToLankId(linkPosition.LinkIdentification.LinkId)
	nat.WithLank(lankId, action)
}
