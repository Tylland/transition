package reporter

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"

	"github.com/tylland/transition/specs"
	rm "github.com/tylland/transition/specs/reportmodel"
)

type MarkdownReporter struct {
	outputFileName string
	outputFile     *os.File
}

func NewMarkdownReporter() *MarkdownReporter {
	return &MarkdownReporter{}
}

func (r *MarkdownReporter) Attach(session *specs.Session) {
	r.outputFileName = filepath.Join(session.OutputDir, strings.ToLower(session.Name)+".md")

	file, _ := os.Create(r.outputFileName)

	r.outputFile = file
}

func (r *MarkdownReporter) write(format string, a ...any) {
	fmt.Fprintf(r.outputFile, format, a...)
}

func (r *MarkdownReporter) writeln(format string, a ...any) {
	fmt.Fprintf(r.outputFile, format+"\n", a...)
}

func (r *MarkdownReporter) writeStory(story *rm.Story) {

	r.writeln("# Story")
	r.writeln("> %s", story.Metadata.AsAText())
	r.writeln(">")
	r.writeln("> I want to %s", story.Metadata.IWantTo)
	r.writeln(">")
	r.writeln("> So that %s", story.Metadata.SoThat)
	r.writeln(">")

	var tags []string

	for _, t := range story.Tags {
		tags = append(tags, fmt.Sprintf("&#127991; %s:%s", strings.ToLower(t.Name), strings.ToLower(t.Value)))
	}

	r.writeln(strings.Join(tags, " "))
}

func (r *MarkdownReporter) writeScenariosTOC(story *rm.Story) {

	r.writeln("## Scenarios")

	for _, scenario := range story.Scenarios {
		r.writeScenarioResult(scenario)
	}

	r.writeln("")
	r.writeln("---")
	r.writeln("")
}

func (r *MarkdownReporter) writeScenarioResult(scenario *rm.Scenario) {
	var checked string = "&#x274C;"

	if scenario.Result == rm.Pass {
		checked = "&#9989;"
	}

	parts := strings.Split(scenario.Title, " ")
	tag := html.EscapeString("scenario-" + strings.ToLower(strings.Join(parts, "-")))

	r.writeln("%s %s [view](#%s)  ", checked, scenario.Title, tag)
}

func (r *MarkdownReporter) writeScenario(scenario *rm.Scenario) {
	r.writeln("")
	r.writeln("## Scenario: %s", scenario.Title)
}

func (r *MarkdownReporter) writeSteps(steps []*rm.Step) {
	for _, step := range steps {
		message := ""

		if step.Error != nil {
			message = step.Error.Message
		}

		r.writeln("%-60s %s %s", step.Title, step.Result, message)
	}
}

func (r *MarkdownReporter) writeHeaders(headers []string) {
	for i := 0; i < len(headers); i++ {
		r.write("%-10s", fmt.Sprintf("<%s>", headers[i]))
	}

	r.writeln("")
}

func (r *MarkdownReporter) writeExample(values []any, message string) {
	for i := 0; i < len(values); i++ {
		r.write("%-10s", fmt.Sprintf("|%s", values[i]))
	}

	r.writeln("| %s", message)
}

func (r *MarkdownReporter) writeExamples(examples []*rm.Example) {
	const totalWidth = 60
	const colWidth = 10

	for _, example := range examples {
		exampleStr := ""
		for i := 0; i < len(example.Values); i++ {
			exampleStr += fmt.Sprintf("%-*s", colWidth, "|"+example.Values[i])
		}

		exampleStr += "|"

		message := ""

		if example.Error != nil {
			message = example.Error.Message
		}

		r.writeln("%-*s %s %s", totalWidth, exampleStr, example.Result, message)
	}
}

func (r *MarkdownReporter) writeExampleTable(table *rm.ExampleTable) {
	if table != nil {
		r.writeln("")
		r.writeln("Examples")
		r.writeHeaders(table.Headers)
		r.writeExamples(table.Examples)
	}
}

func (r *MarkdownReporter) Report(model *rm.ReportModel) {
	r.writeln("")

	for _, story := range model.Stories {
		r.writeStory(story)
		r.writeScenariosTOC(story)

		for _, scenario := range story.Scenarios {
			r.writeScenario(scenario)
			r.writeln("``` GHERKIN")
			r.writeSteps(scenario.Steps)
			r.writeExampleTable(scenario.ExampleTable)
			r.writeln("```")
		}

	}

	r.outputFile.Close()

	fmt.Printf("Report saved to file: %s", r.outputFileName)
}
