package driver

import (
	"apiserv/transform"
	"log"
)

// TODO: maybe don't completely shut down app using log.Fatal..
const url = "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"

type RecordReader interface {
	GetBody(url string) ([]byte, error)
}

type RecordTransformer interface {
	GetExportRecords(body []byte) ([]transform.ExportRecords, error)
}

type RecordWriter interface {
	WriteJSON(records []transform.ExportRecords) error
}

// Process aggregates the steps that the service has to do in order to transform the data from the API call.
func Process(rr RecordReader, rt RecordTransformer, rw RecordWriter) {
	body, err := rr.GetBody(url)
	if err != nil {
		log.Fatal(err)
	}
	exportRecords, err := rt.GetExportRecords(body)
	if err != nil {
		log.Fatal(err)
	}
	err = rw.WriteJSON(exportRecords)
	if err != nil {
		log.Fatal(err)
	}
}
