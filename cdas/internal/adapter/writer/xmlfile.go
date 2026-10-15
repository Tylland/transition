package writer

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tylland/transition/internal/core/domain"
	sfera "github.com/tylland/transition/internal/schema/sfera/v3"
)

type xmlFileWriter struct {
	path string
}

func (w *xmlFileWriter) write(r any) error {
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

type SignalRoutesXMLFileWriter struct {
	xmlFileWriter
}

func NewSignalRoutesXMLFileWriter(path string) *SignalRoutesXMLFileWriter {
	return &SignalRoutesXMLFileWriter{xmlFileWriter: xmlFileWriter{path: path}}
}

func (w *SignalRoutesXMLFileWriter) Write(r *domain.SignalRoutes) error {
	return w.write(r)
}

type TrackRoutesXMLFileWriter struct {
	xmlFileWriter
}

func NewTrackRoutesXMLFileWriter(path string) *TrackRoutesXMLFileWriter {
	return &TrackRoutesXMLFileWriter{xmlFileWriter: xmlFileWriter{path: path}}
}

func (w *TrackRoutesXMLFileWriter) Write(r *domain.TrackRoutes) error {
	return w.write(r)
}

// FileWriter writes SegmentProfile objects as XML files to a folder.
type SegmentProfileXMLFileWriter struct {
	FolderPath string
}

func NewSegmentProfileXMLFileWriter(folderPath string) *SegmentProfileXMLFileWriter {
	return &SegmentProfileXMLFileWriter{FolderPath: folderPath}
}

// Write serializes a SegmentProfile to an XML file and returns the file path.
func (w *SegmentProfileXMLFileWriter) Write(segment *sfera.SegmentProfile) (string, error) {
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
