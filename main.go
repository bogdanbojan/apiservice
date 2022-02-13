package main

import "log"

const url = "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"

func main() {
	body, err := getBody(url)
	if err != nil {
		log.Fatal(err)
	}
	err = decode(body)
	if err != nil {
		log.Fatal(err)
	}
}
