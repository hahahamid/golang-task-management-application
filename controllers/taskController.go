package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"task-manager-app/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var tasks = []models.Task{}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = uuid.New().String()
	tasks = append(tasks, task)
	SaveTasks()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTask models.Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Completed = updatedTask.Completed
			SaveTasks()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // Remove the task from the slice
			SaveTasks()
			w.WriteHeader(http.StatusNoContent) // Respond with 204 No Content
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
}

const dataFile = "tasks.json" // Define the file name

// LoadTasks loads tasks from the JSON file
func LoadTasks() {
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []models.Task{} // If the file doesn't exist, start with an empty slice
			return
		}
		panic(err) // Handle other errors
	}
	json.Unmarshal(file, &tasks) // Decode the JSON data into the tasks slice
}

// SaveTasks saves tasks to the JSON file
func SaveTasks() {
	data, err := json.Marshal(tasks) // Encode tasks as JSON
	if err != nil {
		panic(err) // Handle encoding errors
	}
	ioutil.WriteFile(dataFile, data, 0644) // Write the JSON data to the file
}
