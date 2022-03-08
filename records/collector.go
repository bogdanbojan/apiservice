package records

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Record is the initial structure of the JSON from the API call. It is
// unmarshalled here from the byte array given by collector.go with the TransformRecords function.
type Record struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Created   string `json:"created"`
	Balance   string `json:"balance"`
}

type Collector struct{}

// NewCollector constructs a new Collector instance.
func NewCollector() Collector {
	return Collector{}
}

// CollectRecords returns the Records which are fetched from the API call that was made to the const `url` found in cmd/controller.go.
func (Collector) CollectRecords(ctx context.Context, recordsNr int, URL string) ([]Record, error) {
	records, err := getRecords(ctx, URL)
	if err != nil {
		return nil, fmt.Errorf("could not get initial records: %w", err)
	}
	validatedRecords, err := validateRecordsNr(ctx, records, recordsNr, URL)
	if err != nil {
		return nil, fmt.Errorf("could not get validatedRecords: %w", err)
	}
	return validatedRecords, nil

}

// getRecords takes the url string and unmarshals the API response into an array of Record type with the unmarshalBody helper function.
func getRecords(ctx context.Context, URL string) ([]Record, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not use GET on url: %q, err: %w", URL, err)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not establish proper HTTP connection: %w", err)
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("status code error %d \n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not collectrec the body: %w", err)
	}

	err = res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close the body: %w", err)
	}

	ok := json.Valid(body)
	if !ok {
		return nil, errors.New("error: JSON is not valid")
	}

	records, err := unmarshalBody(body)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal the body: %w", err)
	}
	return records, nil
}

// TODO: refactor this method so it does one thing.
// validateRecordsNr checks that the user's records are exactly the number the user wanted. If not, it uses the helper function
// getAdditionalRecords to get more records from the API.
func validateRecordsNr(ctx context.Context, records []Record, recordsNr int, URL string) ([]Record, error) {
	if !isValid(len(records), recordsNr) {
		return records[:recordsNr], nil
	}
	additionalRecords, err := getAdditionalRecords(ctx, records, recordsNr, URL)
	if err != nil {
		return nil, fmt.Errorf("could not get additional records: %w", err)
	}
	return additionalRecords, nil
}

// unmarshalBody takes the byte slice given in the body of the API response and unmarshals it into a slice of Record type.
func unmarshalBody(body []byte) ([]Record, error) {
	var records []Record
	err := json.Unmarshal(body, &records)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal body: %w", err)
	}
	return records, nil
}

// getAdditionalRecords is a helper function that loops until it gets the records bounded by the number set by the user.
func getAdditionalRecords(ctx context.Context, records []Record, recordsNr int, URL string) ([]Record, error) {
	for len(records) < recordsNr {
		addRecords, err := getRecords(ctx, URL)
		if err != nil {
			return nil, fmt.Errorf("could not get addRecords: %w", err)
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
		if !isValid(recordsLen, recordsNr) {
			break
		}
		records = append(records, r)
	}
	return records
}

// isValid checks if the current number of records is smaller than the number set by the user.
func isValid(recordsLen int, recordsNr int) bool {
	return recordsLen < recordsNr
}
