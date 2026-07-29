package segmentprofile

import "github.com/tylland/transition/internal/schema/sfera"

// Writer is the interface for writing SFERA segment profiles.
type Writer interface {
	Write(segment *sfera.SegmentProfile) (string, error)
}
