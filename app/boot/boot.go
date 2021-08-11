package boot

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	appMiddleware "go-resepee-api/app/middleware"
)

type BootApp struct {
	DB        *gorm.DB
	JwtAuth   *appMiddleware.ConfigJWT
	JwtConfig middleware.JWTConfig
	Echo      *echo.Echo
}
