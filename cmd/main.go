package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// all of our routes here [we'll add more later]
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
