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
	fc := input.FileConfiger{}
	fr := read.FileReader{}
	fd := transform.FileDecoder{}
	fw := write.FileWriter{}
	driver.Process(&fc, &fr, &fd, &fw)
}

// TODO:
// package names write, read problematic- too general already used in the standard libraary?
// interface methods not similar to interface name
// decoder is already in standard go library
