package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// writeJSON exports the records after they have been filtered, formatted and grouped into
// a JSON in the current user folder.
func writeJSON(records []ExportRecords) error {
	file, err := json.MarshalIndent(records, " ", "\t")
	if err != nil {
		return fmt.Errorf("cannot marshal json, error: %v", err)
	}
	err = ioutil.WriteFile("test.json", file, 0644)
	if err != nil {
		return fmt.Errorf("cannot write json, error: %v", err)
	}
	return nil
}
