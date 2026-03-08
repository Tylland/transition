package planningdata

import (
	"encoding/xml"
	"fmt"
	"os"

	pd "github.com/tylland/transition/internal/schema/planningdata"
)

// ReadPlanningData deserializes a DagligFärdplan XML file from the given path.
func ReadPlanningData(path string) (*pd.Response, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not open planning data file %s: %w", path, err)
	}

	var response pd.Response
	if err := xml.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("could not parse planning data file %s: %w", path, err)
	}

	return &response, nil
}
