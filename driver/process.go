package driver

import (
	"apiserv/read"
	"apiserv/transform"
	"apiserv/write"
	"log"
)

// TODO: maybe don't completely shut down app using log.Fatal..
const url = "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"

type RecordReader interface {
}

type RecordTransformer interface {
}

type RecordWriter interface {
}

// Process aggregates the steps that the service has to do in order to transform the data from the API call.
func Process() {
	body, err := read.GetBody(url)
	if err != nil {
		log.Fatal(err)
	}
	exportRecords, err := transform.GetExportRecords(body)
	if err != nil {
		log.Fatal(err)
	}
	err = write.WriteJSON(exportRecords)
	if err != nil {
		log.Fatal(err)
	}
}
