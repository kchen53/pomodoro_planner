package routes

import (
	"github.com/gorilla/mux"
	"github.com/kchen53/pomodoro_planner/pkg/controllers"
)

// var RegisterBookStoreRoutes = func(router *mux.Router) {
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.LoginUser).Methods("PUT")
	router.HandleFunc("/user", controllers.SignupUser).Methods("POST")
	router.HandleFunc("/user", controllers.LogoutUser).Methods("DELETE")
	router.HandleFunc("/user", controllers.GetUsername).Methods("GET")
}
