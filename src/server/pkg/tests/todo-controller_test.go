package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

var db *sql.DB

func TestGetTodo(t *testing.T) {

	// Call login
	TESTsetup()
	TESTreset(t)
	TESTpopulate()

	//Create request
	req, err := http.NewRequest("GET", "/todo", nil)
	if err != nil {
		t.Fatal(err)
	}

	//Get Response
	res := httptest.NewRecorder()
	controllers.GetTodo(res, req)

	//Parse Response
	resTodos := []models.Todo{}
	utils.ParseBodyTest(res, &resTodos)

	//Compare response to input
	if len(resTodos) != 1 {
		t.Errorf("Incorrect number of Todo's")
	} else {
		if resTodos[0].ID != 1 {
			t.Errorf("Incorrect ID")
		}
		if resTodos[0].Name != "TEST" {
			t.Errorf("Incorrect task name")
		}
		if resTodos[0].Date != "01-01-2023" {
			t.Errorf("Incorrect date")
		}
		if resTodos[0].Time != 60 {
			t.Errorf("Incorrect time")
		}
		if resTodos[0].Repeat != 0 {
			t.Errorf("Incorrect repeat")
		}
		if resTodos[0].Complete != true {
			t.Errorf("Incorrect completed state")
		}
	}
}

func TestGetTodoByID(t *testing.T) {

	// Call login
	TESTsetup()
	TESTreset(t)
	TESTpopulate()

	//Create request
	req, err := http.NewRequest("GET", "/todo/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(1, 10)
	req = mux.SetURLVars(req, vars)

	//Get Response
	res := httptest.NewRecorder()
	controllers.GetTodoByID(res, req)

	//Parse Response
	resTodo := &models.Todo{}
	utils.ParseBodyTest(res, resTodo)

	//Compare response to input
	if resTodo.ID != 1 {
		t.Errorf("Incorrect ID")
	}
	if resTodo.Name != "TEST" {
		t.Errorf("Incorrect task name")
	}
	if resTodo.Date != "01-01-2023" {
		t.Errorf("Incorrect date")
	}
	if resTodo.Time != 60 {
		t.Errorf("Incorrect time")
	}
	if resTodo.Repeat != 0 {
		t.Errorf("Incorrect repeat")
	}
	if resTodo.Complete != true {
		t.Errorf("Incorrect completed state")
	}
}

func TestDeleteTodo(t *testing.T) {
	// Call login
	TESTsetup()
	TESTreset(t)
	TESTpopulate()

	//Create request
	req, err := http.NewRequest("GET", "/todo/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(1, 10)
	req = mux.SetURLVars(req, vars)

	//Get Response
	res := httptest.NewRecorder()
	controllers.DeleteTodo(res, req)

	//Parse Response
	resTodos := []models.Todo{}
	utils.ParseBodyTest(res, &resTodos)

	//Check Response
	if len(resTodos) != 0 {
		t.Errorf("Incorrect number of Todo's")
	}
}

func TestCreateTodo(t *testing.T) {
	// Call login
	TESTsetup()
	TESTreset(t)

	//Populate array db with a todo
	var todo models.Todo
	todo.ID = 1
	todo.Name = "TEST"
	todo.Date = "01-01-2023"
	todo.Time = 60
	todo.Repeat = 0
	todo.Complete = true
	todoDetails, _ := json.Marshal(todo)

	//Create request
	req, err := http.NewRequest("POST", "/todo/1", bytes.NewReader(todoDetails))
	if err != nil {
		t.Fatal(err)
	}

	//Set Request
	res := httptest.NewRecorder()
	controllers.CreateTodo(res, req)

	//Parse Response
	resTodo := &models.Todo{}
	utils.ParseBodyTest(res, resTodo)

	//Compare response to input
	if resTodo.ID != 1 {
		t.Errorf("Incorrect ID")
	}
	if resTodo.Name != "TEST" {
		t.Errorf("Incorrect task name")
	}
	if resTodo.Date != "01-01-2023" {
		t.Errorf("Incorrect date")
	}
	if resTodo.Time != 60 {
		t.Errorf("Incorrect time")
	}
	if resTodo.Repeat != 0 {
		t.Errorf("Incorrect repeat")
	}
	if resTodo.Complete != true {
		t.Errorf("Incorrect completed state")
	}
}

func TestUpdateTodo(t *testing.T) {

	var todo models.Todo

	// Call login
	TESTsetup()
	TESTreset(t)
	TESTpopulate()

	//Populate array db with a todo
	todo.Name = "TEST"
	todo.Date = "01-02-2023"
	todo.Time = 50
	todo.Repeat = 1
	todo.Complete = false
	todoDetails, _ := json.Marshal(todo)

	//Create request
	req, err := http.NewRequest("PUT", "/todo/1", bytes.NewReader(todoDetails))
	if err != nil {
		t.Fatal(err)
	}

	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(1, 10)
	req = mux.SetURLVars(req, vars)

	//Send Request
	res := httptest.NewRecorder()
	controllers.UpdateTodo(res, req)

	//Parse Response
	resTodo := &models.Todo{}
	utils.ParseBodyTest(res, resTodo)

	//Compare response to input
	if resTodo.ID != 1 {
		t.Errorf("Incorrect ID")
	}
	if resTodo.Name != "TEST" {
		t.Errorf("Incorrect task name")
	}
	if resTodo.Date != "01-02-2023" {
		t.Errorf("Incorrect date")
	}
	if resTodo.Time != 50 {
		t.Errorf("Incorrect time")
	}
	if resTodo.Repeat != 1 {
		t.Errorf("Incorrect repeat")
	}
	if resTodo.Complete != false {
		t.Errorf("Incorrect completed state")
	}
}

func TESTsetup() {
	models.LoginAdmin()
	db = config.GetDB()
}

func TESTreset(t *testing.T) {
	statement, _ := db.Prepare(`
	DELETE FROM todo WHERE userid=0;
	`)
	statement.Exec()
	statement.Close()
}

func TESTpopulate() {
	statement, _ := db.Prepare(`
	INSERT into todo(id, name, date, time, repeat, complete, userid) VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	statement.Exec(1, "TEST", "01-01-2023", 60, 0, true, 0)
	statement.Close()
}
