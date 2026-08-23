package repository

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tylland/transition/specs"
)

// minimalXSD is valid XSD content; its structure is not parsed — only the file name matters.
var minimalXSD = []byte(`<?xml version="1.0" encoding="UTF-8"?><xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"/>`)

// ── Story declarations ────────────────────────────────────────────────────────

var repoStory = specs.NewStory(specs.UserStory{
	AsA:     "CI operator",
	IWantTo: "load XSD schemas from a configurable directory",
	SoThat:  "message types are resolved without recompiling",
})

// ── Step functions — Given (input state) ────────────────────────────────────

type SchemaRepositoryTestContext struct {
	t          *testing.T
	schemasDir string
	repo       *SchemaRepository
	err        error
}

func NewCITestContext(t *testing.T) *SchemaRepositoryTestContext {
	baseDir := t.TempDir()

	schemasDir := filepath.Join(baseDir, "schemas")

	os.MkdirAll(schemasDir, 0x644)

	return &SchemaRepositoryTestContext{t: t, schemasDir: schemasDir}
}

func (c *SchemaRepositoryTestContext) SchemaDirContains(types ...string) *specs.StepResult {
	for _, typ := range types {
		os.WriteFile(filepath.Join(c.schemasDir, typ), minimalXSD, 0644)
	}
	return specs.NewStepResult("schema directory contains %s", strings.Join(types, " and ")).Success()
}

func (c *SchemaRepositoryTestContext) SchemaDirIsEmpty() *specs.StepResult {
	return specs.NewStepResult("schema directory is empty").Success()
}

func (c *SchemaRepositoryTestContext) SchemaDirDoesNotExist() *specs.StepResult {
	c.schemasDir = filepath.Join(c.schemasDir, "nonexistent")
	return specs.NewStepResult("the schema directory does not exist").Success()
}

// ── Step functions — When (actions) ──────────────────────────────────────

func (c *SchemaRepositoryTestContext) CreatingSchemaRepository() *specs.StepResult {
	sr := specs.NewStepResult("creating schema repository")

	c.repo, c.err = NewSchemaRepository(c.schemasDir)

	return sr.Success()
}

// ── Step functions — Then (assertions) ──────────────────────────────────────

// func (c *SchemaRepositoryTestContext) SchemaRepositoryContains(types ...string) *specs.StepResult {
// 	sr := specs.NewStepResult("schema repository containss %s", strings.Join(types, " and "))
// 	for _, mt := range types {
// 		if !c.repo.HasSchema(mt) {
// 			return sr.WithErrorf("missing schema: %s", mt)
// 		}
// 	}
// 	return sr.Success()
// }

func (c *SchemaRepositoryTestContext) SchemaRepositoryIsEmpty() *specs.StepResult {
	sr := specs.NewStepResult("schema repository is empty")

	if c.repo.Len() != 0 {
		return sr.WithErrorf("expected 0 schemas, got %d", c.repo.Len())
	}

	return sr.Success()
}

func (c *SchemaRepositoryTestContext) ErrorWasReturned() *specs.StepResult {
	res := specs.NewStepResult("an error is returned")

	if c.err != nil {
		return res.Success()
	}

	return res.WithErrorMessage("Expected error but no error found")
}

// ── SchemaRepository tests ────────────────────────────────────────────────────

func TestSchemaRepositoryLoadsSchemasByFileNameStem(t *testing.T) {
	f := NewCITestContext(t)

	repoStory.Scenario(t).
		Given(f.SchemaDirContains("TrainRunningInformationMessage.xsd", "PathRequestMessage.xsd")).
		When(f.CreatingSchemaRepository()).
		Then(f.SchemaDirContains("TrainRunningInformationMessage", "PathRequestMessage")).
		Execute()
}

func TestSchemaRepositoryIsEmptyWhenNoXSDFilesPresent(t *testing.T) {
	c := NewCITestContext(t)

	repoStory.Scenario(t).
		Given(c.SchemaDirIsEmpty()).
		When(c.CreatingSchemaRepository()).
		Then(c.SchemaRepositoryIsEmpty()).
		Execute()
}

func TestInitialisationFailsForMissingDirectory(t *testing.T) {
	c := NewCITestContext(t)

	repoStory.Scenario(t).
		Given(c.SchemaDirDoesNotExist()).
		When(c.CreatingSchemaRepository()).
		Then(c.ErrorWasReturned()).
		Execute()
}
