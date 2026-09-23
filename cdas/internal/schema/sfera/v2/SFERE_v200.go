package v2

import (
	"encoding/xml"
)

// AdditionalSpeedRestriction is Additional Speed Restriction
type AdditionalSpeedRestriction struct {
	ASRFront *bool    `xml:"ASR_Front,attr"`
	ASRSpeed *float64 `xml:"ASR_Speed,attr"`
}

// AdditionalTrainNumber is Additional identifier for train number
type AdditionalTrainNumber *AdditionalTrainNumber

// AdhesionConditionsChange is Change in adhesion conditions
type AdhesionConditionsChange struct {
	NewAdhesionConditions string  `xml:"newAdhesionConditions,attr"`
	LowAdhesionCategory   *string `xml:"lowAdhesionCategory,attr"`
}

// AdvisedSpeed is Explicit speed advice
type AdvisedSpeed struct {
	Speed           *float64 `xml:"speed,attr"`
	MaxAdvisedSpeed *float64 `xml:"maxAdvisedSpeed,attr"`
	MinAdvisedSpeed *float64 `xml:"minAdvisedSpeed,attr"`
	ReasonCode      *string  `xml:"reasonCode,attr"`
	ReasonText      *string  `xml:"reasonText,attr"`
}

// AgainstBooked is Identifies the Delta delay time against the booked schedule in minutes
type AgainstBooked string

// AgainstReferenced is Delay compared to the referenced Date/Time
type AgainstReferenced string

// AllocationCompany is Name of company who is responsible for allocation and maintenance of codes
type AllocationCompany *Numeric44

// ApplicableOTN is Applicable Operational Train Number. In some cases this is different from the TrainIdentification, for example when the OTN changes during the trip.
type ApplicableOTN struct {
	XMLName xml.Name `xml:"Applicable_OTN"`
	OTN     string   `xml:"OTN,attr"`
	*TemporaryConstraintsComplexType
}

// ApplicableIfTrainRunsOnSP is The effects of this signal element are to be taken into account only if this SP is in the SP_List after the current SP.
type ApplicableIfTrainRunsOnSP struct {
	SPID   string             `xml:"SP_ID,attr"`
	SPZone *SPZoneComplexType `xml:"SP_Zone"`
}

// ATOTSContactInfo is Information about ATO-Trackside Contact Information
type ATOTSContactInfo struct {
	XMLName xml.Name  `xml:"ATOTS_ContactInfo"`
	ATOTSID *string   `xml:"ATOTS_ID,attr"`
	NIDC    *NIDCType `xml:"NID_C"`
}

// ATPSystemIdentifier is Identifier of ATP System. If it is a number: ETCS NID_
// NTC from ERA "Assignment of Values to ETCS Variables", section 3.3 (ERA_ERTMS_0
// 40001 v1.27. https://www.era.europa.eu/sites/default/files/activities/docs/ertms_040001_etcs_variables_values_en.pdf). If it is a name: from the ERA "List of Class B Systems", section 3.2 (ERA/
// TD/2011-11 v4.0, https://www.era.europa.eu/sites/default/files/activities/docs/ertms_ccs_class_b_systems_en.pdf)
type ATPSystemIdentifier struct {
	XMLName xml.Name `xml:"ATP_System_Identifier"`
	Value   string   `xml:",chardata"`
}

// ATPSystems is List of ATP systems supported in the Segment Profile. The DAS system can then find out what is supported. ATP System Consequences will be for Operational Train data containing its characteristics necessary to calculate its dynamic capabilities.
type ATPSystems struct {
	XMLName          xml.Name               `xml:"ATP_Systems"`
	ATPSystemsStart  *ATPSystemsComplexType `xml:"ATP_SystemsStart"`
	ATPSystemsChange []*ATPSystemsChange    `xml:"ATP_SystemsChange"`
}

// ATPSystemsChange is Change in ATP systems supported
type ATPSystemsChange struct {
	XMLName  xml.Name `xml:"ATP_SystemsChange"`
	Location float64  `xml:"location,attr"`
	*ATPSystemsComplexType
}

// ATPSystemsStart is ATP systems supported at the start of the Segment Profile.
type ATPSystemsStart *ATPSystemsComplexType

// AverageSpeed is Average speed of the train
type AverageSpeed *AverageSpeed

// AvoidSpeedRange is Speed range (in km/h) to avoid
type AvoidSpeedRange struct {
	RangeMinSpeed float64 `xml:"rangeMinSpeed,attr"`
	RangeMaxSpeed float64 `xml:"rangeMaxSpeed,attr"`
}

// AxleLoadSpeedProfile is [SUBSET-126] Speed Profile restriction to be applied to specific axle loads
type AxleLoadSpeedProfile struct {
	AxleLoadCategory string  `xml:"axleLoadCategory,attr"`
	NewSpeedLevel    float64 `xml:"newSpeedLevel,attr"`
	ALSPFront        *bool   `xml:"ALSP_Front,attr"`
	*SPGenericAreaComplexType
}

// B2GError is Description of error (board to ground). To be used only if the result of the message is "ERROR"
type B2GError struct {
	XMLName xml.Name `xml:"B2G_Error"`
	*ErrorComplexType
}

// B2GMessageResponse is Response to message received, board to ground
type B2GMessageResponse struct {
	XMLName  xml.Name    `xml:"B2G_MessageResponse"`
	Result   string      `xml:"result,attr"`
	B2GError []*B2GError `xml:"B2G_Error"`
}

// B2GReplyPayload is Payload of board to ground reply
type B2GReplyPayload struct {
	XMLName                xml.Name                  `xml:"B2G_ReplyPayload"`
	ReportedDASDrivingMode []*ReportedDASDrivingMode `xml:"Reported_DAS_DrivingMode"`
	B2GMessageResponse     *B2GMessageResponse       `xml:"B2G_MessageResponse"`
	PlaintextMessage       []*PlaintextMessage       `xml:"PlaintextMessage"`
	B2GStatusReport        []*B2GStatusReport        `xml:"B2G_StatusReport"`
	SessionTermination     *SessionTermination       `xml:"SessionTermination"`
}

// InterruptionPointRequest ...
type InterruptionPointRequest struct {
	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
}

// B2GRequest is Request from board to ground
type B2GRequest struct {
	XMLName                  xml.Name                    `xml:"B2G_Request"`
	SPRequest                []*SPRequest                `xml:"SP_Request"`
	TCRequest                []*TCRequest                `xml:"TC_Request"`
	JPRequest                []*JPRequest                `xml:"JP_Request"`
	CDASCAdviceRequest       []*CDASCAdviceRequest       `xml:"C_DAS_C_AdviceRequest"`
	PlaintextMessageRequest  []*PlaintextMessageRequest  `xml:"PlaintextMessageRequest"`
	ForceStatusChangeRequest []*ForceStatusChangeRequest `xml:"ForceStatusChangeRequest"`
	PositionSpeedRequest     []*PositionSpeedRequest     `xml:"PositionSpeedRequest"`
	InterruptionPointRequest []*InterruptionPointRequest `xml:"InterruptionPointRequest"`
}

// B2GStatusReport is Status Report from Board to Ground
type B2GStatusReport struct {
	XMLName                          xml.Name                    `xml:"B2G_StatusReport"`
	Identifier                       *string                     `xml:"identifier,attr"`
	DASDrivingMode                   *string                     `xml:"DAS_drivingMode,attr"`
	ATOState                         *string                     `xml:"ATO_State,attr"`
	JPSPConsistencyError             *bool                       `xml:"JP_SP_ConsistencyError,attr"`
	NextStoppingPointSkip            *bool                       `xml:"nextStoppingPointSkip,attr"`
	SlipSlideReported                *bool                       `xml:"slipSlideReported,attr"`
	OperationalConditionsFulfillment *bool                       `xml:"operationalConditionsFulfillment,attr"`
	TrainMoving                      *bool                       `xml:"trainMoving,attr"`
	UnableToStop                     *bool                       `xml:"unableToStop,attr"`
	DriverID                         *string                     `xml:"driverID,attr"`
	RoutingError                     *bool                       `xml:"routingError,attr"`
	TrainIdentification              *TrainIdentification        `xml:"TrainIdentification"`
	PositionSpeed                    *PositionSpeed              `xml:"PositionSpeed"`
	PreviousTimingPoint              *PreviousTimingPoint        `xml:"PreviousTimingPoint"`
	TimingPointEstimation            []*TimingPointEstimation    `xml:"TimingPointEstimation"`
	AdhesionConditionsChange         *AdhesionConditionsChange   `xml:"AdhesionConditionsChange"`
	ExpectedCurrentData              []*ExpectedCurrentData      `xml:"ExpectedCurrentData"`
	TrainCharacteristicsChange       *TrainCharacteristicsChange `xml:"TrainCharacteristicsChange"`
	PlaintextMessage                 *PlaintextMessage           `xml:"PlaintextMessage"`
	JPInUse                          *JPInUse                    `xml:"JP_InUse"`
	CDASCShown                       *CDASCShown                 `xml:"C_DAS_C_Shown"`
	TrainReadyStatus                 *TrainReadyStatus           `xml:"TrainReadyStatus"`
	DelayNotice                      *DelayNotice                `xml:"DelayNotice"`
	InterruptionPoint                *InterruptionPoint          `xml:"InterruptionPoint"`
	EnergyData                       *EnergyData                 `xml:"EnergyData"`
	StatusReportNSPs                 []*NSPListComplexType       `xml:"StatusReport_NSPs"`
}

// Balise is Equipment on the track at a known position that communicates with the train to exchange different information. The train can recalibrate its positioning systems according to this balise. If the Balise has a NID_C different from the one of the SP, it can be expressed here.
type Balise struct {
	BaliseGroupIdentifier  *uint16   `xml:"baliseGroupIdentifier,attr"`
	BaliseRelativePosition *uint8    `xml:"baliseRelativePosition,attr"`
	LevelCrossingGuarding  *bool     `xml:"levelCrossingGuarding,attr"`
	AmountLevelCrossings   *uint8    `xml:"amountLevelCrossings,attr"`
	NIDC                   *NIDCType `xml:"NID_C"`
	*SPGenericPointComplexType
}

// BFCNSPs is Network Specific Parameters in Braking Force Curve
type BFCNSPs *NSPListComplexType

// BFCTractionType is Traction Type for Braking Force Curve
type BFCTractionType struct {
	XMLName      xml.Name `xml:"BFC_TractionType"`
	TractionType *string  `xml:"tractionType,attr"`
}

// BrakingForceCurve is Braking Force Curve for the train. Several curves may be defined for different traction types and ATP Systems.
type BrakingForceCurve *BrakingForceCurveComplexType

// BrakingForceCurveDataPoint is Data point for the braking force curve (pair of speed and braking force)
type BrakingForceCurveDataPoint struct {
	Speed                    float64  `xml:"speed,attr"`
	BrakingForce             float64  `xml:"brakingForce,attr"`
	MaximumRegenerativeForce *float64 `xml:"maximumRegenerativeForce,attr"`
}

// CDASCAdvice is Explicit Speed Advice to On-Board unit. To be used in C-DAS-C operation.
type CDASCAdvice struct {
	XMLName                     xml.Name             `xml:"C_DAS_C_Advice"`
	Identifier                  *string              `xml:"identifier,attr"`
	AdviceType                  string               `xml:"adviceType,attr"`
	OptimalSpeed                *float64             `xml:"optimalSpeed,attr"`
	DeltaSpeed                  *float64             `xml:"deltaSpeed,attr"`
	LocationText                *string              `xml:"locationText,attr"`
	ReasonText                  *string              `xml:"reasonText,attr"`
	ReasonCode                  *string              `xml:"reasonCode,attr"`
	StartValidity               *string              `xml:"startValidity,attr"`
	EndValidity                 *string              `xml:"endValidity,attr"`
	Percentage                  *float64             `xml:"percentage,attr"`
	EarliestDepartureAdviceTime *string              `xml:"earliestDepartureAdviceTime,attr"`
	LatestDepartureAdviceTime   *string              `xml:"latestDepartureAdviceTime,attr"`
	ReferenceIDAbs              *string              `xml:"referenceIDAbs,attr"`
	StartStationShort           *string              `xml:"startStationShort,attr"`
	StartStationLong            *string              `xml:"startStationLong,attr"`
	StartKm                     *float64             `xml:"startKm,attr"`
	StartLine                   *int                 `xml:"startLine,attr"`
	StartSignal                 *string              `xml:"startSignal,attr"`
	EndStationShort             *string              `xml:"endStationShort,attr"`
	EndStationLong              *string              `xml:"endStationLong,attr"`
	EndKm                       *float64             `xml:"endKm,attr"`
	EndLine                     *int                 `xml:"endLine,attr"`
	EndSignal                   *string              `xml:"endSignal,attr"`
	EndStationShortAbs          *string              `xml:"endStationShortAbs,attr"`
	EndStationLongAbs           *string              `xml:"endStationLongAbs,attr"`
	EndKmAbs                    *float64             `xml:"endKmAbs,attr"`
	EndLineAbs                  *int                 `xml:"endLineAbs,attr"`
	EndSignalAbs                *string              `xml:"endSignalAbs,attr"`
	EndValidityAbs              *string              `xml:"endValidityAbs,attr"`
	TrainIdentification         *TrainIdentification `xml:"TrainIdentification"`
	SegmentProfileInfo          *SegmentProfileInfo  `xml:"SegmentProfileInfo"`
}

// CDASCAdviceRequest is C-DAS-C advice request
type CDASCAdviceRequest struct {
	XMLName             xml.Name             `xml:"C_DAS_C_AdviceRequest"`
	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
	CDASCShown          *CDASCShown          `xml:"C_DAS_C_Shown"`
}

// CDASCShown is Identifier of the C-DAS-C message which is currently being shown to the driver
type CDASCShown struct {
	XMLName    xml.Name `xml:"C_DAS_C_Shown"`
	Identifier *string  `xml:"identifier,attr"`
	MessageID  string   `xml:"message_ID,attr"`
}

// CancelPreviousSignal is This type of signal will be used in the case that the signal will present an aspect that is different from the aspect announced by a previous signal.
type CancelPreviousSignal struct {
	SignalIDPhysical     string   `xml:"signal_ID_Physical,attr"`
	SignalIDObject       *string  `xml:"signal_ID_Object,attr"`
	DistanceBeforeSignal *float64 `xml:"distanceBeforeSignal,attr"`
}

// CantDeficiencySSP is Cant Deficiency Static Speed Profile
type CantDeficiencySSP struct {
	SSPNCCDDIFF *string `xml:"SSP_NC_CDDIFF,attr"`
}

// ChangeReason is Reason for Journey Profile Change
type ChangeReason struct {
	Identifier     *string  `xml:"identifier,attr"`
	StartLocation  *float64 `xml:"startLocation,attr"`
	EndLocation    *float64 `xml:"endLocation,attr"`
	StartTime      *string  `xml:"startTime,attr"`
	EndTime        *string  `xml:"endTime,attr"`
	JPChangeReason *string  `xml:"JP_ChangeReason,attr"`
}

// ClaimType is Gives an indication of the status or occupation of the piece of the segment
type ClaimType interface{}

// CommunicationNetwork is Network used for communication between driver and ground. The driver needs to know it because some procedures in case of an emergency depend on the network type.
type CommunicationNetwork struct {
	CommunicationNetworkType    string  `xml:"communicationNetworkType,attr"`
	CommunicationNetworkChannel *string `xml:"communicationNetworkChannel,attr"`
	*SPGenericAreaComplexType
}

// Company is Identifies a railway company (RU or IM) (UIC RICS Code: https://uic.org/rics)
type Company *Numeric44

// ConnectionTrack is Connection track
type ConnectionTrack struct {
	ConnectionTrackType        string  `xml:"connectionTrackType,attr"`
	ConnectionTrackDescription *string `xml:"connectionTrackDescription,attr"`
	*SPGenericPointComplexType
}

// Contact is Information related to a contact
type Contact struct {
	ContactIdentifier string       `xml:"contactIdentifier,attr"`
	ContactType       *string      `xml:"contactType,attr"`
	ContactRole       *interface{} `xml:"contactRole,attr"`
	ContactNotes      *string      `xml:"contactNotes,attr"`
	MainContact       *bool        `xml:"mainContact,attr"`
}

// ContactList is List of contacts for a specific area
type ContactList struct {
	Contact []*Contact `xml:"Contact"`
	*SPGenericAreaComplexType
}

// Core is It is the main part of identifier and is determined by the company that creates it.
type Core interface{}

// CountryCodeISO is Identifies a County or State by code (ISO 3166-1)
type CountryCodeISO struct {
	*CountryIdentISO
}

// CurrentLimitation is Maximum current to be applied by the train. Expressed in A
type CurrentLimitation struct {
	CurrentLimitationStart  *CurrentLimitationComplexType `xml:"CurrentLimitationStart"`
	CurrentLimitationChange []*CurrentLimitationChange    `xml:"CurrentLimitationChange"`
}

// CurrentLimitationChange is Maximum current change
type CurrentLimitationChange struct {
	Location float64 `xml:"location,attr"`
	*CurrentLimitationComplexType
}

// CurrentLimitationStart is Maximum current at the start of the Segment Profile
type CurrentLimitationStart *CurrentLimitationComplexType

// CurveChange is Curve change
type CurveChange struct {
	Location float64 `xml:"location,attr"`
	*CurveComplexType
}

// Curves is Curves in the Segment Profile
type Curves struct {
	CurveStart  *CurveComplexType `xml:"CurveStart"`
	CurveChange []*CurveChange    `xml:"CurveChange"`
}

// CurveStart is Curve at the start of the Segment Profile
type CurveStart *CurveComplexType

// DASDrivingModeRequest is Request from Ground to Board of DAS driving mode. If the Train Identification attribute is not given, the Board will send the reply according to the train for which the communication layer has been established. In the IM-RU setup between IM DAS-TS and RU DAS-TS, if the communication layer is established for multiple trains, the train identifier must be given.
type DASDrivingModeRequest struct {
	XMLName             xml.Name             `xml:"DAS_DrivingModeRequest"`
	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
}

// DASMode is DAS Operating mode. Triplet of attributes (driving mode, architecture, connectivity)
type DASMode *DASModesComplexType

// DASModesUnavailabilityArea is Area in SP where some operating modes are unavailable
type DASModesUnavailabilityArea struct {
	XMLName                      xml.Name                      `xml:"DAS_ModesUnavailabilityArea"`
	UnavailableDASOperatingModes *UnavailableDASOperatingModes `xml:"Unavailable_DAS_OperatingModes"`
	*SPGenericAreaComplexType
}

// DASOperatingModeSelected is Operating mode selected by the DAS-TS
type DASOperatingModeSelected struct {
	XMLName         xml.Name `xml:"DAS_OperatingModeSelected"`
	DASArchitecture string   `xml:"DAS_architecture,attr"`
	DASConnectivity string   `xml:"DAS_connectivity,attr"`
}

// DASOperatingModesSupported is Operating modes supported by the DAS-OB
type DASOperatingModesSupported *DASModesComplexType

// DateTimeMeasurement is Date and time of the measurement
type DateTimeMeasurement string

