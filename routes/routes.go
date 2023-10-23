package routes

import (
	"net/http"

	"github.com/azkazkazka/task-todo/controllers"
	"github.com/azkazkazka/task-todo/middleware"
	"github.com/azkazkazka/task-todo/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *echo.Echo {
	e := echo.New()
	taskController := &controllers.TaskController{Service: &models.GormTaskService{DB: db}}

	userController := &controllers.UserController{Service: &models.GormUserService{DB: db}}

	protected := e.Group("")
	protected.Use(middleware.ValidateToken)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi!")
	})

	// register & login
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)

	// users
	protected.GET("/user", userController.GetUser)
	protected.PUT("/user", userController.UpdateUser)
	protected.DELETE("/user", userController.DeleteUser)

	// tasks
	protected.GET("/tasks", taskController.FetchAllTasks)
	protected.GET("/tasks/:id", taskController.FetchTask)
	protected.POST("/tasks", taskController.CreateTask)
	protected.PUT("/tasks/:id", taskController.UpdateTask)
	protected.DELETE("/tasks/:id", taskController.DeleteTask)

	return e
}
