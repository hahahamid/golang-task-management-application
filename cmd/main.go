package main

import (
	"log"
	"net/http"
	"task-manager-app/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// all of our routes here [we'll add more later]
	r.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks", controllers.CreateTask).Methods("POST")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
