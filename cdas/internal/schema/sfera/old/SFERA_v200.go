// Code generated from SFERA_v200.xsd. DO NOT EDIT BY HAND.
//
// This file provides a Go object model for the SFERA (Signalling and
// interlocking Functional Exchange for Rail Automation) v2.00 XML
// interface, built on encoding/xml.
//
// Roundtrip design note: every XSD scalar leaf (attribute value or
// element text content) is represented as a Go *defined string type*
// rather than a native bool/int/float64/time.Time. encoding/xml never
// reformats string data, so Unmarshal followed by Marshal reproduces
// the original values byte-for-byte (values, not whitespace/attribute
// order, which XML treats as insignificant). Convenience parse/format
// helpers are provided at the bottom of this file for callers that want
// typed access (bool, int64, float64, time.Time).
package old

// import (
// 	"encoding/xml"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// // ---------------------------------------------------------------------------
// // Simple types (XSD simpleType / scalar element and attribute value types)
// // ---------------------------------------------------------------------------

// // AdhesionCategoryType is a scalar type (XSD base: xs:string). Simple Type for adhesion category
// type AdhesionCategoryType string

// // Enumerated values for AdhesionCategoryType.
// const (
// 	AdhesionCategoryTypeDryRail              AdhesionCategoryType = "Dry Rail"
// 	AdhesionCategoryTypeDryRailMedium        AdhesionCategoryType = "Dry Rail (Medium)"
// 	AdhesionCategoryTypeDryRailLow           AdhesionCategoryType = "Dry Rail (Low)"
// 	AdhesionCategoryTypeLowAdhesion          AdhesionCategoryType = "Low Adhesion"
// 	AdhesionCategoryTypeVeryLowAdhesion      AdhesionCategoryType = "Very Low Adhesion"
// 	AdhesionCategoryTypeExtremelyLowAdhesion AdhesionCategoryType = "Extremely Low Adhesion"
// )

// // CompanyCode is a scalar type (XSD base: Numeric4-4). From TAP/TAF: Identifies the RU, IM or other
// // company involved in the Rail Transport Chain
// type CompanyCode string

// // CountryIdentISO is a scalar type (XSD base: xs:string). ISO 3166-1 alpha code (2 positions)
// type CountryIdentISO string

// // CurrentType is a scalar type (XSD base: xs:decimal). Simple Type for current (A)
// type CurrentType string

// // DelayCode is a scalar type (XSD base: xs:token). Reason for a delay or interruption. UIC Leaflet
// // 450-2, Appendix C. The first digit in the code has to following meaning: 1 Operational planning,
// // Management (IM) 2 Infrastructure installations (IM) 3 Civil engineering causes (IM) 4 Causes of
// // other IM (IM) 5 Commercial causes (RU) 6 Rolling stock (RU) 7 Causes of other RU (RU) 8 External
// // causes 9 Secondary causes New codes added: 23 Power supply equipment 58 Staff 68 Staff 90
// // Dangerous incidents, accidents and hazards 91 Track occupation caused by the lateness of the same
// // train 92 Track occupation caused by the lateness of another train 93 Turn round 94 Connection 95
// // Further investigation needed
// type DelayCode string

// // Enumerated values for DelayCode.
// const (
// 	DelayCode11 DelayCode = "11"
// 	DelayCode10 DelayCode = "10"
// 	DelayCode12 DelayCode = "12"
// 	DelayCode13 DelayCode = "13"
// 	DelayCode14 DelayCode = "14"
// 	DelayCode18 DelayCode = "18"
// 	DelayCode19 DelayCode = "19"
// 	DelayCode20 DelayCode = "20"
// 	DelayCode21 DelayCode = "21"
// 	DelayCode22 DelayCode = "22"
// 	DelayCode23 DelayCode = "23"
// 	DelayCode24 DelayCode = "24"
// 	DelayCode25 DelayCode = "25"
// 	DelayCode28 DelayCode = "28"
// 	DelayCode29 DelayCode = "29"
// 	DelayCode31 DelayCode = "31"
// 	DelayCode30 DelayCode = "30"
// 	DelayCode32 DelayCode = "32"
// 	DelayCode39 DelayCode = "39"
// 	DelayCode40 DelayCode = "40"
// 	DelayCode41 DelayCode = "41"
// 	DelayCode50 DelayCode = "50"
// 	DelayCode51 DelayCode = "51"
// 	DelayCode52 DelayCode = "52"
// 	DelayCode53 DelayCode = "53"
// 	DelayCode54 DelayCode = "54"
// 	DelayCode58 DelayCode = "58"
// 	DelayCode59 DelayCode = "59"
// 	DelayCode60 DelayCode = "60"
// 	DelayCode61 DelayCode = "61"
// 	DelayCode62 DelayCode = "62"
// 	DelayCode63 DelayCode = "63"
// 	DelayCode64 DelayCode = "64"
// 	DelayCode68 DelayCode = "68"
// 	DelayCode70 DelayCode = "70"
// 	DelayCode69 DelayCode = "69"
// 	DelayCode71 DelayCode = "71"
// 	DelayCode80 DelayCode = "80"
// 	DelayCode81 DelayCode = "81"
// 	DelayCode82 DelayCode = "82"
// 	DelayCode83 DelayCode = "83"
// 	DelayCode84 DelayCode = "84"
// 	DelayCode89 DelayCode = "89"
// 	DelayCode90 DelayCode = "90"
// 	DelayCode91 DelayCode = "91"
// 	DelayCode92 DelayCode = "92"
// 	DelayCode93 DelayCode = "93"
// 	DelayCode94 DelayCode = "94"
// 	DelayCode95 DelayCode = "95"
// )

// // DeltaTime is a scalar type (XSD base: xs:string). Time difference delay (+) or ahead of schedule
// // (-) this shall be 1character + 4 Numeric
// type DeltaTime string

// // DeviceType is a scalar type (XSD base: xs:string). Simple type for DAS devices
// type DeviceType string

// // DistanceType is a scalar type (XSD base: xs:double). Simple Type for distance attributes
// type DistanceType string

// // EnergyType is a scalar type (XSD base: xs:decimal). Simple Type for energy (kWh)
// type EnergyType string

// // ForceType is a scalar type (XSD base: xs:decimal). Force Type
// type ForceType string

// // FreeText is a scalar type (XSD base: xs:string). Clear Text in ISO Unicode character set
// type FreeText string

// // GradientDirectionType is a scalar type (XSD base: xs:string). Type for gradient direction
// // (Downhill or Uphill)
// type GradientDirectionType string

// // Enumerated values for GradientDirectionType.
// const (
// 	GradientDirectionTypeDownhill GradientDirectionType = "Downhill"
// 	GradientDirectionTypeUphill   GradientDirectionType = "Uphill"
// )

// // GradientValueType is a scalar type (XSD base: xs:decimal). Type for gradient values (Numeric
// // value, mm by m)
// type GradientValueType string

// // NC_CDTrainType is a scalar type (XSD base: xs:string). Simple Type for Cant Deficiency Train
// // Category (NC_CDTRAIN as defined in SUBSET-026).
// type NC_CDTrainType string

// // Enumerated values for NC_CDTrainType.
// const (
// 	NC_CDTrainType80mm  NC_CDTrainType = "80mm"
// 	NC_CDTrainType100mm NC_CDTrainType = "100mm"
// 	NC_CDTrainType130mm NC_CDTrainType = "130mm"
// 	NC_CDTrainType150mm NC_CDTrainType = "150mm"
// 	NC_CDTrainType165mm NC_CDTrainType = "165mm"
// 	NC_CDTrainType180mm NC_CDTrainType = "180mm"
// 	NC_CDTrainType210mm NC_CDTrainType = "210mm"
// 	NC_CDTrainType225mm NC_CDTrainType = "225mm"
// 	NC_CDTrainType245mm NC_CDTrainType = "245mm"
// 	NC_CDTrainType275mm NC_CDTrainType = "275mm"
// 	NC_CDTrainType300mm NC_CDTrainType = "300mm"
// )

// // NC_TrainType is a scalar type (XSD base: xs:string). Simple Type for Train Category (NC_TRAIN as
// // defined in SUBSET-026).
// type NC_TrainType string

// // Enumerated values for NC_TrainType.
// const (
// 	NC_TrainTypeFreightTrainPPosition NC_TrainType = "Freight_Train_P_Position"
// 	NC_TrainTypeFreightTrainGPosition NC_TrainType = "Freight_Train_G_Position"
// 	NC_TrainTypePassengerTrain        NC_TrainType = "Passenger_Train"
// )

// // NID_C_Type is a scalar type (XSD base: xs:short). Simple type for NID_C
// type NID_C_Type string

// // Numeric1_5 is a scalar type (XSD base: xs:positiveInteger).
// type Numeric1_5 string

// // Numeric2_2 is a scalar type (XSD base: xs:integer).
// type Numeric2_2 string

// // Numeric4_4 is a scalar type (XSD base: xs:integer). Numeric value with 4 digits
// type Numeric4_4 string

// // PercentageType is a scalar type (XSD base: xs:decimal). Percentage Type
// type PercentageType string

// // PowerTypekW is a scalar type (XSD base: xs:double). Simple Type for power (kW)
// type PowerTypekW string

// // PowerTypeW is a scalar type (XSD base: xs:double). Simple Type for power (W)
// type PowerTypeW string

// // RelatedTrainRequestType is a scalar type (XSD base: xs:string).
// type RelatedTrainRequestType string

// // Enumerated values for RelatedTrainRequestType.
// const (
// 	RelatedTrainRequestTypeNone                       RelatedTrainRequestType = "None"
// 	RelatedTrainRequestTypeOwnTrain                   RelatedTrainRequestType = "OwnTrain"
// 	RelatedTrainRequestTypeRelatedTrains              RelatedTrainRequestType = "RelatedTrains"
// 	RelatedTrainRequestTypeOwnTrainAndRelatedTrains   RelatedTrainRequestType = "OwnTrainAndRelatedTrains"
// 	RelatedTrainRequestTypeOwnTrainAndOrRelatedTrains RelatedTrainRequestType = "OwnTrainAndOrRelatedTrains"
// )

// // SpeedType is a scalar type (XSD base: xs:decimal). Speed Type
// type SpeedType string

// // String1_10 is a scalar type (XSD base: xs:string).
// type String1_10 string

// // TractionTypeType is a scalar type (XSD base: xs:string). Simple Type for TractionTypes
// type TractionTypeType string

// // Enumerated values for TractionTypeType.
// const (
// 	TractionTypeType600VDC         TractionTypeType = "600V DC"
// 	TractionTypeType750VDC         TractionTypeType = "750V DC"
// 	TractionTypeType1500VDC        TractionTypeType = "1500V DC"
// 	TractionTypeType3000VDC        TractionTypeType = "3000V DC"
// 	TractionTypeType25000VAC50Hz   TractionTypeType = "25000V AC 50Hz"
// 	TractionTypeType15000VAC16Hz23 TractionTypeType = "15000V AC 16Hz2/3"
// 	TractionTypeTypeSelfPropelled  TractionTypeType = "Self-propelled"
// 	TractionTypeTypePowerless      TractionTypeType = "Powerless"
// )

// // TrainPartType is a scalar type (XSD base: xs:string). Simple type for train part (head, end or
// // last pantograph). To be used for defining which part of the train has to clear a signal for
// // application.
// type TrainPartType string

// // Enumerated values for TrainPartType.
// const (
// 	TrainPartTypeTrainHead           TrainPartType = "TrainHead"
// 	TrainPartTypeTrainEnd            TrainPartType = "TrainEnd"
// 	TrainPartTypeTrainLastPantograph TrainPartType = "TrainLastPantograph"
// )

// // UTCDateTimeType is a scalar type (XSD base: xs:dateTime). Simple Type for date and time in UTC
// type UTCDateTimeType string

// // UUIDv4_Type is a scalar type (XSD base: xs:string). UUIDv4 (random) simpleType as defined in
// // RFC4122
// type UUIDv4_Type string

// // VoltageType is a scalar type (XSD base: xs:int). Simple Type for voltage (V)
// type VoltageType string

// // WeightValueKilo is a scalar type (XSD base: xs:integer). In Kilograms
// type WeightValueKilo string

// // ClaimType is a scalar type (XSD base: xs:string). Gives an indication of the status or occupation
// // of the piece of the segment
// type ClaimType string

// // Enumerated values for ClaimType.
// const (
// 	ClaimTypePlanned  ClaimType = "Planned"
// 	ClaimTypeReserved ClaimType = "Reserved"
// 	ClaimTypeOccupied ClaimType = "Occupied"
// 	ClaimTypeUnknown  ClaimType = "Unknown"
// )

// // Core is a scalar type (XSD base: xs:string). It is the main part of identifier and is determined
// // by the company that creates it.
// type Core string

// // KeeperShortNameVKM is a scalar type (XSD base: xs:string). Free text, short name/vehicle keeper
// // marking of the wagon keeper. Use ERA VKM Register: https://www.era.europa.eu/registers/vkm_en
// type KeeperShortNameVKM string

// // LeavingEnteringStatus is a scalar type (XSD base: xs:string). Gives an indication if a train is
// // entering of leaving your trainpath at a certain location
// type LeavingEnteringStatus string

// // Enumerated values for LeavingEnteringStatus.
// const (
// 	LeavingEnteringStatusEntering LeavingEnteringStatus = "Entering"
// 	LeavingEnteringStatusLeaving  LeavingEnteringStatus = "Leaving"
// 	LeavingEnteringStatusCrossing LeavingEnteringStatus = "Crossing"
// )

// // MessageType is a scalar type (XSD base: xs:string). To indicate the message type transmitted or
// // referred to. The following list was agreed within the sector: 1000 ConsignmentOrderMessage 2001
// // PathCanceledMessage 2002 PathConfirmedMessage 2003 PathDetailsMessage 2004
// // PathDetailsRefusedMessage 2005 PathNotAvailableMessage 2006 PathRequestMessage 2007
// // ReceiptConfirmationMessage -- sector messages (Planning) -- 2500 PathCoordinationMessage 2501
// // PathSectionNotificationMessage -- sector message end -- 3003 TrainCompositionMessage 3006
// // TrainReadyMessage 4001 TrainDelayCauseMessage 4004 TrainRunningForecastMessage 4005
// // TrainRunningInformationMessage 4006 TrainRunningInterruptionMessage -- sector message
// // (Operations) -- 4500 PassengerTrainCompositionProcessMessage 4501 RollingStockRestrictionMessage
// // 4504 ChangeOfTrackMessage 4505 TrainJourneyModificationMessage -- sector message end -- 5001
// // AlertMessage 5002 WagonArrivalNoticeMessage 5003 WagonDeliveryNoticeMessage 5004
// // WagonDepartureNoticeMessage 5006 WagonETI_ETA_Message 5007 WagonExceptionMessage 5008
// // WagonExceptionReasonMessage -- sector message (Wagon interchange) -- 5009
// // WagonInterchangeNoticeMessage 5012 WagonReceivedAtInterchangeMessage 5013
// // WagonRefusedAtInterchangeMessage -- sector message end -- 5014 WagonReleaseNoticeMessage 5015
// // WagonYardArrivalMessage 5016 WagonYardDepartureMessage 6002 LocationFileDatasetMessage 6003
// // RollingStockDatasetMessage --sector (RU-RU) --- 5500 WagonPerformanceMessage -- sector end---
// // 6004 RollingStockDatasetQueryMessage -- Sector Messages TCR begin -- 6500 TCRMessage 6501
// // TCRResponseMessage 6502 TCRCanceledMessage -- Sector Messages TCR end -- -- sector (TrainID)
// // begin -- 8500 UpateLinkMessage 8501 ObjectInfoMessage -- sector end -- 9000 ErrorMessage
// type MessageType string

// // MessageTypeVersion is a scalar type (XSD base: xs:string). Version of the Message Type
// type MessageTypeVersion string

// // ObjectType is a scalar type (XSD base: xs:string). Provides a possibility for differentiantion
// // between the objects: Train, Path, Case Reference and Path Request. In this case, must be TR
// // (Train)
// type ObjectType string

// // Enumerated values for ObjectType.
// const (
// 	ObjectTypeTR ObjectType = "TR"
// 	ObjectTypePA ObjectType = "PA"
// 	ObjectTypeCR ObjectType = "CR"
// 	ObjectTypePR ObjectType = "PR"
// )

// // SignalFunction is a scalar type (XSD base: xs:string). Function of the signal. The list of values
// // is taken from EULynx SignalFunctionTypes
// // (https://dataprep.eulynx.eu/2021-05/index.htm?goto=3:1:1:4:3762), with the addition of
// // laneChange.
// type SignalFunction string

// // Enumerated values for SignalFunction.
// const (
// 	SignalFunctionEntry                      SignalFunction = "entry"
// 	SignalFunctionExit                       SignalFunction = "exit"
// 	SignalFunctionIntermediate               SignalFunction = "intermediate"
// 	SignalFunctionAdvance                    SignalFunction = "advance"
// 	SignalFunctionBlock                      SignalFunction = "block"
// 	SignalFunctionProtection                 SignalFunction = "protection"
// 	SignalFunctionTrainProtection            SignalFunction = "trainProtection"
// 	SignalFunctionGroupMain                  SignalFunction = "groupMain"
// 	SignalFunctionGroupIntermediate          SignalFunction = "groupIntermediate"
// 	SignalFunctionGroupMainAndIntermediate   SignalFunction = "groupMainAndIntermediate"
// 	SignalFunctionStandAloneIndicator        SignalFunction = "standAloneIndicator"
// 	SignalFunctionTrainDestinationOnlySignal SignalFunction = "trainDestinationOnlySignal"
// 	SignalFunctionIntermediateExitSignal     SignalFunction = "intermediateExitSignal"
// 	SignalFunctionEntryAndExit               SignalFunction = "entryAndExit"
// 	SignalFunctionEntryAndBlock              SignalFunction = "entryAndBlock"
// 	SignalFunctionFurther                    SignalFunction = "further"
// 	SignalFunctionLxProtectionSignal         SignalFunction = "lxProtectionSignal"
// 	SignalFunctionDistantOrRepeater          SignalFunction = "distantOrRepeater"
// 	SignalFunctionLaneChange                 SignalFunction = "laneChange"
// )

// // TimetableYear is a scalar type (XSD base: xs:integer). Refers to the timetable period in which
// // the business will be carried out (Not used in SFERA. Only kept for consistency with data elements
// // in TAP/TAF)
// type TimetableYear string

// // TrainReady is a scalar type (XSD base: xs:integer). 0=Not Ready 1=Ready
// type TrainReady string

// // Enumerated values for TrainReady.
// const (
// 	TrainReady0 TrainReady = "0"
// 	TrainReady1 TrainReady = "1"
// )

// // Variant is a scalar type (XSD base: xs:string). The variant shows a relationship between two
// // identifiers referring to the same business case
// type Variant string

// // WagonMaxSpeed is a scalar type (XSD base: xs:int). Maximum allowed speed of the wagon according
// // to the load and entry in the Rolling Stock Databases. In kmh
// type WagonMaxSpeed string

// // WagonTrainPosition is a scalar type (XSD base: xs:int). Identifies the position of a wagon within
// // a train. Sequential number starting with the first wagon at the front of train as N°1.
// type WagonTrainPosition string

// // ---------------------------------------------------------------------------
// // Complex types and message elements
// // ---------------------------------------------------------------------------

// // AdditionalSpeedRestriction: Additional Speed Restriction
// type AdditionalSpeedRestriction struct {
// 	XMLName xml.Name `xml:"AdditionalSpeedRestriction"`
// 	// [SUBSET-126] Qualifier to indicate if a speed limit given for a profile element is to be applied
// 	// until the front of the train (1, no train length delay) or the end of the train (0, train length
// 	// delay) has left the element
// 	ASR_Front string `xml:"ASR_Front,attr,omitempty"`
// 	// [SUBSET-126] Value of the speed level restriction (km/h)
// 	ASR_Speed SpeedType `xml:"ASR_Speed,attr,omitempty"`
// }

// // AdditionalTrainNumber: Additional identifier for train number
// type AdditionalTrainNumber struct {
// 	XMLName xml.Name `xml:"AdditionalTrainNumber"`
// }

// // AdhesionConditionsChange: Change in adhesion conditions
// type AdhesionConditionsChange struct {
// 	XMLName xml.Name `xml:"AdhesionConditionsChange"`
// 	// New adhesion conditions
// 	NewAdhesionConditions string `xml:"newAdhesionConditions,attr"`
// 	// Adhesion category aligned with the definition in EN15595 of the wheel-rail interface for the
// 	// lower boundaries (definition and values taken from SUBSET-126).
// 	LowAdhesionCategory AdhesionCategoryType `xml:"lowAdhesionCategory,attr,omitempty"`
// }

// // AdvisedSpeed: Explicit speed advice
// type AdvisedSpeed struct {
// 	XMLName xml.Name `xml:"AdvisedSpeed"`
// 	// speed (km/h)
// 	Speed SpeedType `xml:"speed,attr,omitempty"`
// 	// Maximum advised speed. Must be higher than "speed".
// 	MaxAdvisedSpeed SpeedType `xml:"maxAdvisedSpeed,attr,omitempty"`
// 	// Minimum advised speed. Must be lower than "speed".
// 	MinAdvisedSpeed SpeedType `xml:"minAdvisedSpeed,attr,omitempty"`
// 	// reasonCode offers coding for standard situations. If reasonCode is null the last reasonCode
// 	// received is still valid. It is possible to use only one of the parameters reasonCode and
// 	// reasonText or to send them in combination. Predefined reasonCodes are: - followTrain: host train
// 	// is following a conflicting (slower) train - trainFollowing: host train is preceding a
// 	// conflicting (faster) train - merge2nd: conflicting (faster) train merging ahead and entering
// 	// common route section first - merge1st: conflicting (slower) train merging ahead, host train
// 	// enters common route section first - overtakeTrain: overtake conflicting (slower) train that has
// 	// (nearly) reached the location of overtaking - followTrainOvertakeTrain: follow, then overtake
// 	// conflicting train that hasn’t yet reached the location of overtaking - beingOvertaken:
// 	// conflicting (faster) train about to overtake - trainFollowingBeingOvertaken: conflicting
// 	// (faster) train following and then overtaking - trainCrossing: train crossing - passTrain:
// 	// passing a train on a single-track line - beingPassed: host train being passed on a single-track
// 	// line - timeSupplementRequired: recovery time is required due to an upcoming temporary speed
// 	// restriction - energyOptimisation: optimisation of energy consumption - endOfReason: the reason
// 	// previously communicated is not valid anymore.
// 	ReasonCode string `xml:"reasonCode,attr,omitempty"`
// 	// Human readable text with the reason why this advice message is sent (e.g. describing the
// 	// conflict or conflict resolution). It is possible to use only one of the parameters reasonCode
// 	// and reasonText or to send them in combination.
// 	ReasonText string `xml:"reasonText,attr,omitempty"`
// }

// // Applicable_OTN: Applicable Operational Train Number. In some cases this is different from the
// // TrainIdentification, for example when the OTN changes during the trip.
// type Applicable_OTN struct {
// 	XMLName xml.Name `xml:"Applicable_OTN"`
// 	TemporaryConstraints_ComplexType
// 	// Operational Train Number (free text). For driver information.
// 	OTN string `xml:"OTN,attr"`
// }

// // ApplicableIfTrainRunsOnSP: The effects of this signal element are to be taken into account only
// // if this SP is in the SP_List after the current SP.
// type ApplicableIfTrainRunsOnSP struct {
// 	XMLName xml.Name `xml:"ApplicableIfTrainRunsOnSP"`
// 	// Segment Profile Identifier
// 	SP_ID string `xml:"SP_ID,attr"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone,omitempty"`
// }

// // ATOTS_ContactInfo: Information about ATO-Trackside Contact Information
// type ATOTS_ContactInfo struct {
// 	XMLName xml.Name `xml:"ATOTS_ContactInfo"`
// 	// [SUBSET-126] Identifier of the ATO-TS
// 	ATOTS_ID string `xml:"ATOTS_ID,attr,omitempty"`
// 	// ERTMS: Identity number of the country or region. From SUBSET-026 chap. 7, 7.5.1.86
// 	NID_C NID_C_Type `xml:"NID_C,omitempty"`
// }

// // ATP_System_Identifier: Identifier of ATP System. If it is a number: ETCS NID_ NTC from ERA
// // "Assignment of Values to ETCS Variables", section 3.3 (ERA_ERTMS_0 40001 v1.27.
// // https://www.era.europa.eu/sites/default/files/activities/docs/ertms_040001_etcs_variables_values_en.pdf).
// // If it is a name: from the ERA "List of Class B Systems", section 3.2 (ERA/ TD/2011-11 v4.0,
// // https://www.era.europa.eu/sites/default/files/activities/docs/ertms_ccs_class_b_systems_en.pdf)
// type ATP_System_Identifier struct {
// 	XMLName xml.Name `xml:"ATP_System_Identifier"`
// 	Value   string   `xml:",chardata"`
// }

