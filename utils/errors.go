package utils

import (
	"github.com/labstack/echo"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func SendErrorResponse(c echo.Context, status int, errResp ErrorResponse) error {
	return c.JSON(status, errResp)
}
