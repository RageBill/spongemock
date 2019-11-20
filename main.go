package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("wElCoMe tO sPoNgEmOcK"))
}

func post(w http.ResponseWriter, r *http.Request) {
	// set http header
	w.Header().Set("Content-Type", "application/json")

	// read request body from request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// parse body to json
	var content Message
	err = json.Unmarshal(body, &content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// prepare response message
	var response = Message{Input: content.Input, Output: ""}
	temp := strings.Split(content.Input, " ")
	for _, word := range temp {
		for _, char := range strings.Split(word, "") {
			response.Output += char
			fmt.Println(response.Output)
		}
	}

	// write to http response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}
