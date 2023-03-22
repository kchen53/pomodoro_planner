package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func Connect() {

	log.Println("Connecting to sqlite database...")
	d, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		// log.Println("Creating sqlite-database.db...")
		// file, err := os.Create("sqlite-database.db")
		// if err != nil {
		// 	//log.Fatal(err.Error())
		// 	panic(err)
		// }
		// file.Close()
		// log.Println("sqlite-database.db created")
		// d, err := sql.Open("sqlite3", "./sqlite-database.db")
		// if err != nil {
		// 	//log.Fatal(err.Error())
		// 	panic(err)
		// }
		// db = d
		panic(err)
	} else {
		db = d
	}

}

func GetDB() *sql.DB {
	return db
}

func createTable(table string) {
	log.Println("Creating table...")
	statement, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Table created")
}

//TODO: INSERT
//TODO: GET
//TODO: DELETE
