package domain

import (
	"fmt"
	"strings"
	"time"

	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

type ActiveTrainSession struct {
	TrainID *sfera.TrainIdentificationComplexType
	Devices map[string]*ActiveDevice
}

func NewActiveTrainSessionKey(trainID *sfera.TrainIdentificationComplexType) string {
	company := string(*trainID.TAFTAPID.TeltsiCompany)
	id := *trainID.TAFTAPID

	return strings.ToLower(fmt.Sprintf("%s.%s.%s.%s.%s.%s",
		id.TeltsiObjectType,
		company,
		id.TeltsiCore,
		id.TeltsiVariant,
		id.TeltsiTimetableYear,
		id.TeltsiStartDate))
}

type ActiveDevice struct {
	CompanyCode  string
	SourceDevice string
	DrivingMode  string
	Started      time.Time
	LastSeen     time.Time
}

func NewActiveDeviceKey(companyCode, sourceDevice string) string {
	return strings.ToLower(fmt.Sprintf("%s.%s", companyCode, sourceDevice))
}
