package routes

import (
	"net/http"

	"github.com/azkazkazka/task-todo/controllers"
	"github.com/azkazkazka/task-todo/db"
	"github.com/azkazkazka/task-todo/middleware"
	"github.com/azkazkazka/task-todo/models"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	con := db.CreateCon()
	taskService := &models.TaskService{DB: con}
	taskController := &controllers.TaskController{Service: taskService}

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
	protected.GET("/tasks", taskController.FetchAllTasks)
	protected.GET("/tasks/:id", taskController.FetchTask)
	protected.POST("/tasks", taskController.CreateTask)
	protected.PUT("/tasks/:id", taskController.UpdateTask)
	protected.DELETE("/tasks/:id", taskController.DeleteTask)

	return e
}
