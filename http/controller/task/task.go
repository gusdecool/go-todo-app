package task

import (
	"fmt"
	"github.com/gusdecool/go-todo-app/db/repository/task"
	"net/http"
)

func List(response http.ResponseWriter, request *http.Request) {
	tasks, err := task.FindAll()

	if err != nil {
		response.Write([]byte(err.Error()))
	}

	fmt.Println(tasks)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	//response.Write(json.Marshal(tasks))
}