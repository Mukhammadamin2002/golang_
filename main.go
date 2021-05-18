package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"Title"`
	Duration string `json:"duration"`
	Date     string `json:"date"`
}

var tasks []Task

func allTasks(w http.ResponseWriter, r *http.Request) {
	// tasks := Tasks{
	// 	Task{Title: "Housework", Duration: "2 hour", Date: "Monday"},
	// 	Task{Title: "Laisure Activity", Duration: "1 hour", Date: "Tuesday"},
	// 	Task{Title: "Go Suburban Area", Duration: "3:30 hour", Date: "Wednesday"},
	// 	Task{Title: "Holiday", Duration: "whole day", Date: "Friday"},
	// 	Task{Title: "Climbing", Duration: "3 hour", Date: "Thursday"},
	// 	Task{Title: "Helping Other People", Duration: "1:30 hour", Date: "Saturday"},
	// 	Task{Title: "Fishing", Duration: "1:50 hour", Date: "Sunday"},
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(10000000))
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)

}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			tasks = append(tasks, task)
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	json.NewEncoder(w).Encode(tasks)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/tasks", allTasks).Methods("GET")
	myRouter.HandleFunc("/task/{id}", getTask).Methods("GET")
	myRouter.HandleFunc("/task", createTask).Methods("POST")
	myRouter.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	tasks = append(tasks, Task{ID: "1", Title: "Housework", Duration: "2 hour", Date: "Monday"})
	tasks = append(tasks, Task{ID: "2", Title: "Laisure Activity", Duration: "1 hour", Date: "Tuesday"})
	tasks = append(tasks, Task{ID: "3", Title: "Go Suburban Area", Duration: "3:30 hour", Date: "Wednesday"})
	tasks = append(tasks, Task{ID: "4", Title: "Holiday", Duration: "whole day", Date: "Friday"})
	tasks = append(tasks, Task{ID: "5", Title: "Climbing", Duration: "3 hour", Date: "Thursday"})
	tasks = append(tasks, Task{ID: "6", Title: "Helping Other People", Duration: "1:30 hour", Date: "Saturday"})
	tasks = append(tasks, Task{ID: "7", Title: "Fishing", Duration: "1:50 hour", Date: "Sunday"})

	handleRequests()
}
