package controller

import (
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository"
	"go-resepee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RecipeCategoryController struct {
	DB *gorm.DB
}

func NewRecipeCategoryController(db *gorm.DB) *RecipeCategoryController {
	return &RecipeCategoryController{
		DB: db,
	}
}

func (controller *RecipeCategoryController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	recipeCategoryRepo := repository.NewRecipeCategoryRepository(ctx, controller.DB)
	recipeCategoryUC := usecase.NewRecipeCategoryUC(ctx, recipeCategoryRepo)
	resp, err := recipeCategoryUC.GetAll()
	if err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusInternalServerError, err)
	}

	return SendSuccess(c, resp, "get_all_category")
}

func (controller *RecipeCategoryController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.RecipeCategoryRequest{}
	if err := c.Bind(&req); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	recipeCategoryRepo := repository.NewRecipeCategoryRepository(ctx, controller.DB)
	recipeCategoryUC := usecase.NewRecipeCategoryUC(ctx, recipeCategoryRepo)
	resp, err := recipeCategoryUC.Store(&req)
	if err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusInternalServerError, err)
	}

	return SendSuccess(c, resp, "recipe_category_created")
}