// // ATP_Systems: List of ATP systems supported in the Segment Profile. The DAS system can then find
// // out what is supported. ATP System Consequences will be for Operational Train data containing its
// // characteristics necessary to calculate its dynamic capabilities.
// type ATP_Systems struct {
// 	XMLName xml.Name `xml:"ATP_Systems"`
// 	// ATP systems supported at the start of the Segment Profile.
// 	ATP_SystemsStart *ATPSystems_ComplexType `xml:"ATP_SystemsStart"`
// 	// Change in ATP systems supported
// 	ATP_SystemsChange []ATP_SystemsChange `xml:"ATP_SystemsChange,omitempty"`
// }

// // ATP_SystemsChange: Change in ATP systems supported
// type ATP_SystemsChange struct {
// 	XMLName xml.Name `xml:"ATP_SystemsChange"`
// 	ATPSystems_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // ATPSystems_ComplexType: Complex Type for ATP Systems
// type ATPSystems_ComplexType struct {
// 	// Identifier of ATP System. If it is a number: ETCS NID_ NTC from ERA "Assignment of Values to
// 	// ETCS Variables", section 3.3 (ERA_ERTMS_0 40001 v1.27.
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/ertms_040001_etcs_variables_values_en.pdf).
// 	// If it is a name: from the ERA "List of Class B Systems", section 3.2 (ERA/ TD/2011-11 v4.0,
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/ertms_ccs_class_b_systems_en.pdf)
// 	ATP_System_Identifier []ATP_System_Identifier `xml:"ATP_System_Identifier,omitempty"`
// }

// // AverageSpeed: Average speed of the train
// type AverageSpeed struct {
// 	XMLName xml.Name `xml:"AverageSpeed"`
// }

// // AvoidSpeedRange: Speed range (in km/h) to avoid
// type AvoidSpeedRange struct {
// 	XMLName xml.Name `xml:"AvoidSpeedRange"`
// 	// Minimum speed in the range of speed to be avoided
// 	RangeMinSpeed SpeedType `xml:"rangeMinSpeed,attr"`
// 	// Maximum speed in the range of speed to be avoided
// 	RangeMaxSpeed SpeedType `xml:"rangeMaxSpeed,attr"`
// }

// // AxleLoadSpeedProfile: [SUBSET-126] Speed Profile restriction to be applied to specific axle loads
// type AxleLoadSpeedProfile struct {
// 	XMLName xml.Name `xml:"AxleLoadSpeedProfile"`
// 	SP_GenericArea_ComplexType
// 	// [SUBSET-126] Axle load from which the speed restriction is applicable
// 	AxleLoadCategory string `xml:"axleLoadCategory,attr"`
// 	// [SUBSET-126] Speed restriction to be applied if the axle load of the train ≥ axleLoadCategory
// 	NewSpeedLevel SpeedType `xml:"newSpeedLevel,attr"`
// 	// [SUBSET-126] Qualifier to indicate if the ALSP is to be applied to the front of the train (no
// 	// train length delay) or to the end of the train (train length delay).
// 	ALSP_Front string `xml:"ALSP_Front,attr,omitempty"`
// }

// // B2G_Error: Description of error (board to ground). To be used only if the result of the message
// // is "ERROR"
// type B2G_Error struct {
// 	XMLName xml.Name `xml:"B2G_Error"`
// 	Error_ComplexType
// }

// // B2G_MessageResponse: Response to message received, board to ground
// type B2G_MessageResponse struct {
// 	XMLName xml.Name `xml:"B2G_MessageResponse"`
// 	// Result of the message
// 	Result string `xml:"result,attr"`
// 	// Description of error (board to ground). To be used only if the result of the message is "ERROR"
// 	B2G_Error []B2G_Error `xml:"B2G_Error,omitempty"`
// }

// // B2G_ReplyPayload: Payload of board to ground reply
// type B2G_ReplyPayload struct {
// 	XMLName xml.Name `xml:"B2G_ReplyPayload"`
// 	// Report of DAS Driving Mode
// 	Reported_DAS_DrivingMode []Reported_DAS_DrivingMode `xml:"Reported_DAS_DrivingMode,omitempty"`
// 	// Response to message received, board to ground
// 	B2G_MessageResponse *B2G_MessageResponse `xml:"B2G_MessageResponse"`
// 	// Plaintext message
// 	PlaintextMessage []PlaintextMessage `xml:"PlaintextMessage,omitempty"`
// 	// Status Report from Board to Ground
// 	B2G_StatusReport []B2G_StatusReport `xml:"B2G_StatusReport,omitempty"`
// 	// Session termination
// 	SessionTermination *SessionTermination `xml:"SessionTermination"`
// }

// // B2G_Request: Request from board to ground
// type B2G_Request struct {
// 	XMLName xml.Name `xml:"B2G_Request"`
// 	// Segment Profile request. If version is not included, latest version is requested.
// 	SP_Request []SP_Request `xml:"SP_Request,omitempty"`
// 	// Train Characteristics request. If version is not included, latest version is requested.
// 	TC_Request []TC_Request `xml:"TC_Request,omitempty"`
// 	// Journey Profile request.
// 	JP_Request []JP_Request `xml:"JP_Request,omitempty"`
// 	// C-DAS-C advice request
// 	C_DAS_C_AdviceRequest []C_DAS_C_AdviceRequest `xml:"C_DAS_C_AdviceRequest,omitempty"`
// 	// Plaintext message request
// 	PlaintextMessageRequest []PlaintextMessageRequest `xml:"PlaintextMessageRequest,omitempty"`
// 	// Request for ForceStatusChange messages
// 	ForceStatusChangeRequest []ForceStatusChangeRequest `xml:"ForceStatusChangeRequest,omitempty"`
// 	// Request from Board to Ground of position and speed. If the Train Identification attribute is not
// 	// given, the Ground will send the reply according to the train for which the communication layer
// 	// has been established. In the IM-RU setup between IM DAS-TS and RU DAS-TS, if the communication
// 	// layer is established for multiple trains, the train identifier must be given.
// 	PositionSpeedRequest     []PositionSpeedRequest     `xml:"PositionSpeedRequest,omitempty"`
// 	InterruptionPointRequest []InterruptionPointRequest `xml:"InterruptionPointRequest,omitempty"`
// }

// // B2G_StatusReport: Status Report from Board to Ground
// type B2G_StatusReport struct {
// 	XMLName xml.Name `xml:"B2G_StatusReport"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// DAS driving mode (see enumeration)
// 	DAS_drivingMode string `xml:"DAS_drivingMode,attr,omitempty"`
// 	// [SUBSET-126] ATO State, from Subset-126. Only relevant when in ATO mode.
// 	ATO_State string `xml:"ATO_State,attr,omitempty"`
// 	// [SUBSET-126] ATO: JP SP Consistency Error, from Subset-126 (boolean)
// 	JP_SP_ConsistencyError string `xml:"JP_SP_ConsistencyError,attr,omitempty"`
// 	// [SUBSET-126] ATO: skip next stopping point, from Subset-126 (boolean)
// 	NextStoppingPointSkip string `xml:"nextStoppingPointSkip,attr,omitempty"`
// 	// [SUBSET-126] ATO: Slip/slide reported by TCMS/Train, from Subset-126 (boolean)
// 	SlipSlideReported string `xml:"slipSlideReported,attr,omitempty"`
// 	// [SUBSET-126] ATO: Operational conditions fulfilment, from Subset-126 (boolean)
// 	OperationalConditionsFulfillment string `xml:"operationalConditionsFulfillment,attr,omitempty"`
// 	// [SUBSET-126] ATO: Train is moving, from Subset-126 (boolean)
// 	TrainMoving string `xml:"trainMoving,attr,omitempty"`
// 	// [SUBSET-126] ATO: Unable to stop at the next Stopping Point, from Subset-126 (boolean)
// 	UnableToStop string `xml:"unableToStop,attr,omitempty"`
// 	// [SUBSET-126] ATO: Driver identifier number as defined in SUBSET-027 v300 §4.2.3.7.
// 	DriverID string `xml:"driverID,attr,omitempty"`
// 	// [SUBSET-126] ATO: Routing Error, from Subset-126 (boolean)
// 	RoutingError string `xml:"routingError,attr,omitempty"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// 	// Position and speed of the train
// 	PositionSpeed *PositionSpeed `xml:"PositionSpeed,omitempty"`
// 	// [SUBSET-126] Information concerning the most recent Timing Point considered as either reached or
// 	// passed
// 	PreviousTimingPoint *PreviousTimingPoint `xml:"PreviousTimingPoint,omitempty"`
// 	// [SUBSET-126] Estimated arrival time to, at least, the next Timing Point
// 	TimingPointEstimation []TimingPointEstimation `xml:"TimingPointEstimation,omitempty"`
// 	// Change in adhesion conditions
// 	AdhesionConditionsChange *AdhesionConditionsChange `xml:"AdhesionConditionsChange,omitempty"`
// 	// Energy data expected in the next timeframe (5 to 15 minutes)
// 	ExpectedCurrentData []ExpectedCurrentData `xml:"ExpectedCurrentData,omitempty"`
// 	// Change in Train Characteristics
// 	TrainCharacteristicsChange *TrainCharacteristicsChange `xml:"TrainCharacteristicsChange,omitempty"`
// 	// Plaintext message
// 	PlaintextMessage *PlaintextMessage `xml:"PlaintextMessage,omitempty"`
// 	// Journey Profile currently in use
// 	JP_InUse *JP_InUse `xml:"JP_InUse,omitempty"`
// 	// Identifier of the C-DAS-C message which is currently being shown to the driver
// 	C_DAS_C_Shown    *C_DAS_C_Shown    `xml:"C_DAS_C_Shown,omitempty"`
// 	TrainReadyStatus *TrainReadyStatus `xml:"TrainReadyStatus,omitempty"`
// 	// Inform the IM/RU Backoffice of the reason for an additional delay encountered by a train. For
// 	// cases where the train is interrupted, or at the points where the train ready message is due
// 	// (origin station and changes in train characteristics or RU).
// 	DelayNotice *DelayNotice `xml:"DelayNotice,omitempty"`
// 	// Describes the interruption points with location and the reason for the interruption. It must be
// 	// sent in all Status reports until the interruption has been lifted.
// 	InterruptionPoint *InterruptionPoint `xml:"InterruptionPoint,omitempty"`
// 	// Real-time energy data
// 	EnergyData *EnergyData `xml:"EnergyData,omitempty"`
// 	// Network Specific Parameters in Status Report
// 	StatusReport_NSPs []NSPList_ComplexType `xml:"StatusReport_NSPs,omitempty"`
// }

// // Balise: Equipment on the track at a known position that communicates with the train to exchange
// // different information. The train can recalibrate its positioning systems according to this
// // balise. If the Balise has a NID_C different from the one of the SP, it can be expressed here.
// type Balise struct {
// 	XMLName xml.Name `xml:"Balise"`
// 	SP_GenericPoint_ComplexType
// 	// [SUBSET-126] Identifier of the balise group
// 	BaliseGroupIdentifier string `xml:"baliseGroupIdentifier,attr,omitempty"`
// 	// [SUBSET-126] Defines the relative position in a balise group
// 	BaliseRelativePosition string `xml:"baliseRelativePosition,attr,omitempty"`
// 	// Determines if the balise is guarding any level crossing
// 	LevelCrossingGuarding string `xml:"levelCrossingGuarding,attr,omitempty"`
// 	// Amount of guarded level crossings (to be used only if levelCrossingGuarding is true)
// 	AmountLevelCrossings string `xml:"amountLevelCrossings,attr,omitempty"`
// 	// ERTMS: Identity number of the country or region. From SUBSET-026 chap. 7, 7.5.1.86
// 	NID_C NID_C_Type `xml:"NID_C,omitempty"`
// }

// // NSPList_ComplexType: Complex Type for a list of Network-Specific Parameters
// type NSPList_ComplexType struct {
// 	// Name of the element
// 	Name string `xml:"name,attr,omitempty"`
// 	// Generic Network-Specific Parameter
// 	NetworkSpecificParameter []NetworkSpecificParameter `xml:"NetworkSpecificParameter,omitempty"`
// }

// // BFC_TractionType: Traction Type for Braking Force Curve
// type BFC_TractionType struct {
// 	XMLName xml.Name `xml:"BFC_TractionType"`
// 	// Traction Type (DC/AC with voltage, self-propelled or powerless)
// 	TractionType TractionTypeType `xml:"tractionType,attr,omitempty"`
// }

// // BrakingForceCurve_ComplexType: Complex Type for Braking Force Curves
// type BrakingForceCurve_ComplexType struct {
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// Traction Type (DC/AC with voltage, self-propelled or powerless)
// 	TractionType TractionTypeType `xml:"tractionType,attr,omitempty"`
// 	// Adhesion category for which the braking force curve is applicable. When no curves are defined
// 	// for better adhesion categories, the curve is also applicable.
// 	BFC_AdhesionCategory AdhesionCategoryType `xml:"BFC_AdhesionCategory,attr,omitempty"`
// 	// Data point for the braking force curve (pair of speed and braking force)
// 	BrakingForceCurveDataPoint []BrakingForceCurveDataPoint `xml:"BrakingForceCurveDataPoint,omitempty"`
// 	// Identifier of ATP System. If it is a number: ETCS NID_ NTC from ERA "Assignment of Values to
// 	// ETCS Variables", section 3.3 (ERA_ERTMS_0 40001 v1.27.
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/ertms_040001_etcs_variables_values_en.pdf).
// 	// If it is a name: from the ERA "List of Class B Systems", section 3.2 (ERA/ TD/2011-11 v4.0,
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/ertms_ccs_class_b_systems_en.pdf)
// 	ATP_System_Identifier []ATP_System_Identifier `xml:"ATP_System_Identifier,omitempty"`
// 	// Traction Type for Braking Force Curve
// 	BFC_TractionType []BFC_TractionType `xml:"BFC_TractionType,omitempty"`
// 	// Network Specific Parameters in Braking Force Curve
// 	BFC_NSPs []NSPList_ComplexType `xml:"BFC_NSPs,omitempty"`
// }

// // BrakingForceCurveDataPoint: Data point for the braking force curve (pair of speed and braking
// // force)
// type BrakingForceCurveDataPoint struct {
// 	XMLName xml.Name `xml:"BrakingForceCurveDataPoint"`
// 	// speed (km/h)
// 	Speed SpeedType `xml:"speed,attr"`
// 	// Braking Force (in N)
// 	BrakingForce ForceType `xml:"brakingForce,attr"`
// 	// Maximum force that can be Applied through regenerative braking. Expressed in N
// 	MaximumRegenerativeForce string `xml:"maximumRegenerativeForce,attr,omitempty"`
// }

// // C_DAS_C_Advice: Explicit Speed Advice to On-Board unit. To be used in C-DAS-C operation.
// type C_DAS_C_Advice struct {
// 	XMLName xml.Name `xml:"C_DAS_C_Advice"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// Type of C-DAS-C Advice
// 	AdviceType string `xml:"adviceType,attr"`
// 	// Optimal speed given to the driver (km/h). It is recommended to fill the parameter optimalSpeed
// 	// according to the following rules, depending on the advice instance being sent: - Stop Advice:
// 	// optimalSpeed must be null - Acceleration Advice: optimalSpeed is the target speed that shall be
// 	// reached at the end of the acceleration phase - Constant Speed Advice: optimalSpeed is the target
// 	// speed for the cruising phase - Coasting Advice: optimalSpeed is the target speed that shall be
// 	// reached at the end of the coasting phase - Operational/Electrical Braking Advice: optimalSpeed
// 	// is the target speed that shall be reached at the end of the braking phase - Departure Advice:
// 	// optimalSpeed is either null or the target speed that shall be reached after departing at the end
// 	// of the acceleration phase There is a special value: 999 indicating "maximum allowed track speed"
// 	OptimalSpeed SpeedType `xml:"optimalSpeed,attr,omitempty"`
// 	// Difference (in km/h) between maximum allowed track speed/most restrictive speed profile and
// 	// optimal speed. The special value 0 indicates "maximum allowed track speed/most restrictive speed
// 	// profile".
// 	DeltaSpeed SpeedType `xml:"deltaSpeed,attr,omitempty"`
// 	// Human readable text of the location information where the advice becomes valid (and/or ends),
// 	// which will be displayed on the OBU. If no location is given the advice is immediately valid. It
// 	// is recommended to fill the parameter location according to the following rules, depending on the
// 	// advice instance being sent: - stop advice: location is the station the train is stopping at
// 	// (e.g. train approaching a scheduled stop) - acceleration advice: location is the location
// 	// (kilometre/landmark etc.) where the acceleration phase starts - constant speed advice: location
// 	// is the location (kilometre/landmark etc.) where the cruising phase starts and/or ends -
// 	// operational/electrical braking advice: location is the location (kilometre/landmark etc.) where
// 	// the braking phase starts - end of advice: location is null - departure advice: location is the
// 	// station the train is departing from
// 	LocationText string `xml:"locationText,attr,omitempty"`
// 	// Human readable text with the reason why this advice message is sent (e.g. describing the
// 	// conflict or conflict resolution). It is possible to use only one of the parameters reasonCode
// 	// and reasonText or to send them in combination.
// 	ReasonText string `xml:"reasonText,attr,omitempty"`
// 	// reasonCode offers coding for standard situations. If reasonCode is null the last reasonCode
// 	// received is still valid. It is possible to use only one of the parameters reasonCode and
// 	// reasonText or to send them in combination. Predefined reasonCodes are: - followTrain: host train
// 	// is following a conflicting (slower) train - trainFollowing: host train is preceding a
// 	// conflicting (faster) train - merge2nd: conflicting (faster) train merging ahead and entering
// 	// common route section first - merge1st: conflicting (slower) train merging ahead, host train
// 	// enters common route section first - overtakeTrain: overtake conflicting (slower) train that has
// 	// (nearly) reached the location of overtaking - followTrainOvertakeTrain: follow, then overtake
// 	// conflicting train that hasn’t yet reached the location of overtaking - beingOvertaken:
// 	// conflicting (faster) train about to overtake - trainFollowingBeingOvertaken: conflicting
// 	// (faster) train following and then overtaking - trainCrossing: train crossing - passTrain:
// 	// passing a train on a single-track line - beingPassed: host train being passed on a single-track
// 	// line - timeSupplementRequired: recovery time is required due to an upcoming temporary speed
// 	// restriction - energyOptimisation: optimisation of energy consumption - endOfReason: the reason
// 	// previously communicated is not valid anymore.
// 	ReasonCode string `xml:"reasonCode,attr,omitempty"`
// 	// Begin of the advice validity. Either both startValidity and endValidity or none of the two must
// 	// be specified. Both parameters should refer to a future point in time and endValidity must be
// 	// after startValidity. If either of the two indicates a past time (w.r.t. the clock of the
// 	// On-Board Unit), which is possible in a situation where Central Unit and OBU are not
// 	// synchronised, transmission delays etc., then the OBU should treat it as valid but past time
// 	// indication.
// 	StartValidity UTCDateTimeType `xml:"startValidity,attr,omitempty"`
// 	// End of the advice validity. Either both startValidity and endValidity or none of the two must be
// 	// specified. Both parameters should refer to a future point in time and endValidity must be after
// 	// startValidity. If either of the two indicates a past time (w.r.t. the clock of the On-Board
// 	// Unit), which is possible in a situation where Central Unit and OBU are not synchronised,
// 	// transmission delays etc., then the OBU should treat it as valid but past time indication.
// 	EndValidity UTCDateTimeType `xml:"endValidity,attr,omitempty"`
// 	// The meaning depends on the type of advice instance being sent. Only valid with certain types of
// 	// advice. - acceleration advice: Percentage of the maximum acceleration - operational braking
// 	// advice: Percentage of the maximum operational braking - electrical braking advice: Percentage of
// 	// the maximum electrical braking
// 	Percentage PercentageType `xml:"percentage,attr,omitempty"`
// 	// Lower bound of time window for train departure. Only to be used if departure advice is sent. If
// 	// earliestDepartureTime is not specified, the advice should be obeyed at any time earlier than
// 	// latestDepartureTime. If both attributes are not specified, the advice should be obeyed
// 	// immediately.
// 	EarliestDepartureAdviceTime UTCDateTimeType `xml:"earliestDepartureAdviceTime,attr,omitempty"`
// 	// Upper bound of time window for train departure. Only to be used if departure advice is sent. If
// 	// latestDepartureTime is not specified, the advice should be obeyed at any time later than
// 	// earliestDepartureTime. If both attributes are not specified, the advice should be obeyed
// 	// immediately.
// 	LatestDepartureAdviceTime UTCDateTimeType `xml:"latestDepartureAdviceTime,attr,omitempty"`
// 	// References the ID of the message, that defined the absolute speed, which is still valid
// 	ReferenceIDAbs UUIDv4_Type `xml:"referenceIDAbs,attr,omitempty"`
// 	// References the start station at the beginning of the message (short form)
// 	StartStationShort string `xml:"startStationShort,attr,omitempty"`
// 	// References the start station at the beginning of the message (long form)
// 	StartStationLong string `xml:"startStationLong,attr,omitempty"`
// 	// References the kilometre at the beginning of the message, preferably linked with a kilometer
// 	// sign near the track
// 	StartKm string `xml:"startKm,attr,omitempty"`
// 	// References the (national) line number at the beginning of the message
// 	StartLine string `xml:"startLine,attr,omitempty"`
// 	// References the name of the start signal of the message, if the start of the message is at or
// 	// near a signal
// 	StartSignal string `xml:"startSignal,attr,omitempty"`
// 	// References the start station at the end of the message (short form)
// 	EndStationShort string `xml:"endStationShort,attr,omitempty"`
// 	// References the start station at the end of the message (long form)
// 	EndStationLong string `xml:"endStationLong,attr,omitempty"`
// 	// References the kilometre at the end of the message, preferably linked with a kilometer sign near
// 	// the track
// 	EndKm string `xml:"endKm,attr,omitempty"`
// 	// References the (national) line number at the end of the message
// 	EndLine string `xml:"endLine,attr,omitempty"`
// 	// References the name of the end signal of the message, if the end of the message is at or near a
// 	// signal
// 	EndSignal string `xml:"endSignal,attr,omitempty"`
// 	// References the station where the optimalSpeed changes next (short form)
// 	EndStationShortAbs string `xml:"endStationShortAbs,attr,omitempty"`
// 	// References the station where the optimalSpeed changes next (long form)
// 	EndStationLongAbs string `xml:"endStationLongAbs,attr,omitempty"`
// 	// References the kilometre where the optimalSpeed changes next, preferably linked with a kilometer
// 	// sign near the track
// 	EndKmAbs string `xml:"endKmAbs,attr,omitempty"`
// 	// References the (national) line number where the optimalSpeed changes next
// 	EndLineAbs string `xml:"endLineAbs,attr,omitempty"`
// 	// References the name of the signal where the optimalSpeed changes next, if this is at or near a
// 	// signal
// 	EndSignalAbs string `xml:"endSignalAbs,attr,omitempty"`
// 	// Time at the end of the optimalSpeed
// 	EndValidityAbs UTCDateTimeType `xml:"endValidityAbs,attr,omitempty"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
// 	// Information about the Segment Profile
// 	SegmentProfileInfo *SegmentProfileInfo `xml:"SegmentProfileInfo,omitempty"`
// }

// // C_DAS_C_AdviceRequest: C-DAS-C advice request
// type C_DAS_C_AdviceRequest struct {
// 	XMLName xml.Name `xml:"C_DAS_C_AdviceRequest"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
// 	// Identifier of the C-DAS-C message which is currently being shown to the driver
// 	C_DAS_C_Shown *C_DAS_C_Shown `xml:"C_DAS_C_Shown,omitempty"`
// }

// // C_DAS_C_Shown: Identifier of the C-DAS-C message which is currently being shown to the driver
// type C_DAS_C_Shown struct {
// 	XMLName xml.Name `xml:"C_DAS_C_Shown"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// Identifier of the SFERA message
// 	Message_ID UUIDv4_Type `xml:"message_ID,attr"`
// }

// // CancelPreviousSignal: This type of signal will be used in the case that the signal will present
// // an aspect that is different from the aspect announced by a previous signal.
// type CancelPreviousSignal struct {
// 	XMLName xml.Name `xml:"CancelPreviousSignal"`
// 	// Identifier of the physical signal on the track
// 	Signal_ID_Physical string `xml:"signal_ID_Physical,attr"`
// 	// Identifier of the signal object
// 	Signal_ID_Object string `xml:"signal_ID_Object,attr,omitempty"`
// 	// The distanceBeforeSignal attribute gives the point from which to restart accelerating. Depending
// 	// on the type of signalling, the position of this point will be defined by the visual
// 	// identification by the driver, the position of a balise, or the time at which the driver will
// 	// receive an updated information for in-cab signalling
// 	DistanceBeforeSignal DistanceType `xml:"distanceBeforeSignal,attr,omitempty"`
// }

// // CantDeficiencySSP: Cant Deficiency Static Speed Profile
// type CantDeficiencySSP struct {
// 	XMLName xml.Name `xml:"CantDeficiencySSP"`
// 	// [SUBSET-126] (If specific_SSP_Category = Cant_Deficiency_SSP) “Cant Deficiency” SSP category for
// 	// which a different value for the static line speed exists. Used together with V_DIFF to permit
// 	// certain trains to go faster or lower than the “international basic static speed” given by
// 	// SSP_Speed.
// 	SSP_NC_CDDIFF NC_CDTrainType `xml:"SSP_NC_CDDIFF,attr,omitempty"`
// }

