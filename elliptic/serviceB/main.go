package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/serviceb", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"welcome form gorilla": true})
	})
	router.HandleFunc("/serviceb/hi", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"hi gorilla": true})
	})
	log.Fatal(http.ListenAndServe(":8000", router))
}
