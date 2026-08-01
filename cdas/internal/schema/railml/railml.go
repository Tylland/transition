// Package railml provides Go types corresponding to the RailML v0.32 XML schema.
// These types are used for XML deserialization of RailML infrastructure data.
package railml

import "encoding/xml"

// TApplicationDirection represents the application direction.
type TApplicationDirection int

const (
	TApplicationDirectionBoth TApplicationDirection = iota
	TApplicationDirectionReverse
	TApplicationDirectionNormal
)

// RailML is the root XML element.
type RailML struct {
	XMLName        xml.Name       `xml:"railML"`
	Infrastructure Infrastructure `xml:"infrastructure"`
}

// Infrastructure is the top-level infrastructure container.
type Infrastructure struct {
	Topology                 Topology                 `xml:"topology"`
	FunctionalInfrastructure FunctionalInfrastructure `xml:"functionalInfrastructure"`
}

// Topology contains net elements and relations.
type Topology struct {
	NetElements  []NetElement  `xml:"netElements>netElement"`
	NetRelations []NetRelation `xml:"netRelations>netRelation"`
}

// NetElement represents a network element (link).
type NetElement struct {
	Id       string             `xml:"id,attr"`
	Length   float64            `xml:"length,attr"`
	Relation []ElementWithIDRef `xml:"relation"`
	Name     []Name             `xml:"name"`
}

// NetRelation represents a relationship between two net elements.
type NetRelation struct {
	Id       string           `xml:"id,attr"`
	ElementA ElementWithIDRef `xml:"elementA"`
	ElementB ElementWithIDRef `xml:"elementB"`
	Name     []Name           `xml:"name"`
}

// ElementWithIDRef is a reference to another element by ID.
type ElementWithIDRef struct {
	Ref string `xml:"ref,attr"`
}

// Name represents a named property with optional description and language.
type Name struct {
	NameProperty string `xml:"name,attr"`
	Description  string `xml:"description,attr"`
	Language     string `xml:"language,attr"`
}

// FunctionalInfrastructure contains all functional infrastructure entities.
type FunctionalInfrastructure struct {
	Balises           []Balise           `xml:"balises>balise"`
	LevelCrossingsIS  []LevelCrossingIS  `xml:"levelCrossingsIS>levelCrossingIS"`
	SignalsIS         []SignalIS         `xml:"signalsIS>signalIS"`
	OperationalPoints []OperationalPoint `xml:"operationalPoints>operationalPoint"`
	Platforms         []Platform         `xml:"platforms>platform"`
	Speeds            []SpeedSection     `xml:"speeds>speed"`
}

// Balise represents a balise entity.
type Balise struct {
	Id           string            `xml:"id,attr"`
	SpotLocation []RTMSpotLocation `xml:"spotLocation"`
	Name         []Name            `xml:"name"`
}

// LevelCrossingIS represents a level crossing.
type LevelCrossingIS struct {
	Id           string                   `xml:"id,attr"`
	Name         []Name                   `xml:"name"`
	AreaLocation []RTMAreaLocation        `xml:"areaLocation"`
	Protection   *LevelCrossingProtection `xml:"protection"`
}

// LevelCrossingProtection describes the protection at a level crossing.
type LevelCrossingProtection struct {
	HasActiveProtection bool `xml:"hasActiveProtection,attr"`
}

// SignalIS represents a signal in the infrastructure.
type SignalIS struct {
	Id                   string              `xml:"id,attr"`
	Name                 []Name              `xml:"name"`
	SpotLocation         []RTMSpotLocation   `xml:"spotLocation"`
	IsAnnouncementSignal *SignalAnnouncement `xml:"isAnnouncementSignal"`
	IsInformationSignal  *SignalInformation  `xml:"isInformationSignal"`
	IsMilepost           *SignalMilepost     `xml:"isMilepost"`
}

// SignalAnnouncement represents an announcement signal type.
type SignalAnnouncement struct{}

// SignalInformation represents an information signal type.
type SignalInformation struct{}

// SignalMilepost represents a milepost signal type.
type SignalMilepost struct {
	ShownValue string `xml:"shownValue,attr"`
}

// OperationalPoint represents an operational point (station, etc.).
type OperationalPoint struct {
	Id           string            `xml:"id,attr"`
	Name         []Name            `xml:"name"`
	SpotLocation []RTMSpotLocation `xml:"spotLocation"`
}

// Platform represents a platform entity.
type Platform struct {
	Id           string            `xml:"id,attr"`
	Name         []Name            `xml:"name"`
	AreaLocation []RTMAreaLocation `xml:"areaLocation"`
	Width        []Length          `xml:"width"`
}

// Length represents a length measurement.
type Length struct {
	Value float64 `xml:"value,attr"`
}

// SpeedSection represents a speed section.
type SpeedSection struct {
	Id             string              `xml:"id,attr"`
	Name           []Name              `xml:"name"`
	LinearLocation []RTMLinearLocation `xml:"linearLocation"`
	MaxSpeed       string              `xml:"maxSpeed,attr"`
}

// RTMSpotLocation represents a spot location on the network.
type RTMSpotLocation struct {
	NetElementRef        string                `xml:"netElementRef,attr"`
	IntrinsicCoord       float64               `xml:"intrinsicCoord,attr"`
	ApplicationDirection TApplicationDirection `xml:"applicationDirection,attr"`
}

// RTMAreaLocation represents an area location on the network.
type RTMAreaLocation struct {
	AssociatedNetElement []RTMAssociatedNetElement `xml:"associatedNetElement"`
}

// RTMLinearLocation represents a linear location on the network.
type RTMLinearLocation struct {
	AssociatedNetElement []RTMAssociatedNetElement `xml:"associatedNetElement"`
	ApplicationDirection TApplicationDirection     `xml:"applicationDirection,attr"`
}

// RTMAssociatedNetElement links area/linear locations to network elements.
type RTMAssociatedNetElement struct {
	NetElementRef       string  `xml:"netElementRef,attr"`
	IntrinsicCoordBegin float64 `xml:"intrinsicCoordBegin,attr"`
	IntrinsicCoordEnd   float64 `xml:"intrinsicCoordEnd,attr"`
}
