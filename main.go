package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define your routes here
	r.HandleFunc("/status", statusHandler).Methods("GET")
	r.HandleFunc("/repair-bay", repairHandler).Methods("GET")
	r.HandleFunc("/teapot", teapotHandler).Methods("GET")

	http.Handle("/", r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
