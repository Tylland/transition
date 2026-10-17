package writer

import (
	"encoding/xml"
	"testing"

	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
	"github.com/tylland/transition/specs"
)

// ── Story declarations ────────────────────────────────────────────────────────

var writerStory = specs.NewStory(specs.UserStory{
	AsA:     "Facility Service",
	IWantTo: "publish segment profiles as xml",
	SoThat:  "segment profiles can be consumed by other services",
})

var b = SegmentBuilder{}
var paniced bool

func MarshalingToXML() *specs.StepResult {
	res := specs.NewStepResult("marshaling to XML")

	defer func() {
		if r := recover(); r != nil {
			res = res.WithErrorf("marshal should not panic: %v", r)
			paniced = true
		}
	}()

	_, err := xml.MarshalIndent(b.segmentProfile, "", "  ")
	if err != nil {
		return res.WithErrorf("marshal should succeed: %v", err)
	}

	return res.Success()
}

func NoPanic() *specs.StepResult {
	res := specs.NewStepResult("no panic is thrown")

	if paniced {
		return res.WithErrorMessage("paniced")
	}

	return res.Success()
}

func TestKMReferenceMarshalDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("marshal should not panic: %v", r)
		}
	}()

	_, err := xml.MarshalIndent(&sfera.KMReference{KmRef: 12.5}, "", "  ")
	if err != nil {
		t.Fatalf("marshal should succeed: %v", err)
	}
}

func TestSegmentProfileXMLMarshalDoesNotPanic(t *testing.T) {
	lineID := "E"
	trackID := "1"

	writerStory.Scenario(t).
		Given(b.SegmentProfile("sp001")).
		And(b.LineAndTrackIdentifier(&lineID, nil, &trackID)).
		When(MarshalingToXML()).
		Then(NoPanic()).
		Execute()
}
