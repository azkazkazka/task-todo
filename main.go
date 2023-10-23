package main

import (
	"github.com/azkazkazka/task-todo/app"
	"github.com/azkazkazka/task-todo/routes"
)

func main() {
	app := app.NewApp()
	e := routes.Init(app.DB)

	e.Logger.Fatal(e.Start(":8080"))
}
