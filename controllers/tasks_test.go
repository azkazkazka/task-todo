package controllers

import (
	"testing"
	"time"

	"github.com/azkazkazka/task-todo/models"
	"github.com/azkazkazka/task-todo/models/mocks/mocktasks"
	"github.com/stretchr/testify/assert"
)

func TestTaskService_FetchAllTasks(t *testing.T) {
    // init mock of ITaskService
	dbMock := mocktasks.NewITaskService(t)
    // create TaskService instance w/ mock service
	mockTaskService := models.TaskService{Service: dbMock}

    // mock data (tasks fetched from the service)
	tasks := []models.TaskResponse{
		{
			ID:               "1",
			UserID:           "user1",
			Title:            "Test task",
			Description:      "Test description",
			DueDate:          time.Now(),
			CompletionStatus: false,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

    // expect FetchAllTasks will be called with "user1" + return mocked tasks
	dbMock.On("FetchAllTasks", "user1").Return(tasks, nil)

    // execute method
	result, err := mockTaskService.Service.FetchAllTasks("user1")

    // assert no error occurred, returned task match mocked data, and expected calls were made
	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	dbMock.AssertExpectations(t)
}

func TestTaskService_FetchTask(t *testing.T) {
	dbMock := mocktasks.NewITaskService(t)
	mockTaskService := models.TaskService{Service: dbMock}

    // mock data (a task fetched from the service)
	task := &models.TaskResponse{
		ID:               "1",
		UserID:           "user1",
		Title:            "Test task",
		Description:      "Test description",
		DueDate:          time.Now(),
		CompletionStatus: false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

    // expect return values and execute the method with assertions like before
	dbMock.On("FetchTask", "1", "user1").Return(task, nil)

	result, err := mockTaskService.Service.FetchTask("1", "user1")

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	dbMock.AssertExpectations(t)
}

func TestTaskService_CreateTask(t *testing.T) {
	dbMock := mocktasks.NewITaskService(t)
	mockTaskService := models.TaskService{Service: dbMock}

    // mock data (a new task)
	newTask := &models.Task{
		UserID:           "user1",
		Title:            "Test task",
		Description:      "Test description",
		DueDate:          time.Now(),
		CompletionStatus: false,
	}

    // expect return values and execute the method with assertions like before
	dbMock.On("CreateTask", newTask).Return(newTask, nil)

	result, err := mockTaskService.Service.CreateTask(newTask)

	assert.NoError(t, err)
	assert.Equal(t, newTask, result)
	dbMock.AssertExpectations(t)
}

func TestTaskService_DeleteTask(t *testing.T) {
    dbMock := mocktasks.NewITaskService(t)
    mockTaskService := models.TaskService{Service: dbMock}

    // define task and user id
    taskID := "1"
    userID := "user1"

    // expect return values and execute the method with assertions like before without checking return value equal or not
    dbMock.On("DeleteTask", taskID, userID).Return(nil, nil)

    _, err := mockTaskService.Service.DeleteTask(taskID, userID)

    assert.NoError(t, err)
    dbMock.AssertExpectations(t)
}

func TestTaskService_UpdateTask(t *testing.T) {
	dbMock := mocktasks.NewITaskService(t)
	mockTaskService := models.TaskService{Service: dbMock}

    // mock data (a task to be updated)
	taskToUpdate := &models.Task{
		ID:               "1",
		UserID:           "user1",
		Title:            "Updated task",
		Description:      "Updated description",
		DueDate:          time.Now(),
		CompletionStatus: false,
	}

    // expect return values and execute the method with assertions like before
	dbMock.On("UpdateTask", taskToUpdate).Return(taskToUpdate, nil)

	result, err := mockTaskService.Service.UpdateTask(taskToUpdate)

	assert.NoError(t, err)
	assert.Equal(t, taskToUpdate, result)
	dbMock.AssertExpectations(t)
}
