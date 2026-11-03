package provider

import (
	"fmt"
	"strings"

	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

type SPKeyProvider struct {
}

func (p *SPKeyProvider) Info(id string) string {
	return fmt.Sprintf("sfera.segmentprofile.info.%s", id)
}

func (p *SPKeyProvider) Data(id string) string {
	return fmt.Sprintf("sfera.segmentprofile.data.%s", id)
}

type JPKeyProvider struct {
}

func (p *JPKeyProvider) Create(trainID *sfera.TrainIdentificationComplexType) string {
	company := string(*trainID.TAFTAPID.TeltsiCompany)
	id := *trainID.TAFTAPID

	return strings.ToLower(fmt.Sprintf("sfera.journeyprofile.%s.%s.%s.%s.%s.%s",
		id.TeltsiObjectType,
		company,
		id.TeltsiCore,
		id.TeltsiVariant,
		id.TeltsiTimetableYear,
		id.TeltsiStartDate))
}
