package read

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//go:embed apiservice.env
var envURL embed.FS

// Record is the initial structure of the JSON from the API call. It is
// unmarshalled here from the byte array given by reader.go with the TransformRecords function.
type Record struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Created   string `json:"created"`
	Balance   string `json:"balance"`
}

type FileReader struct{}

// NewFileReader constructs a new FileReader instance.
func NewFileReader() *FileReader {
	return &FileReader{}
}

// ReadRecords returns the Records which are fetched from the API call that was made to the const `url` found in driver/process.go.
func (fr *FileReader) ReadRecords(ctx context.Context, recordsNr int) ([]Record, error) {
	records, err := getRecords(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get initial records: %w", err)
	}
	validatedRecords, err := validateRecordsNr(records, recordsNr)
	if err != nil {
		return nil, fmt.Errorf("could not get validatedRecords: %w", err)
	}
	return validatedRecords, nil

}

// getRecords takes the url string and unmarshals the API response into an array of Record type with the unmarshalBody helper function.
func getRecords(ctx context.Context) ([]Record, error) {
	client := &http.Client{}
	url, err := getEnvURL()
	if err != nil {
		return nil, fmt.Errorf("could not get the url from the env file: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not use GET on url: %q, err: %w", url, err)
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
		return nil, fmt.Errorf("could not read the body: %w", err)
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

// validateRecordsNr checks that the user's records are exactly the number the user wanted. If not, it uses the helper function
// getAdditionalRecords to get more records from the API.
func validateRecordsNr(records []Record, recordsNr int) ([]Record, error) {
	if isValid(len(records), recordsNr) {
		additionalRecords, err := getAdditionalRecords(records, recordsNr)
		if err != nil {
			return nil, fmt.Errorf("could not get additional records: %w", err)
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
		return nil, fmt.Errorf("could not unmarshal body: %w", err)
	}
	return records, nil
}

// getAdditionalRecords is a helper function that loops until it gets the records bounded by the number set by the user.
func getAdditionalRecords(records []Record, recordsNr int) ([]Record, error) {
	for len(records) < recordsNr {
		ctx := context.Background()
		addRecords, err := getRecords(ctx)
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

// TODO: additional checks?
// getEnvURL fetches the URL from the .env file.
func getEnvURL() (string, error) {
	key := "SOURCE_URL"
	u, err := envURL.ReadFile("apiservice.env")
	if err != nil {
		return "", fmt.Errorf("cannot read url from envfile: %w", err)
	}
	trimmedPrefixURL := strings.TrimPrefix(string(u), key+"=")
	return trimmedPrefixURL, nil
}
