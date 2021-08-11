package controller

import (
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/controller/response"
	"go-resepee-api/app/middleware"
	"go-resepee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthController struct {
	DB      *gorm.DB
	JwtAuth *middleware.ConfigJWT
}

func (ac *AuthController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	authUC := usecase.AuthUC{
		Context: ctx,
		DB:      ac.DB,
		JwtAuth: ac.JwtAuth,
	}

	token, err := authUC.Login(req.Email, req.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return SendError(c, http.StatusNotFound, err)
		}
		logrus.Warn(err.Error())
		return SendError(c, http.StatusInternalServerError, err)
	}

	resp := response.LoginResponse{
		Token: token,
	}

	return SendSuccess(c, resp, "login_success")
}

func (ac *AuthController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	authUC := usecase.AuthUC{
		Context: ctx,
		DB:      ac.DB,
		JwtAuth: ac.JwtAuth,
	}

	user, err := authUC.Register(&req)
	if err != nil {
		logrus.Warn(err.Error())
		return SendError(c, http.StatusInternalServerError, err)
	}

	return SendSuccess(c, user, "account_created")
}
