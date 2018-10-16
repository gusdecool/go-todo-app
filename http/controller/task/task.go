package task

import (
	"encoding/json"
	"github.com/gusdecool/go-todo-app/db/model"
	"github.com/gusdecool/go-todo-app/db/repository/task"
	"github.com/gusdecool/go-todo-app/http/utility"
	"net/http"
	"time"
)

func List(response http.ResponseWriter, request *http.Request) {
	tasks, err := task.FindAll()

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	tasksByte, err := json.Marshal(tasks)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(tasksByte)
}

func Create(response http.ResponseWriter, request *http.Request) {
	var taskModel model.Task
	taskModel.CreatedAt = time.Now()

	err := json.NewDecoder(request.Body).Decode(&taskModel)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	_, err = task.Create(&taskModel)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	taskByte, err := json.Marshal(taskModel)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(taskByte)
}