// DecisiveGradient is Decisive Gradient Area. The decisive gradient is the steepest gradient in front of potential stopping points on a line like signals, level crossings and stations. It is relevant, if it goes downhill, for the calculation of either the minimum braking percentage or the speed limit for a train because in case of reduced braking capacity, the train might not be able to brake properly, therefore, the driver has to adjust the speed. If it goes uphill, it is relevant to calculate the towing load, for example, to make sure, with the available engine power, that the locomotive can get away after it stopped.
type DecisiveGradient struct {
	GradientValue              float64 `xml:"gradientValue,attr"`
	GradientDirection          string  `xml:"gradientDirection,attr"`
	DirectionOfApplicationOnSP string  `xml:"directionOfApplicationOnSP,attr"`
	*SPGenericAreaComplexType
}

// Delay is Delay of the train
type Delay struct {
	Delay         *string `xml:"Delay,attr"`
	DelayCategory *string `xml:"DelayCategory,attr"`
}

// DelayCause is This element identifies the reason for a delay (modified DelayReason)
type DelayCause string

// DelayNotice is Inform the IM/RU Backoffice of the reason for an additional delay encountered by a train. For cases where the train is interrupted, or at the points where the train ready message is due (origin station and changes in train characteristics or RU).
type DelayNotice struct {
	DelayCause        *string        `xml:"DelayCause"`
	PreciseDelayCause *string        `xml:"PreciseDelayCause"`
	TrainLocation     *LocationIdent `xml:"TrainLocation"`
}

// DetailedDescriptionOfLocation ...
type DetailedDescriptionOfLocation string

// DynamicBrakeForceInhibitionArea is [SUBSET-126] Dynamic Brake Force Inhibition Area
type DynamicBrakeForceInhibitionArea *SPGenericAreaComplexType

// EarliestEndDateTime is Forecasted earliest time for end of interruption
type EarliestEndDateTime string

// EndSignalApplication is Location where the signal stops to be applied. If not specified, the signal is applied indefinitely, until superseded by another signal. This may also be in a different SP: then the SP_ID has to be specified (otherwise the current SP is considered).
type EndSignalApplication struct {
	SPID                                *string  `xml:"SP_ID,attr"`
	Location                            *float64 `xml:"location,attr"`
	EndSignalApplicationMustBeClearedBy *string  `xml:"endSignalApplicationMustBeClearedBy,attr"`
}

// EnergyData is Real-time energy data
type EnergyData struct {
	MeasuredVoltage         *int     `xml:"measuredVoltage,attr"`
	MeasuredCurrent         *float64 `xml:"measuredCurrent,attr"`
	Autonomy                *float64 `xml:"autonomy,attr"`
	MaxPower                *float64 `xml:"maxPower,attr"`
	SetValueTractionBraking *string  `xml:"setValueTractionBraking,attr"`
}

// EstimatedVoltage is Estimated voltage of the traction system (V)
type EstimatedVoltage struct {
	EstimatedVoltageStart  *VoltageComplexType       `xml:"EstimatedVoltageStart"`
	EstimatedVoltageChange []*EstimatedVoltageChange `xml:"EstimatedVoltageChange"`
}

// EstimatedVoltageChange is Estimated voltage change
type EstimatedVoltageChange struct {
	Location float64 `xml:"location,attr"`
	*VoltageComplexType
}

// EstimatedVoltageConstraint is Temporary constraints on estimated voltage
type EstimatedVoltageConstraint struct {
	VoltageValue *int `xml:"voltageValue,attr"`
}

// EstimatedVoltageStart is Estimated voltage at the start of the Segment Profile
type EstimatedVoltageStart *VoltageComplexType

// ExpectedCurrentData is Energy data expected in the next timeframe (5 to 15 minutes)
type ExpectedCurrentData struct {
	StartTime                   *string `xml:"startTime,attr"`
	EndTime                     *string `xml:"endTime,attr"`
	ExpectedConsumptionCurrent  float64 `xml:"expectedConsumptionCurrent,attr"`
	ExpectedRegenerativeCurrent float64 `xml:"expectedRegenerativeCurrent,attr"`
}

// ForceStatusChange is Force On-Board DAS to change its Status (mainly used to make a DAS-OB readonly)
type ForceStatusChange struct {
	Identifier          *string              `xml:"identifier,attr"`
	DASDrivingMode      string               `xml:"DAS_drivingMode,attr"`
	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
}

// ForceStatusChangeRequest is Request for ForceStatusChange messages
type ForceStatusChangeRequest struct {
	TrainIdentification          *TrainIdentification          `xml:"TrainIdentification"`
	LastForceStatusChangeMessage *LastForceStatusChangeMessage `xml:"LastForceStatusChangeMessage"`
}

// G2BError is Description of error (ground to board). To be used only if the result of the message is "ERROR"
type G2BError struct {
	XMLName            xml.Name `xml:"G2B_Error"`
	DataFirstAvailable *string  `xml:"dataFirstAvailable,attr"`
	*ErrorComplexType
}

// G2BEventPayload is Payload of ground to board event
type G2BEventPayload struct {
	XMLName                 xml.Name                 `xml:"G2B_EventPayload"`
	JourneyProfile          []*JourneyProfile        `xml:"JourneyProfile"`
	PlaintextMessage        []*PlaintextMessage      `xml:"PlaintextMessage"`
	CDASCAdvice             []*CDASCAdvice           `xml:"C_DAS_C_Advice"`
	ForceStatusChange       []*ForceStatusChange     `xml:"ForceStatusChange"`
	RelatedTrainInformation *RelatedTrainInformation `xml:"RelatedTrainInformation"`
	InterruptionPoint       *InterruptionPoint       `xml:"InterruptionPoint"`
}

// G2BMessageResponse is Response to message received, ground to board
type G2BMessageResponse struct {
	XMLName  xml.Name    `xml:"G2B_MessageResponse"`
	Result   string      `xml:"result,attr"`
	G2BError []*G2BError `xml:"G2B_Error"`
}

// G2BReplyPayload is Payload of ground to board reply
type G2BReplyPayload struct {
	XMLName              xml.Name                `xml:"G2B_ReplyPayload"`
	JourneyProfile       []*JourneyProfile       `xml:"JourneyProfile"`
	SegmentProfile       []*SegmentProfile       `xml:"SegmentProfile"`
	TrainCharacteristics []*TrainCharacteristics `xml:"TrainCharacteristics"`
	G2BMessageResponse   *G2BMessageResponse     `xml:"G2B_MessageResponse"`
	PlaintextMessage     []*PlaintextMessage     `xml:"PlaintextMessage"`
	PositionSpeed        []*PositionSpeed        `xml:"PositionSpeed"`
	CDASCAdvice          []*CDASCAdvice          `xml:"C_DAS_C_Advice"`
	ForceStatusChange    []*ForceStatusChange    `xml:"ForceStatusChange"`
}

// G2BRequest is Request from ground to board
type G2BRequest struct {
	XMLName                   xml.Name                     `xml:"G2B_Request"`
	DASDrivingModeRequest     []*DASDrivingModeRequest     `xml:"DAS_DrivingModeRequest"`
	SessionTerminationRequest []*SessionTerminationRequest `xml:"SessionTerminationRequest"`
	PlaintextMessageRequest   []*PlaintextMessageRequest   `xml:"PlaintextMessageRequest"`
	StatusReportRequest       []*StatusReportRequest       `xml:"StatusReportRequest"`
}

// GeneralJPInformation is Data relevant to the whole JP
type GeneralJPInformation struct {
	XMLName                 xml.Name              `xml:"General_JP_Information"`
	GeneralJPInformationNSP []*NSPListComplexType `xml:"General_JP_Information_NSP"`
}

// GeneralJPInformationNSP is Network-Specific Parameter for General JP Information
type GeneralJPInformationNSP *NSPListComplexType

// GNSSDynamicPosition is GNSS Dynamic Position
type GNSSDynamicPosition *GNSSDynamicPositionComplexType

// GradientAverage is Average gradient in the Segment Profile
type GradientAverage struct {
	GradientAverageStart  *GradientComplexType     `xml:"GradientAverageStart"`
	GradientAverageChange []*GradientAverageChange `xml:"GradientAverageChange"`
}

// GradientAverageChange is Change in average gradient
type GradientAverageChange struct {
	Location float64 `xml:"location,attr"`
	*GradientComplexType
}

// GradientAverageStart is Average gradient at the start of the Segment Profile
type GradientAverageStart *GradientComplexType

// GradientSteepest is Steepest gradient in the Segment Profile
type GradientSteepest struct {
	GradientSteepestStart  *GradientComplexType      `xml:"GradientSteepestStart"`
	GradientSteepestChange []*GradientSteepestChange `xml:"GradientSteepestChange"`
}

// GradientSteepestChange is Change in steepest gradient
type GradientSteepestChange struct {
	Location float64 `xml:"location,attr"`
	*GradientComplexType
}

// GradientSteepestStart is Steepest gradient at the start of the Segment Profile
type GradientSteepestStart *GradientComplexType

// HandshakeAcknowledgement is Handshake between board and ground for SFERA connection or disconnection
type HandshakeAcknowledgement struct {
	TimeoutATOTSResponse               *string                   `xml:"timeout_ATOTS_Response,attr"`
	ReportingTime                      *string                   `xml:"reportingTime,attr"`
	AdditionalInfo                     *string                   `xml:"additionalInfo,attr"`
	MinimumPositionUpdateRequestPeriod *string                   `xml:"minimumPositionUpdateRequestPeriod,attr"`
	ATOTSID                            *string                   `xml:"ATOTS_ID,attr"`
	RelatedTrainSupported              *string                   `xml:"relatedTrainSupported,attr"`
	DASOperatingModeSelected           *DASOperatingModeSelected `xml:"DAS_OperatingModeSelected"`
	SPZone                             *SPZoneComplexType        `xml:"SP_Zone"`
}

// HandshakeReject is Rejection of Handshake from Ground to Board
type HandshakeReject struct {
	ATOTSID               *string            `xml:"ATOTS_ID,attr"`
	HandshakeRejectReason *string            `xml:"handshakeRejectReason,attr"`
	SPZone                *SPZoneComplexType `xml:"SP_Zone"`
}

// HandshakeRequest is Handshake request from board to ground
type HandshakeRequest struct {
	StatusReportsEnabled       *bool                  `xml:"statusReportsEnabled,attr"`
	AdditionalInfo             *string                `xml:"additionalInfo,attr"`
	RelatedTrainRequest        *string                `xml:"relatedTrainRequest,attr"`
	DASOperatingModesSupported []*DASModesComplexType `xml:"DAS_OperatingModesSupported"`
}

// IMID is ID of the Infrastructure Manager. It is a UIC RICS Code.
type IMID *Numeric44

// InternalReferenceIdentifier is The link to the System Reference of the origin system
type InternalReferenceIdentifier string

// Interruption ...
type Interruption struct {
	InterruptionReason          *DelayCode  `xml:"InterruptionReason"`
	InterruptionDateTime        string      `xml:"InterruptionDateTime"`
	InterruptionDescription     []*FreeText `xml:"InterruptionDescription"`
	InternalReferenceIdentifier *string     `xml:"InternalReferenceIdentifier"`
}

// InterruptionDateTime is Date and Time when the Train was interrupted
type InterruptionDateTime string

// InterruptionDescription is The free text description of an interruption
type InterruptionDescription string

// InterruptionDuration is To specifiy the probable duration of the interruption
type InterruptionDuration struct {
	EarliestEndDateTime *string `xml:"EarliestEndDateTime"`
	LatestEndDateTime   *string `xml:"LatestEndDateTime"`
}

// InterruptionLocation is Location of the interruption
type InterruptionLocation *LocationIdent

// InterruptionPoint is To provide any additional information to the RU or next IM (e.g. contact person, next steps, etc)
type InterruptionPoint struct {
	InterruptionLocation          *LocationIdent        `xml:"InterruptionLocation"`
	DetailedDescriptionOfLocation *string               `xml:"DetailedDescriptionOfLocation"`
	Interruption                  *Interruption         `xml:"Interruption"`
	InterruptionDuration          *InterruptionDuration `xml:"InterruptionDuration"`
	Remarks                       []*FreeText           `xml:"Remarks"`
}

// InterruptionReason is This element identifies the reason for an interruption of the train running, with a code according to UIC450-2
type InterruptionReason string

// JourneyProfile is Journey Profile for a specific trip
type JourneyProfile struct {
	JPStatus             *string               `xml:"JP_Status,attr"`
	JPVersion            *string               `xml:"JP_Version,attr"`
	TrainIdentification  *TrainIdentification  `xml:"TrainIdentification"`
	SegmentProfileList   []*SegmentProfileList `xml:"SegmentProfileList"`
	GeneralJPInformation *GeneralJPInformation `xml:"General_JP_Information"`
}

// JPContextInformation is Context information related to Journey Profile to be displayed to drivers
type JPContextInformation struct {
	XMLName                  xml.Name              `xml:"JP_ContextInformation"`
	ApplicableOTN            *ApplicableOTN        `xml:"Applicable_OTN"`
	JPContextInformationNSPs []*NSPListComplexType `xml:"JP_ContextInformation_NSPs"`
}

// JPInUse is Journey Profile currently in use
type JPInUse struct {
	XMLName   xml.Name `xml:"JP_InUse"`
	JPVersion string   `xml:"JP_Version,attr"`
}

// JPRequest is Journey Profile request.
type JPRequest struct {
	XMLName                   xml.Name                                 `xml:"JP_Request"`
	TrainIdentification       *TrainIdentification                     `xml:"TrainIdentification"`
	RequestFromSegmentProfile *SegmentProfileIdentificationComplexType `xml:"RequestFromSegmentProfile"`
	JPInUse                   *JPInUse                                 `xml:"JP_InUse"`
}

// KeeperShortNameVKM is Free text, short name/vehicle keeper marking of the wagon keeper. Use ERA VKM Register: https://www.era.europa.eu/registers/vkm_en
type KeeperShortNameVKM interface{}

// KilometreReferencePoint is Reference Point on the side of the track indicating the passing of a certain distance, usually relative to a "zero" point of the line.
type KilometreReferencePoint struct {
	KmRef          float64 `xml:"kmRef,attr"`
	Ascending      *string `xml:"ascending,attr"`
	AdditionalInfo *string `xml:"additionalInfo,attr"`
	*SPGenericPointComplexType
}

// LastAbsoluteOdometryPoint is Position of the last absolute point for the odometry. In ETCS odometry, this is the last known group of balises.
type LastAbsoluteOdometryPoint struct {
	Location *float64 `xml:"location,attr"`
	*SegmentProfileIdentificationComplexType
}

// LastForceStatusChangeMessage is identifier of the last ForceStatusChange request message sent by the board side of the communication. A ForceStatusChange Reply will only be sent if a ForceStatusChange order has been generated since the last request was sent. Otherwise, a simple "OK" message response is sent.
type LastForceStatusChangeMessage struct {
	Identifier string `xml:"identifier,attr"`
}

// LastPassedKmRef is Last Km reference point passed by the train
type LastPassedKmRef struct {
	KmRef float64 `xml:"kmRef,attr"`
}

// LastPassedSignal is Last signal passed by the train
type LastPassedSignal struct {
	SignalIDPhysical string  `xml:"signal_ID_Physical,attr"`
	SignalIDObject   *string `xml:"signal_ID_Object,attr"`
	AtSignal         *bool   `xml:"AtSignal,attr"`
}

// LastPlaintextMessage is identifier of the last PlainTextMessage message received by the board side of the communication. One or more PlainTextMessage replies will only be sent if more recent advice is available. Otherwise, a simple "OK" message response is sent.
type LastPlaintextMessage struct {
	Identifier string `xml:"identifier,attr"`
}

// LastStatusReportMessage is identifier of the last StatusReport message received by the ground side of the communication. One or more StatusReport replies will only be sent if more recent data is available. Otherwise, a simple "OK" message response is sent.
type LastStatusReportMessage struct {
	Identifier string `xml:"identifier,attr"`
}

// LatestEndDateTime is Forecasted latest time for end of interruption
type LatestEndDateTime string

// LeavingEnteringStatus is Gives an indication if a train is entering of leaving your trainpath at a certain location
type LeavingEnteringStatus interface{}

// LevelCrossingArea is Level crossing area, for driver information
type LevelCrossingArea struct {
	LevelCrossingType        *string `xml:"levelCrossingType,attr"`
	ProtectedLevelCrossing   *bool   `xml:"protectedLevelCrossing,attr"`
	LevelCrossingDescription *string `xml:"levelCrossingDescription,attr"`
	*SPGenericAreaComplexType
}

// LimitedDynamicBrakeForceArea is [SUBSET-126] Dynamic Brake Limitation Area
type LimitedDynamicBrakeForceArea struct {
	DynamicBrakeForceLimit uint16 `xml:"dynamicBrakeForceLimit,attr"`
	*SPGenericAreaComplexType
}

// LineAndTrackIdentifier is Identification of line and track
type LineAndTrackIdentifier struct {
	LineAndTrackIdentifierStart  *LineAndTrackIdentifierComplexType `xml:"LineAndTrackIdentifierStart"`
	LineAndTrackIdentifierChange []*LineAndTrackIdentifierChange    `xml:"LineAndTrackIdentifierChange"`
}

// LineAndTrackIdentifierChange is Change of identification of line and track in the course of the Segment Profile
type LineAndTrackIdentifierChange struct {
	Location float64 `xml:"location,attr"`
	*LineAndTrackIdentifierComplexType
}

// LineAndTrackIdentifierStart is Identification of line and track at the start of the Segment Profile
type LineAndTrackIdentifierStart *LineAndTrackIdentifierComplexType

// LocationPrimaryCode ...
type LocationPrimaryCode int

// LocationSubsidiaryCode is this element identifies a location as a part of primary location e.g. a junction, a signal, a passing loop etc., It is unique when used in combination with a “LocationPrimaryCode”
type LocationSubsidiaryCode struct {
	LocationSubsidiaryTypeCode string `xml:"LocationSubsidiaryTypeCode,attr"`
	*String110
}

// LocationSubsidiaryIdentification is Code, Name and allocation company of Subsidiary Location
type LocationSubsidiaryIdentification struct {
	LocationSubsidiaryCode *LocationSubsidiaryCode `xml:"LocationSubsidiaryCode"`
	AllocationCompany      *Numeric44              `xml:"AllocationCompany"`
	LocationSubsidiaryName *FreeText               `xml:"LocationSubsidiaryName"`
}

// LocationSubsidiaryName is To be completed in an official language of the Country using the ISO Unicode alphabet
type LocationSubsidiaryName string

// LowAdhesion is Low Adhesion Status Report
type LowAdhesion struct {
	LowAdhesionCategory *string `xml:"lowAdhesionCategory,attr"`
}

// MaxSpeed is Maximum speed (in km/h)
type MaxSpeed struct {
	Speed float64 `xml:"speed,attr"`
}

// MaxSpeedAnnouncement is Announcement for a downline maximum speed (in km/h). Note that if the distance to the downline MaxSpeed signal is known, it is recommended to use MaxSpeed together with StartSignalApplication.
type MaxSpeedAnnouncement struct {
	Speed float64 `xml:"speed,attr"`
}

// MaxSpeedTrack is Maximum speed of the train possible at this track
type MaxSpeedTrack string

