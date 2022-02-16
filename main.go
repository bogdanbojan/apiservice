package main

import (
	"apiserv/driver"
	"apiserv/input"
	"apiserv/read"
	"apiserv/transform"
	"apiserv/write"
)

// TODO: reformat naming and introduce unit tests.
// TODO: transform first letter lowercase for case safety.
func main() {
	fc := input.NewFileConfiger()
	fr := read.NewFileReader()
	fd := transform.NewFileDecoder()
	fw := write.NewFileWriter()
	driver.Process(fc, fr, fd, fw)
}

// TODO:
// package names write, read problematic- too general already used in the standard libraary?
// interface methods not similar to interface name
// decoder is already in standard go library
