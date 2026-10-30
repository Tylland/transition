package provider

import (
	"fmt"
)

type SPKeyProvider struct {
}

func (p *SPKeyProvider) Info(id string) string {
	return fmt.Sprintf("sfera.segmentprofile.info.%s", id)
}

func (p *SPKeyProvider) Data(id string) string {
	return fmt.Sprintf("sfera.segmentprofile.data.%s", id)
}
