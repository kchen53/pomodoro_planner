package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func Connect() {
	// fmt.Println("Connecting to mySQL database...")
	// d, err := gorm.Open("mysql", "root:yMm2nI217ABT3jSr@/simpleresttest?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }
	// db = d

	log.Println("Connecting to sqlite database...")
	d, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Println("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db")
		if err != nil {
			//log.Fatal(err.Error())
			panic(err)
		}
		file.Close()
		log.Println("sqlite-database.db created")
		d, err := sql.Open("sqlite3", "./sqlite-database.db")
		if err != nil {
			//log.Fatal(err.Error())
			panic(err)
		}
		db = d
	} else {
		db = d
	}

}

func GetDB() *sql.DB {
	return db
}

//TODO: CREATE TABLE
//TODO: INSERT
//TODO: GET
//TODO: DELETE