// // ChangeReason: Reason for Journey Profile Change
// type ChangeReason struct {
// 	XMLName xml.Name `xml:"ChangeReason"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// [SUBSET-126] Distance of the element's start location from the beginning of the Segment Profile
// 	// (m)
// 	StartLocation DistanceType `xml:"startLocation,attr,omitempty"`
// 	// [SUBSET-126] Distance of the element's end location from the beginning of the Segment Profile
// 	// (m)
// 	EndLocation DistanceType `xml:"endLocation,attr,omitempty"`
// 	// Starting date and time (in UTC)
// 	StartTime UTCDateTimeType `xml:"startTime,attr,omitempty"`
// 	// Ending date and time (in UTC)
// 	EndTime UTCDateTimeType `xml:"endTime,attr,omitempty"`
// 	// Reason for JP Change
// 	JP_ChangeReason string `xml:"JP_ChangeReason,attr,omitempty"`
// }

// // CommunicationNetwork: Network used for communication between driver and ground. The driver needs
// // to know it because some procedures in case of an emergency depend on the network type.
// type CommunicationNetwork struct {
// 	XMLName xml.Name `xml:"CommunicationNetwork"`
// 	SP_GenericArea_ComplexType
// 	// Type of communication network. Standard possibilities are: "GSM-R", "FRMCS", "Public GSM",
// 	// "Public WiFi" and "Private WiFi".
// 	CommunicationNetworkType string `xml:"communicationNetworkType,attr"`
// 	// Channel for communication network
// 	CommunicationNetworkChannel string `xml:"communicationNetworkChannel,attr,omitempty"`
// }

// // ConnectionTrack: Connection track
// type ConnectionTrack struct {
// 	XMLName xml.Name `xml:"ConnectionTrack"`
// 	SP_GenericPoint_ComplexType
// 	// Type of connection track
// 	ConnectionTrackType string `xml:"connectionTrackType,attr"`
// 	// Description of connection track
// 	ConnectionTrackDescription string `xml:"connectionTrackDescription,attr,omitempty"`
// }

// // Contact: Information related to a contact
// type Contact struct {
// 	XMLName xml.Name `xml:"Contact"`
// 	// Identifier to be used to reach the contact on the communication platform defined in ContactType
// 	// (e.g. telephone number, GSM-R number, pseudo...)
// 	ContactIdentifier string `xml:"contactIdentifier,attr"`
// 	// Type of contact. Communication platform to be used to contact this person: e.g. telephone,
// 	// GSM-R, social network, messaging app, etc.
// 	ContactType string `xml:"contactType,attr,omitempty"`
// 	// Role of the contact
// 	ContactRole string `xml:"contactRole,attr,omitempty"`
// 	// Notes related to the contact
// 	ContactNotes string `xml:"contactNotes,attr,omitempty"`
// 	// Specifies if this is the main contact for the area
// 	MainContact string `xml:"mainContact,attr,omitempty"`
// }

// // ContactList: List of contacts for a specific area
// type ContactList struct {
// 	XMLName xml.Name `xml:"ContactList"`
// 	SP_GenericArea_ComplexType
// 	// Information related to a contact
// 	Contact []Contact `xml:"Contact,omitempty"`
// }

// // CountryCodeISO: Identifies a County or State by code (ISO 3166-1)
// type CountryCodeISO struct {
// 	XMLName xml.Name        `xml:"CountryCodeISO"`
// 	Value   CountryIdentISO `xml:",chardata"`
// }

// // CurrentLimitation: Maximum current to be applied by the train. Expressed in A
// type CurrentLimitation struct {
// 	XMLName xml.Name `xml:"CurrentLimitation"`
// 	// Maximum current at the start of the Segment Profile
// 	CurrentLimitationStart *CurrentLimitation_ComplexType `xml:"CurrentLimitationStart"`
// 	// Maximum current change
// 	CurrentLimitationChange []CurrentLimitationChange `xml:"CurrentLimitationChange,omitempty"`
// }

// // CurrentLimitationChange: Maximum current change
// type CurrentLimitationChange struct {
// 	XMLName xml.Name `xml:"CurrentLimitationChange"`
// 	CurrentLimitation_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // CurrentLimitation_ComplexType: Complex Type for Current Limitation
// type CurrentLimitation_ComplexType struct {
// 	// Maximum current value (A)
// 	MaxCurValue CurrentType `xml:"maxCurValue,attr"`
// }

// // CurveChange: Curve change
// type CurveChange struct {
// 	XMLName xml.Name `xml:"CurveChange"`
// 	Curve_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // Curves: Curves in the Segment Profile
// type Curves struct {
// 	XMLName xml.Name `xml:"Curves"`
// 	// Curve at the start of the Segment Profile
// 	CurveStart *Curve_ComplexType `xml:"CurveStart"`
// 	// Curve change
// 	CurveChange []CurveChange `xml:"CurveChange,omitempty"`
// }

// // Curve_ComplexType: Complex Type for Curve
// type Curve_ComplexType struct {
// 	// Radius of the curve (m)
// 	CurveRadius string `xml:"curveRadius,attr"`
// 	// Direction of the curve (Left or Right)
// 	CurveDirection string `xml:"curveDirection,attr,omitempty"`
// }

// // DAS_DrivingModeRequest: Request from Ground to Board of DAS driving mode. If the Train
// // Identification attribute is not given, the Board will send the reply according to the train for
// // which the communication layer has been established. In the IM-RU setup between IM DAS-TS and RU
// // DAS-TS, if the communication layer is established for multiple trains, the train identifier must
// // be given.
// type DAS_DrivingModeRequest struct {
// 	XMLName xml.Name `xml:"DAS_DrivingModeRequest"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// }

// // DAS_ModesComplexType: Complex type for DAS modes
// type DAS_ModesComplexType struct {
// 	// DAS driving mode (see enumeration)
// 	DAS_drivingMode string `xml:"DAS_drivingMode,attr"`
// 	// DAS architecture: DAS-C (ground advice calculation) or DAS-O (board advice calculation)
// 	DAS_architecture string `xml:"DAS_architecture,attr"`
// 	// DAS connectivity (standalone or connected)
// 	DAS_connectivity string `xml:"DAS_connectivity,attr"`
// }

// // DAS_ModesUnavailabilityArea: Area in SP where some operating modes are unavailable
// type DAS_ModesUnavailabilityArea struct {
// 	XMLName xml.Name `xml:"DAS_ModesUnavailabilityArea"`
// 	SP_GenericArea_ComplexType
// 	// Operating modes that are unavailable. If all 3 attributes are included, the triplet is
// 	// unavailable; of only 1 or 2 attributes are included, that mode is unavailable independently of
// 	// the value of the rest of the attributes.
// 	Unavailable_DAS_OperatingModes *Unavailable_DAS_OperatingModes `xml:"Unavailable_DAS_OperatingModes"`
// }

// // DAS_OperatingModeSelected: Operating mode selected by the DAS-TS
// type DAS_OperatingModeSelected struct {
// 	XMLName xml.Name `xml:"DAS_OperatingModeSelected"`
// 	// DAS architecture: DAS-C (ground advice calculation) or DAS-O (board advice calculation)
// 	DAS_architecture string `xml:"DAS_architecture,attr"`
// 	// DAS connectivity (standalone or connected)
// 	DAS_connectivity string `xml:"DAS_connectivity,attr"`
// }

// // DecisiveGradient: Decisive Gradient Area. The decisive gradient is the steepest gradient in front
// // of potential stopping points on a line like signals, level crossings and stations. It is
// // relevant, if it goes downhill, for the calculation of either the minimum braking percentage or
// // the speed limit for a train because in case of reduced braking capacity, the train might not be
// // able to brake properly, therefore, the driver has to adjust the speed. If it goes uphill, it is
// // relevant to calculate the towing load, for example, to make sure, with the available engine
// // power, that the locomotive can get away after it stopped.
// type DecisiveGradient struct {
// 	XMLName xml.Name `xml:"DecisiveGradient"`
// 	SP_GenericArea_ComplexType
// 	// Gradient value (Numeric value, mm by m)
// 	GradientValue GradientValueType `xml:"gradientValue,attr"`
// 	// Direction of the gradient (Downhill or Uphill)
// 	GradientDirection GradientDirectionType `xml:"gradientDirection,attr"`
// 	// Direction of application on the Segment Profile (Nominal, Reverse or Both)
// 	DirectionOfApplicationOnSP string `xml:"directionOfApplicationOnSP,attr"`
// }

// // Delay: Delay of the train
// type Delay struct {
// 	XMLName xml.Name `xml:"Delay"`
// 	// Delay of the train
// 	Delay string `xml:"Delay,attr,omitempty"`
// 	// Category of the delay. Delay could be measured, delay can be expected (for example before
// 	// departure), announcement that the delay will be published (upcoming) or can be unkown
// 	DelayCategory string `xml:"DelayCategory,attr,omitempty"`
// }

// // DelayNotice: Inform the IM/RU Backoffice of the reason for an additional delay encountered by a
// // train. For cases where the train is interrupted, or at the points where the train ready message
// // is due (origin station and changes in train characteristics or RU).
// type DelayNotice struct {
// 	XMLName xml.Name `xml:"DelayNotice"`
// 	// This element identifies the reason for a delay (modified DelayReason)
// 	DelayCause DelayCode `xml:"DelayCause,omitempty"`
// 	// Precise reason for a delay (string)
// 	PreciseDelayCause string         `xml:"PreciseDelayCause,omitempty"`
// 	TrainLocation     *LocationIdent `xml:"TrainLocation,omitempty"`
// }

// // SP_GenericArea_ComplexType: Complex Type for generic Segment Profile Area
// type SP_GenericArea_ComplexType struct {
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// [SUBSET-126] Distance of the element's start location from the beginning of the Segment Profile
// 	// (m)
// 	StartLocation DistanceType `xml:"startLocation,attr,omitempty"`
// 	// [SUBSET-126] Distance of the element's end location from the beginning of the Segment Profile
// 	// (m)
// 	EndLocation DistanceType `xml:"endLocation,attr,omitempty"`
// 	// [SUBSET-126] Specifies if the element starts, ends, starts and ends or covers the whole
// 	// concerning Segment Profile
// 	StartEndQualifier string `xml:"startEndQualifier,attr"`
// }

// // EndSignalApplication: Location where the signal stops to be applied. If not specified, the signal
// // is applied indefinitely, until superseded by another signal. This may also be in a different SP:
// // then the SP_ID has to be specified (otherwise the current SP is considered).
// type EndSignalApplication struct {
// 	XMLName xml.Name `xml:"EndSignalApplication"`
// 	// Segment Profile Identifier
// 	SP_ID string `xml:"SP_ID,attr,omitempty"`
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr,omitempty"`
// 	// What part of the train must clear the signal application before the signal stops being applied
// 	EndSignalApplicationMustBeClearedBy TrainPartType `xml:"endSignalApplicationMustBeClearedBy,attr,omitempty"`
// }

// // EnergyData: Real-time energy data
// type EnergyData struct {
// 	XMLName xml.Name `xml:"EnergyData"`
// 	// Measured voltage at pantograph (in V)
// 	MeasuredVoltage VoltageType `xml:"measuredVoltage,attr,omitempty"`
// 	// Measured current taken from overhead contact line, negative during regeneration (in A)
// 	MeasuredCurrent CurrentType `xml:"measuredCurrent,attr,omitempty"`
// 	// Amount of energy still available on-board for the self-propelled energy source (in kWh)
// 	Autonomy EnergyType `xml:"autonomy,attr,omitempty"`
// 	// Maximal power that can be delivered form energy source available on-board (in kW)
// 	MaxPower PowerTypekW `xml:"maxPower,attr,omitempty"`
// 	// set value by the driver for traction and braking
// 	SetValueTractionBraking string `xml:"setValueTractionBraking,attr,omitempty"`
// }

// // EstimatedVoltage: Estimated voltage of the traction system (V)
// type EstimatedVoltage struct {
// 	XMLName xml.Name `xml:"EstimatedVoltage"`
// 	// Estimated voltage at the start of the Segment Profile
// 	EstimatedVoltageStart *Voltage_ComplexType `xml:"EstimatedVoltageStart"`
// 	// Estimated voltage change
// 	EstimatedVoltageChange []EstimatedVoltageChange `xml:"EstimatedVoltageChange,omitempty"`
// }

// // EstimatedVoltageChange: Estimated voltage change
// type EstimatedVoltageChange struct {
// 	XMLName xml.Name `xml:"EstimatedVoltageChange"`
// 	Voltage_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // EstimatedVoltageConstraint: Temporary constraints on estimated voltage
// type EstimatedVoltageConstraint struct {
// 	XMLName xml.Name `xml:"EstimatedVoltageConstraint"`
// 	// Voltage value (V)
// 	VoltageValue VoltageType `xml:"voltageValue,attr,omitempty"`
// }

// // Voltage_ComplexType: Complex Type for voltage
// type Voltage_ComplexType struct {
// 	// Voltage value (V)
// 	VoltageValue VoltageType `xml:"voltageValue,attr"`
// 	// [SUBSET-126] Country identifier of the traction system. It identifies the information,
// 	// additional to voltageValue, required to fully define the traction system. Use NID_CTRACTION
// 	// numbers from ERA document ERA_ERTMS_040001.
// 	NID_CTRACTION string `xml:"NID_CTRACTION,attr,omitempty"`
// }

// // ExpectedCurrentData: Energy data expected in the next timeframe (5 to 15 minutes)
// type ExpectedCurrentData struct {
// 	XMLName xml.Name `xml:"ExpectedCurrentData"`
// 	// Starting date and time (in UTC)
// 	StartTime UTCDateTimeType `xml:"startTime,attr,omitempty"`
// 	// Ending date and time (in UTC)
// 	EndTime UTCDateTimeType `xml:"endTime,attr,omitempty"`
// 	// Expected consumption current (A)
// 	ExpectedConsumptionCurrent CurrentType `xml:"expectedConsumptionCurrent,attr"`
// 	// Expected regenerative current (A)
// 	ExpectedRegenerativeCurrent CurrentType `xml:"expectedRegenerativeCurrent,attr"`
// }

// // ForceStatusChange: Force On-Board DAS to change its Status (mainly used to make a DAS-OB
// // readonly)
// type ForceStatusChange struct {
// 	XMLName xml.Name `xml:"ForceStatusChange"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// DAS driving mode (see enumeration)
// 	DAS_drivingMode string `xml:"DAS_drivingMode,attr"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// }

// // ForceStatusChangeRequest: Request for ForceStatusChange messages
// type ForceStatusChangeRequest struct {
// 	XMLName xml.Name `xml:"ForceStatusChangeRequest"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
// 	// identifier of the last ForceStatusChange request message sent by the board side of the
// 	// communication. A ForceStatusChange Reply will only be sent if a ForceStatusChange order has been
// 	// generated since the last request was sent. Otherwise, a simple "OK" message response is sent.
// 	LastForceStatusChangeMessage *LastForceStatusChangeMessage `xml:"LastForceStatusChangeMessage,omitempty"`
// }

// // G2B_Error: Description of error (ground to board). To be used only if the result of the message
// // is "ERROR"
// type G2B_Error struct {
// 	XMLName xml.Name `xml:"G2B_Error"`
// 	Error_ComplexType
// 	// Date and time of first availability of the requested data
// 	DataFirstAvailable UTCDateTimeType `xml:"dataFirstAvailable,attr,omitempty"`
// }

// // G2B_EventPayload: Payload of ground to board event
// type G2B_EventPayload struct {
// 	XMLName xml.Name `xml:"G2B_EventPayload"`
// 	// Journey Profile for a specific trip
// 	JourneyProfile []JourneyProfile `xml:"JourneyProfile,omitempty"`
// 	// Plaintext message
// 	PlaintextMessage []PlaintextMessage `xml:"PlaintextMessage,omitempty"`
// 	// Explicit Speed Advice to On-Board unit. To be used in C-DAS-C operation.
// 	C_DAS_C_Advice []C_DAS_C_Advice `xml:"C_DAS_C_Advice,omitempty"`
// 	// Force On-Board DAS to change its Status (mainly used to make a DAS-OB readonly)
// 	ForceStatusChange []ForceStatusChange `xml:"ForceStatusChange,omitempty"`
// 	// Contains information about the own train and the related trains, i.e. trains that share the same
// 	// trainpath (infrastructure) as the own train.
// 	RelatedTrainInformation *RelatedTrainInformation `xml:"RelatedTrainInformation"`
// 	// Describes the interruption points with location and the reason for the interruption. It must be
// 	// sent in all Status reports until the interruption has been lifted.
// 	InterruptionPoint *InterruptionPoint `xml:"InterruptionPoint"`
// }

// // G2B_MessageResponse: Response to message received, ground to board
// type G2B_MessageResponse struct {
// 	XMLName xml.Name `xml:"G2B_MessageResponse"`
// 	// Result of the message
// 	Result string `xml:"result,attr"`
// 	// Description of error (ground to board). To be used only if the result of the message is "ERROR"
// 	G2B_Error []G2B_Error `xml:"G2B_Error,omitempty"`
// }

// // G2B_ReplyPayload: Payload of ground to board reply
// type G2B_ReplyPayload struct {
// 	XMLName xml.Name `xml:"G2B_ReplyPayload"`
// 	// Journey Profile for a specific trip
// 	JourneyProfile []JourneyProfile `xml:"JourneyProfile,omitempty"`
// 	// Profile for a section of railway track
// 	SegmentProfile []SegmentProfile `xml:"SegmentProfile,omitempty"`
// 	// Element containing the characteristics of the train
// 	TrainCharacteristics []TrainCharacteristics `xml:"TrainCharacteristics,omitempty"`
// 	// Response to message received, ground to board
// 	G2B_MessageResponse *G2B_MessageResponse `xml:"G2B_MessageResponse,omitempty"`
// 	// Plaintext message
// 	PlaintextMessage []PlaintextMessage `xml:"PlaintextMessage,omitempty"`
// 	// Position and speed of the train
// 	PositionSpeed []PositionSpeed `xml:"PositionSpeed,omitempty"`
// 	// Explicit Speed Advice to On-Board unit. To be used in C-DAS-C operation.
// 	C_DAS_C_Advice []C_DAS_C_Advice `xml:"C_DAS_C_Advice,omitempty"`
// 	// Force On-Board DAS to change its Status (mainly used to make a DAS-OB readonly)
// 	ForceStatusChange []ForceStatusChange `xml:"ForceStatusChange,omitempty"`
// }

// // G2B_Request: Request from ground to board
// type G2B_Request struct {
// 	XMLName xml.Name `xml:"G2B_Request"`
// 	// Request from Ground to Board of DAS driving mode. If the Train Identification attribute is not
// 	// given, the Board will send the reply according to the train for which the communication layer
// 	// has been established. In the IM-RU setup between IM DAS-TS and RU DAS-TS, if the communication
// 	// layer is established for multiple trains, the train identifier must be given.
// 	DAS_DrivingModeRequest []DAS_DrivingModeRequest `xml:"DAS_DrivingModeRequest,omitempty"`
// 	// Request of Session Termination
// 	SessionTerminationRequest []SessionTerminationRequest `xml:"SessionTerminationRequest,omitempty"`
// 	// Plaintext message request
// 	PlaintextMessageRequest []PlaintextMessageRequest `xml:"PlaintextMessageRequest,omitempty"`
// 	// Request for Status Report messages
// 	StatusReportRequest []StatusReportRequest `xml:"StatusReportRequest,omitempty"`
// }

// // General_JP_Information: Data relevant to the whole JP
// type General_JP_Information struct {
// 	XMLName xml.Name `xml:"General_JP_Information"`
// 	// Network-Specific Parameter for General JP Information
// 	General_JP_Information_NSP []NSPList_ComplexType `xml:"General_JP_Information_NSP,omitempty"`
// }

// // GNSS_DynamicPosition_ComplexType: Complex Type for GNSS Dynamic Position
// type GNSS_DynamicPosition_ComplexType struct {
// 	// Latitude in signed degrees format (DDD.ddddd) in the WGS84 system.
// 	Latitude string `xml:"latitude,attr"`
// 	// Longitude in signed degrees format (DDD.ddddd) in the WGS84 system.
// 	Longitude string `xml:"longitude,attr"`
// 	// Altitude (m). Considering ETRS89 as reference.
// 	Altitude string `xml:"altitude,attr,omitempty"`
// 	// Direction in which the train is going. Expressed in degrees from North (0 = North, 90 = East,
// 	// 180 = South, 270 = West).
// 	Heading string `xml:"heading,attr,omitempty"`
// 	// Precision of the GNSS position. Expressed in metres.
// 	GNSS_precision string `xml:"GNSS_precision,attr,omitempty"`
// 	// Timestamp. Must be in UTC.
// 	Timestamp UTCDateTimeType `xml:"timestamp,attr,omitempty"`
// 	// Distance of GPS antenna from the front of the train
// 	AntennaDistanceFromFrontOfTrain DistanceType `xml:"antennaDistanceFromFrontOfTrain,attr,omitempty"`
// }

// // GradientAverage: Average gradient in the Segment Profile
// type GradientAverage struct {
// 	XMLName xml.Name `xml:"GradientAverage"`
// 	// Average gradient at the start of the Segment Profile
// 	GradientAverageStart *Gradient_ComplexType `xml:"GradientAverageStart"`
// 	// Change in average gradient
// 	GradientAverageChange []GradientAverageChange `xml:"GradientAverageChange,omitempty"`
// }

// // GradientAverageChange: Change in average gradient
// type GradientAverageChange struct {
// 	XMLName xml.Name `xml:"GradientAverageChange"`
// 	Gradient_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // Gradient_ComplexType: Complex Type for Gradient
// type Gradient_ComplexType struct {
// 	// Gradient value (Numeric value, mm by m)
// 	GradientValue GradientValueType `xml:"gradientValue,attr"`
// 	// Direction of the gradient (Downhill or Uphill)
// 	GradientDirection GradientDirectionType `xml:"gradientDirection,attr"`
// }

// // GradientSteepest: Steepest gradient in the Segment Profile
// type GradientSteepest struct {
// 	XMLName xml.Name `xml:"GradientSteepest"`
// 	// Steepest gradient at the start of the Segment Profile
// 	GradientSteepestStart *Gradient_ComplexType `xml:"GradientSteepestStart"`
// 	// Change in steepest gradient
// 	GradientSteepestChange []GradientSteepestChange `xml:"GradientSteepestChange,omitempty"`
// }

// // GradientSteepestChange: Change in steepest gradient
// type GradientSteepestChange struct {
// 	XMLName xml.Name `xml:"GradientSteepestChange"`
// 	Gradient_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // HandshakeAcknowledgement: Handshake between board and ground for SFERA connection or
// // disconnection
// type HandshakeAcknowledgement struct {
// 	XMLName xml.Name `xml:"HandshakeAcknowledgement"`
// 	// [SUBSET-126] Minimum time the ATO-OB shall wait before repeating a request if there was no
// 	// response
// 	Timeout_ATOTS_Response string `xml:"timeout_ATOTS_Response,attr,omitempty"`
// 	// [SUBSET-126] Reporting time cycle for triggering a Status Report for data reporting purposes
// 	ReportingTime string `xml:"reportingTime,attr,omitempty"`
// 	// Additional details
// 	AdditionalInfo string `xml:"additionalInfo,attr,omitempty"`
// 	// Minimum time to be waited between two position requests by board
// 	MinimumPositionUpdateRequestPeriod string `xml:"minimumPositionUpdateRequestPeriod,attr,omitempty"`
// 	// [SUBSET-126] Identifier of the ATO-TS
// 	ATOTS_ID string `xml:"ATOTS_ID,attr,omitempty"`
// 	// Information on whether related train information is supported
// 	RelatedTrainSupported RelatedTrainRequestType `xml:"relatedTrainSupported,attr,omitempty"`
// 	// Operating mode selected by the DAS-TS
// 	DAS_OperatingModeSelected *DAS_OperatingModeSelected `xml:"DAS_OperatingModeSelected"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone,omitempty"`
// }

// // HandshakeReject: Rejection of Handshake from Ground to Board
// type HandshakeReject struct {
// 	XMLName xml.Name `xml:"HandshakeReject"`
// 	// [SUBSET-126] Identifier of the ATO-TS
// 	ATOTS_ID string `xml:"ATOTS_ID,attr,omitempty"`
// 	// Reason for handshake rejection
// 	HandshakeRejectReason string `xml:"handshakeRejectReason,attr,omitempty"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone,omitempty"`
// }

// // HandshakeRequest: Handshake request from board to ground
// type HandshakeRequest struct {
// 	XMLName xml.Name `xml:"HandshakeRequest"`
// 	// Flag notifying whether the DAS is able to send Status Reports to ground.
// 	StatusReportsEnabled string `xml:"statusReportsEnabled,attr,omitempty"`
// 	// Additional details
// 	AdditionalInfo string `xml:"additionalInfo,attr,omitempty"`
// 	// Request for related train information
// 	RelatedTrainRequest RelatedTrainRequestType `xml:"relatedTrainRequest,attr,omitempty"`
// 	// Operating modes supported by the DAS-OB
// 	DAS_OperatingModesSupported []DAS_ModesComplexType `xml:"DAS_OperatingModesSupported,omitempty"`
// }

