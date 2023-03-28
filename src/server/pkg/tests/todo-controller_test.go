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

/*func TestGetTodoByID(t *testing.T) {
	//Populate array db with a todo
	var todo models.Todo
	todo.ID = 1
	todo.Task = "Test Name"
	todo.Due = "Test Day"
	todo.Complete = true
	todo.CreateTodo()

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
	GetTodoByID(res, req)

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

func TestDeleteTodo(t *testing.T) {
	//Populate array db with a todo
	var todo models.Todo
	todo.ID = 1
	todo.Task = "Test Name"
	todo.Due = "Test Day"
	todo.Complete = true
	todo.CreateTodo()

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
	GetTodoByID(res, req)

	//Parse Response
	resTodos := []models.Todo{}
	utils.ParseBodyTest(res, &resTodos)

	//Check Response
	if len(resTodos) != 0 {
		t.Errorf("Incorrect number of Todo's")
	}
}

func TestCreateTodo(t *testing.T) {
	//Populate array db with a todo
	var todo models.Todo
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
	CreateTodo(res, req)

	//Check Models.list for todo
	if !itemInList(todo.ID, models.GetList()) {
		t.Errorf("Not in List")
	}
}

func TestUpdateTodo(t *testing.T) {
	//Populate array db with a todo
	var todo models.Todo
	todo.ID = 1
	todo.Task = "Test Name"
	todo.Due = "Test Day"
	todo.Complete = true
	todo.CreateTodo()

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
	UpdateTodo(res, req)

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

// func Test(t *testing.T) {
// 	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res := httptest.NewRecorder()
// 	GetTodo(res, req)

// 	exp := "Hello World"
// 	act := res.Body.String()
// 	if exp != act {
// 		t.Fatalf("Expected %s gog %s", exp, act)
// 	}
// }
