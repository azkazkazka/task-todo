package models

import (
    "github.com/stretchr/testify/mock"
)

type MockTaskService struct {
    mock.Mock
}

func (m *MockTaskService) FetchAllTasks(userID string) (interface{}, error) {
    args := m.Called(userID)
    return args.Get(0).(interface{}), args.Error(1)
}

func (m *MockTaskService) FetchTask(taskID, userID string) (interface{}, error) {
    args := m.Called(taskID, userID)
    return args.Get(0).(interface{}), args.Error(1)
}

func (m *MockTaskService) CreateTask(task *Task) (interface{}, error) {
    args := m.Called(task)
    return args.Get(0).(interface{}), args.Error(1)
}

func (m *MockTaskService) DeleteTask(taskID, userID string) (interface{}, error) {
    args := m.Called(taskID, userID)
    return args.Get(0).(interface{}), args.Error(1)
}

func (m *MockTaskService) UpdateTask(task *Task) (interface{}, error) {
    args := m.Called(task)
    return args.Get(0).(interface{}), args.Error(1)
}
