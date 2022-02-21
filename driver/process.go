package driver

import (
	"apiserv/input"
	"apiserv/read"
	"apiserv/transform"
	"context"
	"log"
)

// TODO: maybe don't completely shut down app using log.Fatal..

type RecordsConfiger interface {
	ConfigRecords() (*input.FileConfiger, error)
}

type RecordsReader interface {
	ReadRecords(ctx context.Context, recordsNr int) ([]read.Record, error)
}

type RecordsTransformer interface {
	TransformRecords(records []read.Record) ([]transform.ExportRecords, error)
}

type RecordsWriter interface {
	WriteRecords(records []transform.ExportRecords, filePath string) error
}

// Process aggregates the steps that the service has to do in order to transform the data from the API call.
func Process(ctx context.Context, rc RecordsConfiger, rr RecordsReader, rt RecordsTransformer, rw RecordsWriter) {
	config, err := rc.ConfigRecords()
	filePath := config.FilePath
	recordsNr := config.NrOfRecords
	if err != nil {
		log.Fatal(err)
	}
	records, err := rr.ReadRecords(ctx, recordsNr)
	if err != nil {
		log.Fatal(err)
	}
	exportRecords, err := rt.TransformRecords(records)
	if err != nil {
		log.Fatal(err)
	}
	err = rw.WriteRecords(exportRecords, filePath)
	if err != nil {
		log.Fatal(err)
	}
}
