package signalroute

import "github.com/tylland/transition/internal/schema/routes"

// Writer is the interface for writing SFERA signal routes.
type Writer interface {
	Write(r *routes.SignalRoutes) error
}

// NewFileWriter creates a FileWriter that writes to the given file path.
func NewFileWriter(path string) *FileWriter {
	return &FileWriter{path: path}
}
