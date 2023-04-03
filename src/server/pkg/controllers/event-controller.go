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

var Event models.Event

func GetEvent(w http.ResponseWriter, r *http.Request) {
	newEvents := models.GetAllEvent()
	res, _ := json.Marshal(newEvents)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["i"]
	ID, err := strconv.ParseInt(eventID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	eventDetails := models.GetEventByID(int(ID))
	res, _ := json.Marshal(eventDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	createEvent := &models.Event{}
	utils.ParseBody(r, createEvent)
	t := createEvent.CreateEvent()
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["i"]
	ID, err := strconv.ParseInt(eventID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	event := models.DeleteEvent(int(ID))
	res, _ := json.Marshal(event)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	updateEvent := &models.Event{}
	utils.ParseBody(r, updateEvent)

	vars := mux.Vars(r)
	eventID := vars["i"]
	ID, err := strconv.ParseInt(eventID, 0, 0)
	if err != nil {
		fmt.Println("error in parsing")
	}
	updateEvent.ID = int(ID)

	t := updateEvent.UpdateEvent()
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
