package sfera

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tylland/transition/internal/schema/sfera"
)

// SegmentProfileFileWriter writes SegmentProfile objects as XML files to a folder.
type SegmentProfileFileWriter struct {
	FolderPath string
}

// NewSegmentProfileFileWriter creates a new SegmentProfileFileWriter.
func NewSegmentProfileFileWriter(folderPath string) *SegmentProfileFileWriter {
	return &SegmentProfileFileWriter{FolderPath: folderPath}
}

// Write serializes a SegmentProfile to an XML file and returns the file path.
func (w *SegmentProfileFileWriter) Write(segment *sfera.SegmentProfile) (string, error) {
	fileName := strings.ReplaceAll(segment.SPID, ":", "-") + ".xml"
	filePath := filepath.Join(w.FolderPath, fileName)

	data, err := xml.MarshalIndent(segment, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal segment profile: %w", err)
	}

	output := append([]byte(xml.Header), data...)
	if err := os.WriteFile(filePath, output, 0644); err != nil {
		return "", fmt.Errorf("failed to write segment profile to %s: %w", filePath, err)
	}

	return filePath, nil
}
