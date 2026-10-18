package writer

import (
	"fmt"

	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
	"github.com/tylland/transition/specs"
)

type SegmentBuilder struct {
	segmentProfile *sfera.SegmentProfile
}

func (b *SegmentBuilder) SegmentProfile(spid string) *specs.StepResult {
	res := specs.NewStepResult("segment profile %q", spid)

	b.segmentProfile = &sfera.SegmentProfile{
		SPVersionMajor: "1",
		SPVersionMinor: new("0"),
		SPLength:       123.45,
		SPName:         new("test-profile"),
		SPStatus:       sfera.SPStatusValid,
		SegmentProfileIdentificationComplexType: &sfera.SegmentProfileIdentificationComplexType{
			SPID: spid,
		},
		SPContextInformation: &sfera.SPContextInformation{},
	}

	return res.Success()
}

func (b *SegmentBuilder) LineAndTrackIdentifier(lineID, lineName, trackID *string) *specs.StepResult {
	label := ""

	if lineID != nil {
		label = fmt.Sprintf("line %q", *lineID)
	}

	if lineName != nil {
		if label != "" {
			label += "line "
		}
		label += fmt.Sprintf("(%q)", *lineName)
	}

	if trackID != nil {
		label = fmt.Sprintf("track %q and %s", *trackID, label)
	}

	res := specs.NewStepResult("%s", label)

	b.segmentProfile.SPContextInformation.LineAndTrackIdentifier = &sfera.LineAndTrackIdentifier{
		LineAndTrackIdentifierStart: &sfera.LineAndTrackIdentifierComplexType{
			LineIdentifier:  lineID,
			LineName:        lineName,
			TrackIdentifier: trackID,
		},
	}

	return res.Success()
}
