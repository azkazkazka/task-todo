package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/azkazkazka/task-todo/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestFetchAllTasks(t *testing.T) {
	mockTaskService := new(models.MockTaskService)

	mockTasks := []models.TaskResponse{
		{ID: "1", UserID: "user1", Title: "Test Task 1"},
		{ID: "2", UserID: "user1", Title: "Test Task 2"},
	}
	mockTaskService.On("FetchAllTasks", "user1").Return(mockTasks, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userID", "user1")

	controller := &TaskController{
        Service: mockTaskService,
    }

	if assert.NoError(t, controller.FetchAllTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Successfully fetched all task")
	}

	mockTaskService.AssertExpectations(t)
}
