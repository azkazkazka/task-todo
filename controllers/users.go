package controllers

import (
	"net/http"

	"github.com/azkazkazka/task-todo/auth"
	"github.com/azkazkazka/task-todo/models"
	"github.com/azkazkazka/task-todo/utils"
	"github.com/labstack/echo"
)

func Register(c echo.Context) error {
	user := &models.UserRequest{}
	if err := c.Bind(user); err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	data, err := models.Register(user)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to register account",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully registered account")
}

func Login(c echo.Context) error {
	loginRequest := &models.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	user, err := models.AuthenticateUser(loginRequest)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to authenticate user",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to generate token",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	data := map[string]interface{}{
		"user":  user,
		"token": token,
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully logged in")
}

func GetUser(c echo.Context) error {
	userID := c.Get("userID").(string)

	data, err := models.GetUser(userID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get user",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully get user")
}


func DeleteUser(c echo.Context) error {
	userID := c.Get("userID").(string)

	data, err := models.DeleteUser(userID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete user",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully deleted user")
}

func UpdateUser(c echo.Context) error {
	user := &models.UpdateRequest{}
	if err := c.Bind(user); err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid user payload",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	user.ID = c.Get("userID").(string)

	data, err := models.UpdateUser(user)
	if err != nil {
		errResp := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update user",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data, "Successfully updated user")
}
