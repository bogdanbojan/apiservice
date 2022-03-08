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
	fr := records.NewFileReader()
	fd := records.NewFileDecoder()
	fw := records.NewFileWriter()
	err := Process(ctx, cfg, fr, fd, fw)
	if err != nil {
		log.Fatal(err)
	}
}