// MessageHeader is Header of the IRS90940 message. Contains metadata related to the message
type MessageHeader struct {
	SFERAVersionMajor     string               `xml:"SFERA_version_major,attr"`
	SFERAVersionMinor     string               `xml:"SFERA_version_minor,attr"`
	MessageID             string               `xml:"message_ID,attr"`
	Timestamp             string               `xml:"timestamp,attr"`
	SourceDevice          string               `xml:"sourceDevice,attr"`
	DestinationDevice     *string              `xml:"destinationDevice,attr"`
	CorrelationID         *string              `xml:"correlation_ID,attr"`
	NIDOPERATIONAL        *string              `xml:"NID_OPERATIONAL,attr"`
	NIDENGINE             *string              `xml:"NID_ENGINE,attr"`
	EuropeanVehicleNumber *string              `xml:"EuropeanVehicleNumber,attr"`
	TrainIdentification   *TrainIdentification `xml:"TrainIdentification"`
	MessageReference      *MessageReference    `xml:"MessageReference"`
	Sender                *Sender              `xml:"Sender"`
	Recipient             *Recipient           `xml:"Recipient"`
}

// MessageIdentifier is Identification of the Message
type MessageIdentifier string

// MessageReference is Generated by the common Interface
type MessageReference struct {
	MessageType        interface{} `xml:"MessageType"`
	MessageTypeVersion interface{} `xml:"MessageTypeVersion"`
	MessageIdentifier  string      `xml:"MessageIdentifier"`
	MessageDateTime    string      `xml:"MessageDateTime"`
}

// MessageType is To indicate the message type transmitted or referred to. The following list was agreed within the sector:
// 1000        ConsignmentOrderMessage
//
// 2001        PathCanceledMessage
// 2002        PathConfirmedMessage
// 2003        PathDetailsMessage
// 2004        PathDetailsRefusedMessage
// 2005        PathNotAvailableMessage
// 2006        PathRequestMessage
// 2007        ReceiptConfirmationMessage
//
// -- sector messages (Planning) --
// 2500      PathCoordinationMessage
// 2501      PathSectionNotificationMessage
// -- sector message end --
//
// 3003        TrainCompositionMessage
// 3006        TrainReadyMessage
// 4001        TrainDelayCauseMessage
// 4004        TrainRunningForecastMessage
// 4005        TrainRunningInformationMessage
// 4006        TrainRunningInterruptionMessage
//
// -- sector message (Operations) --
// 4500      PassengerTrainCompositionProcessMessage
// 4501      RollingStockRestrictionMessage
// 4504      ChangeOfTrackMessage
// 4505      TrainJourneyModificationMessage
// -- sector message  end --
//
// 5001        AlertMessage
// 5002        WagonArrivalNoticeMessage
// 5003        WagonDeliveryNoticeMessage
// 5004        WagonDepartureNoticeMessage
// 5006        WagonETI_ETA_Message
// 5007        WagonExceptionMessage
// 5008        WagonExceptionReasonMessage
//
// -- sector message (Wagon interchange) --
// 5009        WagonInterchangeNoticeMessage
// 5012        WagonReceivedAtInterchangeMessage
// 5013        WagonRefusedAtInterchangeMessage
// -- sector message  end --
//
// 5014        WagonReleaseNoticeMessage
// 5015        WagonYardArrivalMessage
// 5016        WagonYardDepartureMessage
//
// 6002        LocationFileDatasetMessage
// 6003        RollingStockDatasetMessage
//
// --sector (RU-RU) ---
// 5500        WagonPerformanceMessage
// -- sector end---
//
// 6004        RollingStockDatasetQueryMessage
//
// -- Sector Messages TCR begin --
// 6500        TCRMessage
// 6501TCRResponseMessage
// 6502TCRCanceledMessage
// -- Sector Messages TCR end --
//
// -- sector (TrainID) begin --
// 8500        UpateLinkMessage
// 8501        ObjectInfoMessage
// -- sector end --
//
// 9000        ErrorMessage
type MessageType interface{}

// MessageTypeVersion is Version of the Message Type
type MessageTypeVersion interface{}

// NetworkSpecificArea is Network-Specific Area
type NetworkSpecificArea struct {
	Name                     *string                     `xml:"name,attr"`
	NetworkSpecificParameter []*NetworkSpecificParameter `xml:"NetworkSpecificParameter"`
	*SPGenericAreaComplexType
}

// NetworkSpecificCharacteristics is Network-Specific Characteristics
type NetworkSpecificCharacteristics struct {
	Name        string       `xml:"name,attr"`
	Description *interface{} `xml:"description,attr"`
	NSPStart    *NSPStart    `xml:"NSP_Start"`
	NSPChange   []*NSPChange `xml:"NSP_Change"`
}

// NetworkSpecificConstraint is Network-Specific Temporary Constraint
type NetworkSpecificConstraint *NSPListComplexType

// NetworkSpecificParameter is Generic Network-Specific Parameter
type NetworkSpecificParameter struct {
	Name        string       `xml:"name,attr"`
	Value       string       `xml:"value,attr"`
	Description *interface{} `xml:"description,attr"`
}

// NetworkSpecificPoint is Network-Specific Point
type NetworkSpecificPoint struct {
	Name                     *string                     `xml:"name,attr"`
	NetworkSpecificParameter []*NetworkSpecificParameter `xml:"NetworkSpecificParameter"`
	*SPGenericPointComplexType
}

// NextSP is Data for the next Segment Profile if it is in a different "Zone"
type NextSP struct {
	SPID                       *string            `xml:"SP_ID,attr"`
	DirectionOfApplicationOnSP *string            `xml:"directionOfApplicationOnSP,attr"`
	SPZone                     *SPZoneComplexType `xml:"SP_Zone"`
	ATOTSContactInfo           *ATOTSContactInfo  `xml:"ATOTS_ContactInfo"`
}

// NextTimingPointTrainInformation is Next timing points for the train
type NextTimingPointTrainInformation struct {
	TPID                string             `xml:"TP_ID,attr"`
	TPAbbreviation      *string            `xml:"TP_abbreviation,attr"`
	TPStopSkipPass      *string            `xml:"TP_StopSkipPass,attr"`
	ForecastedArrival   *string            `xml:"forecastedArrival,attr"`
	ForecastedDeparture *string            `xml:"forecastedDeparture,attr"`
	SPZone              *SPZoneComplexType `xml:"SP_Zone"`
}

// NIDC is ERTMS: Identity number of the country or region. From SUBSET-026 chap. 7, 7.5.1.86
type NIDC int16

// NoStoppingArea is It is an area shown to drivers where they should not stop the train. It is typically the case in some of the tunnels, or some of the bridges, or some portions of them.
type NoStoppingArea *SPGenericAreaComplexType

// NSPChange is Change in the Network-Specific Characteristics
type NSPChange struct {
	XMLName                  xml.Name                    `xml:"NSP_Change"`
	Location                 float64                     `xml:"location,attr"`
	NetworkSpecificParameter []*NetworkSpecificParameter `xml:"NetworkSpecificParameter"`
}

// NSPStart is Start of the Network-Specific Characteristics
type NSPStart struct {
	XMLName                  xml.Name                    `xml:"NSP_Start"`
	NetworkSpecificParameter []*NetworkSpecificParameter `xml:"NetworkSpecificParameter"`
}

// ObjectType is Provides a possibility for differentiantion between the objects: Train, Path, Case Reference and Path Request. In this case, must be TR (Train)
type ObjectType interface{}

// OtherSpecificSSP is Other Specific Static Speed Profile
type OtherSpecificSSP struct {
	SSPNCDIFF         *string `xml:"SSP_NC_DIFF,attr"`
	TrainCategoryCode *string `xml:"trainCategoryCode,attr"`
}

// OTNID is Operational Train Number (OTN) Identifier
type OTNID *OTNIDComplexType

// OwnTrain is Contains the identification of the own train
type OwnTrain struct {
	TrainIdentification      *TrainIdentification      `xml:"TrainIdentification"`
	TrainLocationInformation *TrainLocationInformation `xml:"TrainLocationInformation"`
}

// PermittedBrakingDistanceArea is Permitted Braking Distance Area
type PermittedBrakingDistanceArea struct {
	PermittedBrakingDistance *float64 `xml:"permittedBrakingDistance,attr"`
	PBDBrake                 *string  `xml:"PBD_Brake,attr"`
	GradientPBD              *float64 `xml:"gradient_PBD,attr"`
	GradientDirectionPBD     *string  `xml:"gradientDirection_PBD,attr"`
	*SPGenericAreaComplexType
}

// PlaintextMessage is Plaintext message
type PlaintextMessage struct {
	Identifier          *string              `xml:"identifier,attr"`
	MessageString       string               `xml:"messageString,attr"`
	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
}

// PlaintextMessageRequest is Plaintext message request
type PlaintextMessageRequest struct {
	TrainIdentification  *TrainIdentification  `xml:"TrainIdentification"`
	LastPlaintextMessage *LastPlaintextMessage `xml:"LastPlaintextMessage"`
}

// Platform is Area alongside a railway track providing convenient access to trains
type Platform *SPGenericAreaComplexType

// PositionSpeed is Position and speed of the train
type PositionSpeed struct {
	Location                    float64                         `xml:"location,attr"`
	Speed                       *float64                        `xml:"speed,attr"`
	LocationPositiveErrorMargin *float64                        `xml:"location_PositiveErrorMargin,attr"`
	LocationNegativeErrorMargin *interface{}                    `xml:"location_NegativeErrorMargin,attr"`
	TrainIdentification         *TrainIdentification            `xml:"TrainIdentification"`
	GNSSDynamicPosition         *GNSSDynamicPositionComplexType `xml:"GNSS_DynamicPosition"`
	LastAbsoluteOdometryPoint   *LastAbsoluteOdometryPoint      `xml:"LastAbsoluteOdometryPoint"`
	*SegmentProfileIdentificationComplexType
}

// PositionSpeedRequest is Request from Board to Ground of position and speed. If the Train Identification attribute is not given, the Ground will send the reply according to the train for which the communication layer has been established. In the IM-RU setup between IM DAS-TS and RU DAS-TS, if the communication layer is established for multiple trains, the train identifier must be given.
type PositionSpeedRequest struct {
	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
}

// PowerAdvice is Advice on minimising or maximising the power consumption
type PowerAdvice struct {
	PowerConsumptionChange string `xml:"powerConsumptionChange,attr"`
}

// PowerConstraint is Power Constraint signal
type PowerConstraint struct {
	PowerConstraintInstruction string `xml:"powerConstraintInstruction,attr"`
}

// PreciseDelayCause is Precise reason for a delay (string)
type PreciseDelayCause string

// PreviousTimingPoint is [SUBSET-126] Information concerning the most recent Timing Point considered as either reached or passed
type PreviousTimingPoint struct {
	TPPassStopDepart   *string `xml:"TP_Pass_Stop_Depart,attr"`
	TPAccurateStopping *string `xml:"TP_Accurate_Stopping,attr"`
	*TimingPointIdentificationComplexType
}

// PreviousTimingPointTrainInformation is Previous timing points passed by the train
type PreviousTimingPointTrainInformation struct {
	TPID           string             `xml:"TP_ID,attr"`
	TPAbbreviation *string            `xml:"TP_abbreviation,attr"`
	TPStopSkipPass *string            `xml:"TP_StopSkipPass,attr"`
	ArrivalTime    *string            `xml:"arrivalTime,attr"`
	DepartureTime  *string            `xml:"departureTime,attr"`
	SPZone         *SPZoneComplexType `xml:"SP_Zone"`
}

// PrimaryLocationName is Location Name in an offication language of the Country using the ISO Unicode alphabet
type PrimaryLocationName string

// RatedVoltage is Rated voltage of the traction system (V)
type RatedVoltage struct {
	RatedVoltageStart  *VoltageComplexType   `xml:"RatedVoltageStart"`
	RatedVoltageChange []*RatedVoltageChange `xml:"RatedVoltageChange"`
}

// RatedVoltageChange is Rated voltage change
type RatedVoltageChange struct {
	Location float64 `xml:"location,attr"`
	*VoltageComplexType
}

// RatedVoltageStart is Rated voltage at the start of the Segment Profile
type RatedVoltageStart *VoltageComplexType

// Recipient is Receiver of the message
type Recipient struct {
	CIInstanceNumber *int `xml:"CI_InstanceNumber,attr"`
	*CompanyCode
}

// RegenerationTotalCurrent is Expected maximal current for regeneration that can be returned to the overhead contact line
type RegenerationTotalCurrent *TotalCurrentComplexType

// RelatedTrainInformation is Contains information about the own train and the related trains, i.e. trains that share the same trainpath (infrastructure) as the own train.
type RelatedTrainInformation struct {
	RestrictionReason *string          `xml:"RestrictionReason,attr"`
	OwnTrain          *OwnTrain        `xml:"OwnTrain"`
	TrainsBeforeOwn   *TrainsBeforeOwn `xml:"TrainsBeforeOwn"`
	TrainsAfterOwn    *TrainsAfterOwn  `xml:"TrainsAfterOwn"`
}

// Remarks is Free Form Text
type Remarks string

// ReportedDASDrivingMode is Report of DAS Driving Mode
type ReportedDASDrivingMode struct {
	XMLName             xml.Name             `xml:"Reported_DAS_DrivingMode"`
	DASDrivingMode      string               `xml:"DAS_drivingMode,attr"`
	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
}

// RequestFromSegmentProfile is Segment Profile from which the JP is requested
type RequestFromSegmentProfile *SegmentProfileIdentificationComplexType

// RollingStockApplicability is Criteria for applicability related to rolling stock
type RollingStockApplicability struct {
	RollingStockType      string `xml:"rollingStockType,attr"`
	MinimumNumberOfRSType *uint8 `xml:"minimumNumberOfRSType,attr"`
	MaximumNumberOfRSType *uint8 `xml:"maximumNumberOfRSType,attr"`
}

// SameDirection is Only for trains in front of the own train, gives an indication if the train is going in the same direction as you (true) or the opposite direction (false).
type SameDirection bool

// SegmentClaim is Provides information about the occupation of a Segment Profile (or part of it) by the train.
type SegmentClaim struct {
	SPVersionMajor          uint32                   `xml:"SP_VersionMajor,attr"`
	SPVersionMinor          uint32                   `xml:"SP_VersionMinor,attr"`
	SPDirection             *string                  `xml:"SP_Direction,attr"`
	StartLocation           *float64                 `xml:"startLocation,attr"`
	EndLocation             *float64                 `xml:"endLocation,attr"`
	StartEndQualifier       *string                  `xml:"startEndQualifier,attr"`
	SegmentOccupationStatus *SegmentOccupationStatus `xml:"SegmentOccupationStatus"`
	*SegmentProfileIdentificationComplexType
}

// SegmentOccupationStatus is Status of the train in the segment
type SegmentOccupationStatus struct {
	ClaimType             interface{}  `xml:"ClaimType"`
	LeavingEnteringStatus *interface{} `xml:"LeavingEnteringStatus"`
	SameDirection         *bool        `xml:"SameDirection"`
}

// SegmentProfile is Profile for a section of railway track
type SegmentProfile struct {
	SPVersionMajor        uint32                `xml:"SP_VersionMajor,attr"`
	SPVersionMinor        uint32                `xml:"SP_VersionMinor,attr"`
	SPLength              float64               `xml:"SP_Length,attr"`
	SPName                *string               `xml:"SP_Name,attr"`
	SPStatus              SPStatus              `xml:"SP_Status,attr"`
	DistanceEoAOffset     *float64              `xml:"distance_EoA_Offset,attr"`
	UTCOffset             *uint8                `xml:"UTC_Offset,attr"`
	SPAltitude            *float64              `xml:"SP_Altitude,attr"`
	SPGauge               *int                  `xml:"SP_Gauge,attr"`
	SPMaxLateralDeviation *float64              `xml:"SP_MaxLateralDeviation,attr"`
	SPPoints              *SPPoints             `xml:"SP_Points"`
	SPAreas               *SPAreas              `xml:"SP_Areas"`
	SPCharacteristics     *SPCharacteristics    `xml:"SP_Characteristics"`
	SPContextInformation  *SPContextInformation `xml:"SP_ContextInformation"`
	NextSP                *NextSP               `xml:"NextSP"`
	*SegmentProfileIdentificationComplexType
}

// SegmentProfileInfo is Information about the Segment Profile
type SegmentProfileInfo struct {
	StartLocation *float64 `xml:"startLocation,attr"`
	EndLocation   *float64 `xml:"endLocation,attr"`
	NextSP        *NextSP  `xml:"NextSP"`
	*SegmentProfileIdentificationComplexType
}

// SegmentProfileList is Reference to a Segment Profile traversed by the journey
type SegmentProfileList struct {
	SPVersionMajor             uint32                        `xml:"SP_VersionMajor,attr"`
	SPVersionMinor             uint32                        `xml:"SP_VersionMinor,attr"`
	SPDirection                *string                       `xml:"SP_Direction,attr"`
	TimingPointConstraints     []*TimingPointConstraints     `xml:"TimingPointConstraints"`
	TemporaryConstraints       []*TemporaryConstraints       `xml:"TemporaryConstraints"`
	TemporarySignalConstraints []*TemporarySignalConstraints `xml:"TemporarySignalConstraints"`
	TrainCharacteristicsRef    []*TrainCharacteristicsRef    `xml:"TrainCharacteristicsRef"`
	ChangeReason               *ChangeReason                 `xml:"ChangeReason"`
	JPContextInformation       *JPContextInformation         `xml:"JP_ContextInformation"`
	*SegmentProfileIdentificationComplexType
}

// Sender is The sender of the message
type Sender struct {
	CIInstanceNumber *int `xml:"CI_InstanceNumber,attr"`
	*CompanyCode
}

// SessionTermination is Session termination
type SessionTermination struct {
	SessionTerminationReason *string `xml:"sessionTerminationReason,attr"`
}

// SessionTerminationRequest is Request of Session Termination
type SessionTerminationRequest struct {
	AdditionalInfo      *string              `xml:"additionalInfo,attr"`
	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
}

// SFERAB2GEventMessage is IRS90940 Board to Ground Event Message. In standard SFERA implementations via the Common Interface, the value "SF00" will be used for the MessageType.
type SFERAB2GEventMessage struct {
	XMLName            xml.Name            `xml:"SFERA_B2G_EventMessage"`
	MessageHeader      *MessageHeader      `xml:"MessageHeader"`
	B2GStatusReport    []*B2GStatusReport  `xml:"B2G_StatusReport"`
	SessionTermination *SessionTermination `xml:"SessionTermination"`
}

// SFERAB2GReplyMessage is IRS90940 Board to Ground Reply Message. In standard SFERA implementations via the Common Interface, the value "SF02" will be used for the MessageType.
type SFERAB2GReplyMessage struct {
	XMLName         xml.Name         `xml:"SFERA_B2G_ReplyMessage"`
	MessageHeader   *MessageHeader   `xml:"MessageHeader"`
	B2GReplyPayload *B2GReplyPayload `xml:"B2G_ReplyPayload"`
}

// SFERAB2GRequestMessage is IRS90940 Board to Ground Request Message. In standard SFERA implementations via the Common Interface, the value "SF01" will be used for the MessageType.
type SFERAB2GRequestMessage struct {
	XMLName          xml.Name          `xml:"SFERA_B2G_RequestMessage"`
	MessageHeader    *MessageHeader    `xml:"MessageHeader"`
	HandshakeRequest *HandshakeRequest `xml:"HandshakeRequest"`
	B2GRequest       *B2GRequest       `xml:"B2G_Request"`
}

