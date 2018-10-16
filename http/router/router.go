package router

import (
	"github.com/gorilla/mux"
	"github.com/gusdecool/go-todo-app/http/controller/task"
	"log"
	"net/http"
)

func Register() {
	router := mux.NewRouter()

	registerTask(router)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

func registerTask(router *mux.Router) {
	router.HandleFunc("/task", task.List).Methods("GET")
	router.HandleFunc("/task", task.Create).Methods("POST")
}
