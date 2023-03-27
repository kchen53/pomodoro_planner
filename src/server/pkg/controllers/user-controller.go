package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kchen53/pomodoro_planner/pkg/models"
	"github.com/kchen53/pomodoro_planner/pkg/utils"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	LoginUser := &models.User{}
	utils.ParseBody(r, LoginUser)
	success := LoginUser.Login()
	res, _ := json.Marshal(success)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SignupUser(w http.ResponseWriter, r *http.Request) {
	NewUser := &models.User{}
	utils.ParseBody(r, NewUser)
	success := NewUser.Signup()
	res, _ := json.Marshal(success)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	success := models.Logout()
	res, _ := json.Marshal(success)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUsername(w http.ResponseWriter, r *http.Request) {
	username := models.GetUsername()
	res, _ := json.Marshal(username)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
