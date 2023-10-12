package controllers

import (
	"net/http"
	"time"

	"github.com/azkazkazka/task-todo/models"
	"github.com/azkazkazka/task-todo/utils"
	"github.com/labstack/echo"
)

type TaskController struct {
	Service models.ITaskService
}

func (tc *TaskController) FetchAllTasks(c echo.Context) error {
	userID := c.Get("userID").(string)

	data, err := tc.Service.FetchAllTasks(userID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to fetch all tasks",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully fetched all task")
}

func (tc *TaskController) FetchTask(c echo.Context) error {
	taskID := c.Param("id")
	userID := c.Get("userID").(string)

	data, err := tc.Service.FetchTask(taskID, userID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to fetch task",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully fetched task")
}

func (tc *TaskController) CreateTask(c echo.Context) error {
	task := &models.Task{}
	task.UserID = c.Get("userID").(string)

	if err := c.Bind(task); err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	parsedDueDate, err := time.Parse(time.RFC3339, task.DueDateString)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid due date format (accepts RFC3339)",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}
	task.DueDate = parsedDueDate

	data, err := tc.Service.CreateTask(task)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create task",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully created task")
}

func (tc *TaskController) DeleteTask(c echo.Context) error {
	taskID := c.Param("id")
	userID := c.Get("userID").(string)

	data, err := tc.Service.DeleteTask(taskID, userID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete task",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully deleted task")
}

func (tc *TaskController) UpdateTask(c echo.Context) error {
	task := &models.Task{}
	task.UserID = c.Get("userID").(string)

	if err := c.Bind(task); err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid task payload",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	task.ID = c.Param("id")

	data, err := tc.Service.UpdateTask(task)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update task",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully updated task")
}
