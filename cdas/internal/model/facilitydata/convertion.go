package facilitydata

import (
	"github.com/tylland/transition/cdas"
	dm "github.com/tylland/transition/cdas/internal/core/domain/network"
	fd "github.com/tylland/transition/cdas/internal/schema/facilitydata"
)

// FacilityDataExtensions provides conversion utilities from FacilityData schema types
// to domain model types.

// CreateNodeId converts a FacilityData NodeId to a domain model NodId.
func CreateNodeId(nodeId fd.NodeId) dm.NodeID {
	return dm.NewNodeId(nodeId.LocationNumber, nodeId.NodeNumber)
}

// CreateNodeIdForNade creates a domain NodId from a FacilityData Node.
func CreateNodeIdForNade(node fd.Node) dm.NodeID {
	return CreateNodeId(node.Identification.NodeId)
}

// CreateLinkId converts a FacilityData LinkId to a domain model LinkId.
func CreateLinkId(linkId fd.LinkId) dm.LinkID {
	franNodId := CreateNodeId(linkId.NodeFrom)
	tillNodId := CreateNodeId(linkId.NodeTo)
	return dm.NewLinkId(franNodId, tillNodId, linkId.Linknumber)
}

// CreateCreateLinkId creates a domain LinkId from a FacilityData Link.
func CreateLinkIdForLink(link fd.Link) dm.LinkID {
	return CreateLinkId(link.Identification.LinkId)
}

// CreateValidityDirection converts a FacilityData Direction to domain Gilltighetsriktning.
func CreateValidityDirection(direction fd.Direction) dm.ValidityDirection {
	switch direction {
	case fd.DirectionIn:
		return dm.ValidityDirectionWith
	case fd.DirectionAgainst:
		return dm.ValidityDirectionAgainst
	case fd.DirectionAny:
		return dm.ValidityDirectionBoth
	default:
		return dm.ValidityDirectionNone
	}
}

// CreateSignalId converts a FacilityData BusinessId to a domain SignalId.
func CreateSignalId(businessId fd.BusinessId) *dm.SignalId {
	return dm.NewSignalId(businessId.LocationSignature, businessId.Name, businessId.ObjectType)
}

// CreateLinkPosition creates a Linkposition from a FacilityData LinkPosition on a given Link.
func CreateLinkPosition(link *dm.Link, position fd.LinkPosition, direction dm.ValidityDirection) *dm.LinkPosition {
	return cdas.Must(dm.NewLinkPosition(link, position.StartDistanceFromNode/link.Length, direction))
}

// CreateLinkExtent creates a Linkutbredning from a FacilityData LinkPosition on a given Link.
func CreateLinkExtent(link *dm.Link, position fd.LinkPosition, direction dm.ValidityDirection) *dm.LinkExtent {
	return &dm.LinkExtent{
		Link:      link,
		Start:     position.StartDistanceFromNode / link.Length,
		End:       position.EndDistanceFromNode / link.Length,
		Direction: direction,
	}
}

// FDWithLink executes an action on a Link identified by a FacilityData LinkPosition.
func FDWithLink(network *dm.Network, position fd.LinkPosition, action func(*dm.Link)) {
	linkId := CreateLinkId(position.LinkIdentification.LinkId)
	network.WithLink(linkId, action)
}
