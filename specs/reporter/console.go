package reporter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tylland/transition/specs"
	rm "github.com/tylland/transition/specs/reportmodel"
)

type ConsoleReporter struct {
	outputFileName string
	outputFile     *os.File
}

func NewConsoleReporter() *ConsoleReporter {
	return &ConsoleReporter{}
}

func (r *ConsoleReporter) Attach(session *specs.Session) {
	r.outputFileName = filepath.Join(session.OutputDir, strings.ToLower(session.Name)+".console.out")
}

func (r *ConsoleReporter) write(format string, a ...any) {
	fmt.Fprintf(r.outputFile, format, a...)
}

func (r *ConsoleReporter) writeln(format string, a ...any) {
	fmt.Fprintf(r.outputFile, format+"\n", a...)
}

func (r *ConsoleReporter) writeStory(story *rm.Story) {
	var tags []string

	for _, t := range story.Tags {
		tags = append(tags, fmt.Sprintf("@%s:%s", strings.ToLower(t.Name), strings.ToLower(t.Value)))
	}

	r.writeln("")
	r.writeln(strings.Join(tags, " "))
	r.writeln("%s", story.Metadata.AsAText())
	r.writeln("I want to %s", story.Metadata.IWantTo)
	r.writeln("So that %s", story.Metadata.SoThat)
	r.writeln("")
}

func (r *ConsoleReporter) writeScenario(scenario *rm.Scenario) {
	r.writeln("")
	r.writeln("%-60s %s", "Scenario: "+scenario.Title, scenario.Result)
}

func (r *ConsoleReporter) writeSteps(steps []*rm.Step) {
	for _, step := range steps {
		message := ""

		if step.Error != nil {
			message = step.Error.Message
		}

		r.writeln("%-60s %s %s", step.Title, step.Result, message)
	}
}

func (r *ConsoleReporter) writeHeaders(headers []string) {
	r.writeln("Examples")

	for i := 0; i < len(headers); i++ {
		r.write("|%-10s", headers[i])
	}
	r.writeln("|")
}

func (r *ConsoleReporter) writeExamples(examples []*rm.Example) {
	const totalWidth = 60
	const colWidth = 10

	for _, example := range examples {
		exampleStr := ""
		for i := 0; i < len(example.Values); i++ {
			exampleStr += fmt.Sprintf("|%-*v", colWidth, example.Values[i])
		}

		exampleStr += "|"

		message := ""

		if example.Error != nil {
			message = example.Error.Message
		}

		r.writeln("%-*s %s %s", totalWidth, exampleStr, example.Result, message)
	}
}

func (r *ConsoleReporter) writeExampleTable(table *rm.ExampleTable) {
	if table != nil {
		r.writeHeaders(table.Headers)
		r.writeExamples(table.Examples)
	}
}

func (r *ConsoleReporter) Report(model *rm.ReportModel) {
	file, _ := os.Create(r.outputFileName)

	r.outputFile = file

	r.writeln("")

	for _, story := range model.Stories {
		r.writeStory(story)

		for _, scenario := range story.Scenarios {
			r.writeScenario(scenario)
			r.writeSteps(scenario.Steps)
			r.writeExampleTable(scenario.ExampleTable)
			r.writeln("")
		}

	}

	r.outputFile.Close()

	content, err := os.ReadFile(r.outputFileName)

	if err == nil {
		fmt.Print(string(content))
	}

}
