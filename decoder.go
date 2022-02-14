package main

import (
	"encoding/json"
	"fmt"
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

func decode(body []byte) ([]ExportRecords, error) {
	var records []Record

	err := json.Unmarshal(body, &records)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal body, err: %v", err)
	}

	exportRecords := processExport(records)

	return exportRecords, nil
}

func processExport(records []Record) []ExportRecords {
	uniqueRecords := removeDuplicates(records)
	mappedExports := mapExport(uniqueRecords)
	formattedExports := formatExport(mappedExports)
	sortedExports := sortRecords(formattedExports)

	return sortedExports
}

// Formats the mapping into a slice that contains formatted ExportRecords types.
func formatExport(mappedExports map[string][]Record) []ExportRecords {
	var exportRecords []ExportRecords
	for k, v := range mappedExports {
		exportRecords = append(exportRecords, ExportRecords{
			Index:        k,
			Records:      v,
			TotalRecords: len(v),
		})
	}
	return exportRecords
}

// Maps the records to the first letter of FirstName.
func mapExport(records []Record) map[string][]Record {
	mappedExports := make(map[string][]Record)

	for _, r := range records {
		firstLetter := r.First[:1]
		mappedExports[firstLetter] = append(mappedExports[firstLetter], r)
	}

	return mappedExports
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
