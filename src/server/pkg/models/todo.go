package models

import (
	"database/sql"
	"log"

	"github.com/kchen53/pomodoro_planner/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Todo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`   //YYYY-MM-DD
	Time     int    `json:"time"`   //seconds
	Repeat   int    `json:"repeat"` //binaryflags: 0:6 = MTWRFSN
	Complete bool   `json:"complete"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	Login(User{0, "Admin", "test"})
	// config.CreateTable(`
	// CREATE TABLE todo (
	// 	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	// 	"name" TEXT NOT NULL,
	// 	"date" TEXT NOT NULL,
	// 	"time" integer,
	// 	"repeat" integer NOT NULL,
	// 	"complete" integer NOT NULL
	// );
	// `)
}

//1:43:03
func (t *Todo) CreateTodo() *Todo {
	log.Println("Inserting", t.Name, "...")
	statement, err := db.Prepare(`
	INSERT INTO todo(name, date, time, repeat, complete, userid) VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Println("Insertion: Failed to prepare statement")
		log.Println(err)
		return t
	}
	defer statement.Close()
	_, err = statement.Exec(t.Name, t.Date, t.Time, t.Repeat, t.Complete, currentUser)
	if err != nil {
		log.Println("Insertion: Failed to execute statement")
		log.Println(err)
		return t
	}
	statement.Close()
	return t
}

func GetAllTodo() []Todo {
	Todos := make([]Todo, 0)
	log.Println("Getting all todos...")
	rows, err := db.Query(`
	SELECT *
	FROM todo
	ORDER BY id;
	`)
	if err != nil {
		log.Println("Query: Failed to execute query")
		log.Println(err)
		return nil
	}
	defer rows.Close()

	ignore := 0
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Name, &t.Date, &t.Time, &t.Repeat, &t.Complete, &ignore); err != nil {
			log.Println("Query: Failed to read query")
			log.Println(err)
			return nil
		}
		Todos = append(Todos, t)
	}
	log.Println("Queried", len(Todos), "Todos")
	//Todos = append(Todos, Todo{"Test", "0000-00-00", 0, 0, false})
	return Todos
}

func GetTodoByID(id int) Todo {
	var getTodo Todo
	log.Println("Getting todo", id, "...")
	rows, err := db.Query(`
	SELECT *
	FROM todo
	WHERE id=?;
	`, id)
	if err != nil {
		log.Println("Query: Failed to execute query")
		log.Println(err)
		return getTodo
	}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&getTodo.ID, &getTodo.Name, &getTodo.Date, &getTodo.Time, &getTodo.Repeat, &getTodo.Complete, nil); err != nil {
		log.Println("Query: Failed to read query")
		log.Println(err)
		return getTodo
	}
	return getTodo
}

func DeleteTodo(id int) Todo {
	var todo Todo
	log.Println("Deleting todo", id, "...")
	statement, err := db.Prepare(`
	DELETE FROM todo WHERE id=?;
	`)
	if err != nil {
		log.Println("Deletion: failed to prepare statement")
		log.Println(err)
		return todo
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		log.Println("Deletion: failed to execute statement")
		log.Println(err)
		return todo
	}
	todo.ID = id
	return todo
}

func (t *Todo) UpdateTodo() *Todo {
	log.Println("Updating todo", t.ID, "...")
	statement, err := db.Prepare(`
	UPDATE todo
	SET name = ?, date = ?, time = ?, repeat = ?, complete = ?
	WHERE id=?;
	`)
	if err != nil {
		log.Println("Update: Failed to prepare statement")
		log.Println(err)
		return t
	}
	defer statement.Close()
	_, err = statement.Exec(t.Name, t.Date, t.Time, t.Repeat, t.Complete, t.ID)
	if err != nil {
		log.Println("Update: Failed to execute statement")
		log.Println(err)
		return t
	}
	return t
}
