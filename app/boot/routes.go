package boot

import (
	"go-resepee-api/app/controller"

	"github.com/labstack/echo/v4/middleware"
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
	materialRoute.GET("", materialController.Get, middleware.JWTWithConfig(boot.JwtConfig))
	materialRoute.POST("", materialController.Store, middleware.JWTWithConfig(boot.JwtConfig))

	recipeCategoryRoute := baseRoute.Group("/categories")
	recipeCategoryController := controller.NewRecipeCategoryController(boot.DB)
	recipeCategoryRoute.GET("", recipeCategoryController.GetAll, middleware.JWTWithConfig(boot.JwtConfig))
	recipeCategoryRoute.POST("", recipeCategoryController.Store, middleware.JWTWithConfig(boot.JwtConfig))

	recipeRoute := baseRoute.Group("/recipes")
	recipeController := controller.NewRecipeController(boot.DB)
	recipeRoute.GET("", recipeController.GetAll, middleware.JWTWithConfig(boot.JwtConfig))
	recipeRoute.POST("", recipeController.Store, middleware.JWTWithConfig(boot.JwtConfig))
	recipeRoute.GET("/:id", recipeController.FindByID, middleware.JWTWithConfig(boot.JwtConfig))

	reviewRoute := baseRoute.Group("/reviews")
	reviewController := controller.NewReviewController(boot.DB)
	reviewRoute.GET("", reviewController.FindByRecipeID, middleware.JWTWithConfig(boot.JwtConfig))
	reviewRoute.POST("", reviewController.Store, middleware.JWTWithConfig(boot.JwtConfig))

	fileRoute := baseRoute.Group("/files")
	fileController := controller.NewFileController(boot.DB)
	fileRoute.POST("", fileController.Store, middleware.JWTWithConfig(boot.JwtConfig))
}
