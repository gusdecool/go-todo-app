package main

import (
	"flag"
	"fmt"
	"github.com/gusdecool/go-todo-app/db/migration"
)

func main() {
	runMigration := flag.Bool("run-migration", false, "run migration? default: false")
	flag.Parse()

	if *runMigration == true {
		err := migration.Migrate()

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//router.Register()
}