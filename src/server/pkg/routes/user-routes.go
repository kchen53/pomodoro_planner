package routes

import (
	"github.com/gorilla/mux"
	"github.com/kchen53/pomodoro_planner/pkg/controllers"
)

// var RegisterBookStoreRoutes = func(router *mux.Router) {
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.LoginUser).Methods("PUT")
	router.HandleFunc("/signup", controllers.SignupUser).Methods("POST")
	router.HandleFunc("/signout", controllers.LogoutUser).Methods("PUT")
	router.HandleFunc("/userdata", controllers.GetUsername).Methods("GET")
}
