package segmentprofile

import sfera "github.com/tylland/transition/internal/schema/sfera/v3"

// Writer is the interface for writing SFERA segment profiles.
type Writer interface {
	Write(segment *sfera.SegmentProfile) (string, error)
}
