package sfera

import (
	"fmt"

	dm "github.com/tylland/transition/internal/domainmodel"
	"github.com/tylland/transition/internal/schema/sfera"
)

// CreateProfileId creates a SFERA segment profile ID from a LankId and direction.
func CreateProfileId(linkId dm.LinkId, direction dm.LinkDirection) string {
	riktningIndex := int(direction)
	return fmt.Sprintf("%v:%d", linkId, riktningIndex)
}

// GetGradientDirection returns the SFERA gradient direction for a Slope.
func GetGradientDirection(slope *dm.Slope, direction dm.LinkDirection) sfera.GradientDirectionType {
	gradient := slope.Gradient
	if direction == dm.LinkDirectionAgainst {
		gradient = -gradient
	}

	if gradient >= 0 {
		return sfera.GradientDirectionUphill
	}
	return sfera.GradientDirectionDownhill
}
