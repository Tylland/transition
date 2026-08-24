package repository

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lestrrat-go/helium"
	"github.com/lestrrat-go/helium/xsd"
	"github.com/tylland/transition/commoninterface/message/telematics"
)

type ErrSchemaNotRegistred struct {
	msgRef *telematics.MessageReference
}

func (e ErrSchemaNotRegistred) Error() string {
	return fmt.Sprintf("no schema registered for message type %q version %s", e.msgRef.MessageType, e.msgRef.MessageTypeVersion)
}

// SchemaRepository holds TAF/TAP TSI XSD schemas indexed by MessageType and MessageTypeVersion (found in MessageHeader.MessageReference).
type SchemaRepository struct {
	schemas map[string]map[string]*SchemaEntry // type → version → raw XSD bytes
}

type SchemaEntry struct {
	Validator *xsd.Validator
	Schema    *xsd.Schema
}

// NewSchemaRepository loads all .xsd files from a type/version/schema.xsd structure at initialisation time.
func NewSchemaRepository(dir string) (*SchemaRepository, error) {
	categoryEntries, err := os.ReadDir(dir)

	if err != nil {
		return nil, fmt.Errorf("schema repository: read directory %q: %w", dir, err)
	}

	repo := &SchemaRepository{schemas: make(map[string]map[string]*SchemaEntry)}

	for _, categoryEntry := range categoryEntries {
		if !categoryEntry.IsDir() {
			return nil, fmt.Errorf("schema repository: %s is not a directory", categoryEntry.Name())
		}

		category := categoryEntry.Name()
		categoryDir := filepath.Join(dir, category)

		versionEntries, err := os.ReadDir(categoryDir)
		if err != nil {
			return nil, fmt.Errorf("schema repository: read directory %q: %w", categoryDir, err)
		}

		for _, versionEntry := range versionEntries {
			if !versionEntry.IsDir() {
				return nil, fmt.Errorf("schema repository: %s is not a directory", versionEntry.Name())
			}

			msgTypeVer := versionEntry.Name()

			versionDir := filepath.Join(categoryDir, msgTypeVer)

			schemaInfo, err := ReadSchemaFile(filepath.Join(versionDir, "schema.json"))

			if err != nil {
				return nil, fmt.Errorf("schema repository: read schema file %q: %w", filepath.Join(versionDir, "schema.json"), err)
			}

			ctx := context.Background()

			root, err := os.OpenRoot(versionDir)

			if err != nil {
				return nil, fmt.Errorf("open directory root %q: %w\n", versionDir, err)
			}

			defer root.Close()

			src, err := root.ReadFile(schemaInfo.SchemaFile)
			if err != nil {
				return nil, err
			}

			doc, err := helium.NewParser().Parse(ctx, src)
			if err != nil {
				return nil, err
			}

			schema, err := xsd.NewCompiler().
				FS(root.FS()). // resolves xs:include / xs:import
				Compile(ctx, doc)

			if err != nil {
				return nil, fmt.Errorf("failed to compile schema: %w\n", err)
			}

			validator := xsd.NewValidator(schema)

			for _, info := range schemaInfo.MessageTypes {
				if _, ok := repo.schemas[info.Code]; !ok {
					repo.schemas[info.Code] = make(map[string]*SchemaEntry)
				}

				repo.schemas[info.Code][msgTypeVer] = &SchemaEntry{Validator: &validator, Schema: schema}
			}
		}
	}

	return repo, nil
}

// HasSchema reports whether a schema is registered for messageType.
func (r *SchemaRepository) HasSchema(msgRef *telematics.MessageReference) bool {
	_, exists := r.SchemaEntry(msgRef)

	return exists
}

func (r *SchemaRepository) SchemaEntry(msgRef *telematics.MessageReference) (*SchemaEntry, bool) {
	if versions, ok := r.schemas[msgRef.MessageType]; ok {
		if entry, ok := versions[msgRef.MessageTypeVersion]; ok {
			return entry, true
		}
	}

	return nil, false
}

// Len returns the number of registered schemas.
func (r *SchemaRepository) Len() int {
	return len(r.schemas)
}

func (r *SchemaRepository) Validate(msgRef *telematics.MessageReference, payload string) error {
	if msgRef == nil {
		hdr, err := telematics.ParseHeader(payload)

		if err != nil {
			return fmt.Errorf("validate: failed to parse schema")
		}

		if hdr == nil || hdr.MessageReference == nil {
			return fmt.Errorf("validate: message reference not found in payload header")
		}

		msgRef = hdr.MessageReference
	}

	schemaEntry, hasSchema := r.SchemaEntry(msgRef)

	if !hasSchema {
		return ErrSchemaNotRegistred{msgRef: msgRef}
	}

	parser := helium.NewParser()
	doc, err := parser.Parse(context.Background(), []byte(payload))

	if err != nil {
		return fmt.Errorf("validate: failed to parse document (payload)")
	}

	if err := schemaEntry.Validator.Validate(context.Background(), doc); err != nil {
		return fmt.Errorf("xml does not conform to schema type %q version %q: %w", msgRef.MessageType, msgRef.MessageTypeVersion, err)
	}

	return nil
}
