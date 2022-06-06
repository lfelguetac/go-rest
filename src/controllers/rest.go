package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type AllTask []Task

var tareas = AllTask{
	{
		ID:      1,
		Name:    "Tarea 1",
		Content: "Contenido de la tarea 1",
	},
	{
		ID:      2,
		Name:    "Tarea 2",
		Content: "Contenido de la tarea 2",
	},
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Linares")

}
func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Funcionando")
	json.NewEncoder(w).Encode(tareas)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var newtask Task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error al leer el body")
	}

	json.Unmarshal(reqBody, &newtask)

	newtask.ID = len(tareas) + 1
	_ = append(tareas, newtask)

	json.NewEncoder(w).Encode("OK")

}

func GetTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Error al convertir el id")
		return
	}

	for _, task := range tareas {
		if task.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Error al convertir el id")
		return
	}

	var newtask Task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error al leer el body")
	}

	json.Unmarshal(reqBody, &newtask)

	for i, task := range tareas {
		if task.ID == taskId {
			tareas = append(tareas[:i], tareas[i+1:]...)
			newtask.ID = taskId
			_ = append(tareas, newtask)
			json.NewEncoder(w).Encode("OK")
			return
		}
	}
	json.NewEncoder(w).Encode("No se encontro el id")
}
