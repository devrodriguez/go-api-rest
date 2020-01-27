package routes

import (
	"encoding/json"
	"net/http"

	"../models"
)

func GetTasks(res http.ResponseWriter, req *http.Request) {
	var tasks []models.Task

	tasks = append(tasks, models.Task{ID: 1, Description: "Do homework"})
	tasks = append(tasks, models.Task{ID: 2, Description: "Do excersice"})
	tasks = append(tasks, models.Task{ID: 3, Description: "Take a shower"})
	tasks = append(tasks, models.Task{ID: 4, Description: "Find new hobbies"})

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(tasks)
}

func GetTask(res http.ResponseWriter, req *http.Request) {

}
