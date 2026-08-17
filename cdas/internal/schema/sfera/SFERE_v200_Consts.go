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
