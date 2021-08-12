package boot

import (
	"go-resepee-api/app/controller"
)

func (boot *BootApp) RegisterRoutes() {
	e := boot.Echo
	baseRoute := e.Group("/api/v1")

	authRoute := baseRoute.Group("/auth")
	authController := controller.NewAuthController(boot.DB, boot.JwtAuth)
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)
}
