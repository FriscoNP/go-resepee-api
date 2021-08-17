package controller

import (
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/controller/response"
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

	req := request.RecipeRequest{}
	if err := c.Bind(&req); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	// start transaction
	tx := controller.DB.Begin()
	recipeUC := usecase.NewRecipeUC(ctx, tx)
	recipe, recipeMaterials, cookSteps, err := recipeUC.Store(&req)
	if err != nil {
		// rollback if error
		tx.Rollback()
		return SendError(c, http.StatusInternalServerError, err)
	}
	// commit transaction
	tx.Commit()

	recipeMaterialResponses := []response.RecipeMaterialResponse{}
	for _, recipeMaterial := range recipeMaterials {
		recipeMaterialResponse := response.RecipeMaterialResponse{
			Material: response.MaterialResponse{
				ID: int(recipeMaterial.ID),
			},
			Amount: recipeMaterial.Amount,
		}
		recipeMaterialResponses = append(recipeMaterialResponses, recipeMaterialResponse)
	}

	cookStepResponses := []response.CookStepResponse{}
	for _, cookStep := range cookSteps {
		cookStepResponse := response.CookStepResponse{
			ID:          int(cookStep.ID),
			Description: cookStep.Description,
			Order:       cookStep.Order,
		}
		cookStepResponses = append(cookStepResponses, cookStepResponse)
	}

	resp := response.RecipeResponse{
		ID:               int(recipe.ID),
		Title:            recipe.Title,
		Description:      recipe.Description,
		ThumbnailFileID:  recipe.ThumbnailFileID,
		ThumbnailFile:    response.FileResponse{},
		RecipeCategoryID: recipe.RecipeCategoryID,
		RecipeCategory:   response.RecipeCategoryResponse{},
		UserID:           recipe.UserID,
		User:             response.RecipeUserResponse{},
		AverageRating:    recipe.AverageRating,
		Materials:        recipeMaterialResponses,
		CookSteps:        cookStepResponses,
	}

	return SendSuccess(c, resp, "recipe_created")
}

func (controller *RecipeController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	recipeID, err := strconv.Atoi(id)
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	recipeUC := usecase.NewRecipeUC(ctx, controller.DB)
	recipe, recipeMaterials, cookSteps, err := recipeUC.FindByID(recipeID)
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	recipeMaterialResponses := []response.RecipeMaterialResponse{}
	for _, recipeMaterial := range recipeMaterials {
		recipeMaterialResponse := response.RecipeMaterialResponse{
			Material: response.MaterialResponse{
				ID: int(recipeMaterial.ID),
			},
			Amount: recipeMaterial.Amount,
		}
		recipeMaterialResponses = append(recipeMaterialResponses, recipeMaterialResponse)
	}

	cookStepResponses := []response.CookStepResponse{}
	for _, cookStep := range cookSteps {
		cookStepResponse := response.CookStepResponse{
			ID:          int(cookStep.ID),
			Description: cookStep.Description,
			Order:       cookStep.Order,
		}
		cookStepResponses = append(cookStepResponses, cookStepResponse)
	}

	resp := response.RecipeResponse{
		ID:               int(recipe.ID),
		Title:            recipe.Title,
		Description:      recipe.Description,
		ThumbnailFileID:  recipe.ThumbnailFileID,
		ThumbnailFile:    response.FileResponse{},
		RecipeCategoryID: recipe.RecipeCategoryID,
		RecipeCategory:   response.RecipeCategoryResponse{},
		UserID:           recipe.UserID,
		User:             response.RecipeUserResponse{},
		AverageRating:    recipe.AverageRating,
		Materials:        recipeMaterialResponses,
		CookSteps:        cookStepResponses,
	}

	return SendSuccess(c, resp, "get_detail_recipe")
}
