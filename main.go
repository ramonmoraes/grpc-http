package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type SampleJSON struct {
	Name    string
	Age     int
	Hoobies []string
}

func preLoadJSON() []byte {
	jsonPath := "./sample.json"
	bytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func main() {
	fmt.Println("Hello world")
}
