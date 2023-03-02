package controllers

import (
	"github.com/kchen53/pomodoro_planner/pkg/models"
	"github.com/kchen53/pomodoro_planner/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"
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
	req, err := http.NewRequest("POST", "/todo", nil)
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
