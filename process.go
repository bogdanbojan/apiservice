package main

import (
	"log"
)

// TODO: maybe don't completely shut down app using log.Fatal..
const url = "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"

// process aggregates the steps that the service has to do in order to transform the data from the API call.
func process() {
	body, err := getBody(url)
	if err != nil {
		log.Fatal(err)
	}
	exportRecords, err := getExportRecords(body)
	if err != nil {
		log.Fatal(err)
	}
	err = writeJSON(exportRecords)
	if err != nil {
		log.Fatal(err)
	}
}
