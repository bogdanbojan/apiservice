package cmd

import (
	"apiserv/config"
	"apiserv/records"
	"context"
	"fmt"
)

type RecordsConfiger interface {
	ConfigRecords() (*config.FileConfiger, error)
}

type RecordsReader interface {
	ReadRecords(ctx context.Context, recordsNr int) ([]records.Record, error)
}

type RecordsTransformer interface {
	TransformRecords(records []records.Record) ([]records.ExportRecords, error)
}

type RecordsWriter interface {
	WriteRecords(records []records.ExportRecords, filePath string) error
}

// Process aggregates the steps that the service has to do in order to transformrec the data from the API call.
func Process(ctx context.Context, rc RecordsConfiger, rr RecordsReader, rt RecordsTransformer, rw RecordsWriter) error {
	config, err := rc.ConfigRecords()
	filePath := config.FilePath
	recordsNr := config.RecordsNr
	if err != nil {
		return fmt.Errorf("could not config records")
	}
	records, err := rr.ReadRecords(ctx, recordsNr)
	if err != nil {
		return fmt.Errorf("could not collectrec records")
	}
	exportRecords, err := rt.TransformRecords(records)
	if err != nil {
		return fmt.Errorf("could not export records")
	}
	err = rw.WriteRecords(exportRecords, filePath)
	if err != nil {
		return fmt.Errorf("could not writerec records")
	}
	return nil
}
