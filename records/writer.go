package records

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// TODO: WriteRecords should do one thing.

type FileWriter struct{}

// NewFileWriter constructs a new FileWriter instance.
func NewFileWriter() *FileWriter {
	return &FileWriter{}
}

// WriteRecords exports the records after they have been filtered, formatted and grouped into
// a JSON in the current user folder.
func (fw *FileWriter) WriteRecords(records []ExportRecords, filePath string) error {
	file, err := marshalRecords(records)
	if err != nil {
		return err
	}
	err = writeJSON(file, filePath)
	if err != nil {
		return err
	}
	return nil
}

func marshalRecords(records []ExportRecords) ([]byte, error) {
	file, err := json.MarshalIndent(records, " ", "\t")
	if err != nil {
		return nil, fmt.Errorf("cannot marshal json: %w", err)
	}
	return file, nil
}

func writeJSON(file []byte, filePath string) error {
	err := ioutil.WriteFile(filePath, file, 0644)
	if err != nil {
		return fmt.Errorf("cannot writerec json: %w", err)
	}
	return nil
}