// SFERAG2BEventMessage is IRS90940 Ground to Board Event Message. In standard SFERA implementations via the Common Interface, the value "SF10" will be used for the MessageType.
type SFERAG2BEventMessage struct {
	XMLName         xml.Name         `xml:"SFERA_G2B_EventMessage"`
	MessageHeader   *MessageHeader   `xml:"MessageHeader"`
	G2BEventPayload *G2BEventPayload `xml:"G2B_EventPayload"`
}

// SFERAG2BReplyMessage is IRS90940 Ground to Board Reply Message. In standard SFERA implementations via the Common Interface, the value "SF12" will be used for the MessageType.
type SFERAG2BReplyMessage struct {
	XMLName                  xml.Name                  `xml:"SFERA_G2B_ReplyMessage"`
	MessageHeader            *MessageHeader            `xml:"MessageHeader"`
	HandshakeAcknowledgement *HandshakeAcknowledgement `xml:"HandshakeAcknowledgement"`
	HandshakeReject          *HandshakeReject          `xml:"HandshakeReject"`
	G2BReplyPayload          *G2BReplyPayload          `xml:"G2B_ReplyPayload"`
}

// SFERAG2BRequestMessage is IRS90940 Ground to Board Request Message. In standard SFERA implementations via the Common Interface, the value "SF11" will be used for the MessageType.
type SFERAG2BRequestMessage struct {
	XMLName       xml.Name       `xml:"SFERA_G2B_RequestMessage"`
	MessageHeader *MessageHeader `xml:"MessageHeader"`
	G2BRequest    *G2BRequest    `xml:"G2B_Request"`
}

// Signal is Description of the impact on allowed speed, driving reactions of the driver, and traction type induced by permanent signals. The impact described in the SP is the one given by the least restrictive Aspect that can be seen by the driver passing on the specific SP. If the TMS anticipates that the train will meet the signal at an aspect different from the least restrictive aspect, or the signal is cancelled, use TemporarySignalConstraints under JP / SPList (if the signal can be modified or cancelled, the identifier is mandatory).
type Signal *SignalComplexType

// SignalATPSystems is List of ATP systems to which the signal is applied (default: all ATP systems supported in the SP)
type SignalATPSystems *ATPSystemsComplexType

// SignalID is Signal identification. The identification may include the identifier of the physical signal of the track and the identifier of the single signal object that contains the specific information (and possibly application). For a single physical signal there may be several signal objects. The pair (signal_ID_Physical, signal_ID_Object) must be unique within an SP. the signal_ID_Object must be documented if 2 or more signal objects are related to the signal_ID_Physical.
type SignalID struct {
	XMLName          xml.Name `xml:"Signal_ID"`
	SignalIDPhysical string   `xml:"signal_ID_Physical,attr"`
	SignalIDObject   *string  `xml:"signal_ID_Object,attr"`
	Location         float64  `xml:"location,attr"`
}

// SignalNSPs is Network Specific Parameters in Signal
type SignalNSPs *NSPListComplexType

// SignalApplication is Constraints for the application of the signal
type SignalApplication struct {
	StartSignalApplication    *StartSignalApplication      `xml:"StartSignalApplication"`
	EndSignalApplication      *EndSignalApplication        `xml:"EndSignalApplication"`
	SignalATPSystems          *ATPSystemsComplexType       `xml:"Signal_ATP_Systems"`
	SignalTrainCategory       []*SignalTrainCategory       `xml:"SignalTrainCategory"`
	SignalTrainLength         *SignalTrainLength           `xml:"SignalTrainLength"`
	ApplicableIfTrainRunsOnSP []*ApplicableIfTrainRunsOnSP `xml:"ApplicableIfTrainRunsOnSP"`
}

// SignalFunction is Function of the signal. The list of values is taken from EULynx SignalFunctionTypes (https://dataprep.eulynx.eu/2021-05/index.htm?goto=3:1:1:4:3762), with the addition of laneChange.
type SignalFunction interface{}

// SignalInformation is Information contained by the signal that impacts the driving profile
type SignalInformation struct {
	StopBeforePassingSignal *StopBeforePassingSignal              `xml:"StopBeforePassingSignal"`
	MaxSpeed                *MaxSpeed                             `xml:"MaxSpeed"`
	MaxSpeedAnnouncement    *MaxSpeedAnnouncement                 `xml:"MaxSpeedAnnouncement"`
	AvoidSpeedRange         *AvoidSpeedRange                      `xml:"AvoidSpeedRange"`
	TractionTypeTransition  *TractionTypeTransition               `xml:"TractionTypeTransition"`
	PowerConstraint         *PowerConstraint                      `xml:"PowerConstraint"`
	StationAheadSign        *TimingPointIdentificationComplexType `xml:"StationAheadSign"`
	CancelPreviousSignal    *CancelPreviousSignal                 `xml:"CancelPreviousSignal"`
	SignalInformationNSPs   *NSPListComplexType                   `xml:"SignalInformation_NSPs"`
}

// SignalInformationNSPs is Network-Specific Parameters in SignalInformation
type SignalInformationNSPs *NSPListComplexType

// SignalPhysicalCharacteristics is Physical Characteristics of the Signal
type SignalPhysicalCharacteristics struct {
	SignalType                        *string                              `xml:"signalType,attr"`
	HorizontalOffset                  *float64                             `xml:"horizontalOffset,attr"`
	VerticalOffset                    *float64                             `xml:"verticalOffset,attr"`
	SignalPhysicalCharacteristicsNSPs []*SignalPhysicalCharacteristicsNSPs `xml:"SignalPhysicalCharacteristics_NSPs"`
}

// SignalPhysicalCharacteristicsNSPs is Network Specific Parameters in Signal Physical Characteristics
type SignalPhysicalCharacteristicsNSPs struct {
	XMLName                  xml.Name                    `xml:"SignalPhysicalCharacteristics_NSPs"`
	Name                     *string                     `xml:"name,attr"`
	NetworkSpecificParameter []*NetworkSpecificParameter `xml:"NetworkSpecificParameter"`
}

// SignalCantDeficiencyTrainCategory is Cant Deficiency train category to which the signal is applied
type SignalCantDeficiencyTrainCategory struct {
	XMLName   xml.Name `xml:"Signal_CantDeficiencyTrainCategory"`
	NCCDTrain string   `xml:"NC_CDTrain,attr"`
}

// SignalOtherTrainCategory is Other train category to which the signal is applied
type SignalOtherTrainCategory struct {
	XMLName           xml.Name `xml:"Signal_OtherTrainCategory"`
	NCTrain           *string  `xml:"NC_Train,attr"`
	TrainCategoryCode *string  `xml:"trainCategoryCode,attr"`
}

// SignalTrainCategory is Train Categories to which the signal is applied
type SignalTrainCategory struct {
	SignalCantDeficiencyTrainCategory *SignalCantDeficiencyTrainCategory `xml:"Signal_CantDeficiencyTrainCategory"`
	SignalOtherTrainCategory          *SignalOtherTrainCategory          `xml:"Signal_OtherTrainCategory"`
}

// SignalTrainLength is Minimum and maximum train length to which the signal is applied.
type SignalTrainLength struct {
	MinimumTrainLength *float64 `xml:"minimumTrainLength,attr"`
	MaximumTrainLength *float64 `xml:"maximumTrainLength,attr"`
}

// SPAreas is Notable areas present in the Segment Profile (e.g. stations, platforms, tunnels...)
type SPAreas struct {
	XMLName                                xml.Name                        `xml:"SP_Areas"`
	Platform                               []*SPGenericAreaComplexType     `xml:"Platform"`
	Tunnel                                 []*Tunnel                       `xml:"Tunnel"`
	DecisiveGradient                       []*DecisiveGradient             `xml:"DecisiveGradient"`
	NoStoppingArea                         []*SPGenericAreaComplexType     `xml:"NoStoppingArea"`
	AxleLoadSpeedProfile                   []*AxleLoadSpeedProfile         `xml:"AxleLoadSpeedProfile"`
	PermittedBrakingDistanceArea           []*PermittedBrakingDistanceArea `xml:"PermittedBrakingDistanceArea"`
	SwitchOffRegenerativeBrakeArea         []*SPGenericAreaComplexType     `xml:"SwitchOffRegenerativeBrakeArea"`
	SwitchOffEddyCurrentBrakeArea          []*SPGenericAreaComplexType     `xml:"SwitchOffEddyCurrentBrakeArea"`
	SwitchOffEddyCurrentEmergencyBrakeArea []*SPGenericAreaComplexType     `xml:"SwitchOffEddyCurrentEmergencyBrakeArea"`
	SwitchOffMagneticShoeBrakeArea         []*SPGenericAreaComplexType     `xml:"SwitchOffMagneticShoeBrakeArea"`
	DASModesUnavailabilityArea             []*DASModesUnavailabilityArea   `xml:"DAS_ModesUnavailabilityArea"`
	DynamicBrakeForceInhibitionArea        []*SPGenericAreaComplexType     `xml:"DynamicBrakeForceInhibitionArea"`
	LimitedDynamicBrakeForceArea           []*LimitedDynamicBrakeForceArea `xml:"LimitedDynamicBrakeForceArea"`
	TAPTAFLocation                         []*TAPTAFLocation               `xml:"TAP_TAF_Location"`
	NetworkSpecificArea                    []*NetworkSpecificArea          `xml:"NetworkSpecificArea"`
}

// SPCharacteristics is Characteristics of the Segment Profile (e.g. gradient, curve, cant, voltage...)
type SPCharacteristics struct {
	XMLName                        xml.Name                          `xml:"SP_Characteristics"`
	StaticSpeedProfile             []*StaticSpeedProfile             `xml:"StaticSpeedProfile"`
	GradientAverage                *GradientAverage                  `xml:"GradientAverage"`
	GradientSteepest               *GradientSteepest                 `xml:"GradientSteepest"`
	Curves                         *Curves                           `xml:"Curves"`
	SuperElevation                 *SuperElevation                   `xml:"SuperElevation"`
	RatedVoltage                   *RatedVoltage                     `xml:"RatedVoltage"`
	EstimatedVoltage               *EstimatedVoltage                 `xml:"EstimatedVoltage"`
	CurrentLimitation              *CurrentLimitation                `xml:"CurrentLimitation"`
	ATPSystems                     *ATPSystems                       `xml:"ATP_Systems"`
	NetworkSpecificCharacteristics []*NetworkSpecificCharacteristics `xml:"NetworkSpecificCharacteristics"`
}

// SPContextInformation is Context information related to Journey Profile to be displayed to drivers
type SPContextInformation struct {
	XMLName                        xml.Name                          `xml:"SP_ContextInformation"`
	LevelCrossingArea              []*LevelCrossingArea              `xml:"LevelCrossingArea"`
	ConnectionTrack                []*ConnectionTrack                `xml:"ConnectionTrack"`
	ContactList                    []*ContactList                    `xml:"ContactList"`
	CommunicationNetwork           []*CommunicationNetwork           `xml:"CommunicationNetwork"`
	KilometreReferencePoint        []*KilometreReferencePoint        `xml:"KilometreReferencePoint"`
	LineAndTrackIdentifier         *LineAndTrackIdentifier           `xml:"LineAndTrackIdentifier"`
	NetworkSpecificPoint           []*NetworkSpecificPoint           `xml:"NetworkSpecificPoint"`
	NetworkSpecificArea            []*NetworkSpecificArea            `xml:"NetworkSpecificArea"`
	NetworkSpecificCharacteristics []*NetworkSpecificCharacteristics `xml:"NetworkSpecificCharacteristics"`
}

// SPPoints is Notable points present in the Segment Profile (e.g. timing points, signals, balises)
type SPPoints struct {
	XMLName                      xml.Name                        `xml:"SP_Points"`
	TimingPoint                  []*TimingPoint                  `xml:"TimingPoint"`
	Signal                       []*SignalComplexType            `xml:"Signal"`
	Balise                       []*Balise                       `xml:"Balise"`
	VirtualBalise                []*VirtualBalise                `xml:"VirtualBalise"`
	UnprotectedLevelCrossingStop []*UnprotectedLevelCrossingStop `xml:"UnprotectedLevelCrossingStop"`
	NetworkSpecificPoint         []*NetworkSpecificPoint         `xml:"NetworkSpecificPoint"`
}

// SPRequest is Segment Profile request. If version is not included, latest version is requested.
type SPRequest struct {
	XMLName        xml.Name `xml:"SP_Request"`
	SPVersionMajor *uint32  `xml:"SP_VersionMajor,attr"`
	SPVersionMinor *uint32  `xml:"SP_VersionMinor,attr"`
	*SegmentProfileIdentificationComplexType
}

// SPZone is Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be provided, especially if conversion to SUBSET-126 is required.
type SPZone *SPZoneComplexType

// SpecificSSP is Specific Static Speed Profile
type SpecificSSP *SpecificSSPComplexType

// StartDate is Is only used in the operational phase and refers to the date where the single train will start the train journey
type StartDate string

// StartSignalApplication is Location where the signal starts to be applied. If not specified, the signal is applied starting from the signal location. This may also be in a different SP: then the SP_ID has to be specified (otherwise the current SP is considered).
type StartSignalApplication struct {
	SPID                                  *string  `xml:"SP_ID,attr"`
	Location                              *float64 `xml:"location,attr"`
	StartSignalApplicationMustBeClearedBy *string  `xml:"startSignalApplicationMustBeClearedBy,attr"`
}

// StaticSpeedProfile is Static Speed Profile. It is possible to make different SSPs for different ATP systems if needed.
type StaticSpeedProfile struct {
	StaticSpeedProfileStart  *SSPComplexType             `xml:"StaticSpeedProfileStart"`
	StaticSpeedProfileChange []*StaticSpeedProfileChange `xml:"StaticSpeedProfileChange"`
	ATPSystemIdentifier      []*ATPSystemIdentifier      `xml:"ATP_System_Identifier"`
}

// StaticSpeedProfileChange is Change of Static Speed Profile
type StaticSpeedProfileChange struct {
	Location float64 `xml:"location,attr"`
	*SSPComplexType
}

// StaticSpeedProfileStart is Start of the Static Speed Profile
type StaticSpeedProfileStart *SSPComplexType

// StationAheadSign is Sign on the side of the track at stopping distance from  the next station, from which point the driver will start to brake to be at the permitted speed in the station. The station is identified by the Timing Point ID.
type StationAheadSign *TimingPointIdentificationComplexType

// StatusReportNSPs is Network Specific Parameters in Status Report
type StatusReportNSPs *NSPListComplexType

// StatusReportRequest is Request for Status Report messages
type StatusReportRequest struct {
	TrainIdentification     *TrainIdentification     `xml:"TrainIdentification"`
	LastStatusReportMessage *LastStatusReportMessage `xml:"LastStatusReportMessage"`
}

// StopBeforePassingSignal is Indicates if the train must stop before crossing the signal, for example if the driver must receive an order from the signalling responsible.
type StopBeforePassingSignal struct {
	StopTime *string `xml:"stopTime,attr"`
	Value    bool    `xml:",chardata"`
}

// StoppingPointApplicability is Criteria for the applicability of a stopping point
type StoppingPointApplicability struct {
	MinimumNumberOfCars                  *int16                                     `xml:"minimumNumberOfCars,attr"`
	MaximumNumberOfCars                  *int16                                     `xml:"maximumNumberOfCars,attr"`
	MinimumTrainLength                   *float64                                   `xml:"minimumTrainLength,attr"`
	MaximumTrainLength                   *float64                                   `xml:"maximumTrainLength,attr"`
	RollingStockApplicability            []*RollingStockApplicability               `xml:"RollingStockApplicability"`
	StoppingPointApplicableIfPrecedingSP []*SegmentProfileIdentificationComplexType `xml:"StoppingPointApplicableIfPrecedingSP"`
}

// StoppingPointApplicableIfPrecedingSP is The stopping point is applicable if this is a preceding SP
type StoppingPointApplicableIfPrecedingSP *SegmentProfileIdentificationComplexType

// StoppingPointInformation is Stopping Point Information
type StoppingPointInformation struct {
	DepartureTime                string      `xml:"departureTime,attr"`
	EarliestDepartureTime        *string     `xml:"earliestDepartureTime,attr"`
	PlannedDepartureTime         *string     `xml:"plannedDepartureTime,attr"`
	TrainHold                    *bool       `xml:"trainHold,attr"`
	MinimumDwellTime             *string     `xml:"minimumDwellTime,attr"`
	LeaveASAP                    *bool       `xml:"leaveASAP,attr"`
	DepartureAuthorisationType   *string     `xml:"departureAuthorisationType,attr"`
	ChangedStoppingLocationAlert *string     `xml:"changedStoppingLocationAlert,attr"`
	OpeningDoorSide              *string     `xml:"openingDoorSide,attr"`
	CentralisedOpening           *bool       `xml:"centralisedOpening,attr"`
	AutomaticClosing             *bool       `xml:"automaticClosing,attr"`
	DaylightSaving               *bool       `xml:"daylightSaving,attr"`
	RelaxedCoupler               *bool       `xml:"relaxedCoupler,attr"`
	StopType                     []*StopType `xml:"StopType"`
}

// StoppingPointLocation is If the Timing Point is a Stopping Point, point of the track where the head of trains respecting the specified stoppingPointApplicability parameters will stop.
type StoppingPointLocation struct {
	Identifier                 *string                     `xml:"identifier,attr"`
	StopLocationTolerance      *float64                    `xml:"stopLocationTolerance,attr"`
	DistanceSTPReached         *float64                    `xml:"distance_STP_Reached,attr"`
	StoppingZone               *StoppingZoneComplexType    `xml:"StoppingZone"`
	StoppingPointApplicability *StoppingPointApplicability `xml:"StoppingPointApplicability"`
}

// StoppingZone is Stopping zone, common to all alternative TimingPoints for a track/platform. it could be applied in or out of stations (e.g. storage tracks). It is possible to reference a Stopping Zone from Timing Point Constraints.
type StoppingZone *StoppingZoneComplexType

// StoppingZoneReference is Reference to a StoppingZone in the TimingPoint of the Segment Profile
type StoppingZoneReference *StoppingZoneComplexType

// StopType is This field can be used to inform the driver of the reason of a stop. It cannot be used with passing points. The driver may use this for passenger information and to inform the dispatcher on deviations. There will be no other effect on the DAS behavior apart from the information display.
type StopType struct {
	StopTypePurpose string  `xml:"stopTypePurpose,attr"`
	StopTypeDetails *string `xml:"stopTypeDetails,attr"`
	PlannedStop     *bool   `xml:"plannedStop,attr"`
	MandatoryStop   *bool   `xml:"mandatoryStop,attr"`
}

// SuperElevation is Superelevation values in the Segment Profile (Difference in rail height between the inside rail and the outside rail in a curve, in m)
type SuperElevation struct {
	SuperElevationStart  *SuperElevationStart    `xml:"SuperElevationStart"`
	SuperElevationChange []*SuperElevationChange `xml:"SuperElevationChange"`
}

// SuperElevationChange is Superelevation change
type SuperElevationChange struct {
	Location            float64 `xml:"location,attr"`
	SuperElevationValue float64 `xml:"superElevationValue,attr"`
}

