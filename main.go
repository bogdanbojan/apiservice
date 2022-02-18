package main

import (
	"apiserv/driver"
	"apiserv/input"
	"apiserv/read"
	"apiserv/transform"
	"apiserv/write"
	"context"
)

// TODO: unit tests for each step.
// TODO: package names write, read, decoder problematic- too general already used in the standard library?
func main() {
	ctx := context.Background()
	fc := input.NewFileConfiger()
	fr := read.NewFileReader()
	fd := transform.NewFileDecoder()
	fw := write.NewFileWriter()
	driver.Process(ctx, fc, fr, fd, fw)
}
