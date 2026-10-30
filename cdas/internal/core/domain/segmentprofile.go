package domain

import sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"

type SegmentProfileInfo struct {
	Major    int
	Minor    int
	Checksum string
	Segment  *sfera.SegmentProfile
}
