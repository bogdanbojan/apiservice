package main

import (
	"apiserv/driver"
	"apiserv/input"
	"apiserv/read"
	"apiserv/transform"
	"apiserv/write"
	"context"
)

func main() {
	ctx := context.Background()
	fc := input.NewFileConfiger()
	fr := read.NewFileReader()
	fd := transform.NewFileDecoder()
	fw := write.NewFileWriter()
	driver.Process(ctx, fc, fr, fd, fw)
}
