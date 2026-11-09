package service

import "github.com/tylland/transition/cdas/internal/core/port"

/*
The journey service reads planning data and creates journey profiles. It matches
planning data to facility data to determine which segment profiles to include
and in what order. It calculates timing points from the scheduled timetable and
bypasses; each timing point references the facility through target points in
the segment profiles. The service manages the lifecycle of journey profiles,
creating them and removing them when they become historical.
*/
type JourneyService struct {
	spReader port.SegmentProfileReader
	jpWriter port.JourneyProfileWriter
}

func NewJourneyService(spReader port.SegmentProfileReader,
	jpWriter port.JourneyProfileWriter) *JourneyService {

	return &JourneyService{spReader: spReader, jpWriter: jpWriter}
}
