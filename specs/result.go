package specs

type TestResult = uint

const (
	Unknown   TestResult = iota
	Failed    TestResult = iota
	Succeeded TestResult = iota
)
