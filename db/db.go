package db

import (
	"fmt"

	"github.com/azkazkazka/task-todo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB
var err error

func Init() {
	conf := config.GetConfig()

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		conf.DB_HOST,
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_NAME,
		conf.DB_PORT,
		conf.DB_SSLMODE,
	)

	postgresDB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func CreateCon() *gorm.DB {
	return postgresDB
}
