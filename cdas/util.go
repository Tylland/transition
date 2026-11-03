package cdas

import (
	"strings"

	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}

	return obj
}

const (
	TimestampFormat = "2006-01-02T15:04:05.000Z07:00"
	YearFormat      = "2006"
	DateFormat      = "2006-01-02"
)

const (
	SegmentProfileBucket = "SEGMENT_PROFILES"
	JourneyProfileBucket = "JOURNEY_PROFILES"
)

func CompanyCode(s string) *sfera.TeltsiCompanyCode {
	cc := sfera.TeltsiCompanyCode(PadLeft(s, "0", 4))
	return &cc
}

func Core(s string) string {
	return PadLeft(s, "-", 12)
}

func PadLeft(str, tkn string, lenght int) string {
	if n := lenght - len(str); n > 0 {
		return strings.Repeat(tkn, n) + str
	}

	return str
}