// SuperElevationStart is Superelevation at the start of the Segment Profile
type SuperElevationStart struct {
	SuperElevationValue float64 `xml:"superElevationValue,attr"`
}

// SwitchOffEddyCurrentBrakeArea is [SUBSET-126] Switch Off Eddy Current Brake for Service Brake Area
type SwitchOffEddyCurrentBrakeArea *SPGenericAreaComplexType

// SwitchOffEddyCurrentEmergencyBrakeArea is [SUBSET-126] Switch Off Eddy Current Brake for Emergency Brake Area
type SwitchOffEddyCurrentEmergencyBrakeArea *SPGenericAreaComplexType

// SwitchOffMagneticShoeBrakeArea is [SUBSET-126] Switch Off Magnetic Shoe Brake Area
type SwitchOffMagneticShoeBrakeArea *SPGenericAreaComplexType

// SwitchOffRegenerativeBrakeArea is [SUBSET-126] Switch Off Regenerative Brake Area
type SwitchOffRegenerativeBrakeArea *SPGenericAreaComplexType

// TAPTAFID is TAP/TAF Identifier
type TAPTAFID *TAPTAFIDComplexType

// TAPTAFLocation ...
type TAPTAFLocation struct {
	XMLName             xml.Name       `xml:"TAP_TAF_Location"`
	TAPTAFLocationIdent *LocationIdent `xml:"TAP_TAF_LocationIdent"`
	*SPGenericAreaComplexType
}

// TAPTAFLocationIdent ...
type TAPTAFLocationIdent *LocationIdent

// TCChangeFeatures is Features for Train Characteristics Change
type TCChangeFeatures struct {
	XMLName                 xml.Name                   `xml:"TC_ChangeFeatures"`
	RemainingEnginePower    *float64                   `xml:"remainingEnginePower,attr"`
	UnavailableTractionType []*UnavailableTractionType `xml:"UnavailableTractionType"`
	*TCFeaturesComplexType
}

// TCFeatures is List of train's characteristics, necessary to calculate its dynamic capabilities: length, weight, max speed etc.
type TCFeatures *TCFeaturesComplexType

// TCFeaturesNSPs is Network Specific Parameters in TC Features
type TCFeaturesNSPs *NSPListComplexType

// TCRequest is Train Characteristics request. If version is not included, latest version is requested.
type TCRequest struct {
	XMLName        xml.Name `xml:"TC_Request"`
	TCVersionMajor *uint32  `xml:"TC_VersionMajor,attr"`
	TCVersionMinor *uint32  `xml:"TC_VersionMinor,attr"`
	*TCIdentifierComplexType
}

// TCRUID is Train Characteristics - Identifier of the RU operating the train (use UIC RICS Code: https://uic.org/rics)
type TCRUID *Numeric44

// TemporaryConstraintReason is Context information explaining background of temporary constraint
type TemporaryConstraintReason struct {
	Language                     *string      `xml:"language,attr"`
	MessageString                string       `xml:"messageString,attr"`
	OriginalLanguage             *bool        `xml:"originalLanguage,attr"`
	TemporaryConstraintReference *interface{} `xml:"temporaryConstraintReference,attr"`
}

// TemporaryConstraints is Temporary Constraints in the Segment Profile
type TemporaryConstraints struct {
	TemporaryConstraintType      string                        `xml:"temporaryConstraintType,attr"`
	AdditionalSpeedRestriction   *AdditionalSpeedRestriction   `xml:"AdditionalSpeedRestriction"`
	LowAdhesion                  *LowAdhesion                  `xml:"LowAdhesion"`
	TractionTotalCurrent         *TotalCurrentComplexType      `xml:"TractionTotalCurrent"`
	RegenerationTotalCurrent     *TotalCurrentComplexType      `xml:"RegenerationTotalCurrent"`
	PowerAdvice                  *PowerAdvice                  `xml:"PowerAdvice"`
	EstimatedVoltageConstraint   *EstimatedVoltageConstraint   `xml:"EstimatedVoltageConstraint"`
	Wind                         *Wind                         `xml:"Wind"`
	UnavailableDASOperatingModes *UnavailableDASOperatingModes `xml:"Unavailable_DAS_OperatingModes"`
	AdvisedSpeed                 *AdvisedSpeed                 `xml:"AdvisedSpeed"`
	NetworkSpecificConstraint    *NSPListComplexType           `xml:"NetworkSpecificConstraint"`
	TemporaryConstraintReason    []*TemporaryConstraintReason  `xml:"TemporaryConstraintReason"`
	*TemporaryConstraintsComplexType
}

// TemporarySignalConstraints is Description of the impact on allowed speed, driving reactions of the driver, and traction type induced by temporary signals, permanent signals with an aspect different from the least restrictive aspect described in SP / SP_Points / Signal, or temporarily cancelled signals on the ground.
type TemporarySignalConstraints struct {
	StartTime                     *string `xml:"startTime,attr"`
	EndTime                       *string `xml:"endTime,attr"`
	TemporarySignalConstraintType string  `xml:"temporarySignalConstraintType,attr"`
	*SignalComplexType
}

// TFCNSPs is Network Specific Parameters in Traction Force Curve
type TFCNSPs *NSPListComplexType

// TimetableYear is Refers to the timetable period in which the business will be carried out (Not used in SFERA. Only kept for consistency with data elements in TAP/TAF)
type TimetableYear interface{}

// TimingPoint is A location and stopping accuracy defined in the Segment Profile for which a type (Stopping or Passing Point) and specific time may be identified in the Journey Profile. This time may be an arrival time, a departure time, or in the case of a train not scheduled to stop at that location, the passing time.
type TimingPoint struct {
	TPID                  string                 `xml:"TP_ID,attr"`
	Location              float64                `xml:"location,attr"`
	TPAbbreviation        *string                `xml:"TP_abbreviation,attr"`
	TPType                *string                `xml:"TP_type,attr"`
	TPPriority            *uint8                 `xml:"TP_priority,attr"`
	StationTrack          *string                `xml:"stationTrack,attr"`
	TPName                []*TPName              `xml:"TP_Name"`
	StoppingPointLocation *StoppingPointLocation `xml:"StoppingPointLocation"`
	TimingPointNSPs       []*TimingPointNSPs     `xml:"TimingPoint_NSPs"`
}

// TimingPointNSPs is Network Specific Parameters in Timing Point
type TimingPointNSPs struct {
	XMLName xml.Name `xml:"TimingPoint_NSPs"`
	*NSPListComplexType
}

// TimingPointConstraints is Constraints at specific Timing Points in the Segment Profile (arrival and departure times)
type TimingPointConstraints struct {
	TPLatestArrivalTime        *string                   `xml:"TP_latestArrivalTime,attr"`
	TPPlannedLatestArrivalTime *string                   `xml:"TP_PlannedLatestArrivalTime,attr"`
	ArrivalWindow              *string                   `xml:"arrivalWindow,attr"`
	AimEarlyArrival            *bool                     `xml:"aimEarlyArrival,attr"`
	TPAlignment                *string                   `xml:"TP_alignment,attr"`
	TPStopSkipPass             *string                   `xml:"TP_StopSkipPass,attr"`
	TPInformation              *string                   `xml:"TP_Information,attr"`
	DaylightSaving             *bool                     `xml:"daylightSaving,attr"`
	TimingPointReference       *TimingPointReference     `xml:"TimingPointReference"`
	StoppingPointInformation   *StoppingPointInformation `xml:"StoppingPointInformation"`
	TimingPointConstraintsNSPs []*NSPListComplexType     `xml:"TimingPointConstraints_NSPs"`
}

// TimingPointConstraintsNSPs is Network Specific Parameters in Timing Point Constraints
type TimingPointConstraintsNSPs *NSPListComplexType

// TimingPointEstimation is [SUBSET-126] Estimated arrival time to, at least, the next Timing Point
type TimingPointEstimation struct {
	TPID              string             `xml:"TP_ID,attr"`
	ForecastedArrival string             `xml:"forecastedArrival,attr"`
	SPZone            *SPZoneComplexType `xml:"SP_Zone"`
}

// TimingPointReference is Reference to a Timing Point in the Segment Profile. It could be either the ID of a specific TimingPoint or a Stopping Zone. In the latter case, it is up to the DAS-OB to choose the exact Timing Point.
type TimingPointReference struct {
	TPIDReference         *TimingPointIdentificationComplexType `xml:"TP_ID_Reference"`
	StoppingZoneReference *StoppingZoneComplexType              `xml:"StoppingZoneReference"`
}

// TotalLoadWeight is The total weight of the transportation unit on the freight wagon. This is the booked or actual weight of goods including packing and carrier's equipment
type TotalLoadWeight int

// TPIDReference is Reference to the ID of a TimingPoint in the Segment Profile
type TPIDReference *TimingPointIdentificationComplexType

// TPName is Name of the Timing Point
type TPName struct {
	XMLName  xml.Name `xml:"TP_Name"`
	Name     string   `xml:"name,attr"`
	Language *string  `xml:"language,attr"`
}

// TractionForceCurve is Traction Force Curve for the train. Several curves may be defined for different traction types.
type TractionForceCurve *TractionForceCurveComplexType

// TractionForceCurveDataPoint is Data point for the traction force curve (pair of speed and traction force)
type TractionForceCurveDataPoint struct {
	Speed         float64 `xml:"speed,attr"`
	TractionForce float64 `xml:"tractionForce,attr"`
}

// TractionTotalCurrent is Temporary Constraints on Total current for traction and HVAC consumption
type TractionTotalCurrent *TotalCurrentComplexType

// TractionTypeTransition is Traction type transition signal
type TractionTypeTransition struct {
	TractionTypeTransitionElectric *TractionTypeTransitionComplexType `xml:"TractionTypeTransitionElectric"`
	TractionTypeTransitionBimodal  *TractionTypeTransitionComplexType `xml:"TractionTypeTransitionBimodal"`
}

// TractionTypeTransitionBimodal is Traction type transition signal for bimodal trains
type TractionTypeTransitionBimodal *TractionTypeTransitionComplexType

// TractionTypeTransitionElectric is Traction type transition signal for electric trains
type TractionTypeTransitionElectric *TractionTypeTransitionComplexType

// Train is Contains the identification of a related train and the segments
type Train struct {
	*TrainType
}

// TrainATPSystems is ATP system supported by the train. The DAS system can then find out what is supported. ATP System Consequences will be for Operational Train data containing its characteristics necessary to calculate its dynamic capabilities. Use the ERA "List of Class B Systems": ERA/
// TD/2011-11 v3.0. http://www.era.europa.eu/Document-Register/Documents/ERA_TD_2011-11%20v30.pdf
type TrainATPSystems struct {
	XMLName xml.Name `xml:"Train_ATP_Systems"`
	*ATPSystemsComplexType
}

// TrainCharacteristics is Element containing the characteristics of the train
type TrainCharacteristics struct {
	TCVersionMajor uint32                 `xml:"TC_VersionMajor,attr"`
	TCVersionMinor uint32                 `xml:"TC_VersionMinor,attr"`
	TCFeatures     *TCFeaturesComplexType `xml:"TC_Features"`
	*TCIdentifierComplexType
}

// TrainCharacteristicsChange is Change in Train Characteristics
type TrainCharacteristicsChange struct {
	TCChangeFeatures *TCChangeFeatures `xml:"TC_ChangeFeatures"`
	*TCIdentifierComplexType
}

// TrainCharacteristicsRef is Reference to a Train Characteristics element (sent separately). All TCs are taken from that element; some can be overridden locally (e.g. length, weight, max speed).
type TrainCharacteristicsRef struct {
	TCVersionMajor              uint32                `xml:"TC_VersionMajor,attr"`
	TCVersionMinor              uint32                `xml:"TC_VersionMinor,attr"`
	Location                    float64               `xml:"location,attr"`
	TrainLength                 *int                  `xml:"trainLength,attr"`
	TrainWeight                 *int                  `xml:"trainWeight,attr"`
	TrainMaxSpeed               *int                  `xml:"trainMaxSpeed,attr"`
	TrainCharacteristicsRefNSPs []*NSPListComplexType `xml:"TrainCharacteristicsRef_NSPs"`
	*TCIdentifierComplexType
}

// TrainCharacteristicsRefNSPs is Network Specific Parameters in Train Characteristics Reference
type TrainCharacteristicsRefNSPs *NSPListComplexType

// TrainDelay is Delay compared to the referenced Date/Time
type TrainDelay struct {
	AgainstBooked     *string `xml:"AgainstBooked"`
	AgainstReferenced *string `xml:"AgainstReferenced"`
}

// TrainIdentification is Identification of the train
type TrainIdentification struct {
	*TrainIdentifierComplexType
}

// TrainInformation ...
type TrainInformation struct {
	TrainLength           *int                   `xml:"trainLength,attr"`
	TrainCategoryCode     *string                `xml:"trainCategoryCode,attr"`
	Company               *Numeric44             `xml:"Company"`
	TrainNumber           *string                `xml:"TrainNumber"`
	StartDate             *string                `xml:"StartDate"`
	AdditionalTrainNumber *AdditionalTrainNumber `xml:"AdditionalTrainNumber"`
	ATPSystemIdentifier   *ATPSystemIdentifier   `xml:"ATP_System_Identifier"`
}

// TrainLocation ...
type TrainLocation *LocationIdent

// TrainLocationInformation is Information about the current location and occupation of the train
type TrainLocationInformation struct {
	LineIdentifier                      *string                                `xml:"lineIdentifier,attr"`
	SegmentClaim                        []*SegmentClaim                        `xml:"SegmentClaim"`
	PositionSpeed                       *PositionSpeed                         `xml:"PositionSpeed"`
	NextTimingPointTrainInformation     []*NextTimingPointTrainInformation     `xml:"NextTimingPointTrainInformation"`
	PreviousTimingPointTrainInformation []*PreviousTimingPointTrainInformation `xml:"PreviousTimingPointTrainInformation"`
	LastPassedKmRef                     *LastPassedKmRef                       `xml:"LastPassedKmRef"`
	LastPassedSignal                    *LastPassedSignal                      `xml:"LastPassedSignal"`
	AverageSpeed                        *AverageSpeed                          `xml:"AverageSpeed"`
	MaxSpeedTrack                       *string                                `xml:"MaxSpeedTrack"`
	Delay                               *string                                `xml:"Delay"`
	DateTimeMeasurement                 *string                                `xml:"DateTimeMeasurement"`
}

// TrainNumber is Operational Train Number
type TrainNumber string

// TrainReady is 0=Not Ready 1=Ready
type TrainReady interface{}

// TrainReadyStatus ...
type TrainReadyStatus struct {
	TrainReady        interface{}    `xml:"TrainReady"`
	TrainLocation     *LocationIdent `xml:"TrainLocation"`
	TrainDelay        *TrainDelay    `xml:"TrainDelay"`
	DelayCause        *string        `xml:"DelayCause"`
	PreciseDelayCause *string        `xml:"PreciseDelayCause"`
	TrainReadyTime    *string        `xml:"TrainReadyTime"`
}

// TrainReadyTime is It indicates date/time when the train will be ready for departure. IM contract will define if this element can be used. Only to be used if this time is in future (sent in advance). Purpose of this element is to short the train stay.
type TrainReadyTime string

// TrainsAfterOwn is Trains that share the same trainpath as your train, which are behind you. These trains will use the same infrastructure after you.
type TrainsAfterOwn struct {
	Train []*Train `xml:"Train"`
}

// TrainsBeforeOwn is Trains that share the same trainpath as your train, which will use the same infrastructure before you.
type TrainsBeforeOwn struct {
	Train []*Train `xml:"Train"`
}

// Tunnel is Tunnel
type Tunnel struct {
	TunnelCategory *string  `xml:"tunnelCategory,attr"`
	TunnelFactor   *float64 `xml:"tunnelFactor,attr"`
	*SPGenericAreaComplexType
}

// UnavailableDASOperatingModes is Operating modes that are unavailable. If all 3 attributes are included, the triplet is unavailable; of only 1 or 2 attributes are included, that mode is unavailable independently of the value of the rest of the attributes.
type UnavailableDASOperatingModes struct {
	XMLName         xml.Name `xml:"Unavailable_DAS_OperatingModes"`
	DASDrivingMode  *string  `xml:"DAS_drivingMode,attr"`
	DASArchitecture *string  `xml:"DAS_architecture,attr"`
	DASConnectivity *string  `xml:"DAS_connectivity,attr"`
}

// UnavailableTractionType is Traction Type of the TrainCharacteristics of the train that has become completely unavailable (e.g. due to a failure of the pantograph)
type UnavailableTractionType struct {
	TractionType string `xml:"tractionType,attr"`
}

// UnprotectedLevelCrossingStop is Stop in rear of an unprotected level crossing
type UnprotectedLevelCrossingStop struct {
	DirectionOfApplicationOnSP *string `xml:"directionOfApplicationOnSP,attr"`
	*SPGenericPointComplexType
}

// Variant is The variant shows a relationship between two identifiers referring to the same business case
type Variant interface{}

// VirtualBalise is A VirtualBalise is a location in the Segment Profile for which the longitude and latitude is known.
type VirtualBalise struct {
	VirtualBalisePosition *GNSSStaticPositionComplexType `xml:"VirtualBalisePosition"`
	*SPGenericPointComplexType
}

// VirtualBalisePosition is GNSS position of Virtual Balise
type VirtualBalisePosition *GNSSStaticPositionComplexType

// WagonData is Wagon relevant data for the wagons within a running train
type WagonData struct {
	WagonNumber        *string          `xml:"WagonNumber"`
	WagonTrainPosition *interface{}     `xml:"WagonTrainPosition"`
	WagonMaxSpeed      *interface{}     `xml:"WagonMaxSpeed"`
	TotalLoadWeight    *int             `xml:"TotalLoadWeight"`
	WagonWeightEmpty   *WeightValueKilo `xml:"WagonWeightEmpty"`
}

// WagonMaxSpeed is Maximum allowed speed of the wagon according to the load and entry in the Rolling Stock Databases.  In kmh
type WagonMaxSpeed interface{}

// WagonNumber is Identifies uniquely the wagon by its number
type WagonNumber string

// WagonTrainPosition is Identifies the position of a wagon within a train.  Sequential number starting with the first wagon at the front of train as N°1.
type WagonTrainPosition interface{}

// WagonWeightEmpty is The weight of an empty wagon according to the entry in the rolling stock database
type WagonWeightEmpty int

// Wind is Temporary constraints due to wind factors
type Wind struct {
	WindSpeed   float64 `xml:"windSpeed,attr"`
	WindHeading uint16  `xml:"windHeading,attr"`
}

// ATPSystemsComplexType is Complex Type for ATP Systems
type ATPSystemsComplexType struct {
	XMLName             xml.Name               `xml:"ATPSystems_ComplexType"`
	ATPSystemIdentifier []*ATPSystemIdentifier `xml:"ATP_System_Identifier"`
}

// BrakingForceCurveComplexType is Complex Type for Braking Force Curves
type BrakingForceCurveComplexType struct {
	XMLName                    xml.Name                      `xml:"BrakingForceCurve_ComplexType"`
	Identifier                 *string                       `xml:"identifier,attr"`
	TractionType               *string                       `xml:"tractionType,attr"`
	BFCAdhesionCategory        *string                       `xml:"BFC_AdhesionCategory,attr"`
	BrakingForceCurveDataPoint []*BrakingForceCurveDataPoint `xml:"BrakingForceCurveDataPoint"`
	ATPSystemIdentifier        []*ATPSystemIdentifier        `xml:"ATP_System_Identifier"`
	BFCTractionType            []*BFCTractionType            `xml:"BFC_TractionType"`
	BFCNSPs                    []*NSPListComplexType         `xml:"BFC_NSPs"`
}

