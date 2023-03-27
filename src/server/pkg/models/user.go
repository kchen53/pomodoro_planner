package models

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var currentUser int = 1
var loggedIn bool = false

type User struct {
	ID       int
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Login() bool {
	log.Println("Loggin in", u.Username, "...")
	//Check credintials - Query users
	rows, err := db.Query(`
	SELECT userid
	FROM users
	WHERE username = ? AND password = ?;
	`, u.Username, u.Password)
	if err != nil {
		log.Println("Login: Failed to execute query")
		log.Println(err)
		return false
	}
	defer rows.Close()
	if rows.Next() {
		//User exists => set credentials
		if err = rows.Scan(&u.ID); err != nil {
			log.Println("Login: Failed to read query")
			log.Println(err)
			return false
		}
		currentUser = u.ID
		loggedIn = true
		return true
	}
	log.Println("Login: User does not exist")
	return false
}

func (u *User) Signup() bool {
	log.Println("Signing up", u.Username, "...")
	//Check if user already exists
	rows, err := db.Query(`
	SELECT username
	FROM users
	WHERE username = ?;
	`, u.Username)
	if err != nil {
		log.Println("Signup: Failed to execute query")
		log.Println(err)
		return false
	}
	defer rows.Close()
	if rows.Next() {
		//User already exists => signup fails
		log.Println("Signup: User already exists")
		return false
	}
	//Create user
	statement, err := db.Prepare(`
	INSERT INTO users(username, password) VALUES (?, ?)
	`)
	if err != nil {
		log.Println("Signup: Failed to prepare statement")
		log.Println(err)
		return false
	}
	defer statement.Close()
	result, err := statement.Exec(u.Username, u.Password)
	if err != nil {
		log.Println("Signup: Failed to execute statement")
		log.Println(err)
		return false
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Signup: Failed to read result")
		log.Println(err)
		return false
	}
	log.Println(id)
	currentUser = int(id)
	loggedIn = true
	return true
}

func Logout() bool {
	log.Println("Signing out...")
	currentUser = 1
	loggedIn = false
	return true
}

func GetUserId() int {
	if loggedIn {
		return currentUser
	}
	return 1
}

func SetUserIDTest() {
	currentUser = 0
	loggedIn = true
}

func GetUsername() string {
	log.Println("Retrieving username...")
	rows, err := db.Query(`
	SELECT username
	FROM users
	WHERE userid = ?;
	`, GetUserId())
	if err != nil {
		log.Println("Username: Failed to execute query")
		log.Println(err)
		return ""
	}
	defer rows.Close()
	rows.Next()
	var username string
	if err = rows.Scan(&username); err != nil {
		log.Println("Username: Failed to read query")
		log.Println(err)
		return ""
	}
	return username
}
