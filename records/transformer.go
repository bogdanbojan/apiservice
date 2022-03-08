package records

import (
	"fmt"
	"sort"
)

// ExportRecords holds the JSON structure for the exported data. It groups the records obtained by the API
// with an Index (first letter of the FirstName) and holds all the records that satisfy that condition in Records.
type ExportRecords struct {
	Index        string   `json:"index"`
	Records      []Record `json:"records"`
	TotalRecords int      `json:"total_records"`
}

type Transformer struct{}

// NewTransformer constructs a new Transformer instance.
func NewTransformer() Transformer {
	return Transformer{}
}

// TransformRecords unmarshals the JSON body data that is fetched from the API call. Then it processes
// it returning an ExportRecords object.
func (Transformer) TransformRecords(records []Record) ([]ExportRecords, error) {
	fmt.Println("Nr of records: ", len(records))
	exportRecords := processExport(records)

	return exportRecords, nil
}

// processExport transforms the data, filtering out duplicates, groups them, and then
// formats the records and, at last, sorts them.
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

// TODO: make sure you check for index out of bounds.
// Maps the records to the first letter of FirstName.
func mapExport(records []Record) map[string][]Record {
	mappedExports := make(map[string][]Record)

	for _, r := range records {
		if r.FirstName != "" {
			firstLetter := r.FirstName[:1]
			mappedExports[firstLetter] = append(mappedExports[firstLetter], r)
		}
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
