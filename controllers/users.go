package controllers

import (
	"net/http"

	"github.com/azkazkazka/task-todo/auth"
	"github.com/azkazkazka/task-todo/models"
	"github.com/azkazkazka/task-todo/utils"
	"github.com/labstack/echo"
)

type UserController struct {
	Service models.IUserService
}

func (uc *UserController) Register(c echo.Context) error {
	user := &models.UserRequest{}
	if err := c.Bind(user); err != nil {
		errResp := utils.ErrorResponse{
			Message: "Bad Request",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusBadRequest, errResp)
	}

	data, err := uc.Service.Register(user)
	if err != nil {
		errResp := utils.ErrorResponse{
			Message: "Failed to register account",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusInternalServerError, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data)
}

func (uc *UserController) Login(c echo.Context) error {
	loginRequest := &models.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		errResp := utils.ErrorResponse{
			Message: "Bad Request",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusBadRequest, errResp)
	}

	user, err := uc.Service.AuthenticateUser(loginRequest)
	if err != nil {
		errResp := utils.ErrorResponse{
			Message: "Failed to authenticate user",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusInternalServerError, errResp)
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Message: "Failed to generate token",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusInternalServerError, errResp)
	}

	data := map[string]interface{}{
		"user":  user,
		"token": token,
	}

	return utils.SendResponse(c, http.StatusOK, data)
}

func (uc *UserController) GetUser(c echo.Context) error {
	userID := c.Get("userID").(string)

	data, err := uc.Service.GetUser(userID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Message: "Failed to get user",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusInternalServerError, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data)
}


func (uc *UserController) DeleteUser(c echo.Context) error {
	userID := c.Get("userID").(string)

	data, err := uc.Service.DeleteUser(userID)
	if err != nil {
		errResp := utils.ErrorResponse{
			Message: "Failed to delete user",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusInternalServerError, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data)
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	user := &models.UpdateRequest{}
	if err := c.Bind(user); err != nil {
		errResp := utils.ErrorResponse{
			Message: "Invalid user payload",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusBadRequest, errResp)
	}

	user.ID = c.Get("userID").(string)

	data, err := uc.Service.UpdateUser(user)
	if err != nil {
		errResp := utils.ErrorResponse{
			Message: "Failed to update user",
			Details: err.Error(),
		}
		return utils.SendErrorResponse(c, http.StatusInternalServerError, errResp)
	}

	return utils.SendResponse(c, http.StatusOK, data)
}
