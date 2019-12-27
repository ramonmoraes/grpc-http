package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ramonmoraes/grpc-http/grpcserver"
	"github.com/ramonmoraes/grpc-http/httpserver"
)

var EXTRACTOR_TYPE = os.Getenv("EXTRACTOR_TYPE")

func preLoadJSON() []byte {
	jsonPath := "./sample.json"
	b, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func httpRequest(json []byte) string {
	res, err := http.Post("http://localhost:8080/", "application/json", bytes.NewReader(json))
	if err != nil {
		log.Fatal(err)
	}
	text, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(text)
}

func benchmarkFunc(b []byte, function func(b []byte) string) []string {
	i := 0
	results := []string{}
	for i < 10000 {
		res := function(b)
		results = append(results, res)
		i++
	}
	return results
}

func main() {
	fmt.Println("----- Start -----")
	if EXTRACTOR_TYPE == "http" {
		httpserver.Serve()
	}

	if EXTRACTOR_TYPE == "grpc" {
		grpcserver.Serve()
	}
	if EXTRACTOR_TYPE == "" {
		fmt.Println(("Empty extractor type"))
	}
	// jsonB := preLoadJSON()
	// benchmarkFunc(jsonB, httpRequest)
	fmt.Println("----- End   -----")
}
