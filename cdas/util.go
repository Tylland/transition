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

func CompanyCode(s string) *sfera.TeltsiCompanyCode {
	s = PadLeft(s, "0", 4)

	return new(sfera.TeltsiCompanyCode(s))
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
