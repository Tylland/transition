package trackroute

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/tylland/transition/internal/schema/routes"
)

type FileWriter struct {
	path string
}

// WriteRoutesFile writes a Routes object as an XML file.
func (w *FileWriter) Write(r *routes.TrackRoutes) error {
	data, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal routes: %w", err)
	}

	output := append([]byte(xml.Header), data...)
	if err := os.WriteFile(w.path, output, 0644); err != nil {
		return fmt.Errorf("failed to write routes to %s: %w", w.path, err)
	}

	return nil
}
