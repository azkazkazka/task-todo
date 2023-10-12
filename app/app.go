package app

import (
	"gorm.io/gorm"

	"github.com/azkazkazka/task-todo/db"
)

type App struct {
	DB *gorm.DB
}

func NewApp() *App {
	db.Init()
	dbApp := db.CreateCon()
	return &App{
		DB: dbApp,
	}
}
