package facilitydata

import (
	"encoding/xml"
	"fmt"
	"os"

	fd "github.com/tylland/transition/cdas/internal/schema/facilitydata"
)

// ReadFacilityData deserializes a FacilityData XML file from the given path.
func ReadXML(path string) (*fd.FacilityData, error) {
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
