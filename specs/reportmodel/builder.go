package reportmodel

type Builder[T any] interface {
	Build() T
}

type Builders[TResult any] struct {
	lookup   map[string]Builder[TResult]
	builders []Builder[TResult]
}

func NewBuilders[TResult any]() Builders[TResult] {
	return Builders[TResult]{lookup: make(map[string]Builder[TResult])}
}

func (b *Builders[TResult]) add(id string, builder Builder[TResult]) {
	b.lookup[id] = builder
	b.builders = append(b.builders, builder)
}

func (b *Builders[TResult]) With(id string, create func(id string) Builder[TResult]) Builder[TResult] {
	if b.lookup[id] == nil {
		b.add(id, create(id))
	}

	return b.lookup[id]
}

func (b *Builders[TResult]) Build() []TResult {
	var result []TResult

	for _, builder := range b.builders {
		result = append(result, builder.Build())
	}

	return result
}

type ReportModelBuilder struct {
	stories Builders[*Story]
}

func NewReportModelBuilder() *ReportModelBuilder {
	return &ReportModelBuilder{stories: NewBuilders[*Story]()}
}

func (b *ReportModelBuilder) WithStory(id string) *StoryBuilder {
	return b.stories.With(id, b.createStoryBuilder).(*StoryBuilder)
}

func (b *ReportModelBuilder) createStoryBuilder(id string) Builder[*Story] {
	return &StoryBuilder{id: id, metadata: StoryMetadata{}, scenarios: NewBuilders[*Scenario]()}
}

func (b *ReportModelBuilder) Build() *ReportModel {
	return &ReportModel{Stories: b.stories.Build()}
}

type StoryBuilder struct {
	id        string
	metadata  StoryMetadata
	scenarios Builders[*Scenario]
	tags      map[string]string
}

func (b *StoryBuilder) WithNarrative(asA, iWantTo, soThat string) *StoryBuilder {
	b.metadata.AsA = asA
	b.metadata.IWantTo = iWantTo
	b.metadata.SoThat = soThat

	return b
}

func (b *StoryBuilder) WithTags(tags map[string]string) *StoryBuilder {
	b.tags = tags

	return b
}

func (b *StoryBuilder) WithScenario(id string) *ScenarioBuilder {
	return b.scenarios.With(id, newScenarioBuilder).(*ScenarioBuilder)
}

func (b *StoryBuilder) Build() *Story {
	result := Pass

	scenarios := b.scenarios.Build()

	for _, scenario := range scenarios {
		if scenario.Result == Fail {
			result = Fail
			break
		}
	}

	return &Story{Namespace: "", Scenarios: scenarios, Result: result, Metadata: b.metadata, Tags: NewTags(b.tags)}
}

type ScenarioBuilder struct {
	id           string
	title        string
	steps        Builders[*Step]
	exampleTable Builder[*ExampleTable]
}

func newScenarioBuilder(id string) Builder[*Scenario] {
	return &ScenarioBuilder{id: id, steps: NewBuilders[*Step]()}
}

func (b *ScenarioBuilder) WithTitle(title string) *ScenarioBuilder {
	b.title = title

	return b
}

func (b *ScenarioBuilder) WithStep(id string) *StepBuilder {
	builder := b.steps.With(id, newStepBuilder).(*StepBuilder)

	return builder
}

func (b *ScenarioBuilder) WithExample() *ExampleBuilder {

	if b.exampleTable == nil {
		b.exampleTable = &ExampleBuilder{}
	}

	return b.exampleTable.(*ExampleBuilder)
}

func (b *ScenarioBuilder) Build() *Scenario {
	result := Pass

	steps := b.steps.Build()

	for index, step := range steps {
		step.Index = index
		if step.Result == Fail {
			result = Fail
			break
		}
	}

	var exampleTable *ExampleTable = nil

	if b.exampleTable != nil {
		exampleTable = b.exampleTable.Build()
	}

	return &Scenario{Id: b.id, Title: b.title, Steps: steps, ExampleTable: exampleTable, Result: result}
}

type StepBuilder struct {
	id      string
	title   string
	asserts bool
	result  Result
	error   *Error
}

func newStepBuilder(id string) Builder[*Step] {
	return &StepBuilder{id: id}
}

func (b *StepBuilder) WithTitle(title string) *StepBuilder {
	b.title = title
	return b
}

func (b *StepBuilder) WithAsserts(asserts bool) *StepBuilder {
	b.asserts = asserts
	return b
}

func (b *StepBuilder) WithResult(ressult Result) *StepBuilder {
	b.result = ressult
	return b
}

func (b *StepBuilder) WithError(error *Error) *StepBuilder {
	b.error = error
	return b
}

func (b *StepBuilder) Build() *Step {
	step := &Step{Id: b.id, Title: b.title, Asserts: b.asserts, Result: b.result, Error: b.error}

	return step
}

type ExampleBuilder struct {
	id       string
	Headers  []string
	Examples []*Example
}

// func newExampleBuilder(id string) Builder[*Example] {
// 	return &ExampleBuilder{}
// }

func (b *ExampleBuilder) WithHeader(headers []string) *ExampleBuilder {
	b.Headers = headers

	return b
}

func (b *ExampleBuilder) AddExample(values []string, result Result, err error) *ExampleBuilder {
	b.Examples = append(b.Examples, &Example{Values: values, Result: result, Error: NewError(err)})

	return b
}

func (b *ExampleBuilder) Build() *ExampleTable {
	return &ExampleTable{Headers: b.Headers, Examples: b.Examples}
}

func NewError(err error) *Error {
	if err != nil {
		return &Error{Message: err.Error()}
	}

	return nil
}

func NewTags(t map[string]string) []Tag {
	var tags []Tag

	for k, v := range t {
		tags = append(tags, Tag{Name: k, Value: v})
	}

	return tags
}
