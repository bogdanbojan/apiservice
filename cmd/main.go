package main

import (
	"apiserv/config"
	"apiserv/records"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()
	rc := records.NewCollector()
	rt := records.NewTransformer()
	rw := records.NewWriter()
	err := Process(ctx, cfg, rc, rt, rw)
	if err != nil {
		log.Fatal(err)
	}
}
