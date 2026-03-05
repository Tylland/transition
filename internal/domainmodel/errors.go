package domainmodel

import "fmt"

// ErrLankriktningOkand is returned when a Lankriktning value is unknown.
type ErrLankriktningOkand struct {
	ParamName string
}

func (e *ErrLankriktningOkand) Error() string {
	return fmt.Sprintf("Länkriktning okänd: %s", e.ParamName)
}

// NewErrLankriktningOkand creates a new ErrLankriktningOkand error.
func NewErrLankriktningOkand(paramName string) *ErrLankriktningOkand {
	return &ErrLankriktningOkand{ParamName: paramName}
}
