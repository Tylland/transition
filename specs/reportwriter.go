package specs

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	rm "github.com/tylland/transition/specs/reportmodel"
)

type ReportWriter interface {
	Attach(session *Session)
	WriteStory(story *Story)
	WriteScenario(story *Story, title string)
	WriteStep(prefix StepPrefix, title string, result rm.Result, err error)
	WriteExampleHeaders(headers []ExampleHeader)
	WriteExample(values []any, result TestResult, err error)
	Finish() *rm.ReportModel
}

type JsonReportWriter struct {
	outputFileName string
	builder        *rm.ReportModelBuilder
	story          *rm.StoryBuilder
	scenario       *rm.ScenarioBuilder
}

func NewJsonReportWriter() *JsonReportWriter {
	return &JsonReportWriter{builder: rm.NewReportModelBuilder()}
}

func hashID(content string) string {
	h := sha1.New()
	h.Write([]byte(content))

	return hex.EncodeToString(h.Sum(nil))
}

func (r *JsonReportWriter) Attach(session *Session) {

	r.outputFileName = filepath.Join(session.OutputDir, strings.ToLower(session.Name)+".report.json")
}

func (r *JsonReportWriter) WriteStory(story *Story) {
	r.builder.WithStory(hashID(story.String())).
		WithNarrative(story.UserStory.AsA, story.UserStory.IWantTo, story.UserStory.SoThat).
		WithTags(story.tags)
}

func (r *JsonReportWriter) WriteScenario(story *Story, title string) {
	storyID := hashID(story.String())
	scenarioID := hashID(title)

	r.scenario = r.builder.WithStory(storyID).WithScenario(scenarioID)
	r.scenario.WithTitle(title)
}

func (r *JsonReportWriter) WriteStep(prefix StepPrefix, title string, result rm.Result, err error) {
	stepID := hashID(title)

	//	fmt.Print(title + "  " + stepID + "  ")

	r.scenario.WithStep(stepID).
		WithTitle(title).
		WithAsserts(prefix == Then).
		WithResult(result).
		WithError(rm.NewError(err))
}

func (r *JsonReportWriter) WriteExampleHeaders(headers []ExampleHeader) {
	var titles []string

	for _, header := range headers {
		titles = append(titles, header.Title)
	}

	r.scenario.WithExample().WithHeader(titles)
}

func (r *JsonReportWriter) WriteExample(values []any, res TestResult, err error) {
	result := rm.Skip
	var vals []string

	for _, v := range values {
		vals = append(vals, fmt.Sprintf("%v", v))
	}

	if res == Succeeded {
		result = rm.Pass
	} else if res == Failed {
		result = rm.Fail
	}

	r.scenario.WithExample().AddExample(vals, result, err)
}

func (r *JsonReportWriter) save(reportFileName string, model *rm.ReportModel) {
	json, err := json.MarshalIndent(model, "", "  ")

	dir := path.Dir(reportFileName)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0777)

		if err != nil {
			fmt.Print(err.Error())
		}
	}

	if err == nil {
		err = os.WriteFile(reportFileName, json, 0666)
	}
}

func (r *JsonReportWriter) Finish() *rm.ReportModel {
	model := r.builder.Build()

	r.save(r.outputFileName, model)

	fmt.Printf("Test results saved to file: %s", r.outputFileName)

	return model
}
