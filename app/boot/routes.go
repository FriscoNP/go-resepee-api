package boot

import (
	"go-resepee-api/app/controller"
)

func (boot *BootApp) RegisterRoutes() {
	e := boot.Echo
	baseRoute := e.Group("/api/v1")

	authController := controller.AuthController{
		DB:      boot.DB,
		JwtAuth: boot.JwtAuth,
	}
	authRoute := baseRoute.Group("/auth")
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)
}
