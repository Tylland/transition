package specs

import (
	"strings"
	"testing"
	"unicode"
)

type Scenario struct {
	context *ScenarioContext
}

func camelAppend(str string) string {
	w := []rune(str)
	for i := len(w) - 1; i > 1; i-- {
		if unicode.IsUpper(w[i]) {
			w = append(w[:i], append([]rune{' '}, w[i:]...)...)
		}
	}
	return string(w)
}

func NewScenario(story *Story, t *testing.T) *Scenario {
	return &Scenario{context: &ScenarioContext{t: t, story: story, title: camelAppend(strings.TrimPrefix(t.Name(), "Test"))}}
}

func (s *Scenario) WithDebug(value bool) *Scenario {
	s.context.debug = value

	return s
}

func (s *Scenario) Given(action interface{}, args ...interface{}) *GivenStepBlock {
	return NewGivenStepBlock(s.context, NewStep(Given, action, args...))
}

type GivenStepBlock struct {
	context *ScenarioContext
}

func NewGivenStepBlock(context *ScenarioContext, step Step) *GivenStepBlock {
	context.Add(step)

	return &GivenStepBlock{context: context}
}

func (sb *GivenStepBlock) And(action interface{}, args ...interface{}) *GivenStepBlock {
	sb.context.Add(NewStep(And, action, args...))

	return sb
}

func (sb *GivenStepBlock) When(action interface{}, args ...interface{}) *WhenStepBlock {

	return NewWhenStepBlock(sb.context, NewStep(When, action, args...))
}

type WhenStepBlock struct {
	context *ScenarioContext
}

func NewWhenStepBlock(context *ScenarioContext, step Step) *WhenStepBlock {
	context.Add(step)

	return &WhenStepBlock{context: context}
}

func (sb *WhenStepBlock) Then(action interface{}, args ...interface{}) *ThenStepBlock {
	return NewThenStepBlock(sb.context, NewStep(Then, action, args...))
}

type ThenStepBlock struct {
	context *ScenarioContext
}

func NewThenStepBlock(context *ScenarioContext, step Step) *ThenStepBlock {
	context.Add(step)

	return &ThenStepBlock{context: context}
}

func (sb *ThenStepBlock) And(action interface{}, args ...interface{}) *ThenStepBlock {
	sb.context.Add(NewStep(And, action, args...))

	return sb
}

func (sb *ThenStepBlock) WithVariables(variables ...interface{}) *ExampleBlock {

	table := ExampleTable{headers: Headers(variables)}

	sb.context.examples = table

	return NewExampleBlock(sb.context)
}

func (sb *ThenStepBlock) Execute() {
	sb.context.Execute(sb.context.t, false)
}
