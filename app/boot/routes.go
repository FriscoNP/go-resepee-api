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

	materialRoute := baseRoute.Group("/materials")
	materialController := controller.NewMaterialController(boot.DB, boot.JwtAuth)
	materialRoute.GET("/", materialController.Get)
	materialRoute.POST("/", materialController.Store)

	recipeCategoryRoute := baseRoute.Group("/categories")
	recipeCategoryController := controller.NewRecipeCategoryController(boot.DB)
	recipeCategoryRoute.GET("/", recipeCategoryController.GetAll)
	recipeCategoryRoute.POST("/", recipeCategoryController.Store)
}
