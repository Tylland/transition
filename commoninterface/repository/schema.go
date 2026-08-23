package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
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
	schemas map[string]map[string][]byte // messageType → raw XSD bytes
}

// NewSchemaRepository loads all .xsd files from a type/version/schema.xsd structure at initialisation time.
func NewSchemaRepository(dir string) (*SchemaRepository, error) {
	typeEntries, err := os.ReadDir(dir)

	if err != nil {
		return nil, fmt.Errorf("schema repository: read directory %q: %w", dir, err)
	}

	repo := &SchemaRepository{schemas: make(map[string]map[string][]byte)}

	for _, typeEntry := range typeEntries {
		if !typeEntry.IsDir() {
			return nil, fmt.Errorf("schema repository: %s is not a directory", typeEntry.Name())
		}

		msgType := typeEntry.Name()

		versions := make(map[string][]byte)

		typeDir := filepath.Join(dir, msgType)

		versionEntries, err := os.ReadDir(typeDir)
		if err != nil {
			return nil, fmt.Errorf("schema repository: read directory %q: %w", typeDir, err)
		}

		for _, versionEntry := range versionEntries {
			if !versionEntry.IsDir() {
				return nil, fmt.Errorf("schema repository: %s is not a directory", versionEntry.Name())
			}

			msgTypeVer := versionEntry.Name()

			versionDir := filepath.Join(typeDir, msgTypeVer)

			if err != nil {
				return nil, fmt.Errorf("schema repository: read directory %q: %w", versionDir, err)
			}

			schemas, err := filepath.Glob(versionDir + "/*.xsd")

			if err != nil {
				return nil, fmt.Errorf("schema repository: finding schemas %q: %w", versionDir, err)
			}

			if len(schemas) == 0 {
				return nil, fmt.Errorf("schema repository: no schema found for %q version %q in %q", msgType, msgTypeVer, versionDir)
			}

			if len(schemas) > 1 {
				return nil, fmt.Errorf("schema repository: %s schemas found for %q version %q in %q",
					len(schemas), msgType, msgTypeVer, versionDir)
			}

			data, err := os.ReadFile(schemas[0])

			if err != nil {
				return nil, fmt.Errorf("schema repository: read schema for %q version %q in %q: %w",
					msgType, msgTypeVer, schemas[0], err)
			}

			versions[msgTypeVer] = data
		}

		repo.schemas[msgType] = versions

	}

	return repo, nil
}

// HasSchema reports whether a schema is registered for messageType.
func (r *SchemaRepository) HasSchema(msgRef telematics.MessageReference) bool {
	if versions, ok := r.schemas[msgRef.MessageType]; ok {
		if _, ok := versions[msgRef.MessageTypeVersion]; ok {
			return true
		}
	}

	return false
}

func (r *SchemaRepository) Schema(msgRef *telematics.MessageReference) ([]byte, bool) {
	if versions, ok := r.schemas[msgRef.MessageType]; ok {
		if data, ok := versions[msgRef.MessageTypeVersion]; ok {
			return data, true
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

		msgRef = hdr.MessageReference
	}

	data, hasSchema := r.Schema(msgRef)

	if !hasSchema {
		return ErrSchemaNotRegistred{msgRef: msgRef}
	}

	schema, err := xsd.Parse(data)

	if err != nil {
		return fmt.Errorf("validate: failed to parse schema")
	}

	doc, err := libxml2.Parse([]byte(payload))

	if err != nil {
		return fmt.Errorf("validate: failed to parse document (payload)")
	}

	if err := schema.Validate(doc); err != nil {
		return fmt.Errorf("xml does not conform to schema type %q version %q: %w", msgRef.MessageType, msgRef.MessageTypeVersion, err)
	}

	return nil
}
