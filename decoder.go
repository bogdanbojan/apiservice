package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type Record struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Created string `json:"created"`
	Balance string `json:"balance"`
}

type ExportRecords struct {
	Index        string   `json:"index"`
	Records      []Record `json:"records"`
	TotalRecords int      `json:"total-records"`
}

func formatExport(records []Record) []ExportRecords {
	var exportRecords []ExportRecords
	mappedExports := make(map[string][]Record)

	for _, r := range records {
		firstLetter := r.First[:1]
		mappedExports[firstLetter] = append(mappedExports[firstLetter], r)
	}

	for k, v := range mappedExports {
		exportRecords = append(exportRecords, ExportRecords{
			Index:        k,
			Records:      v,
			TotalRecords: len(v),
		})
	}

	sortRecords(exportRecords)

	return exportRecords
}

func decode(body []byte) error {
	var records []Record

	// unmarshal the json
	err := json.Unmarshal(body, &records)
	if err != nil {
		return fmt.Errorf("could not unmarshal body, err: %v", err)
	}

	// steps for record processing
	// remove duplicates
	uniqueRecords := removeDuplicates(records)

	// sort records
	//sortedRecords := sortRecords(uniqueRecords)

	// format records
	exportRecords := formatExport(uniqueRecords)

	//marshal the json and write it
	err = writeJSON(exportRecords)
	if err != nil {
		return fmt.Errorf("cannot write json, error: %v", err)
	}

	return nil
}

// Sorts records after first name.
func sortRecords(records []ExportRecords) []ExportRecords {
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].Index < records[j].Index
	})

	return records
}

func removeDuplicates(records []Record) []Record {
	var uniqueRecords []Record // filter wo allocation
	recordMap := make(map[Record]struct{})

	for _, record := range records {
		if _, ok := recordMap[record]; !ok {
			uniqueRecords = append(uniqueRecords, record)
			recordMap[record] = struct{}{}
		}
	}
	return uniqueRecords
}

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
