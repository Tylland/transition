// Package facilitydata provides Go types corresponding to the FacilityData v1.82 XML schema.
// These types are used for XML deserialization of Swedish railway facility data exports.
package facilitydata

import (
	"encoding/xml"
	"fmt"
)

// Direction represents the direction of an object on a link.
type Direction int

const (
	DirectionIn      Direction = iota // Normal direction
	DirectionAgainst                  // Against direction
	DirectionAny                      // Both directions
	DirectionUnknown                  // Unknown direction
)

// UnmarshalXML decodes a Direction from its XML string representation.
func (d *Direction) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := decoder.DecodeElement(&s, &start); err != nil {
		return err
	}
	switch s {
	case "In":
		*d = DirectionIn
	case "Against":
		*d = DirectionAgainst
	case "Any":
		*d = DirectionAny
	case "Unknown":
		*d = DirectionUnknown
	default:
		return fmt.Errorf("unknown Direction value: %q", s)
	}
	return nil
}

// DivisionTypes represents the type of a division.
type DivisionTypes int

const (
	DivisionTypesTrafficManagementSystem DivisionTypes = iota
	DivisionTypesOverSpeed
	DivisionTypesTracknumber
	DivisionTypesInfrastructureManager
	DivisionTypesInfrastructureOwner
	DivisionTypesElectrification
	DivisionTypesPassage
	DivisionTypesTrafficDirection
	DivisionTypesLineCategory
	DivisionTypesKilometerRegion
	DivisionTypesTrackPortion
	DivisionTypesNetworkStatus
	DivisionTypesTrackUpDownSingle
	DivisionTypesTrackType
	DivisionTypesTrafficControlNet
	DivisionTypesCapacityPlanningNet
	DivisionTypesOperatingStation
	DivisionTypesTrafficControlArea
)

// UnmarshalXML decodes a DivisionTypes from its XML string representation.
func (d *DivisionTypes) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := decoder.DecodeElement(&s, &start); err != nil {
		return err
	}
	switch s {
	case "TrafficManagementSystem":
		*d = DivisionTypesTrafficManagementSystem
	case "OverSpeed":
		*d = DivisionTypesOverSpeed
	case "Tracknumber":
		*d = DivisionTypesTracknumber
	case "InfrastructureManager":
		*d = DivisionTypesInfrastructureManager
	case "InfrastructureOwner":
		*d = DivisionTypesInfrastructureOwner
	case "Electrification":
		*d = DivisionTypesElectrification
	case "Passage":
		*d = DivisionTypesPassage
	case "TrafficDirection":
		*d = DivisionTypesTrafficDirection
	case "LineCategory":
		*d = DivisionTypesLineCategory
	case "KilometerRegion":
		*d = DivisionTypesKilometerRegion
	case "TrackPortion":
		*d = DivisionTypesTrackPortion
	case "NetworkStatus":
		*d = DivisionTypesNetworkStatus
	case "TrackUpDownSingle":
		*d = DivisionTypesTrackUpDownSingle
	case "TrackType":
		*d = DivisionTypesTrackType
	case "TrafficControlNet":
		*d = DivisionTypesTrafficControlNet
	case "CapacityPlanningNet":
		*d = DivisionTypesCapacityPlanningNet
	case "OperatingStation":
		*d = DivisionTypesOperatingStation
	case "TrafficControlArea":
		*d = DivisionTypesTrafficControlArea
	default:
		// Unknown division types are left at zero value
		*d = -1
	}
	return nil
}

// FacilityData is the root XML element.
type FacilityData struct {
	XMLName  xml.Name `xml:"FacilityData"`
	BaseRoot Root     `xml:"BaseRoot"`
}

// Root contains the main data sections.
type Root struct {
	Infrastructure Infrastructure `xml:"Infrastructure"`
	Facility       Facility       `xml:"Facility"`
	Divisions      []Division     `xml:"Divisions>Division"`
}

// Infrastructure contains nodes, links, and locations.
type Infrastructure struct {
	Links     []Link     `xml:"Links>Link"`
	Nodes     []Node     `xml:"Nodes>Node"`
	Locations []Location `xml:"Locations>Location"`
}

// Facility contains signaling facilities, geometries, and other facilities.
type Facility struct {
	SignalingFacilities SignalingFacilities `xml:"SignalingFacilities"`
	Geometries          Geometries          `xml:"Geometries"`
	OtherFacilities     OtherFacilities     `xml:"OtherFacilities"`
}