// // Interruption models the corresponding XSD type/element.
// type Interruption struct {
// 	XMLName xml.Name `xml:"Interruption"`
// 	// This element identifies the reason for an interruption of the train running, with a code
// 	// according to UIC450-2
// 	InterruptionReason DelayCode `xml:"InterruptionReason,omitempty"`
// 	// Date and Time when the Train was interrupted
// 	InterruptionDateTime string `xml:"InterruptionDateTime"`
// 	// The free text description of an interruption
// 	InterruptionDescription []FreeText `xml:"InterruptionDescription,omitempty"`
// 	// The link to the System Reference of the origin system
// 	InternalReferenceIdentifier FreeText `xml:"InternalReferenceIdentifier,omitempty"`
// }

// // InterruptionDuration: To specifiy the probable duration of the interruption
// type InterruptionDuration struct {
// 	XMLName xml.Name `xml:"InterruptionDuration"`
// 	// Forecasted earliest time for end of interruption
// 	EarliestEndDateTime string `xml:"EarliestEndDateTime,omitempty"`
// 	// Forecasted latest time for end of interruption
// 	LatestEndDateTime string `xml:"LatestEndDateTime,omitempty"`
// }

// // LocationIdent: Indication of the Railway or Customer Location
// type LocationIdent struct {
// 	// Identifies a County or State by code (ISO 3166-1)
// 	CountryCodeISO      *CountryCodeISO `xml:"CountryCodeISO"`
// 	LocationPrimaryCode Numeric1_5      `xml:"LocationPrimaryCode"`
// 	// Location Name in an offication language of the Country using the ISO Unicode alphabet
// 	PrimaryLocationName FreeText `xml:"PrimaryLocationName,omitempty"`
// 	// Code, Name and allocation company of Subsidiary Location
// 	LocationSubsidiaryIdentification *LocationSubsidiaryIdentification `xml:"LocationSubsidiaryIdentification,omitempty"`
// }

// // InterruptionPoint: Describes the interruption points with location and the reason for the
// // interruption. It must be sent in all Status reports until the interruption has been lifted.
// type InterruptionPoint struct {
// 	XMLName xml.Name `xml:"InterruptionPoint"`
// 	// Location of the interruption
// 	InterruptionLocation          *LocationIdent `xml:"InterruptionLocation"`
// 	DetailedDescriptionOfLocation FreeText       `xml:"DetailedDescriptionOfLocation,omitempty"`
// 	Interruption                  *Interruption  `xml:"Interruption"`
// 	// To specifiy the probable duration of the interruption
// 	InterruptionDuration *InterruptionDuration `xml:"InterruptionDuration,omitempty"`
// 	// Free Form Text
// 	Remarks []FreeText `xml:"Remarks,omitempty"`
// }

// // JourneyProfile: Journey Profile for a specific trip
// type JourneyProfile struct {
// 	XMLName xml.Name `xml:"JourneyProfile"`
// 	// [SUBSET-126] Status of the Journey Profile: "Valid": JP containing the data requested.
// 	// "Unavailable": JP specifies that the requested part of the Journey Profile is currently not
// 	// available yet (but all the previously sent JPs are still applicable). "Invalid": JP specifies
// 	// that the TP identifier asserted in the JPReq does not belong to the preceding JP already sent to
// 	// the ATO-OB. "Update": JP specifies that the Journey Profile has been updated by the TMS within
// 	// the current visibility of the ATO-OB.
// 	JP_Status string `xml:"JP_Status,attr,omitempty"`
// 	// Version of the Journey Proifile
// 	JP_Version string `xml:"JP_Version,attr,omitempty"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
// 	// Reference to a Segment Profile traversed by the journey
// 	SegmentProfileList []SegmentProfileList `xml:"SegmentProfileList,omitempty"`
// 	// Data relevant to the whole JP
// 	General_JP_Information *General_JP_Information `xml:"General_JP_Information,omitempty"`
// }

// // JP_ContextInformation: Context information related to Journey Profile to be displayed to drivers
// type JP_ContextInformation struct {
// 	XMLName xml.Name `xml:"JP_ContextInformation"`
// 	// Applicable Operational Train Number. In some cases this is different from the
// 	// TrainIdentification, for example when the OTN changes during the trip.
// 	Applicable_OTN             *Applicable_OTN       `xml:"Applicable_OTN,omitempty"`
// 	JP_ContextInformation_NSPs []NSPList_ComplexType `xml:"JP_ContextInformation_NSPs,omitempty"`
// }

// // JP_InUse: Journey Profile currently in use
// type JP_InUse struct {
// 	XMLName xml.Name `xml:"JP_InUse"`
// 	// Version of the Journey Proifile
// 	JP_Version string `xml:"JP_Version,attr"`
// }

// // JP_Request: Journey Profile request.
// type JP_Request struct {
// 	XMLName xml.Name `xml:"JP_Request"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
// 	// Segment Profile from which the JP is requested
// 	RequestFromSegmentProfile *SegmentProfileIdentification_ComplexType `xml:"RequestFromSegmentProfile,omitempty"`
// 	// Journey Profile currently in use
// 	JP_InUse *JP_InUse `xml:"JP_InUse,omitempty"`
// }

// // KilometreReferencePoint: Reference Point on the side of the track indicating the passing of a
// // certain distance, usually relative to a "zero" point of the line.
// type KilometreReferencePoint struct {
// 	XMLName xml.Name `xml:"KilometreReferencePoint"`
// 	SP_GenericPoint_ComplexType
// 	// KM reference point (in km)
// 	KmRef string `xml:"kmRef,attr"`
// 	// Ascending or Descending
// 	Ascending string `xml:"ascending,attr,omitempty"`
// 	// Additional details
// 	AdditionalInfo string `xml:"additionalInfo,attr,omitempty"`
// }

// // LastAbsoluteOdometryPoint: Position of the last absolute point for the odometry. In ETCS
// // odometry, this is the last known group of balises.
// type LastAbsoluteOdometryPoint struct {
// 	XMLName xml.Name `xml:"LastAbsoluteOdometryPoint"`
// 	SegmentProfileIdentification_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr,omitempty"`
// }

// // LastForceStatusChangeMessage: identifier of the last ForceStatusChange request message sent by
// // the board side of the communication. A ForceStatusChange Reply will only be sent if a
// // ForceStatusChange order has been generated since the last request was sent. Otherwise, a simple
// // "OK" message response is sent.
// type LastForceStatusChangeMessage struct {
// 	XMLName xml.Name `xml:"LastForceStatusChangeMessage"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr"`
// }

// // LastPassedKmRef: Last Km reference point passed by the train
// type LastPassedKmRef struct {
// 	XMLName xml.Name `xml:"LastPassedKmRef"`
// 	// KM reference point (in km)
// 	KmRef string `xml:"kmRef,attr"`
// }

// // LastPassedSignal: Last signal passed by the train
// type LastPassedSignal struct {
// 	XMLName xml.Name `xml:"LastPassedSignal"`
// 	// Identifier of the physical signal on the track
// 	Signal_ID_Physical string `xml:"signal_ID_Physical,attr"`
// 	// Identifier of the signal object
// 	Signal_ID_Object string `xml:"signal_ID_Object,attr,omitempty"`
// 	AtSignal         string `xml:"AtSignal,attr,omitempty"`
// }

// // LastPlaintextMessage: identifier of the last PlainTextMessage message received by the board side
// // of the communication. One or more PlainTextMessage replies will only be sent if more recent
// // advice is available. Otherwise, a simple "OK" message response is sent.
// type LastPlaintextMessage struct {
// 	XMLName xml.Name `xml:"LastPlaintextMessage"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr"`
// }

// // LastStatusReportMessage: identifier of the last StatusReport message received by the ground side
// // of the communication. One or more StatusReport replies will only be sent if more recent data is
// // available. Otherwise, a simple "OK" message response is sent.
// type LastStatusReportMessage struct {
// 	XMLName xml.Name `xml:"LastStatusReportMessage"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr"`
// }

// // LevelCrossingArea: Level crossing area, for driver information
// type LevelCrossingArea struct {
// 	XMLName xml.Name `xml:"LevelCrossingArea"`
// 	SP_GenericArea_ComplexType
// 	// Type of level crossing
// 	LevelCrossingType string `xml:"levelCrossingType,attr,omitempty"`
// 	// Is the level crossing protected
// 	ProtectedLevelCrossing string `xml:"protectedLevelCrossing,attr,omitempty"`
// 	// Description of level crossing
// 	LevelCrossingDescription string `xml:"levelCrossingDescription,attr,omitempty"`
// }

// // LimitedDynamicBrakeForceArea: [SUBSET-126] Dynamic Brake Limitation Area
// type LimitedDynamicBrakeForceArea struct {
// 	XMLName xml.Name `xml:"LimitedDynamicBrakeForceArea"`
// 	SP_GenericArea_ComplexType
// 	// [SUBSET-126] Maximum Brake Force (kN) to be applied by the Dynamic Brake
// 	DynamicBrakeForceLimit string `xml:"dynamicBrakeForceLimit,attr"`
// }

// // LineAndTrackIdentifier: Identification of line and track
// type LineAndTrackIdentifier struct {
// 	XMLName xml.Name `xml:"LineAndTrackIdentifier"`
// 	// Identification of line and track at the start of the Segment Profile
// 	LineAndTrackIdentifierStart *LineAndTrackIdentifier_ComplexType `xml:"LineAndTrackIdentifierStart"`
// 	// Change of identification of line and track in the course of the Segment Profile
// 	LineAndTrackIdentifierChange []LineAndTrackIdentifierChange `xml:"LineAndTrackIdentifierChange,omitempty"`
// }

// // LineAndTrackIdentifierChange: Change of identification of line and track in the course of the
// // Segment Profile
// type LineAndTrackIdentifierChange struct {
// 	XMLName xml.Name `xml:"LineAndTrackIdentifierChange"`
// 	LineAndTrackIdentifier_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // LineAndTrackIdentifier_ComplexType: Complex Type for Line and Track Identifier
// type LineAndTrackIdentifier_ComplexType struct {
// 	// Identification of the line. For example “115000” is the Line Identifier for the line from
// 	// Strasbourg to Basel.
// 	LineIdentifier string `xml:"lineIdentifier,attr,omitempty"`
// 	// Name of the line describing the connection stations. For example “Ligne de Strasbourg à Bale."
// 	LineName string `xml:"lineName,attr,omitempty"`
// 	// Identification of track on a line (ex : “V1” or “Voie 1” or “Line 1”)
// 	TrackIdentifier string `xml:"trackIdentifier,attr,omitempty"`
// }

// // LocationSubsidiaryCode: this element identifies a location as a part of primary location e.g. a
// // junction, a signal, a passing loop etc., It is unique when used in combination with a
// // “LocationPrimaryCode”
// type LocationSubsidiaryCode struct {
// 	XMLName xml.Name `xml:"LocationSubsidiaryCode"`
// 	// New codes added: 42 DIUM stations - Places of acceptance/delivery Station open into
// 	// international traffic of goods (tariff point included in DIUM) – consignment acceptance/delivery
// 	// station (loading points are excluded and covered by TypeCode 37). 43 Passengers cars public
// 	// loading Is a type of physical location on the open access network where passengers can put their
// 	// car on a carrying train 44 Passengers cars private loading Is a type of physical location
// 	// outside the open access network where passengers can put their car on a carrying train 45 Sewage
// 	// dump Place for cleaning purposes - disposal of the waste 46 Refuelling Point Location where
// 	// refuelling takes place 47 Mains Supply Location where energy supply can be provided for the
// 	// rolling stock e.g. preheating 48 Water Supply Location where water supply can be provided for
// 	// the rolling stock 49 Compressed plant Train on a track with motion stabled with external air
// 	// supply for braking systems 50 Indoor cleaning platform Cleaning point -interior 51 Car-wash
// 	// plant Cleaning point -outdoor 52 Short dry-cleaning track Cleaning point 53 Pollution protective
// 	// plate Track where floor that avoids pollution of the earth below 54 Sand-filling station
// 	// Location where sand is filled 55 Repair track Location where a train/wagon/engine can be
// 	// repaired 56 Signal box The location of a building containing signalling equipment 58 OSJD system
// 	// based location
// 	LocationSubsidiaryTypeCode string     `xml:"LocationSubsidiaryTypeCode,attr"`
// 	Value                      String1_10 `xml:",chardata"`
// }

// // LocationSubsidiaryIdentification: Code, Name and allocation company of Subsidiary Location
// type LocationSubsidiaryIdentification struct {
// 	XMLName xml.Name `xml:"LocationSubsidiaryIdentification"`
// 	// this element identifies a location as a part of primary location e.g. a junction, a signal, a
// 	// passing loop etc., It is unique when used in combination with a “LocationPrimaryCode”
// 	LocationSubsidiaryCode *LocationSubsidiaryCode `xml:"LocationSubsidiaryCode"`
// 	// Name of company who is responsible for allocation and maintenance of codes
// 	AllocationCompany CompanyCode `xml:"AllocationCompany"`
// 	// To be completed in an official language of the Country using the ISO Unicode alphabet
// 	LocationSubsidiaryName FreeText `xml:"LocationSubsidiaryName,omitempty"`
// }

// // LowAdhesion: Low Adhesion Status Report
// type LowAdhesion struct {
// 	XMLName xml.Name `xml:"LowAdhesion"`
// 	// Adhesion category aligned with the definition in EN15595 of the wheel-rail interface for the
// 	// lower boundaries (definition and values taken from SUBSET-126).
// 	LowAdhesionCategory AdhesionCategoryType `xml:"lowAdhesionCategory,attr,omitempty"`
// }

// // MaxSpeed: Maximum speed (in km/h)
// type MaxSpeed struct {
// 	XMLName xml.Name `xml:"MaxSpeed"`
// 	// speed (km/h)
// 	Speed SpeedType `xml:"speed,attr"`
// }

// // MaxSpeedAnnouncement: Announcement for a downline maximum speed (in km/h). Note that if the
// // distance to the downline MaxSpeed signal is known, it is recommended to use MaxSpeed together
// // with StartSignalApplication.
// type MaxSpeedAnnouncement struct {
// 	XMLName xml.Name `xml:"MaxSpeedAnnouncement"`
// 	// speed (km/h)
// 	Speed SpeedType `xml:"speed,attr"`
// }

// // MessageHeader: Header of the IRS90940 message. Contains metadata related to the message
// type MessageHeader struct {
// 	XMLName xml.Name `xml:"MessageHeader"`
// 	// SFERA version, major number. There should be backwards compatibility for messages with the same
// 	// major SFERA version.
// 	SFERA_version_major string `xml:"SFERA_version_major,attr"`
// 	// SFERA version, minor number.
// 	SFERA_version_minor string `xml:"SFERA_version_minor,attr"`
// 	// Identifier of the SFERA message
// 	Message_ID UUIDv4_Type `xml:"message_ID,attr"`
// 	// Timestamp. Must be in UTC.
// 	Timestamp UTCDateTimeType `xml:"timestamp,attr"`
// 	// ID of the device sending the message. For any given company, this ID must be unique. Maximum 128
// 	// characters.
// 	SourceDevice DeviceType `xml:"sourceDevice,attr"`
// 	// ID of the device to which the message is addressed. For any given company, this ID must be
// 	// unique. Maximum 128 characters. For G2B messages, if this attribute is empty, the message is
// 	// directed to all devices in the train. For B2G messages, this attribute can be ignored.
// 	DestinationDevice DeviceType `xml:"destinationDevice,attr,omitempty"`
// 	// ID of the message for which the current message is a reply
// 	Correlation_ID UUIDv4_Type `xml:"correlation_ID,attr,omitempty"`
// 	// [SUBSET-126] Operational train running number. NID_OPERATIONAL consists of up to 8 digits which
// 	// are entered left adjusted into the data field, the leftmost digit is the digit to be entered
// 	// first. In case the NID_OPERATIONAL is shorter than 8 digits, the remaining space is to be filled
// 	// with special character “F”.
// 	NID_OPERATIONAL string `xml:"NID_OPERATIONAL,attr,omitempty"`
// 	// [SUBSET-126] ETCS identity number uniquely defined for ERTMS/ETCS purposes. Values to be
// 	// assigned. 24 binary positions.
// 	NID_ENGINE string `xml:"NID_ENGINE,attr,omitempty"`
// 	// Number on 12 digits allocated to each rail vehicle to uniquely identify it from any other rail
// 	// vehicle
// 	EuropeanVehicleNumber string `xml:"EuropeanVehicleNumber,attr,omitempty"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// 	// This element identifies the message
// 	MessageReference *MessageReference `xml:"MessageReference,omitempty"`
// 	// The sender of the message
// 	Sender *Sender `xml:"Sender"`
// 	// Receiver of the message
// 	Recipient *Recipient `xml:"Recipient"`
// }

// // MessageReference: This element identifies the message
// type MessageReference struct {
// 	XMLName xml.Name `xml:"MessageReference"`
// 	// To indicate the message type transmitted or referred to. The following list was agreed within
// 	// the sector: 1000 ConsignmentOrderMessage 2001 PathCanceledMessage 2002 PathConfirmedMessage 2003
// 	// PathDetailsMessage 2004 PathDetailsRefusedMessage 2005 PathNotAvailableMessage 2006
// 	// PathRequestMessage 2007 ReceiptConfirmationMessage -- sector messages (Planning) -- 2500
// 	// PathCoordinationMessage 2501 PathSectionNotificationMessage -- sector message end -- 3003
// 	// TrainCompositionMessage 3006 TrainReadyMessage 4001 TrainDelayCauseMessage 4004
// 	// TrainRunningForecastMessage 4005 TrainRunningInformationMessage 4006
// 	// TrainRunningInterruptionMessage -- sector message (Operations) -- 4500
// 	// PassengerTrainCompositionProcessMessage 4501 RollingStockRestrictionMessage 4504
// 	// ChangeOfTrackMessage 4505 TrainJourneyModificationMessage -- sector message end -- 5001
// 	// AlertMessage 5002 WagonArrivalNoticeMessage 5003 WagonDeliveryNoticeMessage 5004
// 	// WagonDepartureNoticeMessage 5006 WagonETI_ETA_Message 5007 WagonExceptionMessage 5008
// 	// WagonExceptionReasonMessage -- sector message (Wagon interchange) -- 5009
// 	// WagonInterchangeNoticeMessage 5012 WagonReceivedAtInterchangeMessage 5013
// 	// WagonRefusedAtInterchangeMessage -- sector message end -- 5014 WagonReleaseNoticeMessage 5015
// 	// WagonYardArrivalMessage 5016 WagonYardDepartureMessage 6002 LocationFileDatasetMessage 6003
// 	// RollingStockDatasetMessage --sector (RU-RU) --- 5500 WagonPerformanceMessage -- sector end---
// 	// 6004 RollingStockDatasetQueryMessage -- Sector Messages TCR begin -- 6500 TCRMessage 6501
// 	// TCRResponseMessage 6502 TCRCanceledMessage -- Sector Messages TCR end -- -- sector (TrainID)
// 	// begin -- 8500 UpateLinkMessage 8501 ObjectInfoMessage -- sector end -- 9000 ErrorMessage
// 	MessageType MessageType `xml:"MessageType"`
// 	// Version of the Message Type
// 	MessageTypeVersion MessageTypeVersion `xml:"MessageTypeVersion"`
// 	// Identification of the Message
// 	MessageIdentifier FreeText `xml:"MessageIdentifier"`
// 	// Generated by the common Interface
// 	MessageDateTime string `xml:"MessageDateTime"`
// }

// // NetworkSpecificArea: Network-Specific Area
// type NetworkSpecificArea struct {
// 	XMLName xml.Name `xml:"NetworkSpecificArea"`
// 	SP_GenericArea_ComplexType
// 	// Name of the element
// 	Name string `xml:"name,attr,omitempty"`
// 	// Generic Network-Specific Parameter
// 	NetworkSpecificParameter []NetworkSpecificParameter `xml:"NetworkSpecificParameter,omitempty"`
// }

// // NetworkSpecificCharacteristics: Network-Specific Characteristics
// type NetworkSpecificCharacteristics struct {
// 	XMLName xml.Name `xml:"NetworkSpecificCharacteristics"`
// 	// Name of the element
// 	Name string `xml:"name,attr"`
// 	// Description of the element
// 	Description string `xml:"description,attr,omitempty"`
// 	// Start of the Network-Specific Characteristics
// 	NSP_Start *NSP_Start `xml:"NSP_Start"`
// 	// Change in the Network-Specific Characteristics
// 	NSP_Change []NSP_Change `xml:"NSP_Change,omitempty"`
// }

// // NetworkSpecificParameter: Generic Network-Specific Parameter
// type NetworkSpecificParameter struct {
// 	XMLName xml.Name `xml:"NetworkSpecificParameter"`
// 	// Name of the element
// 	Name string `xml:"name,attr"`
// 	// Value of the element
// 	Value string `xml:"value,attr"`
// 	// Description of the element
// 	Description string `xml:"description,attr,omitempty"`
// }

// // NetworkSpecificPoint: Network-Specific Point
// type NetworkSpecificPoint struct {
// 	XMLName xml.Name `xml:"NetworkSpecificPoint"`
// 	SP_GenericPoint_ComplexType
// 	// Name of the element
// 	Name string `xml:"name,attr,omitempty"`
// 	// Generic Network-Specific Parameter
// 	NetworkSpecificParameter []NetworkSpecificParameter `xml:"NetworkSpecificParameter,omitempty"`
// }

// // NextSP: Data for the next Segment Profile if it is in a different "Zone"
// type NextSP struct {
// 	XMLName xml.Name `xml:"NextSP"`
// 	// Segment Profile Identifier
// 	SP_ID string `xml:"SP_ID,attr,omitempty"`
// 	// Direction of application on the Segment Profile (Nominal, Reverse or Both)
// 	DirectionOfApplicationOnSP string `xml:"directionOfApplicationOnSP,attr,omitempty"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone"`
// 	// Information about ATO-Trackside Contact Information
// 	ATOTS_ContactInfo *ATOTS_ContactInfo `xml:"ATOTS_ContactInfo,omitempty"`
// }

// // NextTimingPointTrainInformation: Next timing points for the train
// type NextTimingPointTrainInformation struct {
// 	XMLName xml.Name `xml:"NextTimingPointTrainInformation"`
// 	// Timing Point Identifier
// 	TP_ID string `xml:"TP_ID,attr"`
// 	// Abbreviation of the Timing Point name
// 	TP_abbreviation string `xml:"TP_abbreviation,attr,omitempty"`
// 	// [SUBSET-126] Specifies if the Timing Point is a Passing Point, an operational Stopping Point or
// 	// a skipped Stopping Point
// 	TP_StopSkipPass string `xml:"TP_StopSkipPass,attr,omitempty"`
// 	// Forecasted arrival at the Timing Point
// 	ForecastedArrival UTCDateTimeType `xml:"forecastedArrival,attr,omitempty"`
// 	// Forecasted departure from the Timing Point
// 	ForecastedDeparture UTCDateTimeType `xml:"forecastedDeparture,attr,omitempty"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone"`
// }

// // NSP_Change: Change in the Network-Specific Characteristics
// type NSP_Change struct {
// 	XMLName xml.Name `xml:"NSP_Change"`
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// 	// Generic Network-Specific Parameter
// 	NetworkSpecificParameter []NetworkSpecificParameter `xml:"NetworkSpecificParameter,omitempty"`
// }

// // NSP_Start: Start of the Network-Specific Characteristics
// type NSP_Start struct {
// 	XMLName xml.Name `xml:"NSP_Start"`
// 	// Generic Network-Specific Parameter
// 	NetworkSpecificParameter []NetworkSpecificParameter `xml:"NetworkSpecificParameter,omitempty"`
// }

// // OtherSpecificSSP: Other Specific Static Speed Profile
// type OtherSpecificSSP struct {
// 	XMLName xml.Name `xml:"OtherSpecificSSP"`
// 	// [SUBSET-126] (If specific_SSP_Category = Other_SSP_*) It is the “other specific” SSP category
// 	// for which a different value for the static line speed exists. Used together with SS126_V_DIFF to
// 	// permit trains belonging to the corresponding “other international” train category to go faster
// 	// or lower than the “international basic static speed” given by SSP_Speed.
// 	SSP_NC_DIFF NC_TrainType `xml:"SSP_NC_DIFF,attr,omitempty"`
// 	// Code for the train category (freeform text, limited to 10 characters)
// 	TrainCategoryCode string `xml:"trainCategoryCode,attr,omitempty"`
// }

// // OTN_ID_ComplexType: Complex Type for the Operational Train Number (OTN) Identifier
// type OTN_ID_ComplexType struct {
// 	// Identifies a railway company (RU or IM) (UIC RICS Code: https://uic.org/rics)
// 	Company CompanyCode `xml:"Company"`
// 	// Operational Train Number
// 	TrainNumber string `xml:"TrainNumber"`
// 	// Additional identifier for train number
// 	AdditionalTrainNumber *AdditionalTrainNumber `xml:"AdditionalTrainNumber,omitempty"`
// 	// Is only used in the operational phase and refers to the date where the single train will start
// 	// the train journey
// 	StartDate string `xml:"StartDate"`
// }

// // OwnTrain: Contains the identification of the own train
// type OwnTrain struct {
// 	XMLName xml.Name `xml:"OwnTrain"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
// 	// Information about the current location and occupation of the train
// 	TrainLocationInformation *TrainLocationInformation `xml:"TrainLocationInformation,omitempty"`
// }

