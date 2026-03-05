package sfera

import "github.com/tylland/transition/internal/schema/sfera"

// SegmentProfileWriter is the interface for writing SFERA segment profiles.
type SegmentProfileWriter interface {
	Write(segment *sfera.SegmentProfile) (string, error)
}
