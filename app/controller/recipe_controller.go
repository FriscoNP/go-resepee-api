package controller

import (
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/middleware"
	"go-resepee-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RecipeController struct {
	DB *gorm.DB
}

func NewRecipeController(db *gorm.DB) *RecipeController {
	return &RecipeController{
		DB: db,
	}
}

func (controller *RecipeController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	recipeUC := usecase.NewRecipeUC(ctx, controller.DB)
	resp, err := recipeUC.GetAll()
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	return SendSuccess(c, resp, "get_all_recipe")
}

func (controller *RecipeController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	jwtUser := middleware.GetJWTUser(c)

	req := request.RecipeRequest{}
	if err := c.Bind(&req); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	// start transaction
	tx := controller.DB.Begin()
	recipeUC := usecase.NewRecipeUC(ctx, tx)
	recipe, err := recipeUC.Store(&req, jwtUser.ID)
	if err != nil {
		// rollback if error
		tx.Rollback()
		return SendError(c, http.StatusInternalServerError, err)
	}
	// commit transaction
	tx.Commit()

	return SendSuccess(c, recipe, "recipe_created")
}

func (controller *RecipeController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	recipeID, err := strconv.Atoi(id)
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	recipeUC := usecase.NewRecipeUC(ctx, controller.DB)
	recipe, err := recipeUC.FindByID(recipeID)
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	return SendSuccess(c, recipe, "get_detail_recipe")
}