// // PermittedBrakingDistanceArea: Permitted Braking Distance Area
// type PermittedBrakingDistanceArea struct {
// 	XMLName xml.Name `xml:"PermittedBrakingDistanceArea"`
// 	SP_GenericArea_ComplexType
// 	// [SUBSET-126] Permitted Braking Distance value (in m)
// 	PermittedBrakingDistance DistanceType `xml:"permittedBrakingDistance,attr,omitempty"`
// 	// [SUBSET-126] Whether the permitted braking distance is to be achieved with the Service Brake or
// 	// Emergency Brake
// 	PBD_Brake string `xml:"PBD_Brake,attr,omitempty"`
// 	// [SUBSET-126] A single gradient value applicable for the calculation of the Permitted Braking
// 	// Distance
// 	Gradient_PBD GradientValueType `xml:"gradient_PBD,attr,omitempty"`
// 	// [SUBSET-126] Direction of the gradient for the Permitted Braking Distance
// 	GradientDirection_PBD GradientDirectionType `xml:"gradientDirection_PBD,attr,omitempty"`
// }

// // PlaintextMessage: Plaintext message
// type PlaintextMessage struct {
// 	XMLName xml.Name `xml:"PlaintextMessage"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// String containing the message
// 	MessageString string `xml:"messageString,attr"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// }

// // PlaintextMessageRequest: Plaintext message request
// type PlaintextMessageRequest struct {
// 	XMLName xml.Name `xml:"PlaintextMessageRequest"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
// 	// identifier of the last PlainTextMessage message received by the board side of the communication.
// 	// One or more PlainTextMessage replies will only be sent if more recent advice is available.
// 	// Otherwise, a simple "OK" message response is sent.
// 	LastPlaintextMessage *LastPlaintextMessage `xml:"LastPlaintextMessage,omitempty"`
// }

// // PositionSpeed: Position and speed of the train
// type PositionSpeed struct {
// 	XMLName xml.Name `xml:"PositionSpeed"`
// 	SegmentProfileIdentification_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// 	// speed (km/h)
// 	Speed SpeedType `xml:"speed,attr,omitempty"`
// 	// Error margin of the position of the head of the train. This margin added to the location
// 	// attribute gives the most downline position of the train possible at the moment of the measure.
// 	// Expressed in m
// 	Location_PositiveErrorMargin DistanceType `xml:"location_PositiveErrorMargin,attr,omitempty"`
// 	// Error margin of the position of the head of the train. This margin subtracted from the location
// 	// attribute gives the least downline position of the train possible at the moment of the measure.
// 	// Expressed in m
// 	Location_NegativeErrorMargin string `xml:"location_NegativeErrorMargin,attr,omitempty"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// 	// GNSS Dynamic Position
// 	GNSS_DynamicPosition *GNSS_DynamicPosition_ComplexType `xml:"GNSS_DynamicPosition,omitempty"`
// 	// Position of the last absolute point for the odometry. In ETCS odometry, this is the last known
// 	// group of balises.
// 	LastAbsoluteOdometryPoint *LastAbsoluteOdometryPoint `xml:"LastAbsoluteOdometryPoint,omitempty"`
// }

// // PositionSpeedRequest: Request from Board to Ground of position and speed. If the Train
// // Identification attribute is not given, the Ground will send the reply according to the train for
// // which the communication layer has been established. In the IM-RU setup between IM DAS-TS and RU
// // DAS-TS, if the communication layer is established for multiple trains, the train identifier must
// // be given.
// type PositionSpeedRequest struct {
// 	XMLName xml.Name `xml:"PositionSpeedRequest"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// }

// // PowerAdvice: Advice on minimising or maximising the power consumption
// type PowerAdvice struct {
// 	XMLName xml.Name `xml:"PowerAdvice"`
// 	// Change in power consumption: maximise, minimise or "no advice" ("no advice" cancels a previous
// 	// advice)
// 	PowerConsumptionChange string `xml:"powerConsumptionChange,attr"`
// }

// // PowerConstraint: Power Constraint signal
// type PowerConstraint struct {
// 	XMLName xml.Name `xml:"PowerConstraint"`
// 	// Instruction for Power Constraint
// 	PowerConstraintInstruction string `xml:"powerConstraintInstruction,attr"`
// }

// // PreviousTimingPoint: [SUBSET-126] Information concerning the most recent Timing Point considered
// // as either reached or passed
// type PreviousTimingPoint struct {
// 	XMLName xml.Name `xml:"PreviousTimingPoint"`
// 	TimingPointIdentification_ComplexType
// 	// [SUBSET-126] Qualifier to indicate if train has stopped at the TP, has departed from the TP or
// 	// has passed the TP.
// 	TP_Pass_Stop_Depart string `xml:"TP_Pass_Stop_Depart,attr,omitempty"`
// 	// [SUBSET-126] (If TP_Pass_Stop_Depart = Stopped) This qualifier specifies if the train has
// 	// stopped accurately or not at the Operational Stopping Point.
// 	TP_Accurate_Stopping string `xml:"TP_Accurate_Stopping,attr,omitempty"`
// }

// // PreviousTimingPointTrainInformation: Previous timing points passed by the train
// type PreviousTimingPointTrainInformation struct {
// 	XMLName xml.Name `xml:"PreviousTimingPointTrainInformation"`
// 	// Timing Point Identifier
// 	TP_ID string `xml:"TP_ID,attr"`
// 	// Abbreviation of the Timing Point name
// 	TP_abbreviation string `xml:"TP_abbreviation,attr,omitempty"`
// 	// [SUBSET-126] Specifies if the Timing Point is a Passing Point, an operational Stopping Point or
// 	// a skipped Stopping Point
// 	TP_StopSkipPass string `xml:"TP_StopSkipPass,attr,omitempty"`
// 	// Arrival time at the Timing Point
// 	ArrivalTime UTCDateTimeType `xml:"arrivalTime,attr,omitempty"`
// 	// Expected time of departure at Timing Point expressed in UTC
// 	DepartureTime UTCDateTimeType `xml:"departureTime,attr,omitempty"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone"`
// }

// // RatedVoltage: Rated voltage of the traction system (V)
// type RatedVoltage struct {
// 	XMLName xml.Name `xml:"RatedVoltage"`
// 	// Rated voltage at the start of the Segment Profile
// 	RatedVoltageStart *Voltage_ComplexType `xml:"RatedVoltageStart"`
// 	// Rated voltage change
// 	RatedVoltageChange []RatedVoltageChange `xml:"RatedVoltageChange,omitempty"`
// }

// // RatedVoltageChange: Rated voltage change
// type RatedVoltageChange struct {
// 	XMLName xml.Name `xml:"RatedVoltageChange"`
// 	Voltage_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // Recipient: Receiver of the message
// type Recipient struct {
// 	XMLName xml.Name `xml:"Recipient"`
// 	// Number of a Common Interface Instance for the same Company
// 	CI_InstanceNumber Numeric2_2  `xml:"CI_InstanceNumber,attr,omitempty"`
// 	Value             CompanyCode `xml:",chardata"`
// }

// // TotalCurrent_ComplexType: Complex Type for Total Current temporary constraints
// type TotalCurrent_ComplexType struct {
// 	// Maximum current value (A)
// 	MaxCurValue CurrentType `xml:"maxCurValue,attr"`
// }

// // RelatedTrainInformation: Contains information about the own train and the related trains, i.e.
// // trains that share the same trainpath (infrastructure) as the own train.
// type RelatedTrainInformation struct {
// 	XMLName xml.Name `xml:"RelatedTrainInformation"`
// 	// Reason why the related train information is restricted (e.g. missing train location system, or
// 	// train in another country)
// 	RestrictionReason string `xml:"RestrictionReason,attr,omitempty"`
// 	// Contains the identification of the own train
// 	OwnTrain *OwnTrain `xml:"OwnTrain"`
// 	// Trains that share the same trainpath as your train, which will use the same infrastructure
// 	// before you.
// 	TrainsBeforeOwn *TrainsBeforeOwn `xml:"TrainsBeforeOwn,omitempty"`
// 	// Trains that share the same trainpath as your train, which are behind you. These trains will use
// 	// the same infrastructure after you.
// 	TrainsAfterOwn *TrainsAfterOwn `xml:"TrainsAfterOwn,omitempty"`
// }

// // Reported_DAS_DrivingMode: Report of DAS Driving Mode
// type Reported_DAS_DrivingMode struct {
// 	XMLName xml.Name `xml:"Reported_DAS_DrivingMode"`
// 	// DAS driving mode (see enumeration)
// 	DAS_drivingMode string `xml:"DAS_drivingMode,attr"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// }

// // SegmentProfileIdentification_ComplexType: Complex type containing the data for the identification
// // of Segment Profiles
// type SegmentProfileIdentification_ComplexType struct {
// 	// Segment Profile Identifier
// 	SP_ID string `xml:"SP_ID,attr"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone"`
// }

// // RollingStockApplicability: Criteria for applicability related to rolling stock
// type RollingStockApplicability struct {
// 	XMLName xml.Name `xml:"RollingStockApplicability"`
// 	// Free-form text describing the rolling stock. Used for driver information.
// 	RollingStockType string `xml:"rollingStockType,attr"`
// 	// Minimum number of a certain type of rolling stock
// 	MinimumNumberOfRSType string `xml:"minimumNumberOfRSType,attr,omitempty"`
// 	// Maximum number of a certain type of rolling stock
// 	MaximumNumberOfRSType string `xml:"maximumNumberOfRSType,attr,omitempty"`
// }

// // SegmentClaim: Provides information about the occupation of a Segment Profile (or part of it) by
// // the train.
// type SegmentClaim struct {
// 	XMLName xml.Name `xml:"SegmentClaim"`
// 	SegmentProfileIdentification_ComplexType
// 	// Identifier of the Segment Profile version, major number.
// 	SP_VersionMajor string `xml:"SP_VersionMajor,attr"`
// 	// Identifier of the Segment Profile version, minor number.
// 	SP_VersionMinor string `xml:"SP_VersionMinor,attr"`
// 	// Travelling direction of the Segment Profile (Nominal or Reverse). If the travelling direction is
// 	// Reverse, all locations will apply from the end of the segment. Note that Signals and Stopping
// 	// Points only apply in the nominal direction of the SP.
// 	SP_Direction string `xml:"SP_Direction,attr,omitempty"`
// 	// [SUBSET-126] Distance of the element's start location from the beginning of the Segment Profile
// 	// (m)
// 	StartLocation DistanceType `xml:"startLocation,attr,omitempty"`
// 	// [SUBSET-126] Distance of the element's end location from the beginning of the Segment Profile
// 	// (m)
// 	EndLocation DistanceType `xml:"endLocation,attr,omitempty"`
// 	// [SUBSET-126] Specifies if the element starts, ends, starts and ends or covers the whole
// 	// concerning Segment Profile
// 	StartEndQualifier string `xml:"startEndQualifier,attr,omitempty"`
// 	// Status of the train in the segment
// 	SegmentOccupationStatus *SegmentOccupationStatus `xml:"SegmentOccupationStatus"`
// }

// // SegmentOccupationStatus: Status of the train in the segment
// type SegmentOccupationStatus struct {
// 	XMLName xml.Name `xml:"SegmentOccupationStatus"`
// 	// Gives an indication of the status or occupation of the piece of the segment
// 	ClaimType ClaimType `xml:"ClaimType"`
// 	// Gives an indication if a train is entering of leaving your trainpath at a certain location
// 	LeavingEnteringStatus LeavingEnteringStatus `xml:"LeavingEnteringStatus,omitempty"`
// 	// Only for trains in front of the own train, gives an indication if the train is going in the same
// 	// direction as you (true) or the opposite direction (false).
// 	SameDirection string `xml:"SameDirection,omitempty"`
// }

// // SegmentProfile: Profile for a section of railway track
// type SegmentProfile struct {
// 	XMLName xml.Name `xml:"SegmentProfile"`
// 	SegmentProfileIdentification_ComplexType
// 	// Identifier of the Segment Profile version, major number.
// 	SP_VersionMajor string `xml:"SP_VersionMajor,attr"`
// 	// Identifier of the Segment Profile version, minor number.
// 	SP_VersionMinor string `xml:"SP_VersionMinor,attr"`
// 	// Length of the section of railway covered by the Segment Profile (m)
// 	SP_Length DistanceType `xml:"SP_Length,attr"`
// 	// Human-readable name of the Segment Profile. To be used only for debugging purposes.
// 	SP_Name string `xml:"SP_Name,attr,omitempty"`
// 	// [SUBSET-126] Status of the Segment Profile (Valid or Invalid). From SUBSET-125: 10.1.3.23 The
// 	// Segment Profile shall indicate its status to the ATO-OB (Segment Profile Status): a) Valid:
// 	// specifies that the data sent correspond to a SPReq sent; b) Invalid: the requested Segment
// 	// Profile is not available in the ATO-TS.
// 	SP_Status string `xml:"SP_Status,attr,omitempty"`
// 	// [SUBSET-126] Distance to stop the train in rear of the End of Authority (in m)
// 	Distance_EoA_Offset DistanceType `xml:"distance_EoA_Offset,attr,omitempty"`
// 	// [SUBSET-126] Offset to add to the UTC time in order to calculate the local time. Unsigned value
// 	// Resolution: 15 min 0 = UTC - 14:00 56 = UTC ± 0 112 = UTC + 14:00 113 - 127 = Spare
// 	UTC_Offset string `xml:"UTC_Offset,attr,omitempty"`
// 	// [SUBSET-126] Altitude (in m) at the beginning of the SP. Considering ETRS89 as reference.
// 	// Starting at -1000m.
// 	SP_Altitude string `xml:"SP_Altitude,attr,omitempty"`
// 	// Gauge of the segment profile (in mm). Is needed to calculate rolling resistance.
// 	SP_Gauge string `xml:"SP_Gauge,attr,omitempty"`
// 	// Maximum lateral deviation in Segment Profile (in metres). Defines the maximum allowed lateral
// 	// deviation of of any middle point on the used track to straight lines connecting the virtual
// 	// balises
// 	SP_MaxLateralDeviation string `xml:"SP_MaxLateralDeviation,attr,omitempty"`
// 	// Notable points present in the Segment Profile (e.g. timing points, signals, balises)
// 	SP_Points *SP_Points `xml:"SP_Points,omitempty"`
// 	// Notable areas present in the Segment Profile (e.g. stations, platforms, tunnels...)
// 	SP_Areas *SP_Areas `xml:"SP_Areas,omitempty"`
// 	// Characteristics of the Segment Profile (e.g. gradient, curve, cant, voltage...)
// 	SP_Characteristics *SP_Characteristics `xml:"SP_Characteristics,omitempty"`
// 	// Context information related to Journey Profile to be displayed to drivers
// 	SP_ContextInformation *SP_ContextInformation `xml:"SP_ContextInformation,omitempty"`
// 	// Data for the next Segment Profile if it is in a different "Zone"
// 	NextSP *NextSP `xml:"NextSP,omitempty"`
// }

// // SegmentProfileInfo: Information about the Segment Profile
// type SegmentProfileInfo struct {
// 	XMLName xml.Name `xml:"SegmentProfileInfo"`
// 	SegmentProfileIdentification_ComplexType
// 	// [SUBSET-126] Distance of the element's start location from the beginning of the Segment Profile
// 	// (m)
// 	StartLocation DistanceType `xml:"startLocation,attr,omitempty"`
// 	// [SUBSET-126] Distance of the element's end location from the beginning of the Segment Profile
// 	// (m)
// 	EndLocation DistanceType `xml:"endLocation,attr,omitempty"`
// 	// Data for the next Segment Profile if it is in a different "Zone"
// 	NextSP *NextSP `xml:"NextSP,omitempty"`
// }

// // SegmentProfileList: Reference to a Segment Profile traversed by the journey
// type SegmentProfileList struct {
// 	XMLName xml.Name `xml:"SegmentProfileList"`
// 	SegmentProfileIdentification_ComplexType
// 	// Identifier of the Segment Profile version, major number.
// 	SP_VersionMajor string `xml:"SP_VersionMajor,attr"`
// 	// Identifier of the Segment Profile version, minor number.
// 	SP_VersionMinor string `xml:"SP_VersionMinor,attr"`
// 	// Travelling direction of the Segment Profile (Nominal or Reverse). If the travelling direction is
// 	// Reverse, all locations will apply from the end of the segment. Note that Signals and Stopping
// 	// Points only apply in the nominal direction of the SP.
// 	SP_Direction string `xml:"SP_Direction,attr,omitempty"`
// 	// Constraints at specific Timing Points in the Segment Profile (arrival and departure times)
// 	TimingPointConstraints []TimingPointConstraints `xml:"TimingPointConstraints,omitempty"`
// 	// Temporary Constraints in the Segment Profile
// 	TemporaryConstraints []TemporaryConstraints `xml:"TemporaryConstraints,omitempty"`
// 	// Description of the impact on allowed speed, driving reactions of the driver, and traction type
// 	// induced by temporary signals, permanent signals with an aspect different from the least
// 	// restrictive aspect described in SP / SP_Points / Signal, or temporarily cancelled signals on the
// 	// ground.
// 	TemporarySignalConstraints []TemporarySignalConstraints `xml:"TemporarySignalConstraints,omitempty"`
// 	// Reference to a Train Characteristics element (sent separately). All TCs are taken from that
// 	// element; some can be overridden locally (e.g. length, weight, max speed).
// 	TrainCharacteristicsRef []TrainCharacteristicsRef `xml:"TrainCharacteristicsRef,omitempty"`
// 	// Reason for Journey Profile Change
// 	ChangeReason *ChangeReason `xml:"ChangeReason,omitempty"`
// 	// Context information related to Journey Profile to be displayed to drivers
// 	JP_ContextInformation *JP_ContextInformation `xml:"JP_ContextInformation,omitempty"`
// }

// // Sender: The sender of the message
// type Sender struct {
// 	XMLName xml.Name `xml:"Sender"`
// 	// Number of a Common Interface Instance for the same Company
// 	CI_InstanceNumber Numeric2_2  `xml:"CI_InstanceNumber,attr,omitempty"`
// 	Value             CompanyCode `xml:",chardata"`
// }

// // SessionTermination: Session termination
// type SessionTermination struct {
// 	XMLName xml.Name `xml:"SessionTermination"`
// 	// Reason for session termination
// 	SessionTerminationReason string `xml:"sessionTerminationReason,attr,omitempty"`
// }

// // SessionTerminationRequest: Request of Session Termination
// type SessionTerminationRequest struct {
// 	XMLName xml.Name `xml:"SessionTerminationRequest"`
// 	// Additional details
// 	AdditionalInfo string `xml:"additionalInfo,attr,omitempty"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// }

// // SFERA_B2G_EventMessage: IRS90940 Board to Ground Event Message. In standard SFERA implementations
// // via the Common Interface, the value "SF00" will be used for the MessageType.
// type SFERA_B2G_EventMessage struct {
// 	XMLName xml.Name `xml:"SFERA_B2G_EventMessage"`
// 	// Header of the IRS90940 message. Contains metadata related to the message
// 	MessageHeader *MessageHeader `xml:"MessageHeader"`
// 	// Status Report from Board to Ground
// 	B2G_StatusReport []B2G_StatusReport `xml:"B2G_StatusReport,omitempty"`
// 	// Session termination
// 	SessionTermination *SessionTermination `xml:"SessionTermination"`
// }

// // SFERA_B2G_ReplyMessage: IRS90940 Board to Ground Reply Message. In standard SFERA implementations
// // via the Common Interface, the value "SF02" will be used for the MessageType.
// type SFERA_B2G_ReplyMessage struct {
// 	XMLName xml.Name `xml:"SFERA_B2G_ReplyMessage"`
// 	// Header of the IRS90940 message. Contains metadata related to the message
// 	MessageHeader *MessageHeader `xml:"MessageHeader"`
// 	// Payload of board to ground reply
// 	B2G_ReplyPayload *B2G_ReplyPayload `xml:"B2G_ReplyPayload"`
// }

// // SFERA_B2G_RequestMessage: IRS90940 Board to Ground Request Message. In standard SFERA
// // implementations via the Common Interface, the value "SF01" will be used for the MessageType.
// type SFERA_B2G_RequestMessage struct {
// 	XMLName xml.Name `xml:"SFERA_B2G_RequestMessage"`
// 	// Header of the IRS90940 message. Contains metadata related to the message
// 	MessageHeader *MessageHeader `xml:"MessageHeader"`
// 	// Handshake request from board to ground
// 	HandshakeRequest *HandshakeRequest `xml:"HandshakeRequest"`
// 	// Request from board to ground
// 	B2G_Request *B2G_Request `xml:"B2G_Request"`
// }

// // SFERA_G2B_EventMessage: IRS90940 Ground to Board Event Message. In standard SFERA implementations
// // via the Common Interface, the value "SF10" will be used for the MessageType.
// type SFERA_G2B_EventMessage struct {
// 	XMLName xml.Name `xml:"SFERA_G2B_EventMessage"`
// 	// Header of the IRS90940 message. Contains metadata related to the message
// 	MessageHeader *MessageHeader `xml:"MessageHeader"`
// 	// Payload of ground to board event
// 	G2B_EventPayload *G2B_EventPayload `xml:"G2B_EventPayload"`
// }

// // SFERA_G2B_ReplyMessage: IRS90940 Ground to Board Reply Message. In standard SFERA implementations
// // via the Common Interface, the value "SF12" will be used for the MessageType.
// type SFERA_G2B_ReplyMessage struct {
// 	XMLName xml.Name `xml:"SFERA_G2B_ReplyMessage"`
// 	// Header of the IRS90940 message. Contains metadata related to the message
// 	MessageHeader *MessageHeader `xml:"MessageHeader"`
// 	// Handshake between board and ground for SFERA connection or disconnection
// 	HandshakeAcknowledgement *HandshakeAcknowledgement `xml:"HandshakeAcknowledgement"`
// 	// Rejection of Handshake from Ground to Board
// 	HandshakeReject *HandshakeReject `xml:"HandshakeReject"`
// 	// Payload of ground to board reply
// 	G2B_ReplyPayload *G2B_ReplyPayload `xml:"G2B_ReplyPayload"`
// }

// // SFERA_G2B_RequestMessage: IRS90940 Ground to Board Request Message. In standard SFERA
// // implementations via the Common Interface, the value "SF11" will be used for the MessageType.
// type SFERA_G2B_RequestMessage struct {
// 	XMLName xml.Name `xml:"SFERA_G2B_RequestMessage"`
// 	// Header of the IRS90940 message. Contains metadata related to the message
// 	MessageHeader *MessageHeader `xml:"MessageHeader"`
// 	// Request from ground to board
// 	G2B_Request *G2B_Request `xml:"G2B_Request"`
// }

// // Signal_ComplexType: Complex type for the definition of a signal
// type Signal_ComplexType struct {
// 	// Direction of application on the Segment Profile (Nominal, Reverse or Both)
// 	DirectionOfApplicationOnSP string `xml:"directionOfApplicationOnSP,attr,omitempty"`
// 	// Signal identification. The identification may include the identifier of the physical signal of
// 	// the track and the identifier of the single signal object that contains the specific information
// 	// (and possibly application). For a single physical signal there may be several signal objects.
// 	// The pair (signal_ID_Physical, signal_ID_Object) must be unique within an SP. the
// 	// signal_ID_Object must be documented if 2 or more signal objects are related to the
// 	// signal_ID_Physical.
// 	Signal_ID *Signal_ID `xml:"Signal_ID"`
// 	// Constraints for the application of the signal
// 	SignalApplication *SignalApplication `xml:"SignalApplication,omitempty"`
// 	// Information contained by the signal that impacts the driving profile
// 	SignalInformation *SignalInformation `xml:"SignalInformation,omitempty"`
// 	// Function of the signal. The list of values is taken from EULynx SignalFunctionTypes
// 	// (https://dataprep.eulynx.eu/2021-05/index.htm?goto=3:1:1:4:3762), with the addition of
// 	// laneChange.
// 	SignalFunction []SignalFunction `xml:"SignalFunction,omitempty"`
// 	// Physical Characteristics of the Signal
// 	SignalPhysicalCharacteristics *SignalPhysicalCharacteristics `xml:"SignalPhysicalCharacteristics,omitempty"`
// 	// Network Specific Parameters in Signal
// 	Signal_NSPs []NSPList_ComplexType `xml:"Signal_NSPs,omitempty"`
// }

