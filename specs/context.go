package specs

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	rm "github.com/tylland/transition/specs/reportmodel"
)

type ScenarioContext struct {
	t        *testing.T
	story    *Story
	title    string
	steps    []Step
	examples ExampleTable
	debug    bool
}

func (sc *ScenarioContext) Add(step Step) {
	sc.steps = append(sc.steps, step)
}

func (sc *ScenarioContext) HeaderTitle(arg interface{}) string {
	examples := sc.examples

	for i := 0; i < len(examples.headers); i++ {
		pvariable := reflect.ValueOf(examples.headers[i].variable).Pointer()
		parg := reflect.ValueOf(arg).Pointer()

		if sc.debug {
			fmt.Printf("Variable for %s found at %v  - %v\n", examples.headers[i].Title, pvariable, parg)
		}

		if pvariable == parg {
			return examples.headers[i].Title
		}
	}

	return "Not found"
}

func (sc *ScenarioContext) Execute(t *testing.T, silent bool) (TestResult, error) {
	scenarioResult := Succeeded
	var messages []string

	steps := []*StepResult{}

	if !silent {
		if !sc.story.reported {
			S.WriteStory(sc.story)
			sc.story.reported = true
		}

		S.WriteScenario(sc.story, sc.title)
		t.Log(sc.title)
	}

	for i := 0; i < len(sc.steps); i++ {
		stepResult := rm.Skip

		step := sc.steps[i]

		title, err := step.Execute(sc)

		if err != nil {
			stepResult = rm.Fail
			scenarioResult = Failed

			message := "Failed - " + err.Error()
			messages = append(messages, message)

			if !silent {
				steps = append(steps, NewStepResult(step.prefix, title, Failed, message))
			}

			t.Error(title + " " + message)
		} else {
			stepResult = rm.Pass

			if !silent {
				steps = append(steps, NewStepResult(step.prefix, title, Succeeded, ""))
			}
		}

		if !silent {
			S.WriteStep(step.prefix, title, stepResult, err)
			t.Log(title)
		}
	}

	if len(messages) > 0 {
		return scenarioResult, errors.New(strings.Join(messages, ","))
	}

	return scenarioResult, nil
}
