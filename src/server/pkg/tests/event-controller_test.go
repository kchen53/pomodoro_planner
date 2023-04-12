package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kchen53/pomodoro_planner/pkg/config"
	"github.com/kchen53/pomodoro_planner/pkg/controllers"
	"github.com/kchen53/pomodoro_planner/pkg/models"
	"github.com/kchen53/pomodoro_planner/pkg/utils"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

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
		if resEvents[0].Date != "01-01-2023" {
			t.Errorf("Incorrect date")
		}
		if resEvents[0].StartTime != 18-30 {
			t.Errorf("Incorrect time")
		}
		if resEvents[0].EndTime != 20-00 {
			t.Errorf("Incorrect repeat")
		}
		if resEvents[0].Repeat != 0 {
			t.Errorf("Incorrect completed state")
		}
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
	INSERT into event(id, name, date, starttime, endtime, repeat, userid) VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	statement.Exec(1, "TEST", "01-01-2023", 18-30, 20-0, 0, 0)
	statement.Close()
}
