package driver

import (
	"apiserv/input"
	"apiserv/read"
	"apiserv/transform"
	"context"
	"fmt"
)

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
func Process(ctx context.Context, rc RecordsConfiger, rr RecordsReader, rt RecordsTransformer, rw RecordsWriter) error {
	config, err := rc.ConfigRecords()
	filePath := config.FilePath
	recordsNr := config.NrOfRecords
	if err != nil {
		return fmt.Errorf("could not config records")
	}
	records, err := rr.ReadRecords(ctx, recordsNr)
	if err != nil {
		return fmt.Errorf("could not read records")
	}
	exportRecords, err := rt.TransformRecords(records)
	if err != nil {
		return fmt.Errorf("could not export records")
	}
	err = rw.WriteRecords(exportRecords, filePath)
	if err != nil {
		return fmt.Errorf("could not write records")
	}
	return nil
}
