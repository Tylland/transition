package sfera

// StartEndQualifier for area positions.

const (
	StartEndQualifierStarts     StartEndQualifier = "Starts"
	StartEndQualifierEnds       StartEndQualifier = "Ends"
	StartEndQualifierStartsEnds StartEndQualifier = "StartsEnds"
	StartEndQualifierWholeSP    StartEndQualifier = "WholeSP"
)

const (
	SpecificSSPCategoryOtherReplaces SpecificSSPCategory = "Other_SSP_Replaces_CDSSP"
)

const (
	NCTrainTypeFreightP  NCTrainType = "Freight_Train_P_Position"
	NCTrainTypeFreightG  NCTrainType = "Freight_Train_G_Position"
	NCTrainTypePassenger NCTrainType = "Passenger_Train"
)

const (
	GradientDirectionDownhill GradientDirectionType = "Downhill"
	GradientDirectionUphill   GradientDirectionType = "Uphill"
)

const (
	SegmentProfileStatusValid   SPStatus = "Valid"
	SegmentProfileStatusInvalid SPStatus = "Invalid"
)

const (
	AdhesionCategoryTypeDryRail              AdhesionCategoryType = "Dry Rail"
	AdhesionCategoryTypeDryRailMedium        AdhesionCategoryType = "Dry Rail (Medium)"
	AdhesionCategoryTypeDryRailLow           AdhesionCategoryType = "Dry Rail (Low)"
	AdhesionCategoryTypeLowAdhesion          AdhesionCategoryType = "Low Adhesion"
	AdhesionCategoryTypeVeryLowAdhesion      AdhesionCategoryType = "Very Low Adhesion"
	AdhesionCategoryTypeExtremelyLowAdhesion AdhesionCategoryType = "Extremely Low Adhesion"
)

const (
	DelayCode11 DelayCode = "11"
	DelayCode10 DelayCode = "10"
	DelayCode12 DelayCode = "12"
	DelayCode13 DelayCode = "13"
	DelayCode14 DelayCode = "14"
	DelayCode18 DelayCode = "18"
	DelayCode19 DelayCode = "19"
	DelayCode20 DelayCode = "20"
	DelayCode21 DelayCode = "21"
	DelayCode22 DelayCode = "22"
	DelayCode23 DelayCode = "23"
	DelayCode24 DelayCode = "24"
	DelayCode25 DelayCode = "25"
	DelayCode28 DelayCode = "28"
	DelayCode29 DelayCode = "29"
	DelayCode31 DelayCode = "31"
	DelayCode30 DelayCode = "30"
	DelayCode32 DelayCode = "32"
	DelayCode39 DelayCode = "39"
	DelayCode40 DelayCode = "40"
	DelayCode41 DelayCode = "41"
	DelayCode50 DelayCode = "50"
	DelayCode51 DelayCode = "51"
	DelayCode52 DelayCode = "52"
	DelayCode53 DelayCode = "53"
	DelayCode54 DelayCode = "54"
	DelayCode58 DelayCode = "58"
	DelayCode59 DelayCode = "59"
	DelayCode60 DelayCode = "60"
	DelayCode61 DelayCode = "61"
	DelayCode62 DelayCode = "62"
	DelayCode63 DelayCode = "63"
	DelayCode64 DelayCode = "64"
	DelayCode68 DelayCode = "68"
	DelayCode70 DelayCode = "70"
	DelayCode69 DelayCode = "69"
	DelayCode71 DelayCode = "71"
	DelayCode80 DelayCode = "80"
	DelayCode81 DelayCode = "81"
	DelayCode82 DelayCode = "82"
	DelayCode83 DelayCode = "83"
	DelayCode84 DelayCode = "84"
	DelayCode89 DelayCode = "89"
	DelayCode90 DelayCode = "90"
	DelayCode91 DelayCode = "91"
	DelayCode92 DelayCode = "92"
	DelayCode93 DelayCode = "93"
	DelayCode94 DelayCode = "94"
	DelayCode95 DelayCode = "95"
)

// Enumerated values for NC_CDTrainType.
const (
	NC_CDTrainType80mm  NCCDTrainType = "80mm"
	NC_CDTrainType100mm NCCDTrainType = "100mm"
	NC_CDTrainType130mm NCCDTrainType = "130mm"
	NC_CDTrainType150mm NCCDTrainType = "150mm"
	NC_CDTrainType165mm NCCDTrainType = "165mm"
	NC_CDTrainType180mm NCCDTrainType = "180mm"
	NC_CDTrainType210mm NCCDTrainType = "210mm"
	NC_CDTrainType225mm NCCDTrainType = "225mm"
	NC_CDTrainType245mm NCCDTrainType = "245mm"
	NC_CDTrainType275mm NCCDTrainType = "275mm"
	NC_CDTrainType300mm NCCDTrainType = "300mm"
)

const (
	NC_TrainTypeFreightTrainPPosition NCTrainType = "Freight_Train_P_Position"
	NC_TrainTypeFreightTrainGPosition NCTrainType = "Freight_Train_G_Position"
	NC_TrainTypePassengerTrain        NCTrainType = "Passenger_Train"
)

const (
	RelatedTrainRequestTypeNone                       RelatedTrainRequestType = "None"
	RelatedTrainRequestTypeOwnTrain                   RelatedTrainRequestType = "OwnTrain"
	RelatedTrainRequestTypeRelatedTrains              RelatedTrainRequestType = "RelatedTrains"
	RelatedTrainRequestTypeOwnTrainAndRelatedTrains   RelatedTrainRequestType = "OwnTrainAndRelatedTrains"
	RelatedTrainRequestTypeOwnTrainAndOrRelatedTrains RelatedTrainRequestType = "OwnTrainAndOrRelatedTrains"
)

const (
	TractionTypeType600VDC         TractionTypeType = "600V DC"
	TractionTypeType750VDC         TractionTypeType = "750V DC"
	TractionTypeType1500VDC        TractionTypeType = "1500V DC"
	TractionTypeType3000VDC        TractionTypeType = "3000V DC"
	TractionTypeType25000VAC50Hz   TractionTypeType = "25000V AC 50Hz"
	TractionTypeType15000VAC16Hz23 TractionTypeType = "15000V AC 16Hz2/3"
	TractionTypeTypeSelfPropelled  TractionTypeType = "Self-propelled"
	TractionTypeTypePowerless      TractionTypeType = "Powerless"
)

const (
	TrainPartTypeTrainHead           TrainPartType = "TrainHead"
	TrainPartTypeTrainEnd            TrainPartType = "TrainEnd"
	TrainPartTypeTrainLastPantograph TrainPartType = "TrainLastPantograph"
)

// const (
// 	ClaimTypePlanned  ClaimType = "Planned"
// 	ClaimTypeReserved ClaimType = "Reserved"
// 	ClaimTypeOccupied ClaimType = "Occupied"
// 	ClaimTypeUnknown  ClaimType = "Unknown"
// )

// const (
// 	LeavingEnteringStatusEntering LeavingEnteringStatus = "Entering"
// 	LeavingEnteringStatusLeaving  LeavingEnteringStatus = "Leaving"
// 	LeavingEnteringStatusCrossing LeavingEnteringStatus = "Crossing"
// )

// const (
// 	ObjectTypeTR ObjectType = "TR"
// 	ObjectTypePA ObjectType = "PA"
// 	ObjectTypeCR ObjectType = "CR"
// 	ObjectTypePR ObjectType = "PR"
// )

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
