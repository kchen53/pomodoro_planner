package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/kchen53/pomodoro_planner/pkg/config"
	"github.com/kchen53/pomodoro_planner/pkg/controllers"
	"github.com/kchen53/pomodoro_planner/pkg/models"
	"github.com/kchen53/pomodoro_planner/pkg/utils"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func TestGetEvent(t *testing.T) {

	// Call login
	Setup()
	Reset(t)
	Populate()

	//Create request
	req, err := http.NewRequest("GET", "/event", nil)
	if err != nil {
		t.Fatal(err)
	}

	//Get Response
	res := httptest.NewRecorder()
	controllers.GetEvent(res, req)

	//Parse Response
	resEvents := []models.Event{}
	utils.ParseBodyTest(res, &resEvents)

	//Compare response to input
	if len(resEvents) != 1 {
		t.Errorf("Incorrect number of Events")
	} else {
		if resEvents[0].ID != 1 {
			t.Errorf("Incorrect ID")
		}
		if resEvents[0].Name != "TEST" {
			t.Errorf("Incorrect task name")
		}
		if resEvents[0].StartTime != "18-30" {
			t.Errorf("Incorrect start time")
		}
		if resEvents[0].EndTime != "20-00" {
			t.Errorf("Incorrect end time")
		}
		if resEvents[0].Color != "blue" {
			t.Errorf("Incorrect color")
		}
	}
}

func TestGetEventByID(t *testing.T) {

	// Call login
	Setup()
	Reset(t)
	Populate()

	//Create request
	req, err := http.NewRequest("GET", "/event/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(1, 10)
	req = mux.SetURLVars(req, vars)

	//Get Response
	res := httptest.NewRecorder()
	controllers.GetEventByID(res, req)

	//Parse Response
	resEvent := &models.Event{}
	utils.ParseBodyTest(res, resEvent)

	//Compare response to input
	if resEvent.ID != 1 {
		t.Errorf("Incorrect ID")
	}
	if resEvent.Name != "TEST" {
		t.Errorf("Incorrect task name")
	}
	if resEvent.StartTime != "18-30" {
		t.Errorf("Incorrect start time")
	}
	if resEvent.EndTime != "20-00" {
		t.Errorf("Incorrect end time")
	}
	if resEvent.Color != "blue" {
		t.Errorf("Incorrect color")
	}
}

func TestDeleteEvent(t *testing.T) {
	// Call login
	Setup()
	Reset(t)
	Populate()

	//Create request
	req, err := http.NewRequest("GET", "/event/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(1, 10)
	req = mux.SetURLVars(req, vars)

	//Get Response
	res := httptest.NewRecorder()
	controllers.DeleteEvent(res, req)

	//Parse Response
	resEvents := []models.Event{}
	utils.ParseBodyTest(res, &resEvents)

	//Check Response
	if len(resEvents) != 0 {
		t.Errorf("Incorrect number of Events")
	}
}

func TestCreateEvent(t *testing.T) {
	// Call login
	Setup()
	Reset(t)

	//Populate variable with an event
	var event models.Event
	event.ID = 1
	event.Name = "TEST"
	event.StartTime = "18-30"
	event.EndTime = "20-00"
	event.Color = "blue"
	eventDetails, _ := json.Marshal(event)

	//Create request
	req, err := http.NewRequest("POST", "/event/1", bytes.NewReader(eventDetails))
	if err != nil {
		t.Fatal(err)
	}

	//Set Request
	res := httptest.NewRecorder()
	controllers.CreateEvent(res, req)

	//Parse Response
	resEvent := &models.Event{}
	utils.ParseBodyTest(res, resEvent)

	//Compare response to input
	if resEvent.ID != 1 {
		t.Errorf("Incorrect ID")
	}
	if resEvent.Name != "TEST" {
		t.Errorf("Incorrect task name")
	}
	if resEvent.StartTime != "18-30" {
		t.Errorf("Incorrect start time")
	}
	if resEvent.EndTime != "20-00" {
		t.Errorf("Incorrect end time")
	}
	if resEvent.Color != "blue" {
		t.Errorf("Incorrect color")
	}
}

func TestUpdateEvent(t *testing.T) {

	var event models.Event

	// Call login
	Setup()
	Reset(t)
	Populate()

	//Populate variable with an event
	event.Name = "TEST3"
	event.StartTime = "17-30"
	event.EndTime = "21-00"
	event.Color = "orange"
	eventDetails, _ := json.Marshal(event)

	//Create request
	req, err := http.NewRequest("PUT", "/todo/1", bytes.NewReader(eventDetails))
	if err != nil {
		t.Fatal(err)
	}

	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(1, 10)
	req = mux.SetURLVars(req, vars)

	//Send Request
	res := httptest.NewRecorder()
	controllers.UpdateEvent(res, req)

	//Parse Response
	resEvent := &models.Event{}
	utils.ParseBodyTest(res, resEvent)

	//Compare response to input
	if resEvent.ID != 1 {
		t.Errorf("Incorrect ID")
	}
	if resEvent.Name != "TEST3" {
		t.Errorf("Incorrect task name")
	}
	if resEvent.StartTime != "17-30" {
		t.Errorf("Incorrect start time")
	}
	if resEvent.EndTime != "21-00" {
		t.Errorf("Incorrect end time")
	}
	if resEvent.Color != "orange" {
		t.Errorf("Incorrect color")
	}
}

func Setup() {
	models.LoginAdmin()
	db = config.GetDB()
}

func Reset(t *testing.T) {
	statement, err := db.Prepare(`
	DELETE FROM event WHERE userid=0;
	`)
	if err != nil {
		log.Println("FAILED TO PREPARE RESET STATEMENT")
		log.Println(err)
	}
	statement.Exec()
	statement.Close()
}

func Populate() {
	statement, _ := db.Prepare(`
	INSERT into event(id, name, start, end, color, userid) VALUES (?, ?, ?, ?, ?, ?)
	`)
	statement.Exec(1, "TEST", "18-30", "20-00", "blue", 0)
	statement.Close()
}
