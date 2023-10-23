package utils

import "github.com/labstack/echo"

func SendResponse(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}
