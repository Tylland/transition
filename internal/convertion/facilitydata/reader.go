package convertion

import (
	"encoding/xml"
	"fmt"
	"os"

	fd "github.com/tylland/transition/internal/schema/facilitydata"
)

// ReadFacilityData deserializes a FacilityData XML file from the given path.
func ReadFacilityData(path string) (*fd.FacilityData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not open facility data file %s: %w", path, err)
	}

	var facilityData fd.FacilityData
	if err := xml.Unmarshal(data, &facilityData); err != nil {
		return nil, fmt.Errorf("could not parse facility data file %s: %w", path, err)
	}

	return &facilityData, nil
}
