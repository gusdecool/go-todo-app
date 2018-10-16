package task

import (
	"github.com/gusdecool/go-todo-app/db/connector"
	"github.com/gusdecool/go-todo-app/db/model"
)

func FindAll() ([]model.Task, error) {
	db, err := connector.Connect()

	defer db.Close()
	var tasks []model.Task

	if err != nil {
		return tasks, err
	}

	db.Find(&tasks)

	return tasks, nil
}