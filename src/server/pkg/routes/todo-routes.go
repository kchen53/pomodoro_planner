package routes

import (
	"github.com/gorilla/mux"
	"github.com/kchen53/pomodoro_planner/pkg/controllers"
)

// var RegisterBookStoreRoutes = func(router *mux.Router) {
func RegisterTodoRoutes(router *mux.Router) {
	router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{i}", controllers.GetTodoByID).Methods("GET")
	router.HandleFunc("/todo/{i}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{i}", controllers.DeleteTodo).Methods("DELETE")
}
