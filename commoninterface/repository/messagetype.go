package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type SchemaFile struct {
	SchemaFile   string               `json:"schemaFile"`
	MessageTypes []*MessageTypeRecord `json:"messageTypes"`
}

type MessageTypeRecord struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func Read(name string) ([]*MessageTypeRecord, error) {
	var records []*MessageTypeRecord

	file, err := os.Open(name)

	if err != nil {
		return nil, fmt.Errorf("open message type file %q: %w", name, err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "--") {
			continue
		}

		values := strings.Split(line, " ")

		records = append(records, &MessageTypeRecord{Code: values[0], Name: values[1]})
	}

	return records, nil
}

func ReadSchemaFile(name string) (*SchemaFile, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		return nil, fmt.Errorf("read schema file %q: %w", name, err)
	}

	file := &SchemaFile{}

	err = json.Unmarshal(data, file)

	if err != nil {
		return nil, fmt.Errorf("deserialize schema file %q: %w", name, err)
	}

	return file, nil
}
