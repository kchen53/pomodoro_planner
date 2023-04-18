package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"bytes"
	"encoding/json"

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
	log.Println("here!")
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
		if resEvents[0].Date != "01-01-2023" {
			t.Errorf("Incorrect date")
		}
		if resEvents[0].StartTime != "18-30" {
			t.Errorf("Incorrect time")
		}
		if resEvents[0].EndTime != "20-00" {
			t.Errorf("Incorrect repeat")
		}
		if resEvents[0].Repeat != 0 {
			t.Errorf("Incorrect completed state")
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
	if resEvent.Date != "01-01-2023" {
		t.Errorf("Incorrect date")
	}
	if resEvent.StartTime != "18-30" {
		t.Errorf("Incorrect time")
	}
	if resEvent.EndTime != "20-00" {
		t.Errorf("Incorrect time")
	}
	if resEvent.Repeat != 0 {
		t.Errorf("Incorrect repeat")
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

	//Populate array db with a todo
	var event models.Event
	event.ID = 1
	event.Name = "TEST"
	event.Date = "01-01-2023"
	event.StartTime = "18-30"
	event.EndTime = "20-00"
	event.Repeat = 0
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
	if resEvent.Date != "01-01-2023" {
		t.Errorf("Incorrect date")
	}
	if resEvent.StartTime != "18-30" {
		t.Errorf("Incorrect time")
	}
	if resEvent.EndTime != "20-00" {
		t.Errorf("Incorrect repeat")
	}
	if resEvent.Repeat != 0 {
		t.Errorf("Incorrect completed state")
	}
}

func Setup() {
	models.LoginAdmin()
	db = config.GetDB()
}

func Reset(t *testing.T) {
	statement, _ := db.Prepare(`
	DELETE FROM event WHERE userid=0;
	`)
	statement.Exec()
	statement.Close()
}

func Populate() {
	statement, _ := db.Prepare(`
	INSERT into event(id, name, date, start_time, end_time, repeat, userid) VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	statement.Exec(1, "TEST", "01-01-2023", "18-30", "20-00", 0, 0)
	statement.Close()
}