// CurrentLimitationComplexType is Complex Type for Current Limitation
type CurrentLimitationComplexType struct {
	XMLName     xml.Name `xml:"CurrentLimitation_ComplexType"`
	MaxCurValue float64  `xml:"maxCurValue,attr"`
}

// CurveComplexType is Complex Type for Curve
type CurveComplexType struct {
	XMLName        xml.Name `xml:"Curve_ComplexType"`
	CurveRadius    float64  `xml:"curveRadius,attr"`
	CurveDirection *string  `xml:"curveDirection,attr"`
}

// DASModesComplexType is Complex type for DAS modes
type DASModesComplexType struct {
	XMLName         xml.Name `xml:"DAS_ModesComplexType"`
	DASDrivingMode  string   `xml:"DAS_drivingMode,attr"`
	DASArchitecture string   `xml:"DAS_architecture,attr"`
	DASConnectivity string   `xml:"DAS_connectivity,attr"`
}

// ErrorComplexType is Complex type for error in response messages
type ErrorComplexType struct {
	XMLName        xml.Name  `xml:"Error_ComplexType"`
	ErrorCode      *string   `xml:"errorCode,attr"`
	XPath          *xml.Name `xml:"XPath,attr"`
	AdditionalInfo *string   `xml:"additionalInfo,attr"`
}

// GNSSDynamicPositionComplexType is Complex Type for GNSS Dynamic Position
type GNSSDynamicPositionComplexType struct {
	XMLName                         xml.Name `xml:"GNSS_DynamicPosition_ComplexType"`
	Latitude                        float64  `xml:"latitude,attr"`
	Longitude                       float64  `xml:"longitude,attr"`
	Altitude                        *float64 `xml:"altitude,attr"`
	Heading                         *float64 `xml:"heading,attr"`
	GNSSPrecision                   *float64 `xml:"GNSS_precision,attr"`
	Timestamp                       *string  `xml:"timestamp,attr"`
	AntennaDistanceFromFrontOfTrain *float64 `xml:"antennaDistanceFromFrontOfTrain,attr"`
}

// GNSSStaticPositionComplexType is Complex Type for GNSS Static Position
type GNSSStaticPositionComplexType struct {
	XMLName   xml.Name `xml:"GNSS_StaticPosition_ComplexType"`
	Latitude  float64  `xml:"latitude,attr"`
	Longitude float64  `xml:"longitude,attr"`
	Altitude  *float64 `xml:"altitude,attr"`
}

// GradientComplexType is Complex Type for Gradient
type GradientComplexType struct {
	XMLName           xml.Name              `xml:"Gradient_ComplexType"`
	GradientValue     float64               `xml:"gradientValue,attr"`
	GradientDirection GradientDirectionType `xml:"gradientDirection,attr"`
}

// LineAndTrackIdentifierComplexType is Complex Type for Line and Track Identifier
type LineAndTrackIdentifierComplexType struct {
	XMLName         xml.Name     `xml:"LineAndTrackIdentifier_ComplexType"`
	LineIdentifier  *string      `xml:"lineIdentifier,attr"`
	LineName        *interface{} `xml:"lineName,attr"`
	TrackIdentifier *interface{} `xml:"trackIdentifier,attr"`
}

// LocationIdent is Indication of the Railway or Customer Location
type LocationIdent struct {
	CountryCodeISO                   *CountryCodeISO                   `xml:"CountryCodeISO"`
	LocationPrimaryCode              int                               `xml:"LocationPrimaryCode"`
	PrimaryLocationName              *string                           `xml:"PrimaryLocationName"`
	LocationSubsidiaryIdentification *LocationSubsidiaryIdentification `xml:"LocationSubsidiaryIdentification"`
}

// NSPListComplexType is Complex Type for a list of Network-Specific Parameters
type NSPListComplexType struct {
	XMLName                  xml.Name                    `xml:"NSPList_ComplexType"`
	Name                     *string                     `xml:"name,attr"`
	NetworkSpecificParameter []*NetworkSpecificParameter `xml:"NetworkSpecificParameter"`
}

// OTNIDComplexType is Complex Type for the Operational Train Number (OTN) Identifier
type OTNIDComplexType struct {
	XMLName               xml.Name               `xml:"OTN_ID_ComplexType"`
	Company               *Numeric44             `xml:"Company"`
	TrainNumber           string                 `xml:"TrainNumber"`
	AdditionalTrainNumber *AdditionalTrainNumber `xml:"AdditionalTrainNumber"`
	StartDate             string                 `xml:"StartDate"`
}

// SegmentProfileIdentificationComplexType is Complex type containing the data for the identification of Segment Profiles
type SegmentProfileIdentificationComplexType struct {
	XMLName xml.Name           `xml:"SegmentProfileIdentification_ComplexType"`
	SPID    string             `xml:"SP_ID,attr"`
	SPZone  *SPZoneComplexType `xml:"SP_Zone"`
}

// SignalComplexType is Complex type for the definition of a signal
type SignalComplexType struct {
	XMLName                       xml.Name                       `xml:"Signal_ComplexType"`
	DirectionOfApplicationOnSP    *string                        `xml:"directionOfApplicationOnSP,attr"`
	SignalID                      *SignalID                      `xml:"Signal_ID"`
	SignalApplication             *SignalApplication             `xml:"SignalApplication"`
	SignalInformation             *SignalInformation             `xml:"SignalInformation"`
	SignalFunction                []interface{}                  `xml:"SignalFunction"`
	SignalPhysicalCharacteristics *SignalPhysicalCharacteristics `xml:"SignalPhysicalCharacteristics"`
	SignalNSPs                    []*NSPListComplexType          `xml:"Signal_NSPs"`
}

// SPGenericAreaComplexType is Complex Type for generic Segment Profile Area
type SPGenericAreaComplexType struct {
	XMLName           xml.Name          `xml:"SP_GenericArea_ComplexType"`
	Identifier        *string           `xml:"identifier,attr"`
	StartLocation     *float64          `xml:"startLocation,attr"`
	EndLocation       *float64          `xml:"endLocation,attr"`
	StartEndQualifier StartEndQualifier `xml:"startEndQualifier,attr"`
}

// SPGenericPointComplexType is Complex Type for generic Segment Profile Point
type SPGenericPointComplexType struct {
	XMLName    xml.Name `xml:"SP_GenericPoint_ComplexType"`
	Identifier *string  `xml:"identifier,attr"`
	Location   float64  `xml:"location,attr"`
}

// SPZoneComplexType is Complex Type for SP Zone. At least one of IM_ID (Infrastructure Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be provided, especially if conversion to SUBSET-126 is required.
type SPZoneComplexType struct {
	IMID *Numeric44 `xml:"IM_ID"`
	NIDC *int16     `xml:"NID_C"`
}

// SpecificSSPComplexType is Complex Type for Specific Static Speed Profile
type SpecificSSPComplexType struct {
	XMLName             xml.Name           `xml:"SpecificSSP_ComplexType"`
	SpecificSSPCategory *string            `xml:"specific_SSP_Category,attr"`
	VDIFF               *float64           `xml:"V_DIFF,attr"`
	CantDeficiencySSP   *CantDeficiencySSP `xml:"CantDeficiencySSP"`
	OtherSpecificSSP    *OtherSpecificSSP  `xml:"OtherSpecificSSP"`
}

// SSPComplexType is Complex type for Static SPeed Profiles
type SSPComplexType struct {
	XMLName            xml.Name                  `xml:"SSP_ComplexType"`
	SSPSpeed           *float64                  `xml:"SSP_Speed,attr"`
	SSPFront           *bool                     `xml:"SSP_Front,attr"`
	SSPVisualReference *string                   `xml:"SSP_VisualReference,attr"`
	SpecificSSP        []*SpecificSSPComplexType `xml:"SpecificSSP"`
}

// StoppingZoneComplexType is Complex type for stopping zone
type StoppingZoneComplexType struct {
	XMLName    xml.Name           `xml:"StoppingZone_ComplexType"`
	Identifier string             `xml:"identifier,attr"`
	Name       *string            `xml:"name,attr"`
	SPZone     *SPZoneComplexType `xml:"SP_Zone"`
}

// TAPTAFIDComplexType is Is only used in the operational phase and refers to the date where the single train will start the train journey
type TAPTAFIDComplexType struct {
	XMLName       xml.Name     `xml:"TAP_TAF_ID_ComplexType"`
	ObjectType    interface{}  `xml:"ObjectType"`
	Company       *Numeric44   `xml:"Company"`
	Core          interface{}  `xml:"Core"`
	Variant       interface{}  `xml:"Variant"`
	TimetableYear *interface{} `xml:"TimetableYear"`
	StartDate     string       `xml:"StartDate"`
}

// TCFeaturesComplexType is Complex type for Train Characteristics features
type TCFeaturesComplexType struct {
	XMLName                    xml.Name                         `xml:"TC_Features_ComplexType"`
	TrainLength                *int                             `xml:"trainLength,attr"`
	TrainMaxSpeed              *int                             `xml:"trainMaxSpeed,attr"`
	TrainWeight                *int                             `xml:"trainWeight,attr"`
	ComfortableAcceleration    *float64                         `xml:"comfortableAcceleration,attr"`
	ComfortableBraking         *float64                         `xml:"comfortableBraking,attr"`
	MaximumRegenerativeForce   *float64                         `xml:"maximumRegenerativeForce,attr"`
	MaximumRegenerativeBraking *float64                         `xml:"maximumRegenerativeBraking,attr"`
	EnginePower                *float64                         `xml:"enginePower,attr"`
	RolloutCoefficientA        *float64                         `xml:"rolloutCoefficientA,attr"`
	RolloutCoefficientB        *float64                         `xml:"rolloutCoefficientB,attr"`
	RolloutCoefficientC        *float64                         `xml:"rolloutCoefficientC,attr"`
	RollingStockType           *string                          `xml:"rollingStockType,attr"`
	BrakedWeightPercentage     *float64                         `xml:"brakedWeightPercentage,attr"`
	BrakingReactionTime        *string                          `xml:"brakingReactionTime,attr"`
	TractionReactionTime       *int                             `xml:"tractionReactionTime,attr"`
	RotatingMassFactor         *float64                         `xml:"rotatingMassFactor,attr"`
	Engines                    *int                             `xml:"engines,attr"`
	Wagons                     *int                             `xml:"wagons,attr"`
	TrpBe                      *float64                         `xml:"trp_be,attr"`
	TrpBw                      *float64                         `xml:"trp_bw,attr"`
	TrpAe                      *float64                         `xml:"trp_ae,attr"`
	TrpAw                      *float64                         `xml:"trp_aw,attr"`
	NCTrain                    *string                          `xml:"NC_Train,attr"`
	NCCDTrain                  *string                          `xml:"NC_CDTrain,attr"`
	TrainCategoryCode          *string                          `xml:"trainCategoryCode,attr"`
	DisplayName                *string                          `xml:"displayName,attr"`
	AxleLoadCategory           *string                          `xml:"axleLoadCategory,attr"`
	PositionOfLastPantograph   *float64                         `xml:"positionOfLastPantograph,attr"`
	TrainATPSystems            *TrainATPSystems                 `xml:"Train_ATP_Systems"`
	TractionForceCurve         []*TractionForceCurveComplexType `xml:"TractionForceCurve"`
	BrakingForceCurve          []*BrakingForceCurveComplexType  `xml:"BrakingForceCurve"`
	WagonData                  []*WagonData                     `xml:"WagonData"`
	TCFeaturesNSPs             []*NSPListComplexType            `xml:"TC_Features_NSPs"`
}

// TCIdentifierComplexType is Type for the identification of Train Characteristics. Contains the ID of TC, and the mandatory identification of the entity responsible (RU or VKM) for the data
type TCIdentifierComplexType struct {
	XMLName            xml.Name     `xml:"TC_Identifier_ComplexType"`
	TCID               string       `xml:"TC_ID,attr"`
	TCRUID             *Numeric44   `xml:"TC_RU_ID"`
	KeeperShortNameVKM *interface{} `xml:"KeeperShortNameVKM"`
}

// TemporaryConstraintsComplexType is Complex type for Temporary Constraints
type TemporaryConstraintsComplexType struct {
	XMLName           xml.Name `xml:"TemporaryConstraints_ComplexType"`
	StartLocation     *float64 `xml:"startLocation,attr"`
	EndLocation       *float64 `xml:"endLocation,attr"`
	StartEndQualifier string   `xml:"startEndQualifier,attr"`
	StartTime         *string  `xml:"startTime,attr"`
	EndTime           *string  `xml:"endTime,attr"`
}

// TimingPointIdentificationComplexType is Complex type containing the data for the identification of Timing Points
type TimingPointIdentificationComplexType struct {
	XMLName xml.Name           `xml:"TimingPointIdentification_ComplexType"`
	TPID    string             `xml:"TP_ID,attr"`
	SPZone  *SPZoneComplexType `xml:"SP_Zone"`
}

// TotalCurrentComplexType is Complex Type for Total Current temporary constraints
type TotalCurrentComplexType struct {
	XMLName     xml.Name `xml:"TotalCurrent_ComplexType"`
	MaxCurValue float64  `xml:"maxCurValue,attr"`
}

// TractionForceCurveComplexType is Complex Type for Traction Force Curves
type TractionForceCurveComplexType struct {
	XMLName                     xml.Name                       `xml:"TractionForceCurve_ComplexType"`
	Identifier                  *string                        `xml:"identifier,attr"`
	TractionType                string                         `xml:"tractionType,attr"`
	TFCEstimatedVoltage         *int                           `xml:"TFC_estimatedVoltage,attr"`
	TFCMaxCurrentValue          *float64                       `xml:"TFC_maxCurrentValue,attr"`
	TractionForceCurveDataPoint []*TractionForceCurveDataPoint `xml:"TractionForceCurveDataPoint"`
	TFCNSPs                     []*NSPListComplexType          `xml:"TFC_NSPs"`
}

// TractionTypeTransitionComplexType is Complex Type for Traction Type Transition
type TractionTypeTransitionComplexType struct {
	XMLName         xml.Name `xml:"TractionTypeTransition_ComplexType"`
	NewTractionType string   `xml:"newTractionType,attr"`
}

// TrainIdentifierComplexType is Complex Type for train identifier. Train identifier can be TAP/TAF or OTN.
type TrainIdentifierComplexType struct {
	XMLName  xml.Name             `xml:"TrainIdentifier_ComplexType"`
	TAPTAFID *TAPTAFIDComplexType `xml:"TAP_TAF_ID"`
	OTNID    *OTNIDComplexType    `xml:"OTN_ID"`
}

// TrainType is Complex type for the "Train" object used in related train information.
type TrainType struct {
	Order                    *interface{}              `xml:"order,attr"`
	TrainInformation         *TrainInformation         `xml:"TrainInformation"`
	TrainLocationInformation *TrainLocationInformation `xml:"TrainLocationInformation"`
}

// VoltageComplexType is Complex Type for voltage
type VoltageComplexType struct {
	VoltageValue int    `xml:"voltageValue,attr"`
	NIDCTRACTION *uint8 `xml:"NID_CTRACTION,attr"`
}

// AdhesionCategoryType is Simple Type for adhesion category
type AdhesionCategoryType string

// CompanyCode is From TAP/TAF: Identifies the RU, IM or other company involved in the Rail Transport Chain
type CompanyCode int

// CountryIdentISO is ISO 3166-1 alpha code (2 positions)
type CountryIdentISO string

// CurrentType is Simple Type for current (A)
type CurrentType float64

// DelayCode is Reason for a delay or interruption.  UIC Leaflet 450-2, Appendix C.
// The first digit in the code has to following meaning:
// 1  Operational planning, Management (IM)
// 2 Infrastructure installations (IM)
// 3 Civil engineering causes (IM)
// 4 Causes of other IM (IM)
// 5 Commercial causes (RU)
// 6 Rolling stock (RU)
// 7 Causes of other RU (RU)
// 8 External causes
// 9 Secondary causes
// New codes added:
// 23Power supply equipment
// 58Staff
// 68Staff
// 90Dangerous incidents, accidents and hazards
// 91Track occupation caused by the lateness of the same train
// 92Track occupation caused by the lateness of another train
// 93Turn round
// 94Connection
// 95Further investigation needed
type DelayCode string

// DeltaTime is Time difference delay (+) or ahead of schedule (-) this shall be 1character + 4 Numeric
type DeltaTime string

// DeviceType is Simple type for DAS devices
type DeviceType string

// DistanceType is Simple Type for distance attributes
type DistanceType float64

// EnergyType is Simple Type for energy (kWh)
type EnergyType float64

// ForceType is Force Type
type ForceType float64

// FreeText is Clear Text in ISO Unicode character set
type FreeText string

// GradientDirectionType is Type for gradient direction (Downhill or Uphill)
type GradientDirectionType string

// GradientValueType is Type for gradient values (Numeric value, mm by m)
type GradientValueType float64

// NCCDTrainType is Simple Type for Cant Deficiency Train Category (NC_CDTRAIN as defined in SUBSET-026).
type NCCDTrainType string

// NCTrainType is Simple Type for Train Category (NC_TRAIN as defined in SUBSET-026).
type NCTrainType string

// NIDCType is Simple type for NID_C
type NIDCType int16

// Numeric15 ...
type Numeric15 int

// Numeric22 ...
type Numeric22 int

// Numeric44 is Numeric value with 4 digits
type Numeric44 string

// PercentageType is Percentage Type
type PercentageType float64

// PowerTypekW is Simple Type for power (kW)
type PowerTypekW float64

// PowerTypeW is Simple Type for power (W)
type PowerTypeW float64

// RelatedTrainRequestType ...
type RelatedTrainRequestType string

// SpeedType is Speed Type
type SpeedType float64

// String110 ...
type String110 string

// TractionTypeType is Simple Type for TractionTypes
type TractionTypeType string

// TrainPartType is Simple type for train part (head, end or last pantograph). To be used for defining which part of the train has to clear a signal for application.
type TrainPartType string

// UTCDateTimeType is Simple Type for date and time in UTC
type UTCDateTimeType string

// UUIDv4Type is UUIDv4 (random) simpleType as defined in RFC4122
type UUIDv4Type string

// VoltageType is Simple Type for voltage (V)
type VoltageType int

// WeightValueKilo is In Kilograms
type WeightValueKilo int

// AdditionalInfo is Additional details
type AdditionalInfo string

// AdviceType is Type of C-DAS-C Advice
type AdviceType string

// AimEarlyArrival is Aim to be as near as possible to the beginning of the arrival window
type AimEarlyArrival bool

// ALSPFront is [SUBSET-126] Qualifier to indicate if the ALSP is to be applied to the front of the train (no train length delay) or to the end of the train (train length delay).
type ALSPFront bool

// Altitude is Altitude (m). Considering ETRS89 as reference.
type Altitude float64

// AmountLevelCrossings is Amount of guarded level crossings (to be used only if levelCrossingGuarding is true)
type AmountLevelCrossings uint8

