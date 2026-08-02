package specs

import (
	"fmt"
	"reflect"
	"strings"
)

type StepPrefix string

const (
	Given StepPrefix = "Given"
	When  StepPrefix = "When"
	Then  StepPrefix = "Then"
	And   StepPrefix = "  And"
)

type StepAction func(args ...interface{}) StepTitle

type StepTitle interface {
	title() string
}

type StepTitleFormat struct {
	Template string
}

func (st *StepTitleFormat) title() string {
	return st.Template
}

type Step struct {
	action interface{}
	args   []interface{}
	prefix StepPrefix
}

func NewStep(prefix StepPrefix, action interface{}, args ...interface{}) Step {
	//	stepArgs := ...args

	return Step{action: action, args: args, prefix: prefix}
}

func SplitCamelCase(input string) string {
	return "" // regex.Replace(input, "([A-Z])", " $1")Trim().ToLower();
}

func (s *Step) Execute(sc *ScenarioContext) (string, error) {
	var stepTitle string = string(s.prefix)
	var stepError error = nil

	v := reflect.ValueOf(s.action)
	if v.Kind() != reflect.Func {
		panic("not a function")
	}

	vargs := make([]reflect.Value, len(s.args))
	displayValues := make([]string, len(s.args))

	for i, arg := range s.args {
		value := reflect.ValueOf(arg)

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
			displayValues[i] = "<" + sc.HeaderTitle(arg) + ">"
		} else {
			displayValues[i] = fmt.Sprint(value.Interface())
		}

		vargs[i] = value
	}

	vrets := v.Call(vargs)

	if len(vrets) >= 1 {
		stepTitle += " " + vrets[0].String()
		parameters := s.Parameters(stepTitle)

		for pos, parameter := range parameters {
			stepTitle = strings.ReplaceAll(stepTitle, parameter, displayValues[pos])
		}
	}

	if len(vrets) >= 2 {
		err := vrets[1].Interface()

		if err != nil {
			stepError = err.(error)
		}
	}

	return stepTitle, stepError
}

func (s *Step) Parameters(input string) []string {
	var parameters []string
	var parameter string

	open := false

	for _, rune := range input {
		if rune == '{' {
			open = true
		}

		if open {
			parameter += string(rune)
		}

		if rune == '}' {
			open = false
			parameters = append(parameters, parameter)
		}
	}

	return parameters
}
