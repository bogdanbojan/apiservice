package read

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Record is the initial structure of the JSON from the API call. It is
// unmarshalled here from the byte array given by reader.go with the RecordsTransform function.
type Record struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Created   string `json:"created"`
	Balance   string `json:"balance"`
}

type FileReader struct{}

// RecordsRead returns the Records which are fetched from the API call that was made to the const `url` found in driver/process.go.
func (rf *FileReader) RecordsRead(url string, recordsNr int) ([]Record, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not use GET on url: %q, err: %v", url, err)
	}

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("status code error %d \n with body %s \n", res.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf("could not read body, error: %v", err)
	}

	err = res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close body, error: %v", err)
	}

	ok := json.Valid(body)
	if !ok {
		return nil, fmt.Errorf("error: JSON is not valid")
	}
	// TODO: Improve errors
	records, err := unmarshalBody(body)
	if err != nil {
		return nil, err
	}
	records, err = validateRecordsNr(records, recordsNr, url)
	if err != nil {
		return nil, err
	}
	return records, nil

}

func validateRecordsNr(records []Record, recordsNr int, url string) ([]Record, error) {
	if isValid(len(records), recordsNr) {
		additionalRecords, err := getAdditionalRecords(records, url, recordsNr)
		if err != nil {
			return nil, err
		}
		return additionalRecords, nil
	} else {
		return records[:recordsNr], nil
	}
}

func unmarshalBody(body []byte) ([]Record, error) {
	var records []Record
	err := json.Unmarshal(body, &records)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal body, err: %v", err)
	}
	return records, nil
}

func getAdditionalRecords(records []Record, url string, recordsNr int) ([]Record, error) {
	rf := FileReader{}
	for len(records) < recordsNr {
		addRecords, err := rf.RecordsRead(url, recordsNr)
		if err != nil {
			return nil, err
		}
		records = addAdditionalRecords(records, addRecords, recordsNr)
	}
	return records, nil
}

func addAdditionalRecords(records []Record, addRecords []Record, recordsNr int) []Record {
	for _, r := range addRecords {
		recordsLen := len(records)
		if isValid(recordsLen, recordsNr) {
			records = append(records, r)
		} else {
			break
		}
	}
	return records
}

func isValid(recordsLen int, recordsNr int) bool {
	return recordsLen < recordsNr
}
