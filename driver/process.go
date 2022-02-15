package driver

import (
	"apiserv/input"
	"apiserv/transform"
	"log"
)

// TODO: maybe don't completely shut down app using log.Fatal..
const url = "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"

type RecordsConfiger interface {
	RecordsConfig() (*input.Configuration, error)
}

type RecordsReader interface {
	RecordsRead(url string) ([]byte, error)
}

type RecordsTransformer interface {
	RecordsTransform(body []byte) ([]transform.ExportRecords, error)
}

type RecordsWriter interface {
	RecordsWrite(records []transform.ExportRecords, filePath string) error
}

// Process aggregates the steps that the service has to do in order to transform the data from the API call.
func Process(rc RecordsConfiger, rr RecordsReader, rt RecordsTransformer, rw RecordsWriter) {
	config, err := rc.RecordsConfig()
	filePath := config.FilePath
	if err != nil {
		log.Fatal(err)
	}
	body, err := rr.RecordsRead(url)
	if err != nil {
		log.Fatal(err)
	}
	exportRecords, err := rt.RecordsTransform(body)
	if err != nil {
		log.Fatal(err)
	}
	err = rw.RecordsWrite(exportRecords, filePath)
	if err != nil {
		log.Fatal(err)
	}
}