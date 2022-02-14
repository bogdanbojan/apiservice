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

func formatExportTest(records []Record) []ExportRecords {
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

func formatExport(records []Record) []ExportRecords {
	//var exportRecords []ExportRecords

	for _, r := range records {
		fmt.Println(r)
	}

	exportRecords := make([]ExportRecords, len(records))
	//for i := 0; i < len(records); i++ {
	//	expRecords := exportRecords[i]
	//
	//	expRecords.Index = getIndex(records[i])
	//
	//	if records[i].First[:1] == records[i+1].First[:1] {
	//		expRecords.Records = getIndexedRecords(records[i:])
	//	} else {
	//		expRecords.Records = append(expRecords.Records, records[i])
	//	}
	//	expRecords.TotalRecords = getTotalRecords(expRecords.Records)
	//
	//}
	//fmt.Println(exportRecords)
	return exportRecords
}

func getIndex(records Record) string {
	return records.First[:1]
}

func getTotalRecords(records []Record) int {
	return len(records)
}

func getIndexedRecords(records []Record) []Record {
	var indexedRecords []Record
	i := 0
	for records[i].First[:1] == records[i+1].First[:1] {
		indexedRecords = append(indexedRecords, records[i])
		i++
	}
	indexedRecords = append(indexedRecords, records[i]) // make sure to add the last one
	return indexedRecords
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

func testingRemoveDuplicates() {
	records := []Record{
		{
			First:   "Thad",
			Last:    "Feest",
			Email:   "Thad.Feest@cleve.name",
			Address: "77881 Schaefer Loaf",
			Created: "July 13, 2021",
			Balance: "$1,950.71",
		},
		{First: "Thad",
			Last:    "Feest",
			Email:   "Thad.Feest@cleve.name",
			Address: "77881 Schaefer Loaf",
			Created: "July 13, 2021",
			Balance: "$1,950.71",
		},
		{
			First:   "Thad",
			Last:    "Adam",
			Email:   "Thad.Feest@cleve.name",
			Address: "77881 Schaefer Loaf",
			Created: "July 13, 2021",
			Balance: "$1,950.71"},
		{
			First:   "Sabrina",
			Last:    "Kuphal",
			Email:   "whiterabbit70@gmail.com",
			Address: "03491 Howard Vista",
			Created: "August 29, 2018",
			Balance: "$6,996.45"},
	}

	formatExportTest(records)
	//removeDuplicates(records)

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
