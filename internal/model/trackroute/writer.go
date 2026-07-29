package trackroute

import "github.com/tylland/transition/internal/schema/routes"

// Writer is the interface for writing SFERA track routes.
type Writer interface {
	Write(r *routes.TrackRoutes) error
}

// NewFileWriter creates a FileWriter that writes to the given file path.
func NewFileWriter(path string) *FileWriter {
	return &FileWriter{path: path}
}
