package controller

import (
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/controller/response"
	"go-resepee-api/app/middleware"
	"go-resepee-api/db/repository"
	"go-resepee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MaterialController struct {
	DB      *gorm.DB
	JwtAuth *middleware.ConfigJWT
}

func NewMaterialController(db *gorm.DB, jwtAuth *middleware.ConfigJWT) *MaterialController {
	return &MaterialController{
		DB:      db,
		JwtAuth: jwtAuth,
	}
}

func (mc *MaterialController) Get(c echo.Context) error {
	ctx := c.Request().Context()

	materialRepo := repository.NewMaterialRepository(ctx, mc.DB)
	materialUC := usecase.NewMaterialUC(ctx, materialRepo)
	materials, err := materialUC.Get()
	if err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusBadRequest, err)
	}

	res := []response.MaterialResponse{}
	for _, material := range materials {
		res = append(res, response.CreateMaterialResponse(&material))
	}

	return SendSuccess(c, res, "get_materials")
}

func (mc *MaterialController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.CreateMaterialRequest{}
	if err := c.Bind(&req); err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusBadRequest, err)
	}

	materialRepo := repository.NewMaterialRepository(ctx, mc.DB)
	materialUC := usecase.NewMaterialUC(ctx, materialRepo)
	material, err := materialUC.Store(&req)
	if err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	return SendSuccess(c, response.CreateMaterialResponse(&material), "material_created")
}
