package controllers

import (
	"net/http"
	"strconv"

	"github.com/azkazkazka/task-todo/models"
	"github.com/azkazkazka/task-todo/utils"
	"github.com/labstack/echo"
)

func FetchUsers(c echo.Context) error {
	res, err := models.FetchUsers()
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := models.CreateUser(user)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID format")
	}

	res, err := models.DeleteUser(uint(userID))
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user payload")
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID format")
	}

	user.ID = uint(userID)

	res, err := models.UpdateUser(user)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}
