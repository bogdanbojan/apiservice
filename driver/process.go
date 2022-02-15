package driver

import (
	"apiserv/input"
	"apiserv/transform"
	"log"
)

// TODO: maybe don't completely shut down app using log.Fatal..
const url = "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"

type RecordConfig interface {
	RecordsConfig() (*input.Configuration, error)
}

type RecordReader interface {
	GetBody(url string) ([]byte, error)
}

type RecordTransformer interface {
	GetExportRecords(body []byte) ([]transform.ExportRecords, error)
}

type RecordWriter interface {
	WriteJSON(records []transform.ExportRecords, filePath string) error
}

// Process aggregates the steps that the service has to do in order to transform the data from the API call.
func Process(rc RecordConfig, rr RecordReader, rt RecordTransformer, rw RecordWriter) {
	config, err := rc.RecordsConfig()
	filePath := config.FilePath
	if err != nil {
		log.Fatal(err)
	}
	body, err := rr.GetBody(url)
	if err != nil {
		log.Fatal(err)
	}
	exportRecords, err := rt.GetExportRecords(body)
	if err != nil {
		log.Fatal(err)
	}
	err = rw.WriteJSON(exportRecords, filePath)
	if err != nil {
		log.Fatal(err)
	}
}
