package models

import (
	"database/sql"
	"log"

	"github.com/kchen53/pomodoro_planner/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Todo struct {
	Name     string `json:"name"`
	Date     string `json:"date"` //YYYY-MM-DD
	Time     int    `json:time`   //seconds
	Repeat   int    `json:repeat` //binaryflags: 0:6 = MTWRFSN
	Complete bool   `json:"complete"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	config.CreateTable(`
	CREATE TABLE todo (
		"name" TEXT NOT NULL PRIMARY KEY AUTOINCREMENT,
		"date" TEXT NOT NULL,
		"time" integer,
		"repeat" integer NOT NULL,
		"complete" integer NOT NULL
	);
	`)
}

//1:43:03
func (t *Todo) CreateToDo() *Todo {
	log.Println("Inserting", t.Name, "...")
	statement, err := db.Prepare(`
	INSERT INTO todo(name, date, time, repeat, complete) VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Println("Failed to insert", t.Name)
		return t
	}
	defer statement.Close()
	_, err = statement.Exec(t.Name, t.Date, t.Time, t.Repeat, t.Complete)
	if err != nil {
		log.Println("Failed to insert", t.Name)
		return t
	}
	statement.Close()
	return t
}

func GetAllTodo() []Todo {
	var Todos []Todo
	rows, err := db.Query(`
	SELECT *
	FROM todo
	ORDER BY id;
	`)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.Name, &t.Date, &t.Time, &t.Repeat, &t.Complete); err != nil {
			log.Println(err)
			return nil
		}
		Todos = append(Todos, t)
	}
	log.Println("Queried", len(Todos), "Todos")
	return Todos
}

func GetTodoByID(Id int64) (*Todo, *sql.DB) {
	var getTodo Todo
	//db := db.Where("ID=?", Id).Find(&getToDo)
	// for _, t := range list {
	// 	if t.Name == Id {
	// 		getToDo = t
	// 	}
	// }
	return &getTodo, db
}

func GetTodoByIDList(Id int64) (*Todo, *[]Todo) {
	var getTodo Todo
	// for _, t := range list {
	// 	if t.ID == Id {
	// 		getToDo = t
	// 	}
	// }
	return &getTodo, &list
}

func DeleteToDo(Id int64) Todo {
	var todo Todo
	//db.Where("ID=?", Id).Delete(todo)
	// for i, t := range list {
	// 	if t.ID == Id {
	// 		list = append(list[:i], list[i+1:]...)
	// 		break
	// 	}
	// }
	return todo
}

func GetList() []Todo {
	return list
}
