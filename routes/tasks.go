package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devrodriguez/go-api-rest/models"
	"github.com/gorilla/mux"
)

var tasks []models.Task

// tasks = append(tasks, models.Task{ID: 1, Description: "Do homework"})
// tasks = append(tasks, models.Task{ID: 2, Description: "Do excersice"})
// tasks = append(tasks, models.Task{ID: 3, Description: "Take a shower"})
// tasks = append(tasks, models.Task{ID: 4, Description: "Find new hobbies"})

func GetTasks(res http.ResponseWriter, req *http.Request) {
	log.Println("GetTasks called")

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(tasks)
}

func GetTask(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}

	json.NewEncoder(res).Encode(&models.Task{})
}

func CreateTask(res http.ResponseWriter, req *http.Request) {
	log.Println("GetTasks called")

	params := mux.Vars(req)
	var task models.Task

	task.ID = params["id"]
	json.NewDecoder(req.Body).Decode(&task)

	tasks = append(tasks, task)

	json.NewEncoder(res).Encode(tasks)
}

func DeleteTask(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}

	json.NewEncoder(res).Encode(tasks)
}
