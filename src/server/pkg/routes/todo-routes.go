package routes

import (
	"github.com/gorilla/mux"
	"github.com/kchen53/pomodoro_planner/pkg/controllers"
)

// var RegisterBookStoreRoutes = func(router *mux.Router) {
func RegisterToDoStoreRoutes(router *mux.Router) {
	router.HandleFunc("/todo", controllers.CreateToDo).Methods("POST")
	router.HandleFunc("/todo", controllers.GetToDo).Methods("GET")
	router.HandleFunc("/todo/{i}", controllers.GetToDoByID).Methods("GET")
	router.HandleFunc("/todo/{i}", controllers.UpdateToDo).Methods("PUT")
	router.HandleFunc("/todo/{i}", controllers.DeleteToDo).Methods("DELETE")
}
