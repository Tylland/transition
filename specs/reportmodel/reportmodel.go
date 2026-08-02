package reportmodel

import (
	"fmt"
	"strings"
	"time"
)

type ReportModel struct {
	Stories []*Story `json:"stories"`
}

type Story struct {
	Namespace string        `json:"namespace"`
	Result    Result        `json:"result"`
	Scenarios []*Scenario   `json:"scenarios"`
	Metadata  StoryMetadata `json:"metadata"`
	Tags      []Tag         `json:"tags"`
}

type StoryMetadata struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	TitlePrefix string `json:"title_prefix"`
	AsA         string `json:"as_a"`
	IWantTo     string `json:"i_want_to"`
	SoThat      string `json:"so_that"`
	ImageUri    string `json:"image_uri"`
	StoryUri    string `json:"story_uri"`
}

func (md *StoryMetadata) isVowel(str string) bool {
	if len(str) == 0 {
		return false
	}

	sl := []rune(strings.ToLower(str))[0]

	return sl == 'a' || sl == 'e' || sl == 'i' || sl == 'o' || sl == 'u' || sl == 'å' || sl == 'ä' || sl == 'ö'
}

func (md *StoryMetadata) AsAText() string {
	text := "As a"

	if md.isVowel(md.AsA) {
		text += "n"
	}

	return text + " " + md.AsA
}

type Scenario struct {
	Id           string        `json:"id"`
	Title        string        `json:"title"`
	Tags         []string      `json:"tags"`
	ExampleTable *ExampleTable `json:"exampleTable"`
	Duration     time.Duration `json:"duration"`
	Steps        []*Step       `json:"steps"`
	Result       Result        `json:"result"`
}

type Step struct {
	Id             string         `json:"id"`
	Index          int            `json:"index"`
	Asserts        bool           `json:"asserts"`
	ShouldReport   bool           `json:"should_report"`
	Title          string         `json:"title"`
	ExecutionOrder ExecutionOrder `json:"execution_order"`
	Result         Result         `json:"result"`
	Error          *Error         `json:"error"`
	Duration       time.Duration  `json:"duration"`
}

type ExampleTable struct {
	Headers  []string   `json:"headers"`
	Examples []*Example `json:"examples"`
}

type Example struct {
	Values   []string      `json:"values"`
	Result   Result        `json:"result"`
	Error    *Error        `json:"error"`
	Duration time.Duration `json:"duration"`
}

func (e Example) String() string {
	var values []string
	for _, v := range e.Values {
		values = append(values, v)
	}

	return fmt.Sprintf("%s", values)
}

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Result string

const (
	Pass Result = "Pass"
	Fail Result = "Fail"
	Skip Result = "Skip"
)

type ExecutionOrder string

const (
	Before ExecutionOrder = "Before"
	After  ExecutionOrder = "After"
)

type Error struct {
	Message string `json:"message"`
	Stack   string `json:"stack"`
}
