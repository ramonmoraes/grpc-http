package httpserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", getNameHandler)
	fmt.Println("Listening at :8080")
	http.ListenAndServe(":8080", nil)
}

type withName struct {
	Name string
}

func getNameHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var jsonRequest withName
	err = json.Unmarshal(body, &jsonRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, jsonRequest.Name)
}