// // Signal_ID: Signal identification. The identification may include the identifier of the physical
// // signal of the track and the identifier of the single signal object that contains the specific
// // information (and possibly application). For a single physical signal there may be several signal
// // objects. The pair (signal_ID_Physical, signal_ID_Object) must be unique within an SP. the
// // signal_ID_Object must be documented if 2 or more signal objects are related to the
// // signal_ID_Physical.
// type Signal_ID struct {
// 	XMLName xml.Name `xml:"Signal_ID"`
// 	// Identifier of the physical signal on the track
// 	Signal_ID_Physical string `xml:"signal_ID_Physical,attr"`
// 	// Identifier of the signal object
// 	Signal_ID_Object string `xml:"signal_ID_Object,attr,omitempty"`
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // SignalApplication: Constraints for the application of the signal
// type SignalApplication struct {
// 	XMLName xml.Name `xml:"SignalApplication"`
// 	// Location where the signal starts to be applied. If not specified, the signal is applied starting
// 	// from the signal location. This may also be in a different SP: then the SP_ID has to be specified
// 	// (otherwise the current SP is considered).
// 	StartSignalApplication *StartSignalApplication `xml:"StartSignalApplication,omitempty"`
// 	// Location where the signal stops to be applied. If not specified, the signal is applied
// 	// indefinitely, until superseded by another signal. This may also be in a different SP: then the
// 	// SP_ID has to be specified (otherwise the current SP is considered).
// 	EndSignalApplication *EndSignalApplication `xml:"EndSignalApplication,omitempty"`
// 	// List of ATP systems to which the signal is applied (default: all ATP systems supported in the
// 	// SP)
// 	Signal_ATP_Systems *ATPSystems_ComplexType `xml:"Signal_ATP_Systems,omitempty"`
// 	// Train Categories to which the signal is applied
// 	SignalTrainCategory []SignalTrainCategory `xml:"SignalTrainCategory,omitempty"`
// 	// Minimum and maximum train length to which the signal is applied.
// 	SignalTrainLength *SignalTrainLength `xml:"SignalTrainLength,omitempty"`
// 	// The effects of this signal element are to be taken into account only if this SP is in the
// 	// SP_List after the current SP.
// 	ApplicableIfTrainRunsOnSP []ApplicableIfTrainRunsOnSP `xml:"ApplicableIfTrainRunsOnSP,omitempty"`
// }

// // SignalInformation: Information contained by the signal that impacts the driving profile
// type SignalInformation struct {
// 	XMLName xml.Name `xml:"SignalInformation"`
// 	// Indicates if the train must stop before crossing the signal, for example if the driver must
// 	// receive an order from the signalling responsible.
// 	StopBeforePassingSignal *StopBeforePassingSignal `xml:"StopBeforePassingSignal"`
// 	// Maximum speed (in km/h)
// 	MaxSpeed *MaxSpeed `xml:"MaxSpeed"`
// 	// Announcement for a downline maximum speed (in km/h). Note that if the distance to the downline
// 	// MaxSpeed signal is known, it is recommended to use MaxSpeed together with
// 	// StartSignalApplication.
// 	MaxSpeedAnnouncement *MaxSpeedAnnouncement `xml:"MaxSpeedAnnouncement"`
// 	// Speed range (in km/h) to avoid
// 	AvoidSpeedRange *AvoidSpeedRange `xml:"AvoidSpeedRange"`
// 	// Traction type transition signal
// 	TractionTypeTransition *TractionTypeTransition `xml:"TractionTypeTransition"`
// 	// Power Constraint signal
// 	PowerConstraint *PowerConstraint `xml:"PowerConstraint"`
// 	// Sign on the side of the track at stopping distance from the next station, from which point the
// 	// driver will start to brake to be at the permitted speed in the station. The station is
// 	// identified by the Timing Point ID.
// 	StationAheadSign *TimingPointIdentification_ComplexType `xml:"StationAheadSign"`
// 	// This type of signal will be used in the case that the signal will present an aspect that is
// 	// different from the aspect announced by a previous signal.
// 	CancelPreviousSignal *CancelPreviousSignal `xml:"CancelPreviousSignal"`
// 	// Network-Specific Parameters in SignalInformation
// 	SignalInformation_NSPs *NSPList_ComplexType `xml:"SignalInformation_NSPs"`
// }

// // SignalPhysicalCharacteristics: Physical Characteristics of the Signal
// type SignalPhysicalCharacteristics struct {
// 	XMLName xml.Name `xml:"SignalPhysicalCharacteristics"`
// 	// Type of signal from which the physical characteristics of the signal can be deduced. This will
// 	// be network specific, as no standardised solution has been identified by the SFERA project.
// 	// Examples: - CTC (Centralized Traffic Control Signal) - Permissive (Permissive Signal) -
// 	// Permanent signal - Maximum Sign - Announcement Sign - Disque - Cantonnement BAL 3 aspects
// 	SignalType string `xml:"signalType,attr,omitempty"`
// 	// Horizontal offset (in m) from the center of the track to the middle of the object. Negative is
// 	// left and positive is right.
// 	HorizontalOffset DistanceType `xml:"horizontalOffset,attr,omitempty"`
// 	// Vertical offset (in m) from the top of the track to the middle of the object.
// 	VerticalOffset DistanceType `xml:"verticalOffset,attr,omitempty"`
// 	// Network Specific Parameters in Signal Physical Characteristics
// 	SignalPhysicalCharacteristics_NSPs []SignalPhysicalCharacteristics_NSPs `xml:"SignalPhysicalCharacteristics_NSPs,omitempty"`
// }

// // SignalPhysicalCharacteristics_NSPs: Network Specific Parameters in Signal Physical
// // Characteristics
// type SignalPhysicalCharacteristics_NSPs struct {
// 	XMLName xml.Name `xml:"SignalPhysicalCharacteristics_NSPs"`
// 	// Name of the element
// 	Name string `xml:"name,attr,omitempty"`
// 	// Generic Network-Specific Parameter
// 	NetworkSpecificParameter []NetworkSpecificParameter `xml:"NetworkSpecificParameter,omitempty"`
// }

// // SignalTrainCategory: Train Categories to which the signal is applied
// type SignalTrainCategory struct {
// 	XMLName xml.Name `xml:"SignalTrainCategory"`
// 	// Cant Deficiency train category to which the signal is applied
// 	Signal_CantDeficiencyTrainCategory *Signal_CantDeficiencyTrainCategory `xml:"Signal_CantDeficiencyTrainCategory"`
// 	// Other train category to which the signal is applied
// 	Signal_OtherTrainCategory *Signal_OtherTrainCategory `xml:"Signal_OtherTrainCategory"`
// }

// // SignalTrainLength: Minimum and maximum train length to which the signal is applied.
// type SignalTrainLength struct {
// 	XMLName xml.Name `xml:"SignalTrainLength"`
// 	// Minimum Train Length for which stopping point is valid. Expressed in m
// 	MinimumTrainLength string `xml:"minimumTrainLength,attr,omitempty"`
// 	// Maximum Train Length for which stopping point is valid. Expressed in m
// 	MaximumTrainLength string `xml:"maximumTrainLength,attr,omitempty"`
// }

// // SP_Areas: Notable areas present in the Segment Profile (e.g. stations, platforms, tunnels...)
// type SP_Areas struct {
// 	XMLName xml.Name `xml:"SP_Areas"`
// 	// Area alongside a railway track providing convenient access to trains
// 	Platform []SP_GenericArea_ComplexType `xml:"Platform,omitempty"`
// 	// Tunnel
// 	Tunnel []Tunnel `xml:"Tunnel,omitempty"`
// 	// Decisive Gradient Area. The decisive gradient is the steepest gradient in front of potential
// 	// stopping points on a line like signals, level crossings and stations. It is relevant, if it goes
// 	// downhill, for the calculation of either the minimum braking percentage or the speed limit for a
// 	// train because in case of reduced braking capacity, the train might not be able to brake
// 	// properly, therefore, the driver has to adjust the speed. If it goes uphill, it is relevant to
// 	// calculate the towing load, for example, to make sure, with the available engine power, that the
// 	// locomotive can get away after it stopped.
// 	DecisiveGradient []DecisiveGradient `xml:"DecisiveGradient,omitempty"`
// 	// It is an area shown to drivers where they should not stop the train. It is typically the case in
// 	// some of the tunnels, or some of the bridges, or some portions of them.
// 	NoStoppingArea []SP_GenericArea_ComplexType `xml:"NoStoppingArea,omitempty"`
// 	// [SUBSET-126] Speed Profile restriction to be applied to specific axle loads
// 	AxleLoadSpeedProfile []AxleLoadSpeedProfile `xml:"AxleLoadSpeedProfile,omitempty"`
// 	// Permitted Braking Distance Area
// 	PermittedBrakingDistanceArea []PermittedBrakingDistanceArea `xml:"PermittedBrakingDistanceArea,omitempty"`
// 	// [SUBSET-126] Switch Off Regenerative Brake Area
// 	SwitchOffRegenerativeBrakeArea []SP_GenericArea_ComplexType `xml:"SwitchOffRegenerativeBrakeArea,omitempty"`
// 	// [SUBSET-126] Switch Off Eddy Current Brake for Service Brake Area
// 	SwitchOffEddyCurrentBrakeArea []SP_GenericArea_ComplexType `xml:"SwitchOffEddyCurrentBrakeArea,omitempty"`
// 	// [SUBSET-126] Switch Off Eddy Current Brake for Emergency Brake Area
// 	SwitchOffEddyCurrentEmergencyBrakeArea []SP_GenericArea_ComplexType `xml:"SwitchOffEddyCurrentEmergencyBrakeArea,omitempty"`
// 	// [SUBSET-126] Switch Off Magnetic Shoe Brake Area
// 	SwitchOffMagneticShoeBrakeArea []SP_GenericArea_ComplexType `xml:"SwitchOffMagneticShoeBrakeArea,omitempty"`
// 	// Area in SP where some operating modes are unavailable
// 	DAS_ModesUnavailabilityArea []DAS_ModesUnavailabilityArea `xml:"DAS_ModesUnavailabilityArea,omitempty"`
// 	// [SUBSET-126] Dynamic Brake Force Inhibition Area
// 	DynamicBrakeForceInhibitionArea []SP_GenericArea_ComplexType `xml:"DynamicBrakeForceInhibitionArea,omitempty"`
// 	// [SUBSET-126] Dynamic Brake Limitation Area
// 	LimitedDynamicBrakeForceArea []LimitedDynamicBrakeForceArea `xml:"LimitedDynamicBrakeForceArea,omitempty"`
// 	TAP_TAF_Location             []TAP_TAF_Location             `xml:"TAP_TAF_Location,omitempty"`
// 	// Network-Specific Area
// 	NetworkSpecificArea []NetworkSpecificArea `xml:"NetworkSpecificArea,omitempty"`
// }

// // SP_Characteristics: Characteristics of the Segment Profile (e.g. gradient, curve, cant,
// // voltage...)
// type SP_Characteristics struct {
// 	XMLName xml.Name `xml:"SP_Characteristics"`
// 	// Static Speed Profile. It is possible to make different SSPs for different ATP systems if needed.
// 	StaticSpeedProfile []StaticSpeedProfile `xml:"StaticSpeedProfile,omitempty"`
// 	// Average gradient in the Segment Profile
// 	GradientAverage *GradientAverage `xml:"GradientAverage,omitempty"`
// 	// Steepest gradient in the Segment Profile
// 	GradientSteepest *GradientSteepest `xml:"GradientSteepest,omitempty"`
// 	// Curves in the Segment Profile
// 	Curves *Curves `xml:"Curves,omitempty"`
// 	// Superelevation values in the Segment Profile (Difference in rail height between the inside rail
// 	// and the outside rail in a curve, in m)
// 	SuperElevation *SuperElevation `xml:"SuperElevation,omitempty"`
// 	// Rated voltage of the traction system (V)
// 	RatedVoltage *RatedVoltage `xml:"RatedVoltage,omitempty"`
// 	// Estimated voltage of the traction system (V)
// 	EstimatedVoltage *EstimatedVoltage `xml:"EstimatedVoltage,omitempty"`
// 	// Maximum current to be applied by the train. Expressed in A
// 	CurrentLimitation *CurrentLimitation `xml:"CurrentLimitation,omitempty"`
// 	// List of ATP systems supported in the Segment Profile. The DAS system can then find out what is
// 	// supported. ATP System Consequences will be for Operational Train data containing its
// 	// characteristics necessary to calculate its dynamic capabilities.
// 	ATP_Systems *ATP_Systems `xml:"ATP_Systems,omitempty"`
// 	// Network-Specific Characteristics
// 	NetworkSpecificCharacteristics []NetworkSpecificCharacteristics `xml:"NetworkSpecificCharacteristics,omitempty"`
// }

// // SP_ContextInformation: Context information related to Journey Profile to be displayed to drivers
// type SP_ContextInformation struct {
// 	XMLName xml.Name `xml:"SP_ContextInformation"`
// 	// Level crossing area, for driver information
// 	LevelCrossingArea []LevelCrossingArea `xml:"LevelCrossingArea,omitempty"`
// 	// Connection track
// 	ConnectionTrack []ConnectionTrack `xml:"ConnectionTrack,omitempty"`
// 	// List of contacts for a specific area
// 	ContactList []ContactList `xml:"ContactList,omitempty"`
// 	// Network used for communication between driver and ground. The driver needs to know it because
// 	// some procedures in case of an emergency depend on the network type.
// 	CommunicationNetwork []CommunicationNetwork `xml:"CommunicationNetwork,omitempty"`
// 	// Reference Point on the side of the track indicating the passing of a certain distance, usually
// 	// relative to a "zero" point of the line.
// 	KilometreReferencePoint []KilometreReferencePoint `xml:"KilometreReferencePoint,omitempty"`
// 	// Identification of line and track
// 	LineAndTrackIdentifier *LineAndTrackIdentifier `xml:"LineAndTrackIdentifier,omitempty"`
// 	// Network-Specific Point
// 	NetworkSpecificPoint []NetworkSpecificPoint `xml:"NetworkSpecificPoint,omitempty"`
// 	// Network-Specific Area
// 	NetworkSpecificArea []NetworkSpecificArea `xml:"NetworkSpecificArea,omitempty"`
// 	// Network-Specific Characteristics
// 	NetworkSpecificCharacteristics []NetworkSpecificCharacteristics `xml:"NetworkSpecificCharacteristics,omitempty"`
// }

// // SP_Points: Notable points present in the Segment Profile (e.g. timing points, signals, balises)
// type SP_Points struct {
// 	XMLName xml.Name `xml:"SP_Points"`
// 	// A location and stopping accuracy defined in the Segment Profile for which a type (Stopping or
// 	// Passing Point) and specific time may be identified in the Journey Profile. This time may be an
// 	// arrival time, a departure time, or in the case of a train not scheduled to stop at that
// 	// location, the passing time.
// 	TimingPoint []TimingPoint `xml:"TimingPoint,omitempty"`
// 	// Description of the impact on allowed speed, driving reactions of the driver, and traction type
// 	// induced by permanent signals. The impact described in the SP is the one given by the least
// 	// restrictive Aspect that can be seen by the driver passing on the specific SP. If the TMS
// 	// anticipates that the train will meet the signal at an aspect different from the least
// 	// restrictive aspect, or the signal is cancelled, use TemporarySignalConstraints under JP / SPList
// 	// (if the signal can be modified or cancelled, the identifier is mandatory).
// 	Signal []Signal_ComplexType `xml:"Signal,omitempty"`
// 	// Equipment on the track at a known position that communicates with the train to exchange
// 	// different information. The train can recalibrate its positioning systems according to this
// 	// balise. If the Balise has a NID_C different from the one of the SP, it can be expressed here.
// 	Balise []Balise `xml:"Balise,omitempty"`
// 	// A VirtualBalise is a location in the Segment Profile for which the longitude and latitude is
// 	// known.
// 	VirtualBalise []VirtualBalise `xml:"VirtualBalise,omitempty"`
// 	// Stop in rear of an unprotected level crossing
// 	UnprotectedLevelCrossingStop []UnprotectedLevelCrossingStop `xml:"UnprotectedLevelCrossingStop,omitempty"`
// 	// Network-Specific Point
// 	NetworkSpecificPoint []NetworkSpecificPoint `xml:"NetworkSpecificPoint,omitempty"`
// }

// // SP_Request: Segment Profile request. If version is not included, latest version is requested.
// type SP_Request struct {
// 	XMLName xml.Name `xml:"SP_Request"`
// 	SegmentProfileIdentification_ComplexType
// 	// Identifier of the Segment Profile version, major number.
// 	SP_VersionMajor string `xml:"SP_VersionMajor,attr,omitempty"`
// 	// Identifier of the Segment Profile version, minor number.
// 	SP_VersionMinor string `xml:"SP_VersionMinor,attr,omitempty"`
// }

// // SP_Zone_ComplexType: Complex Type for SP Zone. At least one of IM_ID (Infrastructure Manager ID)
// // or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide the IM_ID. If
// // the SP is in a zone covered by a NID_C identifier from ERTMS, it should be provided, especially
// // if conversion to SUBSET-126 is required.
// type SP_Zone_ComplexType struct {
// 	// ID of the Infrastructure Manager. It is a UIC RICS Code.
// 	IM_ID CompanyCode `xml:"IM_ID,omitempty"`
// 	// ERTMS: Identity number of the country or region. From SUBSET-026 chap. 7, 7.5.1.86
// 	NID_C NID_C_Type `xml:"NID_C,omitempty"`
// }

// // SpecificSSP_ComplexType: Complex Type for Specific Static Speed Profile
// type SpecificSSP_ComplexType struct {
// 	// [SUBSET-126] Indicates the type of specific SSP category.
// 	Specific_SSP_Category string `xml:"specific_SSP_Category,attr,omitempty"`
// 	// [SUBSET-126] Absolute Positive Speed associated to a train category
// 	V_DIFF SpeedType `xml:"V_DIFF,attr,omitempty"`
// 	// Cant Deficiency Static Speed Profile
// 	CantDeficiencySSP *CantDeficiencySSP `xml:"CantDeficiencySSP,omitempty"`
// 	// Other Specific Static Speed Profile
// 	OtherSpecificSSP *OtherSpecificSSP `xml:"OtherSpecificSSP,omitempty"`
// }

// // StartSignalApplication: Location where the signal starts to be applied. If not specified, the
// // signal is applied starting from the signal location. This may also be in a different SP: then the
// // SP_ID has to be specified (otherwise the current SP is considered).
// type StartSignalApplication struct {
// 	XMLName xml.Name `xml:"StartSignalApplication"`
// 	// Segment Profile Identifier
// 	SP_ID string `xml:"SP_ID,attr,omitempty"`
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr,omitempty"`
// 	// What part of the train must clear the signal application before applying the signal
// 	StartSignalApplicationMustBeClearedBy TrainPartType `xml:"startSignalApplicationMustBeClearedBy,attr,omitempty"`
// }

// // StaticSpeedProfile: Static Speed Profile. It is possible to make different SSPs for different ATP
// // systems if needed.
// type StaticSpeedProfile struct {
// 	XMLName xml.Name `xml:"StaticSpeedProfile"`
// 	// Start of the Static Speed Profile
// 	StaticSpeedProfileStart *SSP_ComplexType `xml:"StaticSpeedProfileStart"`
// 	// Change of Static Speed Profile
// 	StaticSpeedProfileChange []StaticSpeedProfileChange `xml:"StaticSpeedProfileChange,omitempty"`
// 	// Identifier of ATP System. If it is a number: ETCS NID_ NTC from ERA "Assignment of Values to
// 	// ETCS Variables", section 3.3 (ERA_ERTMS_0 40001 v1.27.
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/ertms_040001_etcs_variables_values_en.pdf).
// 	// If it is a name: from the ERA "List of Class B Systems", section 3.2 (ERA/ TD/2011-11 v4.0,
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/ertms_ccs_class_b_systems_en.pdf)
// 	ATP_System_Identifier []ATP_System_Identifier `xml:"ATP_System_Identifier,omitempty"`
// }

// // StaticSpeedProfileChange: Change of Static Speed Profile
// type StaticSpeedProfileChange struct {
// 	XMLName xml.Name `xml:"StaticSpeedProfileChange"`
// 	SSP_ComplexType
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // SSP_ComplexType: Complex type for Static SPeed Profiles
// type SSP_ComplexType struct {
// 	// [SUBSET-126] Basic Static Speed Profile speed.
// 	SSP_Speed SpeedType `xml:"SSP_Speed,attr,omitempty"`
// 	// [SUBSET-126] Qualifier to indicate if the Static Speed Profile is to be applied to the front of
// 	// the train (1, no train length delay) or to the end of the train (0, train length delay).
// 	SSP_Front string `xml:"SSP_Front,attr,omitempty"`
// 	// Visual reference of a speed limit threshold (e.g. a building along the line). This attribute
// 	// shall be used only if the object is not already in the Segment Profile.
// 	SSP_VisualReference string `xml:"SSP_VisualReference,attr,omitempty"`
// 	// Specific Static Speed Profile
// 	SpecificSSP []SpecificSSP_ComplexType `xml:"SpecificSSP,omitempty"`
// }

// // TimingPointIdentification_ComplexType: Complex type containing the data for the identification of
// // Timing Points
// type TimingPointIdentification_ComplexType struct {
// 	// Timing Point Identifier
// 	TP_ID string `xml:"TP_ID,attr"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone,omitempty"`
// }

// // StatusReportRequest: Request for Status Report messages
// type StatusReportRequest struct {
// 	XMLName xml.Name `xml:"StatusReportRequest"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification"`
// 	// identifier of the last StatusReport message received by the ground side of the communication.
// 	// One or more StatusReport replies will only be sent if more recent data is available. Otherwise,
// 	// a simple "OK" message response is sent.
// 	LastStatusReportMessage *LastStatusReportMessage `xml:"LastStatusReportMessage,omitempty"`
// }

// // StopBeforePassingSignal: Indicates if the train must stop before crossing the signal, for example
// // if the driver must receive an order from the signalling responsible.
// type StopBeforePassingSignal struct {
// 	XMLName xml.Name `xml:"StopBeforePassingSignal"`
// 	// Stopping time necessary to complete the procedure, only given if StopBeforePassingSignal is
// 	// “yes”. The time given is the safety procedure time, and is independent from the commercial stop
// 	// time.
// 	StopTime string `xml:"stopTime,attr,omitempty"`
// 	Value    string `xml:",chardata"`
// }

// // StoppingPointApplicability: Criteria for the applicability of a stopping point
// type StoppingPointApplicability struct {
// 	XMLName xml.Name `xml:"StoppingPointApplicability"`
// 	// Minimum number of cars
// 	MinimumNumberOfCars string `xml:"minimumNumberOfCars,attr,omitempty"`
// 	// Maximum number of cars
// 	MaximumNumberOfCars string `xml:"maximumNumberOfCars,attr,omitempty"`
// 	// Minimum Train Length for which stopping point is valid. Expressed in m
// 	MinimumTrainLength string `xml:"minimumTrainLength,attr,omitempty"`
// 	// Maximum Train Length for which stopping point is valid. Expressed in m
// 	MaximumTrainLength string `xml:"maximumTrainLength,attr,omitempty"`
// 	// Criteria for applicability related to rolling stock
// 	RollingStockApplicability []RollingStockApplicability `xml:"RollingStockApplicability,omitempty"`
// 	// The stopping point is applicable if this is a preceding SP
// 	StoppingPointApplicableIfPrecedingSP []SegmentProfileIdentification_ComplexType `xml:"StoppingPointApplicableIfPrecedingSP,omitempty"`
// }

// // StoppingPointInformation: Stopping Point Information
// type StoppingPointInformation struct {
// 	XMLName xml.Name `xml:"StoppingPointInformation"`
// 	// Expected time of departure at Timing Point expressed in UTC
// 	DepartureTime UTCDateTimeType `xml:"departureTime,attr"`
// 	// Earliest departure time desired, expressed in UTC.
// 	EarliestDepartureTime UTCDateTimeType `xml:"earliestDepartureTime,attr,omitempty"`
// 	// The original planned departure time from the Timing Point expressed in UTC. Useful for driver
// 	// information.
// 	PlannedDepartureTime UTCDateTimeType `xml:"plannedDepartureTime,attr,omitempty"`
// 	// [SUBSET-126] The variable defines if the train is requested to be held at the Stopping Point (1)
// 	// or not (0)
// 	TrainHold string `xml:"trainHold,attr,omitempty"`
// 	// [SUBSET-126] Minimum dwell time at given Stopping Point (in seconds)
// 	MinimumDwellTime string `xml:"minimumDwellTime,attr,omitempty"`
// 	// Leave as soon as possible, taking into account also the other variables (boolean).
// 	LeaveASAP string `xml:"leaveASAP,attr,omitempty"`
// 	// Type of authorisation for departure (e.g. SMS)
// 	DepartureAuthorisationType string `xml:"departureAuthorisationType,attr,omitempty"`
// 	// This attribute is to be used when the driver will be asked to stop at a stopping which is
// 	// different from the usual one, and it needs to be brought to his attention.
// 	ChangedStoppingLocationAlert string `xml:"changedStoppingLocationAlert,attr,omitempty"`
// 	// [SUBSET-126] Specifies if the ATO-OB has to manage the train doors opening and on which side the
// 	// passenger exchange doors have to be opened
// 	OpeningDoorSide string `xml:"openingDoorSide,attr,omitempty"`
// 	// [SUBSET-126] This variable defines if the doors are to be opened centralised (1) or by the
// 	// passengers (0)
// 	CentralisedOpening string `xml:"centralisedOpening,attr,omitempty"`
// 	// [SUBSET-126] This variable defines if the ATO-OB has to manage the train doors closing (1) or
// 	// not (0)
// 	AutomaticClosing string `xml:"automaticClosing,attr,omitempty"`
// 	// [SUBSET-126] Specifies some information specific for the TPThis variable defines if the daylight
// 	// saving hour is applicable to calculate the local time
// 	DaylightSaving string `xml:"daylightSaving,attr,omitempty"`
// 	// [SUBSET-126] request to the train to stop with couplers relaxed
// 	RelaxedCoupler string `xml:"relaxedCoupler,attr,omitempty"`
// 	// This field can be used to inform the driver of the reason of a stop. It cannot be used with
// 	// passing points. The driver may use this for passenger information and to inform the dispatcher
// 	// on deviations. There will be no other effect on the DAS behavior apart from the information
// 	// display.
// 	StopType []StopType `xml:"StopType,omitempty"`
// }

