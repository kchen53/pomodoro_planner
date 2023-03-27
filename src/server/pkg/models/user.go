package models

import (
	_ "github.com/mattn/go-sqlite3"
)

var currentUser int

type User struct {
	ID       int
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(u User) bool {
	//Check credintials - Query users

	//Set current User if pass
	currentUser = u.ID
	return true
}
