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
	records, err := getRecords(url)
	if err != nil {
		return nil, err
	}
	validatedRecords, err := validateRecordsNr(records, recordsNr, url)
	if err != nil {
		return nil, err
	}
	return validatedRecords, nil

}

// getRecords takes the url string and unmarshals the API response into an array of Record type with the unmarshalBody helper function.
func getRecords(url string) ([]Record, error) {
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
	return records, nil
}

// validateRecordsNr checks that the user's records are exactly the number the user wanted. If not, it uses the helper function
// getAdditionalRecords to get more records from the API.
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

// unmarshalBody takes the byte slice given in the body of the API response and unmarshals it into a slice of Record type.
func unmarshalBody(body []byte) ([]Record, error) {
	var records []Record
	err := json.Unmarshal(body, &records)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal body, err: %v", err)
	}
	return records, nil
}

// getAdditionalRecords is a helper function that loops until it gets the records bounded by the number set by the user.
func getAdditionalRecords(records []Record, url string, recordsNr int) ([]Record, error) {
	for len(records) < recordsNr {
		addRecords, err := getRecords(url)
		if err != nil {
			return nil, err
		}
		records = addAdditionalRecords(records, addRecords, recordsNr)
	}
	return records, nil
}

// addAdditionalRecords iterates through the records given by getAdditionalRecords and keeps adding them to
// the slice `records` if and only if they are valid. That is, if the current records entries are smaller than
// the users desired number.
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

// isValid checks if the current number of records is smaller than the number set by the user.
func isValid(recordsLen int, recordsNr int) bool {
	return recordsLen < recordsNr
}
