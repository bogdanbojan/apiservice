package main

import (
	"apiserv/driver"
	"apiserv/read"
	"apiserv/transform"
	"apiserv/write"
)

// TODO: reformat naming, additional calls if body length is not sufficient, packages and unit tests.
// TODO: transform first letter lowercase for case safety.
func main() {

	fr := read.FileReader{}
	fd := transform.Decoder{}
	fw := write.FileWriter{}
	driver.Process(&fr, &fd, &fw)
}
