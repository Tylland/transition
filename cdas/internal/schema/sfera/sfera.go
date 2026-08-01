// Package sfera provides Go types corresponding to the SFERA v2.00 XML schema.
// These types represent segment profiles used in the SFERA train control system.
package sfera

import "encoding/xml"

// SegmentProfileStatus represents the status of a segment profile.
type SegmentProfileStatus string

const (
	SegmentProfileStatusValid   SegmentProfileStatus = "Valid"
	SegmentProfileStatusInvalid SegmentProfileStatus = "Invalid"
)

// GradientDirectionType represents the gradient direction.
type GradientDirectionType string

const (
	GradientDirectionDownhill GradientDirectionType = "Downhill"
	GradientDirectionUphill   GradientDirectionType = "Uphill"
)

// NCTrainType represents the ETCS train type.
type NCTrainType string

const (
	NCTrainTypeFreightP  NCTrainType = "Freight_Train_P_Position"
	NCTrainTypeFreightG  NCTrainType = "Freight_Train_G_Position"
	NCTrainTypePassenger NCTrainType = "Passenger_Train"
)

// StartEndQualifier for area positions.
type StartEndQualifier string

const (
	StartEndQualifierStarts     StartEndQualifier = "Starts"
	StartEndQualifierEnds       StartEndQualifier = "Ends"
	StartEndQualifierStartsEnds StartEndQualifier = "StartsEnds"
	StartEndQualifierWholeSP    StartEndQualifier = "WholeSP"
)

// SpecificSSPCategory represents specific SSP categories.
type SpecificSSPCategory string

const (
	SpecificSSPCategoryOtherReplaces SpecificSSPCategory = "Other_SSP_Replaces_CDSSP"
)

// SegmentProfile represents a SFERA segment profile.
type SegmentProfile struct {
	XMLName           xml.Name              `xml:"SegmentProfile"`
	SPZone            SPZone                `xml:"SP_Zone"`
	SPID              string                `xml:"SP_ID,attr"`
	SPPoints          *SPPoints             `xml:"SP_Points,omitempty"`
	SPAreas           *SPAreas              `xml:"SP_Areas,omitempty"`
	SPCharacteristics *SPCharacteristics    `xml:"SP_Characteristics,omitempty"`
	SPContextInfo     *SPContextInformation `xml:"SP_ContextInformation,omitempty"`
	SPVersionMajor    uint                  `xml:"SP_VersionMajor,attr"`
	SPVersionMinor    uint                  `xml:"SP_VersionMinor,attr"`
	SPLength          float64               `xml:"SP_Length,attr"`
	SPName            string                `xml:"SP_Name,attr,omitempty"`
	SPStatus          SegmentProfileStatus  `xml:"SP_Status,attr,omitempty"`
}

// SPZone represents the zone information.
type SPZone struct {
	IMID *uint16 `xml:"IM_ID,attr,omitempty"`
}

// SPPoints contains point-type entities on the segment.
type SPPoints struct {
	TimingPoints                  []TimingPoint                  `xml:"TimingPoint,omitempty"`
	Signals                       []SignalComplexType            `xml:"Signal,omitempty"`
	Balises                       []Balise                       `xml:"Balise,omitempty"`
	UnprotectedLevelCrossingStops []UnprotectedLevelCrossingStop `xml:"UnprotectedLevelCrossingStop,omitempty"`
}

// SPAreas contains area-type entities on the segment.
type SPAreas struct {
	Platforms []SPGenericArea `xml:"Platform,omitempty"`
}

// SPCharacteristics contains segment characteristics.
type SPCharacteristics struct {
	StaticSpeedProfiles []StaticSpeedProfile `xml:"StaticSpeedProfile,omitempty"`
	GradientAverage     *GradientAverage     `xml:"GradientAverage,omitempty"`
}

// SPContextInformation contains contextual information.
type SPContextInformation struct {
	KilometreReferencePoints []KilometreReferencePoint `xml:"KilometreReferencePoint,omitempty"`
	LineAndTrackIdentifier   *LineAndTrackIdentifier   `xml:"LineAndTrackIdentifier,omitempty"`
}

// TimingPoint represents a timing point.
type TimingPoint struct {
	TPNames      []TPName `xml:"TP_Name,omitempty"`
	TPID         string   `xml:"TP_ID,attr"`
	Location     float64  `xml:"Location,attr"`
	StationTrack string   `xml:"StationTrack,attr,omitempty"`
}

