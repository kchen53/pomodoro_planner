package main

import (
	"fmt"
	"log"
	"net/http"

	//"os"
	"pomodoro_planner/src/server/utils"
	//"github.com/gorilla/mux"
)

func main() {
	// r := mux.NewRouter()

	// r.HandleFunc("/hello-world", helloWorld)

	// http.Handle("/", r)

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    ":" + os.Getenv("PORT"),
	// }

	// log.Fatal(srv.ListenAndServe())

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloWorld)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}
