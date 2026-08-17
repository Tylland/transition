package repository

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ErrSchemaNotRegistred struct {
	msgType string
}

func (e ErrSchemaNotRegistred) Error() string {
	return fmt.Sprintf("no schema registered for message type %q", e.msgType)
}

type SchemaRecord struct {
	messageTypeCode int
	messageType     string
	schema          string
}

// SchemaRepository holds TAF/TAP TSI XSD schemas indexed by MessageType name.
// MessageType equals the root element name of the payload XML
// (ERA TD 105 taf_cat_complete.xsd: "This message Type is same as the Common Schema Root Element name").
type SchemaRepository struct {
	codes   map[int]string    // code → messageType
	schemas map[string][]byte // messageType → raw XSD bytes
}

// NewSchemaRepository loads all .xsd files from dir at initialisation time.
// Each file name stem (without extension) becomes the MessageType key.
func NewSchemaRepository(dir string) (*SchemaRepository, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("schema repository: read directory %q: %w", dir, err)
	}

	repo := &SchemaRepository{schemas: make(map[string][]byte)}
	for _, e := range entries {
		if e.IsDir() || !strings.EqualFold(filepath.Ext(e.Name()), ".xsd") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("schema repository: read %s: %w", e.Name(), err)
		}
		msgType := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
		repo.schemas[msgType] = data
	}

	return repo, nil
}

// HasSchema reports whether a schema is registered for messageType.
func (r *SchemaRepository) HasSchema(messageType string) bool {
	_, ok := r.schemas[messageType]
	return ok
}

// Len returns the number of registered schemas.
func (r *SchemaRepository) Len() int {
	return len(r.schemas)
}

// Validate checks that payload is well-formed XML and that its root element name
// matches a MessageType registered in the repository.
func (r *SchemaRepository) Validate(payload string) error {
	msgType, err := xmlRootElement(payload)
	if err != nil {
		return fmt.Errorf("schema validation: malformed XML: %w", err)
	}
	if !r.HasSchema(msgType) {
		return ErrSchemaNotRegistred{msgType: msgType}
	}
	return nil
}

// xmlRootElement returns the local name of the first XML start element in s.
func xmlRootElement(s string) (string, error) {
	dec := xml.NewDecoder(strings.NewReader(s))
	for {
		tok, err := dec.Token()
		if err != nil {
			return "", err
		}
		if se, ok := tok.(xml.StartElement); ok {
			return se.Name.Local, nil
		}
	}
}
