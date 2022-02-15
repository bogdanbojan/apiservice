package main

import (
	"apiserv/driver"
	"apiserv/input"
	"apiserv/read"
	"apiserv/transform"
	"apiserv/write"
)

// TODO: reformat naming, additional calls if body length is not sufficient, packages and unit tests.
// TODO: transform first letter lowercase for case safety.
func main() {
	fc := input.Configer{}
	fr := read.FileReader{}
	fd := transform.Decoder{}
	fw := write.FileWriter{}
	driver.Process(&fc, &fr, &fd, &fw)
}

// TODO:
// package names write, read problematic- too general already used in the standard libraary?
// interface methods not similar to interface name
// decoder is already in standard go library
// add user input packaage