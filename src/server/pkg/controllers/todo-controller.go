package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kchen53/pomodoro_planner/pkg/models"
	"github.com/kchen53/pomodoro_planner/pkg/utils"
)

var Todo models.Todo

func GetTodo(w http.ResponseWriter, r *http.Request) {
	newTodos := models.GetAllTodo()
	res, _ := json.Marshal(newTodos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["i"]
	ID, err := strconv.ParseInt(todoID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todoDetails, _ := models.GetTodoByID(ID)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	CreateTodo := &models.Todo{}
	utils.ParseBody(r, CreateTodo)
	b := CreateTodo.CreateTodo()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["i"]
	ID, err := strconv.ParseInt(todoID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todo := models.DeleteTodo(ID)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// updateTodo := &models.Todo{}
	// utils.ParseBody(r, updateTodo)

	// vars := mux.Vars(r)
	// todoID := vars["i"]
	// ID, err := strconv.ParseInt(todoID, 0, 0)
	// if err != nil {
	// 	fmt.Println("error in parsing")
	// }
	// //todoDetails, db := models.GetTodoByID(ID)
	// todoDetails, _ := models.GetTodoByID(ID) //TODO: DELETE
	// if updateTodo.Task != "" {
	// 	todoDetails.Task = updateTodo.Task
	// }
	// if updateTodo.Due != "" {
	// 	todoDetails.Due = updateTodo.Due
	// }
	// todoDetails.Complete = updateTodo.Complete
	// //db.Save(&todoDetails)
	// res, _ := json.Marshal(todoDetails)
	// w.Header().Set("Content-Type", "pkglication/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}
