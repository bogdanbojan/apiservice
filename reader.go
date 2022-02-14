package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getBody(url string) ([]byte, error) {
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

	return body, nil
}
