package sfera

import (
	"fmt"

	dm "github.com/tylland/transition/internal/domainmodel"
	"github.com/tylland/transition/internal/schema/sfera"
)

// CreateProfileId creates a SFERA segment profile ID from a LankId and direction.
func CreateProfileId(lankId dm.LankId, riktning dm.Lankriktning) string {
	riktningIndex := int(riktning)
	return fmt.Sprintf("%v:%d", lankId, riktningIndex)
}

// GetGradientDirection returns the SFERA gradient direction for a Lutning.
func GetGradientDirection(lutning *dm.Lutning, riktning dm.Lankriktning) sfera.GradientDirectionType {
	grader := lutning.Grader
	if riktning == dm.LankriktningMot {
		grader = -grader
	}

	if grader >= 0 {
		return sfera.GradientDirectionUphill
	}
	return sfera.GradientDirectionDownhill
}
