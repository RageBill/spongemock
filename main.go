package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Message struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

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
	originalString := strings.Split(content.Input, " ")
	for _, word := range originalString {
		// default to upper case if word is only 1 letter long
		if len(word) == 1 {
			response.Output += strings.ToUpper(word)
			response.Output += " "
			continue
		}

		// odd letter = lower case, even letter = upper case, skip the count on special characters
		count := 0
		for _, char := range strings.Split(word, "") {
			// check if special characters
			if IsLetter(char) {
				count += 1
			} else {
				response.Output += char
				continue
			}

			// convert cases base on count of letter
			if count%2 == 0 {
				response.Output += strings.ToUpper(char)
			} else {
				response.Output += strings.ToLower(char)
			}
		}
		response.Output += " "
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
