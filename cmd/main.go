package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", IndexRoute)
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/task", CreateTask).Methods("POST")
	router.HandleFunc("/task", UpdateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
