package routes

import (
	"github.com/gorilla/mux"
	"github.com/kchen53/pomodoro_planner/pkg/controllers"
)

func RegisterEventRoutes(router *mux.Router) {
	router.HandleFunc("/event", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/event", controllers.GetEvent).Methods("GET")
	router.HandleFunc("/event/{i}", controllers.GetEventByID).Methods("GET")
	router.HandleFunc("/event/{i}", controllers.UpdateEvent).Methods("PUT")
	router.HandleFunc("/event/{i}", controllers.DeleteEvent).Methods("DELETE")
}
