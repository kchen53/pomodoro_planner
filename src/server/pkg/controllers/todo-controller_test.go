package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kchen53/pomodoro_planner/pkg/models"
	"github.com/kchen53/pomodoro_planner/pkg/utils"
)

// func Test(t *testing.T) {
// 	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res := httptest.NewRecorder()
// 	GetToDo(res, req)

// 	exp := "Hello World"
// 	act := res.Body.String()
// 	if exp != act {
// 		t.Fatalf("Expected %s gog %s", exp, act)
// 	}
// }

func TestGetToDo(t *testing.T) {
	//Populate array db with a todo
	var todo models.ToDo
	todo.ID = 1
	todo.Task = "Test Name"
	todo.Due = "Test Day"
	todo.Complete = true
	todo.CreateToDo()

	//Create request
	req, err := http.NewRequest("GET", "/todo", nil)
	if err != nil {
		t.Fatal(err)
	}

	//Get Response
	res := httptest.NewRecorder()
	GetToDo(res, req)

	//Parse Response
	resTodos := []models.ToDo{}
	utils.ParseBodyTest(res, &resTodos)

	//Compare response to input
	if len(resTodos) != 1 {
		t.Errorf("Incorrect number of Todo's")
	}
	if resTodos[0].ID != todo.ID {
		t.Errorf("Incorrect ID")
	}
	if resTodos[0].Task != todo.Task {
		t.Errorf("Incorrect task name")
	}
	if resTodos[0].Due != todo.Due {
		t.Errorf("Incorrect due date")
	}
	if resTodos[0].Complete != todo.Complete {
		t.Errorf("Incorrect completed state")
	}
}

func TestGetToDoByID(t *testing.T) {
	//Populate array db with a todo
	var todo models.ToDo
	todo.ID = 1
	todo.Task = "Test Name"
	todo.Due = "Test Day"
	todo.Complete = true
	todo.CreateToDo()

	//Create request
	req, err := http.NewRequest("GET", "/todo/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(todo.ID, 10)
	req = mux.SetURLVars(req, vars)

	//Get Response
	res := httptest.NewRecorder()
	GetToDoByID(res, req)

	//Parse Response
	resTodo := &models.ToDo{}
	utils.ParseBodyTest(res, resTodo)

	//Compare response to input
	if resTodo.ID != todo.ID {
		t.Errorf("Incorrect ID")
	}
	if resTodo.Task != todo.Task {
		t.Errorf("Incorrect task name")
	}
	if resTodo.Due != todo.Due {
		t.Errorf("Incorrect due date")
	}
	if resTodo.Complete != todo.Complete {
		t.Errorf("Incorrect completed state")
	}
}

func TestDeleteToDo(t *testing.T) {
	//Populate array db with a todo
	var todo models.ToDo
	todo.ID = 1
	todo.Task = "Test Name"
	todo.Due = "Test Day"
	todo.Complete = true
	todo.CreateToDo()

	//Create request
	req, err := http.NewRequest("GET", "/todo/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(todo.ID, 10)
	req = mux.SetURLVars(req, vars)

	//Get Response
	res := httptest.NewRecorder()
	GetToDoByID(res, req)

	//Parse Response
	resTodos := []models.ToDo{}
	utils.ParseBodyTest(res, &resTodos)

	//Check Response
	if len(resTodos) != 0 {
		t.Errorf("Incorrect number of Todo's")
	}
}

func TestCreateToDo(t *testing.T) {
	//Populate array db with a todo
	var todo models.ToDo
	todo.ID = 1
	todo.Task = "Test Name"
	todo.Due = "Test Day"
	todo.Complete = true
	todoDetails, _ := json.Marshal(todo)

	//Create request
	req, err := http.NewRequest("POST", "/todo", bytes.NewReader(todoDetails))
	if err != nil {
		t.Fatal(err)
	}

	//Set Request
	res := httptest.NewRecorder()
	CreateToDo(res, req)

	//Check Models.list for todo
	if !itemInList(todo.ID, models.GetList()) {
		t.Errorf("Not in List")
	}
}

/*func TestUpdateToDo(t *testing.T) {
	//Populate array db with a todo
	var todo models.ToDo
	todo.ID = 1
	todo.Task = "Test Name"
	todo.Due = "Test Day"
	todo.Complete = true
	todo.CreateToDo()

	//Populate array db with a todo
	todo.Task = "Test Name 1"
	todo.Due = "Test Day 2"
	todo.Complete = false
	todoDetails, _ := json.Marshal(todo)

	//Create request
	req, err := http.NewRequest("PUT", "/todo/1", bytes.NewReader(todoDetails))
	if err != nil {
		t.Fatal(err)
	}

	//Set variables
	vars := make(map[string]string)
	vars["i"] = strconv.FormatInt(todo.ID, 10)
	req = mux.SetURLVars(req, vars)

	//Send Request
	res := httptest.NewRecorder()
	UpdateToDo(res, req)

	//Check Models.list for todo
	for _, b := range models.GetList() {
		if b.ID == todo.ID {
			if b.Task != todo.Task {
				fmt.Println(b.Task)
				t.Errorf("Wrong Task")
			}
			if b.Due != todo.Due {
				fmt.Println(b.Due)
				t.Errorf("Wrong Due")
			}
			if b.Complete != todo.Complete {
				fmt.Println(b.Complete)
				t.Errorf("Wrong Complete")
			}
		}
	}
}*/

func itemInList(a int64, list []models.ToDo) bool {
	for _, b := range list {
		if b.ID == a {
			return true
		}
	}
	return false
}
