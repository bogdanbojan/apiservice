package transform

import (
	"apiserv/read"
	"fmt"
	"sort"
)

// ExportRecords holds the JSON structure for the exported data. It groups the records obtained by the API
// with an Index (first letter of the FirstName) and holds all the records that satisfy that condition in Records.
type ExportRecords struct {
	Index        string        `json:"index"`
	Records      []read.Record `json:"records"`
	TotalRecords int           `json:"total-records"`
}

type FileDecoder struct{}

// NewFileDecoder constructs a new FileDecoder instance.
func NewFileDecoder() *FileDecoder {
	return &FileDecoder{}
}

// TransformRecords unmarshals the JSON body data that is fetched from the API call. Then it processes
// it returning an ExportRecords object.
func (fd *FileDecoder) TransformRecords(records []read.Record) ([]ExportRecords, error) {

	fmt.Println("Nr of records: ", len(records))

	exportRecords := processExport(records)

	return exportRecords, nil
}

// processExport transforms the data, filtering out duplicates, groups them, and then
// formats the records and, at last, sorts them.
func processExport(records []read.Record) []ExportRecords {
	uniqueRecords := removeDuplicates(records)
	mappedExports := mapExport(uniqueRecords)
	formattedExports := formatExport(mappedExports)
	sortedExports := sortRecords(formattedExports)

	return sortedExports
}

// Formats the mapping into a slice that contains formatted ExportRecords types.
func formatExport(mappedExports map[string][]read.Record) []ExportRecords {
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
func mapExport(records []read.Record) map[string][]read.Record {
	mappedExports := make(map[string][]read.Record)

	for _, r := range records {
		firstLetter := r.FirstName[:1]
		mappedExports[firstLetter] = append(mappedExports[firstLetter], r)
	}

	return mappedExports
}

// sortRecords sorts the given records after the first letter of the persons' first name.
func sortRecords(records []ExportRecords) []ExportRecords {
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].Index < records[j].Index
	})

	return records
}

// removeDuplicates gets rid of any duplicate entries from the provided records that come from the API call.
func removeDuplicates(records []read.Record) []read.Record {
	var uniqueRecords []read.Record // filter wo allocation
	recordMap := make(map[read.Record]struct{})

	for _, record := range records {
		if _, ok := recordMap[record]; !ok {
			uniqueRecords = append(uniqueRecords, record)
			recordMap[record] = struct{}{}
		}
	}
	return uniqueRecords
}
