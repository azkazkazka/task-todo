package utils

import (
	"github.com/labstack/echo"
)

type ErrorResponse struct {
	Status    int    `json:"status"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func SendErrorResponse(c echo.Context, errResp ErrorResponse) error {
    return c.JSON(errResp.Status, errResp)
}