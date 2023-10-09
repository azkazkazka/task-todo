package routes

import (
	"fmt"
	"net/http"

	"github.com/azkazkazka/task-todo/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	fmt.Println("hoho")
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi!")
	})

	// users
	e.GET("/users", controllers.FetchUsers)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	// tasks
	e.GET("/tasks", controllers.FetchTasks)
	e.POST("/tasks", controllers.CreateTask)
	e.PUT("/tasks/:id", controllers.UpdateTask)
	e.DELETE("/tasks/:id", controllers.DeleteTask)

	return e
}
