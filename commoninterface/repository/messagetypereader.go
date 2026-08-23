package repository

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MessageTypeRecord struct {
	Code string
	Name string
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