// TPName represents a timing point name.
type TPName struct {
	Name     string `xml:"Name,attr"`
	Language string `xml:"Language,attr,omitempty"`
}

// SignalComplexType represents a signal on the segment.
type SignalComplexType struct {
	SignalID SignalID `xml:"Signal_ID"`
}

// SignalID identifies a signal.
type SignalID struct {
	SignalIDPhysical string  `xml:"Signal_ID_Physical,attr"`
	SignalIDObject   string  `xml:"Signal_ID_Object,attr,omitempty"`
	Location         float64 `xml:"Location,attr"`
}

// Balise represents a balise on the segment.
type Balise struct {
	Identifier string  `xml:"Identifier,attr,omitempty"`
	Location   float64 `xml:"Location,attr"`
}

// UnprotectedLevelCrossingStop represents an unprotected level crossing.
type UnprotectedLevelCrossingStop struct {
	Identifier string  `xml:"Identifier,attr,omitempty"`
	Location   float64 `xml:"Location,attr"`
}

// SPGenericArea represents a generic area on the segment.
type SPGenericArea struct {
	Identifier        string            `xml:"Identifier,attr,omitempty"`
	StartLocation     *float64          `xml:"StartLocation,attr,omitempty"`
	EndLocation       *float64          `xml:"EndLocation,attr,omitempty"`
	StartEndQualifier StartEndQualifier `xml:"StartEndQualifier,attr"`
}

// StaticSpeedProfile represents a static speed profile.
type StaticSpeedProfile struct {
	Start   *SSPComplexType            `xml:"StaticSpeedProfileStart,omitempty"`
	Changes []StaticSpeedProfileChange `xml:"StaticSpeedProfileChange,omitempty"`
}

// StaticSpeedProfileChange represents a speed profile change.
type StaticSpeedProfileChange struct {
	SSPComplexType
	Location float64 `xml:"Location,attr"`
}

// SSPComplexType represents SSP data.
type SSPComplexType struct {
	SpecificSSPs []SpecificSSP `xml:"SpecificSSP,omitempty"`
	SSPSpeed     *float64      `xml:"SSP_Speed,attr,omitempty"`
}

// SpecificSSP represents a specific speed profile.
type SpecificSSP struct {
	OtherSpecificSSP *OtherSpecificSSP    `xml:"OtherSpecificSSP,omitempty"`
	Category         *SpecificSSPCategory `xml:"Specific_SSP_Category,attr,omitempty"`
}

// OtherSpecificSSP represents other specific SSP data.
type OtherSpecificSSP struct {
	SSPNCDIFF         *NCTrainType `xml:"SSP_NC_DIFF,attr,omitempty"`
	TrainCategoryCode string       `xml:"TrainCategoryCode,attr,omitempty"`
}

// GradientAverage represents average gradient data.
type GradientAverage struct {
	Start   *GradientComplexType    `xml:"GradientAverageStart,omitempty"`
	Changes []GradientAverageChange `xml:"GradientAverageChange,omitempty"`
}

// GradientAverageChange represents a gradient change.
type GradientAverageChange struct {
	GradientComplexType
	Location float64 `xml:"Location,attr"`
}

// GradientComplexType represents gradient data.
type GradientComplexType struct {
	GradientValue     float64               `xml:"GradientValue,attr"`
	GradientDirection GradientDirectionType `xml:"GradientDirection,attr"`
}

// KilometreReferencePoint represents a kilometre reference point.
type KilometreReferencePoint struct {
	Identifier string  `xml:"Identifier,attr,omitempty"`
	Location   float64 `xml:"Location,attr"`
	KmRef      float64 `xml:"KmRef,attr"`
}

// LineAndTrackIdentifier represents line and track identification.
type LineAndTrackIdentifier struct {
	Start *LineAndTrackIdentifierData `xml:"LineAndTrackIdentifierStart,omitempty"`
}

// LineAndTrackIdentifierData represents line and track data.
type LineAndTrackIdentifierData struct {
	LineIdentifier  string `xml:"LineIdentifier,attr,omitempty"`
	LineName        string `xml:"LineName,attr,omitempty"`
	TrackIdentifier string `xml:"TrackIdentifier,attr,omitempty"`
}