// AntennaDistanceFromFrontOfTrain is Distance of GPS antenna from the front of the train
type AntennaDistanceFromFrontOfTrain float64

// ApplicationDistance is How far (in m) after the signal the speed must be maintained.
type ApplicationDistance float64

// ArrivalTime is Arrival time at the Timing Point
type ArrivalTime string

// ArrivalWindow is Acceptable time allowance to be earlier at the TP
type ArrivalWindow string

// Ascending is Ascending or Descending
type Ascending string

// ASRFront is [SUBSET-126] Qualifier to indicate if a speed limit given for a profile element is to be applied  until the front of the train (1, no train length delay) or the end of the train (0, train length delay) has left the element
type ASRFront bool

// ASRSpeed is [SUBSET-126] Value of the speed level restriction (km/h)
type ASRSpeed float64

// ATOState is [SUBSET-126] ATO State, from Subset-126. Only relevant when in ATO mode.
type ATOState string

// ATOTSID is [SUBSET-126] Identifier of the ATO-TS
type ATOTSID string

// AutomaticClosing is [SUBSET-126] This variable defines if the ATO-OB has to manage the train doors closing (1) or not (0)
type AutomaticClosing bool

// Autonomy is Amount of energy still available on-board for the self-propelled energy source (in kWh)
type Autonomy float64

// AxleLoadCategory is [SUBSET-126] Axle load from which the speed restriction is applicable
type AxleLoadCategory string

// BaliseGroupIdentifier is [SUBSET-126] Identifier of the balise group
type BaliseGroupIdentifier uint16

// BaliseRelativePosition is [SUBSET-126] Defines the relative position in a balise group
type BaliseRelativePosition uint8

// BFCAdhesionCategory is Adhesion category for which the braking force curve is applicable. When no curves are defined for better adhesion categories, the curve is also applicable.
type BFCAdhesionCategory string

// BrakingForce is Braking Force (in N)
type BrakingForce float64

// BrakingReactionTime is Delay between braking command and start of braking.
type BrakingReactionTime string

// CDASCAble is 1 = The device is able to receive C-DAS-C messages. 0 (Default) = The device is unable to receive C-DAS-C messages
type CDASCAble bool

// CentralisedOpening is [SUBSET-126] This variable defines if the doors are to be opened centralised (1) or by the passengers (0)
type CentralisedOpening bool

// ChangedStoppingLocationAlert is This attribute is to be used when the driver will be asked to stop at a stopping which is different from the usual one, and it needs to be brought to his attention.
type ChangedStoppingLocationAlert string

// CIInstanceNumber is Number of a Common Interface Instance for the same Company
type CIInstanceNumber int

// ComfortableAcceleration is Comfortable acceleration. Expressed in m/s^2
type ComfortableAcceleration float64

// ComfortableBraking is Deceleration used for comfortable braking. Expressed in m/s^2
type ComfortableBraking float64

// CommunicationNetworkChannel is Channel for communication network
type CommunicationNetworkChannel string

// CommunicationNetworkType is Type of communication network. Standard possibilities are: "GSM-R", "FRMCS", "Public GSM", "Public WiFi" and "Private WiFi".
type CommunicationNetworkType string

// ConnectionTrackDescription is Description of connection track
type ConnectionTrackDescription string

// ConnectionTrackType is Type of connection track
type ConnectionTrackType string

// ContactIdentifier is Identifier to be used to reach the contact on the communication platform defined in ContactType (e.g. telephone number, GSM-R number, pseudo...)
type ContactIdentifier string

// ContactNotes is Notes related to the contact
type ContactNotes string

// ContactRole is Role of the contact
type ContactRole interface{}

// ContactType is Type of contact. Communication platform to be used to contact this person: e.g. telephone, GSM-R, social network, messaging app, etc.
type ContactType string

// CorrelationID is ID of the message for which the current message is a reply
type CorrelationID string

// CurveDirection is Direction of the curve (Left or Right)
type CurveDirection string

// CurveRadius is Radius of the curve (m)
type CurveRadius float64

// DASArchitecture is DAS architecture: DAS-C (ground advice calculation) or DAS-O (board advice calculation)
type DASArchitecture string

// DASBehaviour is Behaviour of DAS (Standalone or Connected)
type DASBehaviour string

// DASConnectivity is DAS connectivity (standalone or connected)
type DASConnectivity string

// DASDrivingMode is DAS driving mode (see enumeration)
type DASDrivingMode string

// DASStatus is DAS Status
type DASStatus string

// DataFirstAvailable is Date and time of first availability of the requested data
type DataFirstAvailable string

// DaylightSaving is [SUBSET-126] Specifies some information specific for the TPThis variable defines if the daylight saving hour is applicable to calculate the local time
type DaylightSaving bool

// DelayCategory is Category of the delay. Delay could be measured, delay can be expected (for example before departure), announcement that the delay will be published (upcoming) or can be unkown
type DelayCategory string

// DeltaSpeed is Difference (in km/h) between maximum allowed track speed/most restrictive speed profile and optimal speed.
// The special value 0 indicates "maximum allowed track speed/most restrictive speed profile".
type DeltaSpeed float64

// DepartureAuthorisationType is Type of authorisation for departure (e.g. SMS)
type DepartureAuthorisationType string

// DepartureTime is Expected time of departure at Timing Point expressed in UTC
type DepartureTime string

// Description is Description of the element
type Description interface{}

// DestinationCompany is ID of the company to which the message is addressed. Must be a UIC RICS Code: : https://uic.org/rics
type DestinationCompany int

// DestinationDevice is ID of the device to which the message is addressed. For any given company, this ID must be unique. Maximum 128 characters. For G2B messages, if this attribute is empty, the message is directed to all devices in the train. For B2G messages, this attribute can be ignored.
type DestinationDevice string

// DirectionOfApplicationOnSP is Direction of application on the Segment Profile (Nominal, Reverse or Both)
type DirectionOfApplicationOnSP string

// DisplayName is Name of train type to be displayed to the driver (e.g. ‘TGV’ of ‘VIRM 6’).
type DisplayName string

// DistanceEoAOffset is [SUBSET-126] Distance to stop the train in rear of the End of Authority (in m)
type DistanceEoAOffset float64

// DistanceSTPReached is [SUBSET-126] Distance (in m) from a Stopping Point to consider it as reached
type DistanceSTPReached float64

// DistanceBeforeSignal is The distanceBeforeSignal attribute gives the point from which to restart accelerating. Depending on the type of signalling, the position of this point will be defined by the visual identification by the driver, the position of a balise, or the time at which the driver will receive an updated information for in-cab signalling
type DistanceBeforeSignal float64

// DistanceToPowerConstraint is Distance (in m) until the Power Constraint is to be applied
type DistanceToPowerConstraint float64

// DistanceToTractionTypeTransition is Distance (in m) from the actual Transition Point
type DistanceToTractionTypeTransition float64

// DriverID is [SUBSET-126] ATO: Driver identifier number as defined in SUBSET-027 v300 §4.2.3.7.
type DriverID string

// DynamicBrakeForceLimit is [SUBSET-126] Maximum Brake Force (kN) to be applied by the Dynamic Brake
type DynamicBrakeForceLimit uint16

// EarliestDepartureAdviceTime is Lower bound of time window for train departure. Only to be used if departure advice is sent. If earliestDepartureTime is not specified, the advice should be obeyed at any time earlier than latestDepartureTime. If both attributes are not specified, the advice should be obeyed immediately.
type EarliestDepartureAdviceTime string

// EarliestDepartureTime is Earliest departure time desired, expressed in UTC.
type EarliestDepartureTime string

// EndKm is References the kilometre at the end of the message, preferably linked with a kilometer sign near the track
type EndKm float64

// EndKmAbs is References the kilometre where the optimalSpeed changes next, preferably linked with a kilometer sign near the track
type EndKmAbs float64

// EndLine is References the (national) line number at the end of the message
type EndLine int

// EndLineAbs is References the (national) line number where the optimalSpeed changes next
type EndLineAbs int

// EndLocation is [SUBSET-126] Distance of the element's end location from the beginning of the Segment Profile (m)
type EndLocation float64

// EndPointOfDownLineSpeedApplication is How far (in m) after the considered signal the speed advice must cease to be respected.
type EndPointOfDownLineSpeedApplication float64

// EndSignal is References the name of the end signal of the message, if the end of the message is at or near a signal
type EndSignal string

// EndSignalAbs is References the name of the signal where the optimalSpeed changes next, if this is at or near a signal
type EndSignalAbs string

// EndSignalApplicationMustBeClearedBy is What part of the train must clear the signal application before the signal stops being applied
type EndSignalApplicationMustBeClearedBy string

// EndStationLong is References the start station at the end of the message (long form)
type EndStationLong string

// EndStationLongAbs is References the station where the optimalSpeed changes next (long form)
type EndStationLongAbs string

// EndStationShort is References the start station at the end of the message (short form)
type EndStationShort string

// EndStationShortAbs is References the station where the optimalSpeed changes next (short form)
type EndStationShortAbs string

// EndTime is Ending date and time (in UTC)
type EndTime string

// EndValidity is End of the advice validity. Either both startValidity and endValidity or none of the two must be specified. Both parameters should refer to a future point in time and endValidity must be after startValidity. If either of the two indicates a past time (w.r.t. the clock of the On-Board Unit), which is possible in a situation where Central Unit and OBU are not synchronised, transmission delays etc., then the OBU should treat it as valid but past time indication.
type EndValidity string

// EndValidityAbs is Time at the end of the optimalSpeed
type EndValidityAbs string

// EnginePower is Maximum power of the engine. Expressed in W
type EnginePower float64

// Engines is Number of engines
type Engines int

// ErrorCode is Error Code
type ErrorCode string

// EuropeanVehicleNumber is Number on 12 digits allocated to each rail vehicle to uniquely identify it from any other rail vehicle
type EuropeanVehicleNumber string

// ExpectedConsumptionCurrent is Expected consumption current (A)
type ExpectedConsumptionCurrent float64

// ExpectedRegenerativeCurrent is Expected regenerative current (A)
type ExpectedRegenerativeCurrent float64

// ForecastedArrival is Forecasted arrival at the Timing Point
type ForecastedArrival string

// ForecastedDeparture is Forecasted departure from the Timing Point
type ForecastedDeparture string

// GNSSPrecision is Precision of the GNSS position. Expressed in metres.
type GNSSPrecision float64

// GradientPBD is [SUBSET-126] A single gradient value applicable for the calculation of the Permitted Braking Distance
type GradientPBD float64

// GradientDirection is Direction of the gradient (Downhill or Uphill)
type GradientDirection string

// GradientDirectionPBD is [SUBSET-126] Direction of the gradient for the Permitted Braking Distance
type GradientDirectionPBD string

// GradientValue is Gradient value (Numeric value, mm by m)
type GradientValue float64

// HandshakeRejectReason is Reason for handshake rejection
type HandshakeRejectReason string

// Heading is Direction in which the train is going. Expressed in degrees from North (0 = North, 90 = East, 180 = South, 270 = West).
type Heading float64

// HorizontalOffset is Horizontal offset (in m) from the center of the track to the middle of the object. Negative is left and positive is right.
type HorizontalOffset float64

// Identifier is Element identifier
type Identifier string

// JPChangeReason is Reason for JP Change
type JPChangeReason string

// JPSPConsistencyError is [SUBSET-126] ATO: JP SP Consistency Error, from Subset-126 (boolean)
type JPSPConsistencyError bool

// JPStatus is [SUBSET-126] Status of the Journey Profile: "Valid": JP containing the data requested. "Unavailable": JP specifies that the requested part of the Journey Profile is currently not available yet (but all the previously sent JPs are still applicable). "Invalid": JP specifies that the TP identifier asserted in the JPReq does not belong to the preceding JP already sent to the ATO-OB. "Update": JP specifies that the Journey Profile has been updated by the TMS within the current visibility of the ATO-OB.
type JPStatus string

// JPVersion is Version of the Journey Proifile
type JPVersion string

// KmRef is KM reference point (in km)
type KmRef float64

// Language is ISO 639-1 language code
type Language string

// LatestDepartureAdviceTime is Upper bound of time window for train departure. Only to be used if departure advice is sent. If latestDepartureTime is not specified, the advice should be obeyed at any time later than earliestDepartureTime. If both attributes are not specified, the advice should be obeyed immediately.
type LatestDepartureAdviceTime string

// Latitude is Latitude in signed degrees format (DDD.ddddd) in the WGS84 system.
type Latitude float64

// LeaveASAP is Leave as soon as possible, taking into account also the other variables (boolean).
type LeaveASAP bool

// LevelCrossingDescription is Description of level crossing
type LevelCrossingDescription string

// LevelCrossingGuarding is Determines if the balise is guarding any level crossing
type LevelCrossingGuarding bool

// LevelCrossingType is Type of level crossing
type LevelCrossingType string

// LineIdentifier is Identification of the line. For example “115000” is the Line Identifier for the line from Strasbourg to Basel.
type LineIdentifier string

// LineName is Name of the line describing the connection stations. For example “Ligne de Strasbourg à Bale."
type LineName interface{}

// Location is Distance of the element from the beginning of the Segment Profile (m)
type Location float64

// LocationNegativeErrorMargin is Error margin of the position of the head of the train. This margin subtracted from the location attribute gives the least downline position of the train possible at the moment of the measure. Expressed in m
type LocationNegativeErrorMargin interface{}

// LocationPositiveErrorMargin is Error margin of the position of the head of the train. This margin added to the location attribute gives the most downline position of the train possible at the moment of the measure. Expressed in m
type LocationPositiveErrorMargin float64

// LocationSubsidiaryTypeCode is An indicator used to show that this location is a subsidiary of another location.
type LocationSubsidiaryTypeCode string

// LocationText is Human readable text of the location information where the advice becomes valid (and/or ends), which will be displayed on the OBU. If no  location is given the advice is immediately valid. It is recommended to fill the parameter location according to the following rules, depending on the advice instance being sent:
// - stop advice: location is the station the train is stopping at (e.g. train approaching a scheduled stop)
// - acceleration advice: location is the location (kilometre/landmark etc.) where the acceleration phase starts
// - constant speed advice: location is the location (kilometre/landmark etc.) where the cruising phase starts and/or ends
// - operational/electrical braking advice: location is the location (kilometre/landmark etc.) where the braking phase starts
// - end of advice: location is null
// - departure advice: location is the station the train is departing from
type LocationText string

// Longitude is Longitude in signed degrees format (DDD.ddddd) in the WGS84 system.
type Longitude float64

// LowAdhesionCategory is Adhesion category aligned with the definition in EN15595 of the wheel-rail interface for the lower boundaries (definition and values taken from SUBSET-126).
type LowAdhesionCategory string

// MainContact is Specifies if this is the main contact for the area
type MainContact bool

// MandatoryStop is Flag indicating whether the stop is mandatory (1) or optional (0)
type MandatoryStop bool

// MaxAdvisedSpeed is Maximum advised speed. Must be higher than "speed".
type MaxAdvisedSpeed float64

// MaxCurValue is Maximum current value (A)
type MaxCurValue float64

// MaximumCantDeficiency is Maximum Cant Deficiency supported by the train (in mm)
type MaximumCantDeficiency int

// MaximumNumberOfCars is Maximum number of cars
type MaximumNumberOfCars int16

// MaximumNumberOfRSType is Maximum number of a certain type of rolling stock
type MaximumNumberOfRSType uint8

// MaximumRegenerativeBraking is Maximum power that can be regenerated. Expressed in W
type MaximumRegenerativeBraking float64

// MaximumRegenerativeForce is Maximum force that can be Applied through regenerative braking. Expressed in N
type MaximumRegenerativeForce float64

// MaximumTrainLength is Maximum Train Length for which stopping point is valid. Expressed in m
type MaximumTrainLength float64

// MaxPower is Maximal power that can be delivered form energy source available on-board (in kW)
type MaxPower float64

// MeasuredCurrent is Measured current taken from overhead contact line, negative during regeneration (in A)
type MeasuredCurrent float64

// MeasuredVoltage is Measured voltage at pantograph (in V)
type MeasuredVoltage int

// MessageID is Identifier of the SFERA message
type MessageID string

// MessageString is String containing the message
type MessageString string

// MinAdvisedSpeed is Minimum advised speed. Must be lower than "speed".
type MinAdvisedSpeed float64

// MinimumDwellTime is [SUBSET-126] Minimum dwell time at given Stopping Point (in seconds)
type MinimumDwellTime string

// MinimumNumberOfCars is Minimum number of cars
type MinimumNumberOfCars int16

// MinimumNumberOfRSType is Minimum number of a certain type of rolling stock
type MinimumNumberOfRSType uint8

// MinimumPositionUpdateRequestPeriod is Minimum time to be waited between two position requests by board
type MinimumPositionUpdateRequestPeriod string

// MinimumTrainLength is Minimum Train Length for which stopping point is valid. Expressed in m
type MinimumTrainLength float64

// Name is Name of the element
type Name string

// NCCDTrain is Cant Deficiency Train Category (NC_CDTRAIN as defined in SUBSET-026).
type NCCDTrain string

// NCTrain is Train Category (NC_TRAIN as defined in SUBSET-026).
type NCTrain string

// NewAdhesionConditions is New adhesion conditions
type NewAdhesionConditions string

// NewSpeedLevel is [SUBSET-126] Speed restriction to be applied if the axle load of the train ≥ axleLoadCategory
type NewSpeedLevel float64

// NewTractionType is New Traction Type
type NewTractionType string

// NextStoppingPointSkip is [SUBSET-126] ATO: skip next stopping point, from Subset-126 (boolean)
type NextStoppingPointSkip bool

// NIDCTRACTION is [SUBSET-126] Country identifier of the traction system. It identifies the information, additional to voltageValue, required to fully define the traction system. Use NID_CTRACTION numbers from ERA document ERA_ERTMS_040001.
type NIDCTRACTION uint8

// NIDENGINE is [SUBSET-126] ETCS identity number uniquely defined for ERTMS/ETCS purposes. Values to be assigned. 24 binary positions.
type NIDENGINE string

// NIDOPERATIONAL is [SUBSET-126] Operational train running number. NID_OPERATIONAL consists of up to 8 digits which are entered left adjusted into the data field, the leftmost digit is the digit to be entered first. In case the NID_OPERATIONAL is shorter than 8 digits, the remaining space is to be filled with special character “F”.
type NIDOPERATIONAL string

// OpeningDoorSide is [SUBSET-126] Specifies if the ATO-OB has to manage the train doors opening and on which side the passenger exchange doors have to be opened
type OpeningDoorSide string

// OperationalConditionsFulfillment is [SUBSET-126] ATO: Operational conditions fulfilment, from Subset-126 (boolean)
type OperationalConditionsFulfillment bool

// OptimalSpeed is Optimal speed given to the
// driver (km/h). It is recommended to fill the parameter optimalSpeed according to the
// following rules, depending on the advice instance being sent:
// - Stop Advice: optimalSpeed must be null
// - Acceleration Advice: optimalSpeed is the target speed that shall be reached at the end of the acceleration phase
// - Constant Speed Advice: optimalSpeed is the target speed for the cruising phase
// - Coasting Advice: optimalSpeed is the target speed that shall be reached at the end of the coasting phase
// - Operational/Electrical Braking Advice: optimalSpeed is the target speed that shall be reached at the end of the braking phase
// - Departure Advice: optimalSpeed is either null or the target speed that shall be reached after departing at the end of the acceleration phase
// There is a special value: 999 indicating "maximum allowed track speed"
type OptimalSpeed float64

