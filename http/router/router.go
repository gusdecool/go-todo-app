package router

import (
	"github.com/gorilla/mux"
	"github.com/gusdecool/go-todo-app/http/controller/task"
	"net/http"
)

func Register() {
	router := mux.NewRouter()

	router.HandleFunc("/task", task.List).Methods("GET")

	http.ListenAndServe(":80", router)
}
