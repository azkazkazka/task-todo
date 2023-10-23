package controllers

import (
	"testing"

	"github.com/azkazkazka/task-todo/models"
	"github.com/azkazkazka/task-todo/models/mocks/mockusers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Register(t *testing.T) {
	// init mock for IUserService
	dbMock := mockusers.NewIUserService(t)
	// create instance of UserService with the mock service
	mockUserService := models.UserService{Service: dbMock}
	
	// mock data (register request)
	userReq := &models.UserRequest{
		Fullname: "Test User",
		Username: "testuser",
		Email:    "test@email.com",
		Password: "password123",
	}

	// expect method will be called with the user request and return a user response
	dbMock.On("Register", userReq).Return(&models.UserResponse{ID: "some-id", Fullname: "Test User", Username: "testuser", Email: "test@email.com"}, nil)

	// execute
	resp, err := mockUserService.Service.Register(userReq)

	// check no errors, response not nil, make sure all expected calls were made
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	dbMock.AssertExpectations(t)
}

func TestUserService_AuthenticateUser(t *testing.T) {
	dbMock := mockusers.NewIUserService(t)
	mockUserService := models.UserService{Service: dbMock}
	
	// mock data (login request)
	loginReq := &models.LoginRequest{
		Username: "testuser",
		Email:    "test@email.com",
		Password: "password123",
	}

	// expect return values and execute the method with assertions like before
	dbMock.On("AuthenticateUser", loginReq).Return(&models.UserResponse{ID: "some-id", Fullname: "Test User", Username: "testuser", Email: "test@email.com"}, nil)

	resp, err := mockUserService.Service.AuthenticateUser(loginReq)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	dbMock.AssertExpectations(t)
}

func TestUserService_FetchAllUsers(t *testing.T) {
	dbMock := mockusers.NewIUserService(t)
	mockUserService := models.UserService{Service: dbMock}

	// mock data (list of users)
	expectedUsers := []models.UserResponse{
		{
			ID:       "some-id-1",
			Fullname: "John Doe",
		},
		{
			ID:       "some-id-2",
			Fullname: "Jane Doe",
		},
	}

	// expect return values and execute the method with assertions like before
	dbMock.On("FetchAllUsers").Return(expectedUsers, nil)

	result, err := mockUserService.Service.FetchAllUsers()

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, expectedUsers, result)
}

func TestUserService_GetUser(t *testing.T) {
	dbMock := mockusers.NewIUserService(t)
	mockUserService := models.UserService{Service: dbMock}

	// mock data (user)
	expectedUser := &models.UserResponse{
		ID:       "some-id-1",
		Fullname: "John Doe",
	}

	// expect return values and execute the method with assertions like before
	dbMock.On("GetUser", mock.Anything).Return(expectedUser, nil)

	// Invoke the GetUser method.
	result, err := mockUserService.Service.GetUser("some-id-1")

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, result)
}

func TestUserService_DeleteUser(t *testing.T) {
	dbMock := mockusers.NewIUserService(t)
	mockUserService := models.UserService{Service: dbMock}

	// expect return values and execute the method with assertions like before
	dbMock.On("DeleteUser", mock.Anything).Return(nil, nil)

	result, err := mockUserService.Service.DeleteUser("some-id-1")

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, nil, result)
}

func TestUserService_UpdateUser(t *testing.T) {
	dbMock := mockusers.NewIUserService(t)
	mockUserService := models.UserService{Service: dbMock}

	// mock data (user update request)
	updateReq := &models.UpdateRequest{
		ID:       "some-id-1",
		Fullname: "John Updated",
	}

	// user data that is expected to be after update
	expectedUser := &models.UserResponse{
		ID:       "some-id-1",
		Fullname: "John Updated",
	}

	// expect return values and execute the method with assertions like before
	dbMock.On("UpdateUser", updateReq).Return(expectedUser, nil)

	// Invoke the UpdateUser method.
	result, err := mockUserService.Service.UpdateUser(updateReq)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, result)
}
