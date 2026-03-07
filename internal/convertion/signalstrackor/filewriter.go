package signalstrackor

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/tylland/transition/internal/schema/signalstrackor"
)

// FileWriter writes Signalstrackor objects as XML files.
type FileWriter struct {
	Path string
}

// NewFileWriter creates a new FileWriter.
func NewFileWriter(path string) *FileWriter {
	return &FileWriter{Path: path}
}

// Write serializes Signalstrackor to an XML file.
func (w *FileWriter) Write(s *signalstrackor.Signalstrackor) error {
	data, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal signalstrackor: %w", err)
	}

	output := append([]byte(xml.Header), data...)
	if err := os.WriteFile(w.Path, output, 0644); err != nil {
		return fmt.Errorf("failed to write signalstrackor to %s: %w", w.Path, err)
	}

	return nil
}
