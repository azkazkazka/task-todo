package utils

import "github.com/labstack/echo"

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SendResponse(c echo.Context, status int, data interface{}, message string) error {
	response := Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
	return c.JSON(status, response)
}
