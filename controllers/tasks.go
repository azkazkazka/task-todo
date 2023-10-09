package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/azkazkazka/task-todo/models"
	"github.com/azkazkazka/task-todo/utils"
	"github.com/labstack/echo"
)

func FetchTasks(c echo.Context) error {
	res, err := models.FetchTasks()
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}

func CreateTask(c echo.Context) error {
	task := &models.Task{}

	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	parsedDueDate, err := time.Parse(time.RFC3339, task.DueDateString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid due_date format. It should be RFC3339 formatted string.")
	}
	task.DueDate = parsedDueDate

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	res, err := models.CreateTask(task)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task ID format")
	}

	res, err := models.DeleteTask(uint(taskID))
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateTask(c echo.Context) error {
	task := &models.Task{}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task payload")
	}

	task.ID = c.Param("id")

	res, err := models.UpdateTask(task)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}
