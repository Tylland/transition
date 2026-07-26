package planningdata

import (
	"encoding/json"
	"fmt"
	"os"

	pd "github.com/tylland/transition/internal/schema/planningdata"
)

// ReadPlanningData deserializes a DagligFärdplan JSON file from the given path.
func ReadPlanningData(path string) (*pd.Response, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not open planning data file %s: %w", path, err)
	}

	var root pd.JSONRoot
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, fmt.Errorf("could not parse planning data file %s: %w", path, err)
	}

	return &root.Response, nil
}
