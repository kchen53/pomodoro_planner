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

var ToDo models.ToDo

func GetToDo(w http.ResponseWriter, r *http.Request) {
	newToDos := models.GetAllToDo()
	res, _ := json.Marshal(newToDos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetToDoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["i"]
	ID, err := strconv.ParseInt(todoID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todoDetails, _ := models.GetToDoByID(ID)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateToDo(w http.ResponseWriter, r *http.Request) {
	CreateToDo := &models.ToDo{}
	utils.ParseBody(r, CreateToDo)
	b := CreateToDo.CreateToDo()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["i"]
	ID, err := strconv.ParseInt(todoID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todo := models.DeleteToDo(ID)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateToDo(w http.ResponseWriter, r *http.Request) {
	updateToDo := &models.ToDo{}
	utils.ParseBody(r, updateToDo)

	vars := mux.Vars(r)
	todoID := vars["i"]
	ID, err := strconv.ParseInt(todoID, 0, 0)
	if err != nil {
		fmt.Println("error in parsing")
	}
	//todoDetails, db := models.GetToDoByID(ID)
	todoDetails, _ := models.GetToDoByID(ID) //TODO: DELETE
	if updateToDo.Task != "" {
		todoDetails.Task = updateToDo.Task
	}
	if updateToDo.Due != "" {
		todoDetails.Due = updateToDo.Due
	}
	todoDetails.Complete = updateToDo.Complete
	//db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
