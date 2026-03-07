package routes

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/tylland/transition/internal/schema/routes"
)

// WriteRoutesFile writes a Routes object as an XML file.
func WriteRoutesFile(r *routes.Routes, path string) error {
	data, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal routes: %w", err)
	}

	output := append([]byte(xml.Header), data...)
	if err := os.WriteFile(path, output, 0644); err != nil {
		return fmt.Errorf("failed to write routes to %s: %w", path, err)
	}

	return nil
}
