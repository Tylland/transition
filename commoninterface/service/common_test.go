package service

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"testing"
	"text/template"

	_ "embed"

	"github.com/tylland/transition/commoninterface/repository"
	"github.com/tylland/transition/specs"
)

// minimalXSD is valid XSD content; its structure is not parsed — only the file name matters.
var minimalXSD = []byte(`<?xml version="1.0" encoding="UTF-8"?><xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"/>`)

const (
	validPayload     = `<TrainRunningInformationMessage><MessageHeader/></TrainRunningInformationMessage>`
	wrongRootPayload = `<PathRequestMessage><Data/></PathRequestMessage>`
	malformedPayload = `<NotClosed`
	messageTemplate  = `<%s><MessageHeader>
								<MessageReference>
									<MessageType>%s</MessageType>
									<MessageTypeVersion>1.0</MessageTypeVersion>
									<MessageIdentifier>796cc4ee-cd51-49c2-a8cf-3bd73b29ddfa</MessageIdentifier>
									<MessageDateTime>2026-02-08T14:25:30.910+05:30</MessageDateTime>
								</MessageReference>
							<Sender>3025</Sender>
							<Recipient>0074</Recipient>
							</MessageHeader>
						</%s>`
)

//go:embed template/testmessage.xsd.tmpl
var schemaTmpl string

const (
	defaultCompanyCode = "0000099"
	defaultMsgID       = "msg-001"
)

type SchemaFile struct {
	SchemaFile   string         `json:"schemaFile"`
	MessageTypes []*MessageType `json:"messageTypes"`
}

type MessageType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type SchemaParams struct {
	MessageName string
}

type CIMessage struct {
	companyCode string
	msgID       string
	body        string
}

type serviceTestContext struct {
	t            *testing.T
	schemasDir   string
	repo         *repository.SchemaRepository
	schemaName   string
	messageTypes []*MessageType
}

func NewServiceTestContext(t *testing.T) *serviceTestContext {
	baseDir := t.TempDir()

	schemasDir := filepath.Join(baseDir, "schemas")

	os.MkdirAll(schemasDir, 0x644)

	return &serviceTestContext{t: t, schemasDir: schemasDir}
}

func (c *serviceTestContext) Schema(schemaName string) *specs.StepResult {
	res := specs.NewStepResult("schema %s", schemaName)

	c.schemaName = schemaName

	return res.Success()
}

func (c *serviceTestContext) MessageType(code, name string) *specs.StepResult {
	res := specs.NewStepResult("message type %s(%s)", name, code)

	c.messageTypes = append(c.messageTypes, &MessageType{Code: code, Name: name})

	return res.Success()
}

func (c *serviceTestContext) RepositoryWithSchemaFor(category, version, name string, messageTypes []*MessageType) *specs.StepResult {
	res := specs.NewStepResult("a schema repository that contains schema(s) for %s %s v%s", category, name, version)

	log.Printf("Start creating repository in directory: %s", c.schemasDir)

	schemaDir := filepath.Join(c.schemasDir, category, version)

	if _, err := os.Stat(schemaDir); os.IsNotExist(err) {
		os.MkdirAll(schemaDir, 0x644)
	}

	schemaFileName := name + ".xsd"

	xsd := specs.Must(c.messageSchema(name))

	os.WriteFile(filepath.Join(schemaDir, schemaFileName), []byte(xsd), 0644)

	schemaFile := SchemaFile{SchemaFile: schemaFileName, MessageTypes: messageTypes}

	data := specs.Must(json.MarshalIndent(schemaFile, "", "  "))

	os.WriteFile(filepath.Join(schemaDir, "schema.json"), data, 0644)

	c.repo = specs.Must(repository.NewSchemaRepository(c.schemasDir))

	return res.Success()
}

func (c *serviceTestContext) RepositoryExists() *specs.StepResult {
	return c.RepositoryWithSchemaFor("test", "1.0", c.schemaName, c.messageTypes)
}

func (c *serviceTestContext) RepositoryLacksSchemaFor(messageType string) *specs.StepResult {
	c.repo = specs.Must(repository.NewSchemaRepository(c.schemasDir))

	return specs.NewStepResult("the schema repository does not contain a schema for %q", messageType).Success()
}

func (c *serviceTestContext) messageSchema(name string) (string, error) {
	params := SchemaParams{MessageName: name}

	tmpl, err := template.New("").Parse(schemaTmpl)

	if err != nil {
		return "", err
	}

	var res bytes.Buffer

	err = tmpl.Execute(&res, params)

	if err != nil {
		return "", err
	}

	return res.String(), nil
}