// OriginalLanguage is Approved version of text (translation when 'false')
type OriginalLanguage bool

// OTN is Operational Train Number (free text). For driver information.
type OTN string

// PBDBrake is [SUBSET-126] Whether the permitted braking distance is to be achieved with the Service Brake or Emergency Brake
type PBDBrake string

// Percentage is The meaning depends on the type of advice instance being sent. Only valid with certain types of advice.
// - acceleration advice: Percentage of the maximum acceleration
// - operational braking advice: Percentage of the maximum operational braking
// - electrical braking advice: Percentage of the maximum electrical braking
type Percentage float64

// PermittedBrakingDistance is [SUBSET-126] Permitted Braking Distance value (in m)
type PermittedBrakingDistance float64

// PlannedDepartureTime is The original planned departure time from the Timing Point expressed in UTC. Useful for driver information.
type PlannedDepartureTime string

// PlannedStop is Flag indicating whether the stop is planned (1) or unplanned (0)
type PlannedStop bool

// PositionOfLastPantograph is Distance from the head of the train of the last point of the rolling stock that has to clear a zone with a power constraint (lowered pantograph, main switch off…) for that power constraint to be lifted. Expressed in m.
type PositionOfLastPantograph float64

// PositionSpeedRequest2 is Request for position/speed update by board
type PositionSpeedRequest2 bool

// PowerConstraintInstruction is Instruction for Power Constraint
type PowerConstraintInstruction string

// PowerConsumptionChange is Change in power consumption: maximise, minimise or "no advice" ("no advice" cancels a previous advice)
type PowerConsumptionChange string

// PreciseDelayReason is Precise reason of the delay
type PreciseDelayReason string

// ProtectedLevelCrossing is Is the level crossing protected
type ProtectedLevelCrossing bool

// RangeMaxSpeed is Maximum speed in the range of speed to be avoided
type RangeMaxSpeed float64

// RangeMinSpeed is Minimum speed in the range of speed to be avoided
type RangeMinSpeed float64

// ReasonCode is reasonCode offers coding for standard situations. If reasonCode is null the last reasonCode received is still valid.
// It is possible to use only one of the parameters reasonCode and reasonText
// or to send them in combination. Predefined reasonCodes are:
// - followTrain: host train is following a conflicting (slower) train
// - trainFollowing: host train is preceding a conflicting (faster) train
// - merge2nd: conflicting (faster) train merging ahead and entering common route section first
// - merge1st: conflicting (slower) train merging ahead, host train enters common route section first
// - overtakeTrain: overtake conflicting (slower) train that has (nearly) reached the location of overtaking
// - followTrainOvertakeTrain: follow, then overtake conflicting train that hasn’t yet reached the location of overtaking
// - beingOvertaken: conflicting (faster) train about to overtake
// - trainFollowingBeingOvertaken: conflicting (faster) train following and then overtaking
// - trainCrossing: train crossing
// - passTrain: passing a train on a single-track line
// - beingPassed: host train being passed on a single-track line
// - timeSupplementRequired: recovery time is required due to an upcoming temporary speed restriction
// - energyOptimisation: optimisation of energy consumption
// - endOfReason: the reason previously communicated is not valid anymore.
type ReasonCode string

// ReasonText is Human readable text with the reason why this advice message is sent (e.g. describing the conflict or conflict resolution).
// It is possible to use only one of the parameters reasonCode and reasonText
// or to send them in combination.
type ReasonText string

// ReferenceIDAbs is References the ID of the message, that defined the absolute speed, which is still valid
type ReferenceIDAbs string

// RelatedTrainRequest is Request for related train information
type RelatedTrainRequest string

// RelatedTrainSupported is Information on whether related train information is supported
type RelatedTrainSupported string

// RelaxedCoupler is [SUBSET-126] request to the train to stop with couplers relaxed
type RelaxedCoupler bool

// RemainingEnginePower is Percentage of power available due to engine problems
type RemainingEnginePower float64

// ReportingTime is [SUBSET-126] Reporting time cycle for triggering a Status Report for data reporting purposes
type ReportingTime string

// RestrictionReason is Reason why the related train information is restricted (e.g. missing train location system, or train in another country)
type RestrictionReason string

// Result is Result of the message
type Result string

// RollingStockType is Free-form text describing the rolling stock. Used for driver information.
type RollingStockType string

// RolloutCoefficientA is Description of how the train accelerates. Coefficient used in order to calculate the resistance force (constant value). Expressed in N
type RolloutCoefficientA float64

// RolloutCoefficientB is Description of how the train accelerates if no traction or braking is applied. Coefficient used in order to calculate the resistance force (value to be multiplied with v). Expressed in N.s/m
type RolloutCoefficientB float64

// RolloutCoefficientC is Description of how the train accelerates if no traction or braking is applied.Coefficient used in order to calculate the resistance force (value to be multiplied with v^2). Expressed in N.s^2/m^2
type RolloutCoefficientC float64

// RotatingMassFactor is Rotating Mass Factor. The rotating mass of the train is equivalent to the rotational inertia of the train, transformed by gear ratio and wheel diameter respectively.
type RotatingMassFactor float64

// RoutingError is [SUBSET-126] ATO: Routing Error, from Subset-126 (boolean)
type RoutingError bool

// SessionTerminationReason is Reason for session termination
type SessionTerminationReason string

// SetValueTractionBraking is set value by the driver for traction and braking
type SetValueTractionBraking string

// SFERAVersionMajor is SFERA version, major number. There should be backwards compatibility for messages with the same major SFERA version.
type SFERAVersionMajor string

// SFERAVersionMinor is SFERA version, minor number.
type SFERAVersionMinor string

// SignalIDObject is Identifier of the signal object
type SignalIDObject string

// SignalIDPhysical is Identifier of the physical signal on the track
type SignalIDPhysical string

// SignalType is Type of signal from which the physical characteristics of the signal can be deduced. This will be network specific, as no standardised solution has been identified by the SFERA project. Examples:
// - CTC (Centralized Traffic Control Signal)
// - Permissive (Permissive Signal)
// - Permanent signal
// - Maximum Sign
// - Announcement Sign
// - Disque
// - Cantonnement BAL 3 aspects
type SignalType string

// SlipSlideReported is [SUBSET-126] ATO: Slip/slide reported by TCMS/Train, from Subset-126 (boolean)
type SlipSlideReported bool

// SourceCompany is ID of the company sending the message. Must be a UIC RICS Code: : https://uic.org/rics
type SourceCompany int

// SourceDevice is ID of the device sending the message. For any given company, this ID must be unique. Maximum 128 characters.
type SourceDevice string

// SPAltitude is [SUBSET-126] Altitude (in m) at the beginning of the SP. Considering ETRS89 as reference. Starting at
// -1000m.
type SPAltitude float64

// SPDirection is Travelling direction of the Segment Profile (Nominal or Reverse). If the travelling direction is Reverse, all locations will apply from the end of the segment. Note that Signals and Stopping Points only apply in the nominal direction of the SP.
type SPDirection string

// SPGauge is Gauge of the segment profile (in mm). Is needed to calculate rolling resistance.
type SPGauge int

// SPID is Segment Profile Identifier
type SPID string

// SPLength is Length of the section of railway covered by the Segment Profile (m)
type SPLength float64

// SPName is Human-readable name of the Segment Profile. To be used only for debugging purposes.
type SPName string

// SPStatus is [SUBSET-126] Status of the Segment Profile (Valid or Invalid). From SUBSET-125: 10.1.3.23 The Segment Profile shall indicate its status to the ATO-OB (Segment Profile Status):
// a) Valid: specifies that the data sent correspond to a SPReq sent;
// b) Invalid: the requested Segment Profile is not available in the ATO-TS.
type SPStatus string

const (
	SPStatusValid   = "Valid"
	SPStatusInvalid = "Invalid"
)

// SPVersionMajor is Identifier of the Segment Profile version, major number.
type SPVersionMajor uint32

// SPVersionMinor is Identifier of the Segment Profile version, minor number.
type SPVersionMinor uint32

// SpecificSSPCategory is [SUBSET-126] Indicates the type of specific SSP category.
type SpecificSSPCategory string

// Speed is speed (km/h)
type Speed float64

// SSPFront is [SUBSET-126] Qualifier to indicate if the Static Speed Profile is to be applied to the front of the train (1, no train length delay) or to the end of the train (0, train length delay).
type SSPFront bool

// SSPNCCDDIFF is [SUBSET-126] (If specific_SSP_Category = Cant_Deficiency_SSP)
//
//	“Cant Deficiency” SSP category for which a different value for the static line speed exists. Used together with V_DIFF to permit certain trains to go faster or lower than the “international basic static speed” given by SSP_Speed.
type SSPNCCDDIFF string

// SSPNCDIFF is [SUBSET-126] (If specific_SSP_Category = Other_SSP_*)
// It is the “other specific” SSP category for which a different value for the static line speed exists. Used together with SS126_V_DIFF to permit trains belonging to the corresponding “other international” train
// category to go faster or lower than the “international basic static speed” given by SSP_Speed.
type SSPNCDIFF string

// SSPSpeed is [SUBSET-126] Basic Static Speed Profile speed.
type SSPSpeed float64

// SSPVisualReference is Visual reference of a speed limit threshold (e.g. a building along the line). This attribute shall be used only if the object is not already in the Segment Profile.
type SSPVisualReference string

// StartEndQualifier is [SUBSET-126] Specifies if the element starts, ends, starts and ends or covers the whole concerning Segment Profile
type StartEndQualifier string

// StartKm is References the kilometre at the beginning of the message, preferably linked with a kilometer sign near the track
type StartKm float64

// StartLine is References the (national) line number at the beginning of the message
type StartLine int

// StartLocation is [SUBSET-126] Distance of the element's start location from the beginning of the Segment Profile (m)
type StartLocation float64

// StartPointOfDownLineSpeedApplication is How far (in m) after the considered signal the speed advice must be respected. 0 = As soon as possible
type StartPointOfDownLineSpeedApplication float64

// StartSignal is References the name of the start signal of the message, if the start of the message is at or near a signal
type StartSignal string

// StartSignalApplicationMustBeClearedBy is What part of the train must clear the signal application before applying the signal
type StartSignalApplicationMustBeClearedBy string

// StartStationLong is References the start station at the beginning of the message (long form)
type StartStationLong string

// StartStationShort is References the start station at the beginning of the message (short form)
type StartStationShort string

// StartTime is Starting date and time (in UTC)
type StartTime string

// StartValidity is Begin of the advice validity. Either both startValidity and endValidity or none of the two must be specified. Both parameters should refer to a future point in time and endValidity must be after startValidity. If either of the two indicates a past time (w.r.t. the clock of the On-Board Unit), which is possible in a situation where Central Unit and OBU are not synchronised, transmission delays etc., then the OBU should treat it as valid but past time indication.
type StartValidity string

// StationTrack is If the Timing Point is a station: identifier of the track where the train will stop or pass through without stopping at the station.
// Examples : “1” “14” “2M” “1bis”
type StationTrack string

// StatusReportsEnabled is Flag notifying whether the DAS is able to send Status Reports to ground.
type StatusReportsEnabled bool

// StopLocationTolerance is [SUBSET-126] Required stopping tolerance (m)
type StopLocationTolerance float64

// StopTime is Stopping time necessary to complete the procedure, only given if StopBeforePassingSignal is “yes”. The time given is the safety procedure time, and is independent from the commercial stop time.
type StopTime string

// StopTypeDetails is Additional details about the stop type
type StopTypeDetails string

// StopTypePurpose is Type of stop and purpose. Stop types can be:
// - Commercial: a stop used by the RU to perform the commercial duties the train was planned for. For a passenger train, this will concern stops where passenger loading or unloading is done. For a freight train, this will concern stops where loading and unloading is done, or when wagons are either added or uncoupled from the train.
// - Non Commercial Regulation: stops integrated for the purpose of guaranteeing the fluidity of traffic.
// - Non Commercial RU: stops used by the RU to fulfil its specific needs. For example, a train driver relay, or a locomotive change could be characterised by this purpose;
// - Non Commercial IM: stops necessary for the IM’s needs;
// - Non Commercial Safety: stops used to perform a safety procedure;
// - Other: any stop that cannot be categorised using one of the 5 other types.
type StopTypePurpose string

// SuperElevationValue is Value of superelevation (mm)
type SuperElevationValue float64

// TCID is Train Characteristics Identifier
type TCID string

// TCVersionMajor is Identifier of the Train Characteristics version, major number.
type TCVersionMajor uint32

// TCVersionMinor is Identifier of the Train Characteristics version, minor number.
type TCVersionMinor uint32

// TemporaryConstraintReference is Human-readable identifier of document informing driver of temporary constraint
type TemporaryConstraintReference interface{}

// TemporaryConstraintType is Reason for the temporary constraint. It is possible to choose one reason among the enumerations; the detailed data has to be put under the related element in the TemporaryConstraint choice. Exceptions are ATO_Inhibition_Zone, DAS_Inhibition_Zone and IM_SDAS_Only, where no additional data is needed.
type TemporaryConstraintType string

// TemporarySignalConstraintType is Type of temporary signal constraints: Temporary Signal (the identifier must then be new, not used by another element), Existing Signal Aspect Modification (the identifier must point to a Signal described in SP / SP_Points if it is present) or Cancelled Signal (the identifier must point to a Signal described in SP / SP_Points).
type TemporarySignalConstraintType string

// TFCEstimatedVoltage is This TractionForceCurve is to be used if the estimated voltage of the considered portion of infrastructure is inferior to this value and superior to a lower estimatedVoltage of another TractionForceCurve (if it exists).
type TFCEstimatedVoltage int

// TFCMaxCurrentValue is This TractionForceCurve is to be used if the maximum current value of the considered portion of infrastructure is inferior to this value and superior to a lower maxCurrentValue of another TractionForceCurve (if it exists).
type TFCMaxCurrentValue float64

// TimeoutATOTSResponse is [SUBSET-126] Minimum time the ATO-OB shall wait before repeating a request if there was no response
type TimeoutATOTSResponse string

// Timestamp is Timestamp. Must be in UTC.
type Timestamp string

// TPAbbreviation is Abbreviation of the Timing Point name
type TPAbbreviation string

// TPAccurateStopping is [SUBSET-126] (If TP_Pass_Stop_Depart = Stopped)
// This qualifier specifies if the train has stopped accurately or not at the Operational Stopping Point.
type TPAccurateStopping string

// TPAlignment is [SUBSET-126] Defines if the TP location is applicable from the front, middle or rear of the train
type TPAlignment string

// TPID is Timing Point Identifier
type TPID string

// TPInformation is [SUBSET-126] Specifies some information specific for the TP
type TPInformation string

// TPLatestArrivalTime is Time of arrival at Timing Point expressed in UTC. Has to be specified always except at departure stations.
type TPLatestArrivalTime string

// TPPassStopDepart is [SUBSET-126] Qualifier to indicate if train has stopped at the TP, has departed from the TP or has passed the TP.
type TPPassStopDepart string

// TPPlannedLatestArrivalTime is The original planned latest arrival time at the Timing Point expressed in UTC. Useful for driver information. Not to be used for calculations.
type TPPlannedLatestArrivalTime string

// TPPriority is TP_Priority is used to indicate the relative significance of the Timing Point in a Journey from the driver’s point of view. This information is intended to be used when choosing the information to display. It does not specify a priority in respecting the timing points when the DAS calculates the speed profile (the attributes TP_latestArrivalTime and arrivalWindow in the Journey Profile are the absolute time constraints on timing points). It has 5 possible values :
// 5: very important station, junction or network border
// 4: important station, junction or shunting area
// 3: station, junction or regional border
// 2: other stopping locations
// 1: all other timing points
type TPPriority uint8

// TPStopSkipPass is [SUBSET-126] Specifies if the Timing Point is a Passing Point, an operational Stopping Point or a skipped Stopping Point
type TPStopSkipPass string

// TPType is Type of Timing Point. Examples:
// Very important station
// Important station
// Station
// Stopping location
// Junction
// Regional border
// Connexion
// Network border
// Signal with detection
// Grid or shunting area
// Service installation
// Service stopping location
// Movable bridge
// Other
type TPType string

// TrackIdentifier is Identification of track on a line (ex : “V1” or “Voie 1” or “Line 1”)
type TrackIdentifier interface{}

// TractionForce is Traction Force (in N)
type TractionForce float64

// TractionReactionTime is Delay between traction command and start of tractioning. Expressed in s
type TractionReactionTime int

// TractionType is Traction Type (DC/AC with voltage, self-propelled or powerless)
type TractionType string

// TrainCategoryCode is Code for the train category (freeform text, limited to 10 characters)
type TrainCategoryCode string

// TrainHold is [SUBSET-126] The variable defines if the train is requested to be held at the Stopping Point (1) or not (0)
type TrainHold bool

// TrainLength is The total length of a train. Expressed in m
type TrainLength int

// TrainMaxSpeed is The maximum possible speed of the train in km/h
type TrainMaxSpeed int

// TrainMoving is [SUBSET-126] ATO: Train is moving, from Subset-126 (boolean)
type TrainMoving bool

// TrainWeight is The sum of all weights of wagons and traction units. Expressed in tonnes (t)
type TrainWeight int

// TrpAe is Tunnel Resistance Prediction coefficient a_e (for engines). Unit: kg.m^(2b_e-1). Typical value 1140. Reference: https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
type TrpAe float64

// TrpAw is Tunnel Resistance Prediction coefficient a_w (for wagons). Unit: kg.m^(2b_w-1). Typical value 662. Reference: https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
type TrpAw float64

// TrpBe is Tunnel Resistance Prediction coefficient b_e (for engines). No unit. Typical value 1.48. Reference: https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
type TrpBe float64

// TrpBw is Tunnel Resistance Prediction coefficient b_w (for wagons). No unit. Typical value 1.75. Reference: https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
type TrpBw float64

// TunnelCategory is [SUBSET-126] Category of the Tunnel (single or double track)
type TunnelCategory string

// TunnelFactor is Aerodynamic cross section of the tunnel in expressed in m^2. Used in predictive tunnel resistance equation developed by Z. Filipovic. See https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
type TunnelFactor float64

// UnableToStop is [SUBSET-126] ATO: Unable to stop at the next Stopping Point, from Subset-126 (boolean)
type UnableToStop bool

// UTCOffset is [SUBSET-126] Offset to add to the UTC time in order to calculate the local time. Unsigned value
// Resolution: 15 min
// 0 = UTC - 14:00
// 56 = UTC ± 0
// 112 = UTC + 14:00
// 113 - 127 = Spare
type UTCOffset uint8

// VDIFF is [SUBSET-126] Absolute Positive Speed associated to a train category
type VDIFF float64

// Value is Value of the element
type Value string

// VerticalOffset is Vertical offset (in m) from the top of the track to the middle of the object.
type VerticalOffset float64

// VoltageValue is Voltage value (V)
type VoltageValue int

// Wagons is Number of wagons
type Wagons int

// WindHeading is Wind heading, in degrees from North
type WindHeading uint16

// WindSpeed is Wind speed
type WindSpeed float64

// XPath is XML Path to identify the location of the error
type XPath xml.Name
