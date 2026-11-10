package dup_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tylland/transition/cdas/internal/infrastructure/dup"
	"github.com/tylland/transition/specs"
)

var requestStory = specs.NewStory(specs.UserStory{
	AsA:     "DUP client",
	IWantTo: "",
	SoThat:  "",
})

const simpleFilter = `<REQUEST>
  <LOGIN authenticationkey="SomeAuthenticationKey" />
  <QUERY objecttype="SomeObjectType" schemaversion="SomeObjectVersion" limit="10">
    <FILTER>
      <EQ name="SomeDataField" value="2" />
    </FILTER>
  </QUERY>
</REQUEST>`

const logicFilter = `<REQUEST>
  <LOGIN authenticationkey="SomeAuthenticationKey" />
  <QUERY objecttype="SomeObjectType" schemaversion="SomeObjectVersion" limit="10">
    <FILTER>
      <OR>
        <AND>
          <EQ name="User.Name" value="kalle anka" />
          <IN name="User.Category" value="a, b, c" />
        </AND>
        <GT name="User.Number" value="133" />
      </OR>
    </FILTER>
  </QUERY>
</REQUEST>`

var (
	request *dup.Request
	query   *dup.Query
	filter  dup.Expression
	xml     []byte
)

func Request(authKey string) *specs.StepResult {
	sr := specs.NewStepResult("request with key %q", authKey)

	request = dup.NewRequest(authKey)

	return sr.Success()
}

func Query(objType string, ver string) *specs.StepResult {
	sr := specs.NewStepResult("query %q", objType)

	query = dup.NewQuery(objType).SchemaVersion(ver)
	request = request.Query(query)

	return sr.Success()
}

func Limit(limit int) *specs.StepResult {
	sr := specs.NewStepResult("with limit %d", limit)

	query = query.Limit(limit)

	return sr.Success()
}

func Filter(expr dup.Expression) *specs.StepResult {
	sr := specs.NewStepResult("filter: %s", expression(expr, 0))

	query = query.Filter(expr)

	return sr.Success()
}

func Serializing() *specs.StepResult {
	sr := specs.NewStepResult("serializing to xml")

	xml, _ = request.XML()

	return sr.Success()
}

func XMLIsGenerated(expected string) *specs.StepResult {
	sr := specs.NewStepResult("generated xml is:\n%s", expected)

	if string(xml) != expected {
		return sr.WithErrorf("Expected:\n%q\n, got:\n%q\n", expected, string(xml))
	}

	return sr.Success()
}

func expression(expr dup.Expression, depth int) string {
	depth++
	res := ""
	switch e := expr.(type) {
	case dup.Comparison:
		res += strings.Repeat("  ", depth)
		res += fmt.Sprintf("%s(%q, %q)\n", e.Operator, e.Name, e.Value)
	case dup.Logical:
		res += strings.Repeat("  ", depth)
		res += fmt.Sprintf("%s(\n", e.Operator)

		for _, child := range e.Children {
			res += expression(child, depth)
		}
		res += ")"
	}

	return res
}

func TestSimpleFilter(t *testing.T) {

	requestStory.Scenario(t).
		Given(Request("SomeAuthenticationKey")).
		And(Query("SomeObjectType", "SomeObjectVersion")).
		And(Limit(10)).
		And(Filter(dup.EQ("SomeDataField", "2"))).
		When(Serializing()).
		Then(XMLIsGenerated(simpleFilter)).
		Execute()

	// filter := dup.EQ("SomeDataField", "2")

	// query := dup.NewQuery("SomeObjectType").
	// 	Limit(10).
	// 	SchemaVersion("SomeObjectVersion").
	// 	Filter(filter)

	// xml, _ := dup.NewRequest("SomeAuthenticationKey").
	// 	Query(query).XML()

	// if string(xml) != simpleFilter {
	// 	t.Logf("Expected:\n%q\n, got:\n%q\n", simpleFilter, string(xml))
	// 	t.Fatal()
	// }
}

func TestLogicFilter(t *testing.T) {
	filter := dup.OR(
		dup.AND(dup.EQ("User.Name", "kalle anka"), dup.IN("User.Category", "a, b, c")),
		dup.GT("User.Number", "133"),
	)

	requestStory.Scenario(t).
		Given(Request("SomeAuthenticationKey")).
		And(Query("SomeObjectType", "SomeObjectVersion")).
		And(Limit(10)).
		And(Filter(filter)).
		When(Serializing()).
		Then(XMLIsGenerated(logicFilter)).
		Execute()

	// query := dup.NewQuery("SomeObjectType").
	// 	Limit(10).
	// 	SchemaVersion("SomeObjectVersion").
	// 	Filter(filter)

	// xml, _ := dup.NewRequest("SomeAuthenticationKey").
	// 	Query(query).XML()

	// if string(xml) != logicFilter {
	// 	t.Logf("Expected:\n%q\n, got:\n%q\n", logicFilter, string(xml))
	// 	t.Fatal()
	// }
}