// // StoppingPointLocation: If the Timing Point is a Stopping Point, point of the track where the head
// // of trains respecting the specified stoppingPointApplicability parameters will stop.
// type StoppingPointLocation struct {
// 	XMLName xml.Name `xml:"StoppingPointLocation"`
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// [SUBSET-126] Required stopping tolerance (m)
// 	StopLocationTolerance DistanceType `xml:"stopLocationTolerance,attr,omitempty"`
// 	// [SUBSET-126] Distance (in m) from a Stopping Point to consider it as reached
// 	Distance_STP_Reached DistanceType `xml:"distance_STP_Reached,attr,omitempty"`
// 	// Stopping zone, common to all alternative TimingPoints for a track/platform. it could be applied
// 	// in or out of stations (e.g. storage tracks). It is possible to reference a Stopping Zone from
// 	// Timing Point Constraints.
// 	StoppingZone *StoppingZone_ComplexType `xml:"StoppingZone,omitempty"`
// 	// Criteria for the applicability of a stopping point
// 	StoppingPointApplicability *StoppingPointApplicability `xml:"StoppingPointApplicability,omitempty"`
// }

// // StoppingZone_ComplexType: Complex type for stopping zone
// type StoppingZone_ComplexType struct {
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr"`
// 	// Name of the element
// 	Name string `xml:"name,attr,omitempty"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone,omitempty"`
// }

// // StopType: This field can be used to inform the driver of the reason of a stop. It cannot be used
// // with passing points. The driver may use this for passenger information and to inform the
// // dispatcher on deviations. There will be no other effect on the DAS behavior apart from the
// // information display.
// type StopType struct {
// 	XMLName xml.Name `xml:"StopType"`
// 	// Type of stop and purpose. Stop types can be: - Commercial: a stop used by the RU to perform the
// 	// commercial duties the train was planned for. For a passenger train, this will concern stops
// 	// where passenger loading or unloading is done. For a freight train, this will concern stops where
// 	// loading and unloading is done, or when wagons are either added or uncoupled from the train. -
// 	// Non Commercial Regulation: stops integrated for the purpose of guaranteeing the fluidity of
// 	// traffic. - Non Commercial RU: stops used by the RU to fulfil its specific needs. For example, a
// 	// train driver relay, or a locomotive change could be characterised by this purpose; - Non
// 	// Commercial IM: stops necessary for the IM’s needs; - Non Commercial Safety: stops used to
// 	// perform a safety procedure; - Other: any stop that cannot be categorised using one of the 5
// 	// other types.
// 	StopTypePurpose string `xml:"stopTypePurpose,attr"`
// 	// Additional details about the stop type
// 	StopTypeDetails string `xml:"stopTypeDetails,attr,omitempty"`
// 	// Flag indicating whether the stop is planned (1) or unplanned (0)
// 	PlannedStop string `xml:"plannedStop,attr,omitempty"`
// 	// Flag indicating whether the stop is mandatory (1) or optional (0)
// 	MandatoryStop string `xml:"mandatoryStop,attr,omitempty"`
// }

// // SuperElevation: Superelevation values in the Segment Profile (Difference in rail height between
// // the inside rail and the outside rail in a curve, in m)
// type SuperElevation struct {
// 	XMLName xml.Name `xml:"SuperElevation"`
// 	// Superelevation at the start of the Segment Profile
// 	SuperElevationStart *SuperElevationStart `xml:"SuperElevationStart"`
// 	// Superelevation change
// 	SuperElevationChange []SuperElevationChange `xml:"SuperElevationChange,omitempty"`
// }

// // SuperElevationChange: Superelevation change
// type SuperElevationChange struct {
// 	XMLName xml.Name `xml:"SuperElevationChange"`
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// 	// Value of superelevation (mm)
// 	SuperElevationValue string `xml:"superElevationValue,attr"`
// }

// // SuperElevationStart: Superelevation at the start of the Segment Profile
// type SuperElevationStart struct {
// 	XMLName xml.Name `xml:"SuperElevationStart"`
// 	// Value of superelevation (mm)
// 	SuperElevationValue string `xml:"superElevationValue,attr"`
// }

// // TAP_TAF_ID_ComplexType: Complex Type for TAP/TAF identifier. Is a modification of TAP/TAF
// // CompositIdentifierOperationalType (ObjectType is always TR and TimetableYear is optional).
// type TAP_TAF_ID_ComplexType struct {
// 	// Provides a possibility for differentiantion between the objects: Train, Path, Case Reference and
// 	// Path Request. In this case, must be TR (Train)
// 	ObjectType ObjectType `xml:"ObjectType"`
// 	// Identifies a railway company (RU or IM) (UIC RICS Code: https://uic.org/rics)
// 	Company CompanyCode `xml:"Company"`
// 	// It is the main part of identifier and is determined by the company that creates it.
// 	Core Core `xml:"Core"`
// 	// The variant shows a relationship between two identifiers referring to the same business case
// 	Variant Variant `xml:"Variant"`
// 	// Refers to the timetable period in which the business will be carried out (Not used in SFERA.
// 	// Only kept for consistency with data elements in TAP/TAF)
// 	TimetableYear TimetableYear `xml:"TimetableYear,omitempty"`
// 	// Is only used in the operational phase and refers to the date where the single train will start
// 	// the train journey
// 	StartDate string `xml:"StartDate"`
// }

// // TAP_TAF_Location models the corresponding XSD type/element.
// type TAP_TAF_Location struct {
// 	XMLName xml.Name `xml:"TAP_TAF_Location"`
// 	SP_GenericArea_ComplexType
// 	TAP_TAF_LocationIdent *LocationIdent `xml:"TAP_TAF_LocationIdent"`
// }

// // TC_ChangeFeatures: Features for Train Characteristics Change
// type TC_ChangeFeatures struct {
// 	XMLName xml.Name `xml:"TC_ChangeFeatures"`
// 	TC_Features_ComplexType
// 	// Percentage of power available due to engine problems
// 	RemainingEnginePower PercentageType `xml:"remainingEnginePower,attr,omitempty"`
// 	// Traction Type of the TrainCharacteristics of the train that has become completely unavailable
// 	// (e.g. due to a failure of the pantograph)
// 	UnavailableTractionType []UnavailableTractionType `xml:"UnavailableTractionType,omitempty"`
// }

// // TC_Features_ComplexType: Complex type for Train Characteristics features
// type TC_Features_ComplexType struct {
// 	// The total length of a train. Expressed in m
// 	TrainLength string `xml:"trainLength,attr,omitempty"`
// 	// The maximum possible speed of the train in km/h
// 	TrainMaxSpeed string `xml:"trainMaxSpeed,attr,omitempty"`
// 	// The sum of all weights of wagons and traction units. Expressed in tonnes (t)
// 	TrainWeight string `xml:"trainWeight,attr,omitempty"`
// 	// Comfortable acceleration. Expressed in m/s^2
// 	ComfortableAcceleration string `xml:"comfortableAcceleration,attr,omitempty"`
// 	// Deceleration used for comfortable braking. Expressed in m/s^2
// 	ComfortableBraking string `xml:"comfortableBraking,attr,omitempty"`
// 	// Maximum force that can be Applied through regenerative braking. Expressed in N
// 	MaximumRegenerativeForce string `xml:"maximumRegenerativeForce,attr,omitempty"`
// 	// Maximum power that can be regenerated. Expressed in W
// 	MaximumRegenerativeBraking string `xml:"maximumRegenerativeBraking,attr,omitempty"`
// 	// Maximum power of the engine. Expressed in W
// 	EnginePower PowerTypeW `xml:"enginePower,attr,omitempty"`
// 	// Description of how the train accelerates. Coefficient used in order to calculate the resistance
// 	// force (constant value). Expressed in N
// 	RolloutCoefficientA string `xml:"rolloutCoefficientA,attr,omitempty"`
// 	// Description of how the train accelerates if no traction or braking is applied. Coefficient used
// 	// in order to calculate the resistance force (value to be multiplied with v). Expressed in N.s/m
// 	RolloutCoefficientB string `xml:"rolloutCoefficientB,attr,omitempty"`
// 	// Description of how the train accelerates if no traction or braking is applied.Coefficient used
// 	// in order to calculate the resistance force (value to be multiplied with v^2). Expressed in
// 	// N.s^2/m^2
// 	RolloutCoefficientC string `xml:"rolloutCoefficientC,attr,omitempty"`
// 	// Free-form text describing the rolling stock. Used for driver information.
// 	RollingStockType string `xml:"rollingStockType,attr,omitempty"`
// 	// Braked weight percentage. Unit: Percent. Acceptable values 0-500. Reference: Introduction to
// 	// ETCS Braking Curves, ERA_ERTMS_040026, Version : 1.4, Date : 14/06/16.
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/introduction_to_etcs_braking_curves_en.doc
// 	BrakedWeightPercentage string `xml:"brakedWeightPercentage,attr,omitempty"`
// 	// Delay between braking command and start of braking.
// 	BrakingReactionTime string `xml:"brakingReactionTime,attr,omitempty"`
// 	// Delay between traction command and start of tractioning. Expressed in s
// 	TractionReactionTime string `xml:"tractionReactionTime,attr,omitempty"`
// 	// Rotating Mass Factor. The rotating mass of the train is equivalent to the rotational inertia of
// 	// the train, transformed by gear ratio and wheel diameter respectively.
// 	RotatingMassFactor string `xml:"rotatingMassFactor,attr,omitempty"`
// 	// Number of engines
// 	Engines string `xml:"engines,attr,omitempty"`
// 	// Number of wagons
// 	Wagons string `xml:"wagons,attr,omitempty"`
// 	// Tunnel Resistance Prediction coefficient b_e (for engines). No unit. Typical value 1.48.
// 	// Reference: https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
// 	Trp_be string `xml:"trp_be,attr,omitempty"`
// 	// Tunnel Resistance Prediction coefficient b_w (for wagons). No unit. Typical value 1.75.
// 	// Reference: https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
// 	Trp_bw string `xml:"trp_bw,attr,omitempty"`
// 	// Tunnel Resistance Prediction coefficient a_e (for engines). Unit: kg.m^(2b_e-1). Typical value
// 	// 1140. Reference: https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
// 	Trp_ae string `xml:"trp_ae,attr,omitempty"`
// 	// Tunnel Resistance Prediction coefficient a_w (for wagons). Unit: kg.m^(2b_w-1). Typical value
// 	// 662. Reference: https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
// 	Trp_aw string `xml:"trp_aw,attr,omitempty"`
// 	// Train Category (NC_TRAIN as defined in SUBSET-026).
// 	NC_Train NC_TrainType `xml:"NC_Train,attr,omitempty"`
// 	// Cant Deficiency Train Category (NC_CDTRAIN as defined in SUBSET-026).
// 	NC_CDTrain NC_CDTrainType `xml:"NC_CDTrain,attr,omitempty"`
// 	// Code for the train category (freeform text, limited to 10 characters)
// 	TrainCategoryCode string `xml:"trainCategoryCode,attr,omitempty"`
// 	// Name of train type to be displayed to the driver (e.g. ‘TGV’ of ‘VIRM 6’).
// 	DisplayName string `xml:"displayName,attr,omitempty"`
// 	// [SUBSET-126] Axle load from which the speed restriction is applicable
// 	AxleLoadCategory string `xml:"axleLoadCategory,attr,omitempty"`
// 	// Distance from the head of the train of the last point of the rolling stock that has to clear a
// 	// zone with a power constraint (lowered pantograph, main switch off…) for that power constraint to
// 	// be lifted. Expressed in m.
// 	PositionOfLastPantograph DistanceType `xml:"positionOfLastPantograph,attr,omitempty"`
// 	// ATP system supported by the train. The DAS system can then find out what is supported. ATP
// 	// System Consequences will be for Operational Train data containing its characteristics necessary
// 	// to calculate its dynamic capabilities. Use the ERA "List of Class B Systems": ERA/ TD/2011-11
// 	// v3.0. http://www.era.europa.eu/Document-Register/Documents/ERA_TD_2011-11%20v30.pdf
// 	Train_ATP_Systems *Train_ATP_Systems `xml:"Train_ATP_Systems,omitempty"`
// 	// Traction Force Curve for the train. Several curves may be defined for different traction types.
// 	TractionForceCurve []TractionForceCurve_ComplexType `xml:"TractionForceCurve,omitempty"`
// 	// Braking Force Curve for the train. Several curves may be defined for different traction types
// 	// and ATP Systems.
// 	BrakingForceCurve []BrakingForceCurve_ComplexType `xml:"BrakingForceCurve,omitempty"`
// 	// Wagon relevant data for the wagons within a running train
// 	WagonData []WagonData `xml:"WagonData,omitempty"`
// 	// Network Specific Parameters in TC Features
// 	TC_Features_NSPs []NSPList_ComplexType `xml:"TC_Features_NSPs,omitempty"`
// }

// // TC_Request: Train Characteristics request. If version is not included, latest version is
// // requested.
// type TC_Request struct {
// 	XMLName xml.Name `xml:"TC_Request"`
// 	TC_Identifier_ComplexType
// 	// Identifier of the Train Characteristics version, major number.
// 	TC_VersionMajor string `xml:"TC_VersionMajor,attr,omitempty"`
// 	// Identifier of the Train Characteristics version, minor number.
// 	TC_VersionMinor string `xml:"TC_VersionMinor,attr,omitempty"`
// }

// // TemporaryConstraintReason: Context information explaining background of temporary constraint
// type TemporaryConstraintReason struct {
// 	XMLName xml.Name `xml:"TemporaryConstraintReason"`
// 	// ISO 639-1 language code
// 	Language string `xml:"language,attr,omitempty"`
// 	// String containing the message
// 	MessageString string `xml:"messageString,attr"`
// 	// Approved version of text (translation when 'false')
// 	OriginalLanguage string `xml:"originalLanguage,attr,omitempty"`
// 	// Human-readable identifier of document informing driver of temporary constraint
// 	TemporaryConstraintReference string `xml:"temporaryConstraintReference,attr,omitempty"`
// }

// // TemporaryConstraints: Temporary Constraints in the Segment Profile
// type TemporaryConstraints struct {
// 	XMLName xml.Name `xml:"TemporaryConstraints"`
// 	TemporaryConstraints_ComplexType
// 	// Reason for the temporary constraint. It is possible to choose one reason among the enumerations;
// 	// the detailed data has to be put under the related element in the TemporaryConstraint choice.
// 	// Exceptions are ATO_Inhibition_Zone, DAS_Inhibition_Zone and IM_SDAS_Only, where no additional
// 	// data is needed.
// 	TemporaryConstraintType string `xml:"temporaryConstraintType,attr"`
// 	// Additional Speed Restriction
// 	AdditionalSpeedRestriction *AdditionalSpeedRestriction `xml:"AdditionalSpeedRestriction,omitempty"`
// 	// Low Adhesion Status Report
// 	LowAdhesion *LowAdhesion `xml:"LowAdhesion,omitempty"`
// 	// Temporary Constraints on Total current for traction and HVAC consumption
// 	TractionTotalCurrent *TotalCurrent_ComplexType `xml:"TractionTotalCurrent,omitempty"`
// 	// Expected maximal current for regeneration that can be returned to the overhead contact line
// 	RegenerationTotalCurrent *TotalCurrent_ComplexType `xml:"RegenerationTotalCurrent,omitempty"`
// 	// Advice on minimising or maximising the power consumption
// 	PowerAdvice *PowerAdvice `xml:"PowerAdvice,omitempty"`
// 	// Temporary constraints on estimated voltage
// 	EstimatedVoltageConstraint *EstimatedVoltageConstraint `xml:"EstimatedVoltageConstraint,omitempty"`
// 	// Temporary constraints due to wind factors
// 	Wind *Wind `xml:"Wind,omitempty"`
// 	// Operating modes that are unavailable. If all 3 attributes are included, the triplet is
// 	// unavailable; of only 1 or 2 attributes are included, that mode is unavailable independently of
// 	// the value of the rest of the attributes.
// 	Unavailable_DAS_OperatingModes *Unavailable_DAS_OperatingModes `xml:"Unavailable_DAS_OperatingModes,omitempty"`
// 	// Explicit speed advice
// 	AdvisedSpeed *AdvisedSpeed `xml:"AdvisedSpeed,omitempty"`
// 	// Network-Specific Temporary Constraint
// 	NetworkSpecificConstraint *NSPList_ComplexType `xml:"NetworkSpecificConstraint,omitempty"`
// 	// Context information explaining background of temporary constraint
// 	TemporaryConstraintReason []TemporaryConstraintReason `xml:"TemporaryConstraintReason,omitempty"`
// }

// // TemporarySignalConstraints: Description of the impact on allowed speed, driving reactions of the
// // driver, and traction type induced by temporary signals, permanent signals with an aspect
// // different from the least restrictive aspect described in SP / SP_Points / Signal, or temporarily
// // cancelled signals on the ground.
// type TemporarySignalConstraints struct {
// 	XMLName xml.Name `xml:"TemporarySignalConstraints"`
// 	Signal_ComplexType
// 	// Starting date and time (in UTC)
// 	StartTime UTCDateTimeType `xml:"startTime,attr,omitempty"`
// 	// Ending date and time (in UTC)
// 	EndTime UTCDateTimeType `xml:"endTime,attr,omitempty"`
// 	// Type of temporary signal constraints: Temporary Signal (the identifier must then be new, not
// 	// used by another element), Existing Signal Aspect Modification (the identifier must point to a
// 	// Signal described in SP / SP_Points if it is present) or Cancelled Signal (the identifier must
// 	// point to a Signal described in SP / SP_Points).
// 	TemporarySignalConstraintType string `xml:"temporarySignalConstraintType,attr"`
// }

// // TimingPoint: A location and stopping accuracy defined in the Segment Profile for which a type
// // (Stopping or Passing Point) and specific time may be identified in the Journey Profile. This time
// // may be an arrival time, a departure time, or in the case of a train not scheduled to stop at that
// // location, the passing time.
// type TimingPoint struct {
// 	XMLName xml.Name `xml:"TimingPoint"`
// 	// Timing Point Identifier
// 	TP_ID string `xml:"TP_ID,attr"`
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// 	// Abbreviation of the Timing Point name
// 	TP_abbreviation string `xml:"TP_abbreviation,attr,omitempty"`
// 	// Type of Timing Point. Examples: Very important station Important station Station Stopping
// 	// location Junction Regional border Connexion Network border Signal with detection Grid or
// 	// shunting area Service installation Service stopping location Movable bridge Other
// 	TP_type string `xml:"TP_type,attr,omitempty"`
// 	// TP_Priority is used to indicate the relative significance of the Timing Point in a Journey from
// 	// the driver’s point of view. This information is intended to be used when choosing the
// 	// information to display. It does not specify a priority in respecting the timing points when the
// 	// DAS calculates the speed profile (the attributes TP_latestArrivalTime and arrivalWindow in the
// 	// Journey Profile are the absolute time constraints on timing points). It has 5 possible values :
// 	// 5: very important station, junction or network border 4: important station, junction or shunting
// 	// area 3: station, junction or regional border 2: other stopping locations 1: all other timing
// 	// points
// 	TP_priority string `xml:"TP_priority,attr,omitempty"`
// 	// If the Timing Point is a station: identifier of the track where the train will stop or pass
// 	// through without stopping at the station. Examples : “1” “14” “2M” “1bis”
// 	StationTrack string `xml:"stationTrack,attr,omitempty"`
// 	// Name of the Timing Point
// 	TP_Name []TP_Name `xml:"TP_Name,omitempty"`
// 	// If the Timing Point is a Stopping Point, point of the track where the head of trains respecting
// 	// the specified stoppingPointApplicability parameters will stop.
// 	StoppingPointLocation *StoppingPointLocation `xml:"StoppingPointLocation,omitempty"`
// 	// Network Specific Parameters in Timing Point
// 	TimingPoint_NSPs []TimingPoint_NSPs `xml:"TimingPoint_NSPs,omitempty"`
// }

// // TimingPoint_NSPs: Network Specific Parameters in Timing Point
// type TimingPoint_NSPs struct {
// 	XMLName xml.Name `xml:"TimingPoint_NSPs"`
// 	NSPList_ComplexType
// }

// // TimingPointConstraints: Constraints at specific Timing Points in the Segment Profile (arrival and
// // departure times)
// type TimingPointConstraints struct {
// 	XMLName xml.Name `xml:"TimingPointConstraints"`
// 	// Time of arrival at Timing Point expressed in UTC. Has to be specified always except at departure
// 	// stations.
// 	TP_latestArrivalTime UTCDateTimeType `xml:"TP_latestArrivalTime,attr,omitempty"`
// 	// The original planned latest arrival time at the Timing Point expressed in UTC. Useful for driver
// 	// information. Not to be used for calculations.
// 	TP_PlannedLatestArrivalTime UTCDateTimeType `xml:"TP_PlannedLatestArrivalTime,attr,omitempty"`
// 	// Acceptable time allowance to be earlier at the TP
// 	ArrivalWindow string `xml:"arrivalWindow,attr,omitempty"`
// 	// Aim to be as near as possible to the beginning of the arrival window
// 	AimEarlyArrival string `xml:"aimEarlyArrival,attr,omitempty"`
// 	// [SUBSET-126] Defines if the TP location is applicable from the front, middle or rear of the
// 	// train
// 	TP_alignment string `xml:"TP_alignment,attr,omitempty"`
// 	// [SUBSET-126] Specifies if the Timing Point is a Passing Point, an operational Stopping Point or
// 	// a skipped Stopping Point
// 	TP_StopSkipPass string `xml:"TP_StopSkipPass,attr,omitempty"`
// 	// [SUBSET-126] Specifies some information specific for the TP
// 	TP_Information string `xml:"TP_Information,attr,omitempty"`
// 	// [SUBSET-126] Specifies some information specific for the TPThis variable defines if the daylight
// 	// saving hour is applicable to calculate the local time
// 	DaylightSaving string `xml:"daylightSaving,attr,omitempty"`
// 	// Reference to a Timing Point in the Segment Profile. It could be either the ID of a specific
// 	// TimingPoint or a Stopping Zone. In the latter case, it is up to the DAS-OB to choose the exact
// 	// Timing Point.
// 	TimingPointReference *TimingPointReference `xml:"TimingPointReference"`
// 	// Stopping Point Information
// 	StoppingPointInformation *StoppingPointInformation `xml:"StoppingPointInformation,omitempty"`
// 	// Network Specific Parameters in Timing Point Constraints
// 	TimingPointConstraints_NSPs []NSPList_ComplexType `xml:"TimingPointConstraints_NSPs,omitempty"`
// }

// // TimingPointEstimation: [SUBSET-126] Estimated arrival time to, at least, the next Timing Point
// type TimingPointEstimation struct {
// 	XMLName xml.Name `xml:"TimingPointEstimation"`
// 	// Timing Point Identifier
// 	TP_ID string `xml:"TP_ID,attr"`
// 	// Forecasted arrival at the Timing Point
// 	ForecastedArrival UTCDateTimeType `xml:"forecastedArrival,attr"`
// 	// Identifier of the zone containing the Segment Profile. At least one of IM_ID (Infrastructure
// 	// Manager ID) or NID_C (from ERTMS) must be provided. It is encouraged to systematically provide
// 	// the IM_ID. If the SP is in a zone covered by a NID_C identifier from ERTMS, it should be
// 	// provided, especially if conversion to SUBSET-126 is required.
// 	SP_Zone *SP_Zone_ComplexType `xml:"SP_Zone"`
// }

// // TimingPointReference: Reference to a Timing Point in the Segment Profile. It could be either the
// // ID of a specific TimingPoint or a Stopping Zone. In the latter case, it is up to the DAS-OB to
// // choose the exact Timing Point.
// type TimingPointReference struct {
// 	XMLName xml.Name `xml:"TimingPointReference"`
// 	// Reference to the ID of a TimingPoint in the Segment Profile
// 	TP_ID_Reference *TimingPointIdentification_ComplexType `xml:"TP_ID_Reference"`
// 	// Reference to a StoppingZone in the TimingPoint of the Segment Profile
// 	StoppingZoneReference *StoppingZone_ComplexType `xml:"StoppingZoneReference"`
// }

// // TP_Name: Name of the Timing Point
// type TP_Name struct {
// 	XMLName xml.Name `xml:"TP_Name"`
// 	// Name of the element
// 	Name string `xml:"name,attr"`
// 	// ISO 639-1 language code
// 	Language string `xml:"language,attr,omitempty"`
// }

