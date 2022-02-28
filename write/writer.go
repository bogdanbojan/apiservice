package write

import (
	"apiserv/transform"
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
func (fw *FileWriter) WriteRecords(records []transform.ExportRecords, filePath string) error {
	file, err := json.MarshalIndent(records, " ", "\t")
	if err != nil {
		return fmt.Errorf("cannot marshal json: %w", err)
	}
	err = ioutil.WriteFile(filePath, file, 0644)
	if err != nil {
		return fmt.Errorf("cannot write json: %w", err)
	}
	return nil
}
