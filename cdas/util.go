package cdas

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