// SignalingFacilities contains information points and kilometer boards.
type SignalingFacilities struct {
	InformationPoints []InformationPoint `xml:"InformationPoints>InformationPoint"`
	KilometerBoards   []KilometerBoard   `xml:"KilometerBoards>KilometerBoard"`
}

// Geometries contains vertical geometries.
type Geometries struct {
	VerticalGeometries []VerticalGeometry `xml:"VerticalGeometries>VerticalGeometry"`
}

// OtherFacilities contains speed profiles.
type OtherFacilities struct {
	SpeedProfiles []SpeedProfile `xml:"SpeedProfiles>SpeedProfile"`
}

// Node represents a network node.
type Node struct {
	Identification NodeIdentification `xml:"Identification"`
}

// NodeIdentification contains the identification of a node.
type NodeIdentification struct {
	NodeId      NodeId `xml:"NodeId"`
	Designation string `xml:"Designation"`
}

// NodeId uniquely identifies a node.
type NodeId struct {
	LocationNumber int `xml:"LocationNumber"`
	NodeNumber     int `xml:"NodeNumber"`
}

// Link represents a network link.
type Link struct {
	Identification LinkIdentification `xml:"Identification"`
	Length         float64            `xml:"Length"`
}

// LinkIdentification contains the identification of a link.
type LinkIdentification struct {
	LinkId LinkId `xml:"LinkId"`
}

// LinkId uniquely identifies a link.
type LinkId struct {
	NodeFrom   NodeId `xml:"NodeFrom"`
	NodeTo     NodeId `xml:"NodeTo"`
	Linknumber int    `xml:"Linknumber"`
}

// LinkPosition represents a position or extent on a link.
type LinkPosition struct {
	LinkIdentification    LinkIdentification `xml:"LinkIdentification"`
	StartDistanceFromNode float64            `xml:"StartDistanceFromNode"`
	EndDistanceFromNode   float64            `xml:"EndDistanceFromNode"`
}

// Location represents a geographical location.
type Location struct {
	Identification Identification `xml:"Identification"`
	Type           string         `xml:"Type"`
	PLCCode        int            `xml:"PLCCode"`
	LinkPositions  []LinkPosition `xml:"LinkPositions>LinkPosition"`
}

// Division represents an infrastructure division.
type Division struct {
	DivisionType DivisionTypes `xml:"DivisionType"`
	Subdivisions []Subdivision `xml:"Subdivisions>Subdivision"`
}

// Subdivision is a part of a division.
type Subdivision struct {
	DivisionCode  string         `xml:"DivisionCode"`
	Description   string         `xml:"Description"`
	LinkPositions []LinkPosition `xml:"LinkPositions>LinkPosition"`
}

// Identification wraps business and global identifiers.
type Identification struct {
	GlobalId   GlobalId   `xml:"GlobalId"`
	BusinessId BusinessId `xml:"BusinessId"`
}

// GlobalId contains a universal identifier.
type GlobalId struct {
	UniversalId string `xml:"UniversalId"`
}

// BusinessId contains business-level identifiers.
type BusinessId struct {
	Name              string `xml:"Name"`
	LocationSignature string `xml:"LocationSignature"`
	ObjectType        string `xml:"ObjectType"`
}

// InformationPoint represents a signaling information point (signal, board, etc.).
type InformationPoint struct {
	Identification Identification `xml:"Identification"`
	LinkPositions  []LinkPosition `xml:"LinkPositions>LinkPosition"`
	Signal         *Signal        `xml:"Signal"`
}

// Signal represents a physical signal.
type Signal struct {
	TypeOfMainSignal string    `xml:"TypeOfMainSignal"`
	Direction        Direction `xml:"Direction"`
}

// KilometerBoard represents a kilometer marker.
type KilometerBoard struct {
	Identification Identification `xml:"Identification"`
	LinkPositions  []LinkPosition `xml:"LinkPositions>LinkPosition"`
	KilometerPost  int            `xml:"KilometerPost"`
}

// SpeedProfile represents a speed profile on a link.
type SpeedProfile struct {
	LinkPositions []LinkPosition `xml:"LinkPositions>LinkPosition"`
	ATrainSpeed   int            `xml:"ATrainSpeed"`
	BTrainSpeed   int            `xml:"BTrainSpeed"`
	CTrainSpeed   int            `xml:"CTrainSpeed"`
	STrainSpeed   int            `xml:"STrainSpeed"`
	Direction     Direction      `xml:"Direction"`
}

// VerticalGeometry represents a vertical geometry section (gradient).
type VerticalGeometry struct {
	LinkPositions []LinkPosition `xml:"LinkPositions>LinkPosition"`
	Gradient      float64        `xml:"Gradient"`
}
