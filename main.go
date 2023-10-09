package main

import (
	"github.com/azkazkazka/task-todo/db"
	"github.com/azkazkazka/task-todo/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
