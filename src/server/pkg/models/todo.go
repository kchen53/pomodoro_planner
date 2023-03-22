package models

import (
	"database/sql"

	"github.com/kchen53/pomodoro_planner/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var list []ToDo

type ToDo struct {
	Name     string `json:"name"`
	Date     string `json:"date"` //YYYY-MM-DD
	Time     int    `json:time`   //seconds
	Repeat   int8   `json:repeat` //binaryflags: 0:6 = MTWRFSN
	Complete bool   `json:"complete"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	config.CreateTable(`
	CREATE TABLE TODO (
		"name" TEXT NOT NULL PRIMARY KEY AUTOINCREMENT,
		"date" TEXT,
		"time" integer,
		"repeat" integer,
		"complete" integer
	);
	`)
}

//1:43:03
func (t *ToDo) CreateToDo() *ToDo {
	//TODO: Insert
	list = append(list, *t)
	return t
}

func GetAllToDo() []ToDo {
	var ToDos []ToDo
	//db.Find(&ToDos)
	ToDos = list
	return ToDos
}

func GetToDoByID(Id int64) (*ToDo, *gorm.DB) {
	var getToDo ToDo
	//db := db.Where("ID=?", Id).Find(&getToDo)
	// for _, t := range list {
	// 	if t.Name == Id {
	// 		getToDo = t
	// 	}
	// }
	return &getToDo, db
}

func GetToDoByIDList(Id int64) (*ToDo, *[]ToDo) {
	var getToDo ToDo
	// for _, t := range list {
	// 	if t.ID == Id {
	// 		getToDo = t
	// 	}
	// }
	return &getToDo, &list
}

func DeleteToDo(Id int64) ToDo {
	var todo ToDo
	//db.Where("ID=?", Id).Delete(todo)
	// for i, t := range list {
	// 	if t.ID == Id {
	// 		list = append(list[:i], list[i+1:]...)
	// 		break
	// 	}
	// }
	return todo
}

func GetList() []ToDo {
	return list
}
