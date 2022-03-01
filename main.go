package main

import (
	"apiserv/driver"
	"apiserv/input"
	"apiserv/read"
	"apiserv/transform"
	"apiserv/write"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	fc := input.NewFileConfiger()
	fr := read.NewFileReader()
	fd := transform.NewFileDecoder()
	fw := write.NewFileWriter()
	err := driver.Process(ctx, fc, fr, fd, fw)
	if err != nil {
		log.Fatal(err)
	}
}
