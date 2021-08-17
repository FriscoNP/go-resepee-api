package controller

import (
	"go-resepee-api/app/controller/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func SendSuccess(c echo.Context, data interface{}, message string) error {
	resp := response.BaseResponse{
		Message: message,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func SendError(c echo.Context, status int, err error) error {
	resp := response.BaseResponse{
		Error: err.Error(),
	}
	return c.JSON(status, resp)
}
