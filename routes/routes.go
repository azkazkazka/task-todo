package routes

import (
	"net/http"

	"github.com/azkazkazka/task-todo/controllers"
	"github.com/azkazkazka/task-todo/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	protected := e.Group("")
	protected.Use(middleware.ValidateToken)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi!")
	})

	// register & login
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// users
	protected.GET("/user", controllers.GetUser)
	protected.PUT("/user", controllers.UpdateUser)
	protected.DELETE("/user", controllers.DeleteUser)

	// tasks
	protected.GET("/tasks", controllers.FetchAllTasks)
	protected.GET("/tasks/:id", controllers.FetchTask)
	protected.POST("/tasks", controllers.CreateTask)
	protected.PUT("/tasks/:id", controllers.UpdateTask)
	protected.DELETE("/tasks/:id", controllers.DeleteTask)

	return e
}