// // TractionForceCurve_ComplexType: Complex Type for Traction Force Curves
// type TractionForceCurve_ComplexType struct {
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// Traction Type (DC/AC with voltage, self-propelled or powerless)
// 	TractionType TractionTypeType `xml:"tractionType,attr"`
// 	// This TractionForceCurve is to be used if the estimated voltage of the considered portion of
// 	// infrastructure is inferior to this value and superior to a lower estimatedVoltage of another
// 	// TractionForceCurve (if it exists).
// 	TFC_estimatedVoltage VoltageType `xml:"TFC_estimatedVoltage,attr,omitempty"`
// 	// This TractionForceCurve is to be used if the maximum current value of the considered portion of
// 	// infrastructure is inferior to this value and superior to a lower maxCurrentValue of another
// 	// TractionForceCurve (if it exists).
// 	TFC_maxCurrentValue CurrentType `xml:"TFC_maxCurrentValue,attr,omitempty"`
// 	// Data point for the traction force curve (pair of speed and traction force)
// 	TractionForceCurveDataPoint []TractionForceCurveDataPoint `xml:"TractionForceCurveDataPoint,omitempty"`
// 	// Network Specific Parameters in Traction Force Curve
// 	TFC_NSPs []NSPList_ComplexType `xml:"TFC_NSPs,omitempty"`
// }

// // TractionForceCurveDataPoint: Data point for the traction force curve (pair of speed and traction
// // force)
// type TractionForceCurveDataPoint struct {
// 	XMLName xml.Name `xml:"TractionForceCurveDataPoint"`
// 	// speed (km/h)
// 	Speed SpeedType `xml:"speed,attr"`
// 	// Traction Force (in N)
// 	TractionForce ForceType `xml:"tractionForce,attr"`
// }

// // TractionTypeTransition: Traction type transition signal
// type TractionTypeTransition struct {
// 	XMLName xml.Name `xml:"TractionTypeTransition"`
// 	// Traction type transition signal for electric trains
// 	TractionTypeTransitionElectric *TractionTypeTransition_ComplexType `xml:"TractionTypeTransitionElectric,omitempty"`
// 	// Traction type transition signal for bimodal trains
// 	TractionTypeTransitionBimodal *TractionTypeTransition_ComplexType `xml:"TractionTypeTransitionBimodal,omitempty"`
// }

// // TractionTypeTransition_ComplexType: Complex Type for Traction Type Transition
// type TractionTypeTransition_ComplexType struct {
// 	// New Traction Type
// 	NewTractionType TractionTypeType `xml:"newTractionType,attr"`
// }

// // Train: Contains the identification of a related train and the segments
// type Train struct {
// 	XMLName xml.Name `xml:"Train"`
// 	TrainType
// }

// // Train_ATP_Systems: ATP system supported by the train. The DAS system can then find out what is
// // supported. ATP System Consequences will be for Operational Train data containing its
// // characteristics necessary to calculate its dynamic capabilities. Use the ERA "List of Class B
// // Systems": ERA/ TD/2011-11 v3.0.
// // http://www.era.europa.eu/Document-Register/Documents/ERA_TD_2011-11%20v30.pdf
// type Train_ATP_Systems struct {
// 	XMLName xml.Name `xml:"Train_ATP_Systems"`
// 	ATPSystems_ComplexType
// }

// // TrainCharacteristics: Element containing the characteristics of the train
// type TrainCharacteristics struct {
// 	XMLName xml.Name `xml:"TrainCharacteristics"`
// 	TC_Identifier_ComplexType
// 	// Identifier of the Train Characteristics version, major number.
// 	TC_VersionMajor string `xml:"TC_VersionMajor,attr"`
// 	// Identifier of the Train Characteristics version, minor number.
// 	TC_VersionMinor string `xml:"TC_VersionMinor,attr"`
// 	// List of train's characteristics, necessary to calculate its dynamic capabilities: length,
// 	// weight, max speed etc.
// 	TC_Features *TC_Features_ComplexType `xml:"TC_Features"`
// }

// // TrainCharacteristicsChange: Change in Train Characteristics
// type TrainCharacteristicsChange struct {
// 	XMLName xml.Name `xml:"TrainCharacteristicsChange"`
// 	TC_Identifier_ComplexType
// 	// Features for Train Characteristics Change
// 	TC_ChangeFeatures *TC_ChangeFeatures `xml:"TC_ChangeFeatures"`
// }

// // TrainCharacteristicsRef: Reference to a Train Characteristics element (sent separately). All TCs
// // are taken from that element; some can be overridden locally (e.g. length, weight, max speed).
// type TrainCharacteristicsRef struct {
// 	XMLName xml.Name `xml:"TrainCharacteristicsRef"`
// 	TC_Identifier_ComplexType
// 	// Identifier of the Train Characteristics version, major number.
// 	TC_VersionMajor string `xml:"TC_VersionMajor,attr"`
// 	// Identifier of the Train Characteristics version, minor number.
// 	TC_VersionMinor string `xml:"TC_VersionMinor,attr"`
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// 	// The total length of a train. Expressed in m
// 	TrainLength string `xml:"trainLength,attr,omitempty"`
// 	// The sum of all weights of wagons and traction units. Expressed in tonnes (t)
// 	TrainWeight string `xml:"trainWeight,attr,omitempty"`
// 	// The maximum possible speed of the train in km/h
// 	TrainMaxSpeed string `xml:"trainMaxSpeed,attr,omitempty"`
// 	// Network Specific Parameters in Train Characteristics Reference
// 	TrainCharacteristicsRef_NSPs []NSPList_ComplexType `xml:"TrainCharacteristicsRef_NSPs,omitempty"`
// }

// // TrainDelay: Identifies the Delta delay time of a train against the booked schedule as well as
// // against the referenced time
// type TrainDelay struct {
// 	XMLName xml.Name `xml:"TrainDelay"`
// 	// Identifies the Delta delay time against the booked schedule in minutes
// 	AgainstBooked DeltaTime `xml:"AgainstBooked,omitempty"`
// 	// Delay compared to the referenced Date/Time
// 	AgainstReferenced DeltaTime `xml:"AgainstReferenced,omitempty"`
// }

// // TrainIdentification: Identification of the train
// type TrainIdentification struct {
// 	XMLName xml.Name `xml:"TrainIdentification"`
// 	TrainIdentifier_ComplexType
// }

// // TrainInformation models the corresponding XSD type/element.
// type TrainInformation struct {
// 	XMLName xml.Name `xml:"TrainInformation"`
// 	// The total length of a train. Expressed in m
// 	TrainLength string `xml:"trainLength,attr,omitempty"`
// 	// Code for the train category (freeform text, limited to 10 characters)
// 	TrainCategoryCode string `xml:"trainCategoryCode,attr,omitempty"`
// 	// Identifies a railway company (RU or IM) (UIC RICS Code: https://uic.org/rics)
// 	Company CompanyCode `xml:"Company,omitempty"`
// 	// Operational Train Number
// 	TrainNumber string `xml:"TrainNumber,omitempty"`
// 	// Is only used in the operational phase and refers to the date where the single train will start
// 	// the train journey
// 	StartDate string `xml:"StartDate,omitempty"`
// 	// Additional identifier for train number
// 	AdditionalTrainNumber *AdditionalTrainNumber `xml:"AdditionalTrainNumber,omitempty"`
// 	// Identifier of ATP System. If it is a number: ETCS NID_ NTC from ERA "Assignment of Values to
// 	// ETCS Variables", section 3.3 (ERA_ERTMS_0 40001 v1.27.
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/ertms_040001_etcs_variables_values_en.pdf).
// 	// If it is a name: from the ERA "List of Class B Systems", section 3.2 (ERA/ TD/2011-11 v4.0,
// 	// https://www.era.europa.eu/sites/default/files/activities/docs/ertms_ccs_class_b_systems_en.pdf)
// 	ATP_System_Identifier *ATP_System_Identifier `xml:"ATP_System_Identifier,omitempty"`
// }

// // TrainLocationInformation: Information about the current location and occupation of the train
// type TrainLocationInformation struct {
// 	XMLName xml.Name `xml:"TrainLocationInformation"`
// 	// Identification of the line. For example “115000” is the Line Identifier for the line from
// 	// Strasbourg to Basel.
// 	LineIdentifier string `xml:"lineIdentifier,attr,omitempty"`
// 	// Provides information about the occupation of a Segment Profile (or part of it) by the train.
// 	SegmentClaim []SegmentClaim `xml:"SegmentClaim,omitempty"`
// 	// Position and speed of the train
// 	PositionSpeed *PositionSpeed `xml:"PositionSpeed,omitempty"`
// 	// Next timing points for the train
// 	NextTimingPointTrainInformation []NextTimingPointTrainInformation `xml:"NextTimingPointTrainInformation,omitempty"`
// 	// Previous timing points passed by the train
// 	PreviousTimingPointTrainInformation []PreviousTimingPointTrainInformation `xml:"PreviousTimingPointTrainInformation,omitempty"`
// 	// Last Km reference point passed by the train
// 	LastPassedKmRef *LastPassedKmRef `xml:"LastPassedKmRef,omitempty"`
// 	// Last signal passed by the train
// 	LastPassedSignal *LastPassedSignal `xml:"LastPassedSignal,omitempty"`
// 	// Average speed of the train
// 	AverageSpeed *AverageSpeed `xml:"AverageSpeed,omitempty"`
// 	// Maximum speed of the train possible at this track
// 	MaxSpeedTrack string `xml:"MaxSpeedTrack,omitempty"`
// 	// Delay of the train
// 	Delay *Delay `xml:"Delay,omitempty"`
// 	// Date and time of the measurement
// 	DateTimeMeasurement UTCDateTimeType `xml:"DateTimeMeasurement,omitempty"`
// }

// // TrainReadyStatus models the corresponding XSD type/element.
// type TrainReadyStatus struct {
// 	XMLName xml.Name `xml:"TrainReadyStatus"`
// 	// 0=Not Ready 1=Ready
// 	TrainReady    TrainReady     `xml:"TrainReady"`
// 	TrainLocation *LocationIdent `xml:"TrainLocation,omitempty"`
// 	// Identifies the Delta delay time of a train against the booked schedule as well as against the
// 	// referenced time
// 	TrainDelay *TrainDelay `xml:"TrainDelay,omitempty"`
// 	// This element identifies the reason for a delay (modified DelayReason)
// 	DelayCause DelayCode `xml:"DelayCause,omitempty"`
// 	// Precise reason for a delay (string)
// 	PreciseDelayCause string `xml:"PreciseDelayCause,omitempty"`
// 	// It indicates date/time when the train will be ready for departure. IM contract will define if
// 	// this element can be used. Only to be used if this time is in future (sent in advance). Purpose
// 	// of this element is to short the train stay.
// 	TrainReadyTime string `xml:"TrainReadyTime,omitempty"`
// }

// // TrainsAfterOwn: Trains that share the same trainpath as your train, which are behind you. These
// // trains will use the same infrastructure after you.
// type TrainsAfterOwn struct {
// 	XMLName xml.Name `xml:"TrainsAfterOwn"`
// 	// Contains the identification of a related train and the segments
// 	Train []Train `xml:"Train,omitempty"`
// }

// // TrainsBeforeOwn: Trains that share the same trainpath as your train, which will use the same
// // infrastructure before you.
// type TrainsBeforeOwn struct {
// 	XMLName xml.Name `xml:"TrainsBeforeOwn"`
// 	// Contains the identification of a related train and the segments
// 	Train []Train `xml:"Train,omitempty"`
// }

// // Tunnel: Tunnel
// type Tunnel struct {
// 	XMLName xml.Name `xml:"Tunnel"`
// 	SP_GenericArea_ComplexType
// 	// [SUBSET-126] Category of the Tunnel (single or double track)
// 	TunnelCategory string `xml:"tunnelCategory,attr,omitempty"`
// 	// Aerodynamic cross section of the tunnel in expressed in m^2. Used in predictive tunnel
// 	// resistance equation developed by Z. Filipovic. See
// 	// https://www.sciencedirect.com/science/article/pii/S2210970616300415 (5)
// 	TunnelFactor string `xml:"tunnelFactor,attr,omitempty"`
// }

// // Unavailable_DAS_OperatingModes: Operating modes that are unavailable. If all 3 attributes are
// // included, the triplet is unavailable; of only 1 or 2 attributes are included, that mode is
// // unavailable independently of the value of the rest of the attributes.
// type Unavailable_DAS_OperatingModes struct {
// 	XMLName xml.Name `xml:"Unavailable_DAS_OperatingModes"`
// 	// DAS driving mode (see enumeration)
// 	DAS_drivingMode string `xml:"DAS_drivingMode,attr,omitempty"`
// 	// DAS architecture: DAS-C (ground advice calculation) or DAS-O (board advice calculation)
// 	DAS_architecture string `xml:"DAS_architecture,attr,omitempty"`
// 	// DAS connectivity (standalone or connected)
// 	DAS_connectivity string `xml:"DAS_connectivity,attr,omitempty"`
// }

// // UnavailableTractionType: Traction Type of the TrainCharacteristics of the train that has become
// // completely unavailable (e.g. due to a failure of the pantograph)
// type UnavailableTractionType struct {
// 	XMLName xml.Name `xml:"UnavailableTractionType"`
// 	// Traction Type (DC/AC with voltage, self-propelled or powerless)
// 	TractionType TractionTypeType `xml:"tractionType,attr"`
// }

// // UnprotectedLevelCrossingStop: Stop in rear of an unprotected level crossing
// type UnprotectedLevelCrossingStop struct {
// 	XMLName xml.Name `xml:"UnprotectedLevelCrossingStop"`
// 	SP_GenericPoint_ComplexType
// 	// Direction of application on the Segment Profile (Nominal, Reverse or Both)
// 	DirectionOfApplicationOnSP string `xml:"directionOfApplicationOnSP,attr,omitempty"`
// }

// // VirtualBalise: A VirtualBalise is a location in the Segment Profile for which the longitude and
// // latitude is known.
// type VirtualBalise struct {
// 	XMLName xml.Name `xml:"VirtualBalise"`
// 	SP_GenericPoint_ComplexType
// 	// GNSS position of Virtual Balise
// 	VirtualBalisePosition *GNSS_StaticPosition_ComplexType `xml:"VirtualBalisePosition,omitempty"`
// }

// // GNSS_StaticPosition_ComplexType: Complex Type for GNSS Static Position
// type GNSS_StaticPosition_ComplexType struct {
// 	// Latitude in signed degrees format (DDD.ddddd) in the WGS84 system.
// 	Latitude string `xml:"latitude,attr"`
// 	// Longitude in signed degrees format (DDD.ddddd) in the WGS84 system.
// 	Longitude string `xml:"longitude,attr"`
// 	// Altitude (m). Considering ETRS89 as reference.
// 	Altitude string `xml:"altitude,attr,omitempty"`
// }

// // WagonData: Wagon relevant data for the wagons within a running train
// type WagonData struct {
// 	XMLName xml.Name `xml:"WagonData"`
// 	// Identifies uniquely the wagon by its number
// 	WagonNumber string `xml:"WagonNumber,omitempty"`
// 	// Identifies the position of a wagon within a train. Sequential number starting with the first
// 	// wagon at the front of train as N°1.
// 	WagonTrainPosition WagonTrainPosition `xml:"WagonTrainPosition,omitempty"`
// 	// Maximum allowed speed of the wagon according to the load and entry in the Rolling Stock
// 	// Databases. In kmh
// 	WagonMaxSpeed WagonMaxSpeed `xml:"WagonMaxSpeed,omitempty"`
// 	// The total weight of the transportation unit on the freight wagon. This is the booked or actual
// 	// weight of goods including packing and carrier's equipment
// 	TotalLoadWeight WeightValueKilo `xml:"TotalLoadWeight,omitempty"`
// 	// The weight of an empty wagon according to the entry in the rolling stock database
// 	WagonWeightEmpty WeightValueKilo `xml:"WagonWeightEmpty,omitempty"`
// }

// // Wind: Temporary constraints due to wind factors
// type Wind struct {
// 	XMLName xml.Name `xml:"Wind"`
// 	// Wind speed
// 	WindSpeed SpeedType `xml:"windSpeed,attr"`
// 	// Wind heading, in degrees from North
// 	WindHeading string `xml:"windHeading,attr"`
// }

// // Error_ComplexType: Complex type for error in response messages
// type Error_ComplexType struct {
// 	// Error Code
// 	ErrorCode string `xml:"errorCode,attr,omitempty"`
// 	// XML Path to identify the location of the error
// 	XPath string `xml:"XPath,attr,omitempty"`
// 	// Additional details
// 	AdditionalInfo string `xml:"additionalInfo,attr,omitempty"`
// }

// // SP_GenericPoint_ComplexType: Complex Type for generic Segment Profile Point
// type SP_GenericPoint_ComplexType struct {
// 	// Element identifier
// 	Identifier string `xml:"identifier,attr,omitempty"`
// 	// Distance of the element from the beginning of the Segment Profile (m)
// 	Location DistanceType `xml:"location,attr"`
// }

// // TC_Identifier_ComplexType: Type for the identification of Train Characteristics. Contains the ID
// // of TC, and the mandatory identification of the entity responsible (RU or VKM) for the data
// type TC_Identifier_ComplexType struct {
// 	// Train Characteristics Identifier
// 	TC_ID string `xml:"TC_ID,attr"`
// 	// Train Characteristics - Identifier of the RU operating the train (use UIC RICS Code:
// 	// https://uic.org/rics)
// 	TC_RU_ID CompanyCode `xml:"TC_RU_ID"`
// 	// Free text, short name/vehicle keeper marking of the wagon keeper. Use ERA VKM Register:
// 	// https://www.era.europa.eu/registers/vkm_en
// 	KeeperShortNameVKM KeeperShortNameVKM `xml:"KeeperShortNameVKM"`
// }

// // TemporaryConstraints_ComplexType: Complex type for Temporary Constraints
// type TemporaryConstraints_ComplexType struct {
// 	// [SUBSET-126] Distance of the element's start location from the beginning of the Segment Profile
// 	// (m)
// 	StartLocation DistanceType `xml:"startLocation,attr,omitempty"`
// 	// [SUBSET-126] Distance of the element's end location from the beginning of the Segment Profile
// 	// (m)
// 	EndLocation DistanceType `xml:"endLocation,attr,omitempty"`
// 	// [SUBSET-126] Specifies if the element starts, ends, starts and ends or covers the whole
// 	// concerning Segment Profile
// 	StartEndQualifier string `xml:"startEndQualifier,attr"`
// 	// Starting date and time (in UTC)
// 	StartTime UTCDateTimeType `xml:"startTime,attr,omitempty"`
// 	// Ending date and time (in UTC)
// 	EndTime UTCDateTimeType `xml:"endTime,attr,omitempty"`
// }

// // TrainIdentifier_ComplexType: Complex Type for train identifier. Train identifier can be TAP/TAF
// // or OTN.
// type TrainIdentifier_ComplexType struct {
// 	// TAP/TAF Identifier
// 	TAP_TAF_ID *TAP_TAF_ID_ComplexType `xml:"TAP_TAF_ID"`
// 	// Operational Train Number (OTN) Identifier
// 	OTN_ID *OTN_ID_ComplexType `xml:"OTN_ID"`
// }

// // TrainType: Complex type for the "Train" object used in related train information.
// type TrainType struct {
// 	// Start from 1, the first is the the nearest and 2 the one after that, etc
// 	Order            string            `xml:"order,attr,omitempty"`
// 	TrainInformation *TrainInformation `xml:"TrainInformation,omitempty"`
// 	// Information about the current location and occupation of the train
// 	TrainLocationInformation *TrainLocationInformation `xml:"TrainLocationInformation,omitempty"`
// }

// // InterruptionPointRequest models the corresponding XSD type/element.
// type InterruptionPointRequest struct {
// 	XMLName xml.Name `xml:"InterruptionPointRequest"`
// 	// Identification of the train
// 	TrainIdentification *TrainIdentification `xml:"TrainIdentification,omitempty"`
// }

// // Signal_CantDeficiencyTrainCategory: Cant Deficiency train category to which the signal is applied
// type Signal_CantDeficiencyTrainCategory struct {
// 	XMLName xml.Name `xml:"Signal_CantDeficiencyTrainCategory"`
// 	// Cant Deficiency Train Category (NC_CDTRAIN as defined in SUBSET-026).
// 	NC_CDTrain NC_CDTrainType `xml:"NC_CDTrain,attr"`
// }

// // Signal_OtherTrainCategory: Other train category to which the signal is applied
// type Signal_OtherTrainCategory struct {
// 	XMLName xml.Name `xml:"Signal_OtherTrainCategory"`
// 	// Train Category (NC_TRAIN as defined in SUBSET-026).
// 	NC_Train NC_TrainType `xml:"NC_Train,attr,omitempty"`
// 	// Code for the train category (freeform text, limited to 10 characters)
// 	TrainCategoryCode string `xml:"trainCategoryCode,attr,omitempty"`
// }

// // ---------------------------------------------------------------------------
// // Convenience parse/format helpers
// // ---------------------------------------------------------------------------
// //
// // All scalar fields in this file are defined string types so that
// // Unmarshal->Marshal never rewrites a value's textual form. These helpers
// // let callers convert to/from native Go types on demand, e.g.:
// //
// //   spd, err := sfera.ParseFloat(string(positionSpeed.Speed))
// //   positionSpeed.Speed = sfera.SpeedType(sfera.FormatFloat(newSpeed))
// //
// //   t, err := sfera.ParseDateTime(string(msgRef.MessageDateTime))
// //   msgRef.MessageDateTime = sfera.FormatDateTime(time.Now().UTC())

// // ParseBool converts an XSD xs:boolean lexical value ("true", "false",
// // "1", or "0") to a Go bool.
// func ParseBool(s string) (bool, error) {
// 	return strconv.ParseBool(s)
// }

// // FormatBool renders b as the canonical XSD xs:boolean lexical value.
// func FormatBool(b bool) string {
// 	return strconv.FormatBool(b)
// }

// // ParseInt converts an XSD integer-family lexical value (xs:integer,
// // xs:int, xs:short, xs:unsignedByte, xs:positiveInteger, ...) to an int64.
// func ParseInt(s string) (int64, error) {
// 	return strconv.ParseInt(strings.TrimSpace(s), 10, 64)
// }

// // FormatInt renders i as an XSD integer lexical value.
// func FormatInt(i int64) string {
// 	return strconv.FormatInt(i, 10)
// }

// // ParseFloat converts an XSD xs:decimal or xs:double lexical value to a
// // float64. Note: re-formatting a parsed float with FormatFloat is not
// // guaranteed to reproduce the exact original text (e.g. trailing zeros,
// // "+" signs); prefer keeping the original defined-string field untouched
// // unless the value is actually being changed.
// func ParseFloat(s string) (float64, error) {
// 	return strconv.ParseFloat(strings.TrimSpace(s), 64)
// }

// // FormatFloat renders f as an XSD xs:decimal/xs:double lexical value using
// // the shortest representation that round-trips exactly through float64.
// func FormatFloat(f float64) string {
// 	return strconv.FormatFloat(f, 'f', -1, 64)
// }

// // ParseDateTime converts an XSD xs:dateTime lexical value (used by
// // UTCDateTimeType and friends) to a time.Time.
// func ParseDateTime(s string) (time.Time, error) {
// 	return time.Parse(time.RFC3339, strings.TrimSpace(s))
// }

// // FormatDateTime renders t as an XSD xs:dateTime lexical value in UTC,
// // e.g. "2026-08-09T12:34:56Z".
// func FormatDateTime(t time.Time) string {
// 	return t.UTC().Format(time.RFC3339)
// }

// // ParseDuration converts an XSD xs:duration lexical value (e.g. "PT30S")
// // is not directly supported by the standard library; this helper handles
// // the common simple case of a pure PT..H..M..S duration string.
// func ParseDuration(s string) (time.Duration, error) {
// 	return time.ParseDuration(xsdDurationToGoDuration(strings.TrimSpace(s)))
// }

// func xsdDurationToGoDuration(s string) string {
// 	// Best-effort conversion of a "PT#H#M#S" XSD duration into the Go
// 	// duration syntax ("#h#m#s"). Falls back to passing the string through
// 	// unchanged (letting time.ParseDuration report the error) for anything
// 	// more exotic (year/month/day components, which time.Duration cannot
// 	// represent).
// 	if !strings.HasPrefix(s, "PT") {
// 		return s
// 	}
// 	body := strings.TrimPrefix(s, "PT")
// 	body = strings.ReplaceAll(body, "H", "h")
// 	body = strings.ReplaceAll(body, "M", "m")
// 	body = strings.ReplaceAll(body, "S", "s")
// 	return body
// }

// // FormatDuration renders d in a form accepted by ParseDuration (a subset
// // of the XSD xs:duration lexical space limited to hours/minutes/seconds).
// func FormatDuration(d time.Duration) string {
// 	s := d.String()
// 	s = strings.ReplaceAll(s, "h", "H")
// 	s = strings.ReplaceAll(s, "m", "M")
// 	s = strings.ReplaceAll(s, "s", "S")
// 	return "PT" + s
// }

// // Unmarshal decodes SFERA XML data into any of the message/element struct
// // types defined in this file, e.g.:
// //
// //	var msg SFERA_B2G_EventMessage
// //	if err := sfera.Unmarshal(data, &msg); err != nil { ... }
// func Unmarshal(data []byte, v interface{}) error {
// 	return xml.Unmarshal(data, v)
// }

// // Marshal encodes v (a pointer to one of the struct types in this file)
// // back to XML. Combined with Unmarshal, this supports full roundtripping
// // of SFERA documents.
// func Marshal(v interface{}) ([]byte, error) {
// 	return xml.Marshal(v)
// }

// // MarshalIndent is like Marshal but indents the output for readability.
// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
// 	return xml.MarshalIndent(v, prefix, indent)
// }
