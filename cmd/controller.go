package main

import (
	"apiserv/config"
	"apiserv/records"
	"context"
	"fmt"
)

type RecordsCollector interface {
	CollectRecords(ctx context.Context, recordsNr int, URL string) ([]records.Record, error)
}

type RecordsTransformer interface {
	TransformRecords(records []records.Record) ([]records.ExportRecords, error)
}

type RecordsWriter interface {
	WriteRecords(records []records.ExportRecords, filePath string) error
}

// Process aggregates the steps that the service has to do in order to transformrec the data from the API call.
func Process(ctx context.Context, cfg *config.Config, rc RecordsCollector, rt RecordsTransformer, rw RecordsWriter) error {
	err := cfg.Init()
	if err != nil {
		return fmt.Errorf("could not config service")
	}
	collectedRecords, err := rc.CollectRecords(ctx, cfg.RecordsNr, cfg.SourceURL)
	if err != nil {
		return fmt.Errorf("could not collectrec records")
	}
	transformedRecords, err := rt.TransformRecords(collectedRecords)
	if err != nil {
		return fmt.Errorf("could not export records")
	}
	err = rw.WriteRecords(transformedRecords, cfg.FilePath)
	if err != nil {
		return fmt.Errorf("could not write records")
	}
	return nil
}
