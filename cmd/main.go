package main

import (
	"log"
	"net/http"
	"task-manager-app/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// all of our routes here [we'll add more later]
	controllers.LoadTasks()
	r.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(r)))

}
