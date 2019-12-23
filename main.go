package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ramonmoraes/grpc-http/httpserver"
)

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
	for i < 100 {
		res := function(b)
		results = append(results, res)
		i++
	}
	return results
}

func main() {
	fmt.Println("----- Start -----")
	jsonB := preLoadJSON()
	go httpserver.Serve()
	benchmarkFunc(jsonB, httpRequest)
	fmt.Println("----- End   -----")
}
