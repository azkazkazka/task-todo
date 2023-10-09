package utils

import (
	"net/http"

	"github.com/labstack/echo"
)

func HandleError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"message": err.Error(),
	})
}
