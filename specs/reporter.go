package specs

import rm "github.com/tylland/transition/specs/reportmodel"

type StepResult struct {
	Prefix  StepPrefix
	Title   string
	Result  TestResult
	Message string
}

func NewStepResult(prefix StepPrefix, title string, result TestResult, message string) *StepResult {
	return &StepResult{Prefix: prefix, Title: title, Result: result, Message: message}
}

type Reporter interface {
	Attach(session *Session)
	Report(reportModel *rm.ReportModel)
}

type ReporterList struct {
	reporters []Reporter
}

func (rl *ReporterList) Add(reporter Reporter) {
	rl.reporters = append(rl.reporters, reporter)
}

func (rl *ReporterList) WithAll(action func(reporter Reporter)) {
	for _, reporter := range rl.reporters {
		action(reporter)
	}
}

func (rl *ReporterList) Report(reportModel *rm.ReportModel) {
	rl.WithAll(func(r Reporter) { r.Report(reportModel) })
}
