package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func SendSuccess(c echo.Context, data interface{}, message string) error {
	resp := BaseResponse{
		Message: message,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func SendError(c echo.Context, status int, err error) error {
	resp := BaseResponse{
		Error: err.Error(),
	}
	return c.JSON(status, resp)
}
