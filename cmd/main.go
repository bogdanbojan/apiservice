package main

import (
	"apiserv/cmd"
	"apiserv/config"
	"apiserv/records"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	fc := config.NewFileConfiger()
	fr := records.NewFileReader()
	fd := records.NewFileDecoder()
	fw := records.NewFileWriter()
	err := cmd.Process(ctx, fc, fr, fd, fw)
	if err != nil {
		log.Fatal(err)
	}
}
