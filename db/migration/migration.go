package migration

import (
	"github.com/gusdecool/go-todo-app/db/connector"
	"github.com/gusdecool/go-todo-app/db/model"
)

func Migrate() error {
	db, err := connector.Connect()
	defer db.Close()

	if err != nil {
		return err
	}

	db.AutoMigrate(&model.Task{})

	return nil
}