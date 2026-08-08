package service

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tylland/transition/commoninterface/repository"
	"github.com/tylland/transition/specs"
)

// minimalXSD is valid XSD content; its structure is not parsed — only the file name matters.
var minimalXSD = []byte(`<?xml version="1.0" encoding="UTF-8"?><xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"/>`)

const (
	validPayload     = `<TrainRunningInformationMessage><MessageHeader/></TrainRunningInformationMessage>`
	wrongRootPayload = `<PathRequestMessage><Data/></PathRequestMessage>`
	malformedPayload = `<NotClosed`
)

type CIMessage struct {
	companyCode string
	msgID       string
	body        string
}

type serviceTestContext struct {
	t          *testing.T
	schemasDir string
	repo       *repository.SchemaRepository
}

func NewServiceTestContext(t *testing.T) *serviceTestContext {
	baseDir := t.TempDir()

	schemasDir := filepath.Join(baseDir, "schemas")

	os.MkdirAll(schemasDir, 0x644)

	return &serviceTestContext{t: t, schemasDir: schemasDir}
}

func (c *serviceTestContext) RepositoryWithSchemaFor(types ...string) *specs.StepResult {
	res := specs.NewStepResult("a schema repository that contains schema(s) for %q", strings.Join(types, " and "))

	for _, msgType := range types {
		os.WriteFile(filepath.Join(c.schemasDir, msgType+"xsd"), minimalXSD, 0644)
	}

	c.repo = specs.Must(repository.NewSchemaRepository(c.schemasDir))

	return res.Success()
}

func (c *serviceTestContext) RepositoryLacksSchemaFor(messageType string) *specs.StepResult {
	c.repo = specs.Must(repository.NewSchemaRepository(c.schemasDir))

	return specs.NewStepResult("the schema repository does not contain a schema for %q", messageType).Success()
}
