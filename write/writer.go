package write

import (
	"apiserv/transform"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type FileWriter struct{}

// WriteJSON exports the records after they have been filtered, formatted and grouped into
// a JSON in the current user folder.
func (fw *FileWriter) WriteJSON(records []transform.ExportRecords) error {
	file, err := json.MarshalIndent(records, " ", "\t")
	if err != nil {
		return fmt.Errorf("cannot marshal json, error: %v", err)
	}
	err = ioutil.WriteFile("test.json", file, 0644) // TODO: make the file name dynamic.
	if err != nil {
		return fmt.Errorf("cannot write json, error: %v", err)
	}
	return nil
}
