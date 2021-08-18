package usecase

import (
	"context"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RecipeUC struct {
	Context context.Context
	DB      *gorm.DB
}

type RecipeUCInterface interface {
	GetAll() (res []entity.Recipe, err error)
	FindByID(id int) (recipe entity.Recipe, err error)
	Store(req *request.RecipeRequest, userID int) (recipe entity.Recipe, err error)
}

func NewRecipeUC(ctx context.Context, db *gorm.DB) RecipeUCInterface {
	return &RecipeUC{
		Context: ctx,
		DB:      db,
	}
}

func (uc *RecipeUC) GetAll() (res []entity.Recipe, err error) {
	recipeRepo := repository.NewRecipeRepository(uc.Context, uc.DB)

	res, err = recipeRepo.GetAll()
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}

func (uc *RecipeUC) FindByID(id int) (recipe entity.Recipe, err error) {
	recipeRepo := repository.NewRecipeRepository(uc.Context, uc.DB)
	recipeMaterialRepo := repository.NewRecipeMaterialRepository(uc.Context, uc.DB)
	cookStepRepository := repository.NewCookStepRepository(uc.Context, uc.DB)

	recipe, err = recipeRepo.FindByID(id)
	if err != nil {
		log.Warn(err.Error())
		return recipe, err
	}

	recipe.RecipeMaterials, err = recipeMaterialRepo.FindByRecipeID(id)
	if err != nil {
		log.Warn(err.Error())
		return recipe, err
	}

	recipe.CookSteps, err = cookStepRepository.FindByRecipeID(id)
	if err != nil {
		log.Warn(err.Error())
		return recipe, err
	}

	return recipe, err
}

func (uc *RecipeUC) Store(req *request.RecipeRequest, userID int) (recipe entity.Recipe, err error) {
	recipeRepo := repository.NewRecipeRepository(uc.Context, uc.DB)
	recipeMaterialRepo := repository.NewRecipeMaterialRepository(uc.Context, uc.DB)
	cookStepRepository := repository.NewCookStepRepository(uc.Context, uc.DB)

	// insert recipe
	newRecipe := entity.Recipe{
		Title:            req.Title,
		Description:      req.Description,
		ThumbnailFileID:  req.ThumbnailFileID,
		RecipeCategoryID: req.RecipeCategoryID,
		UserID:           userID,
	}
	recipe, err = recipeRepo.Store(&newRecipe)
	if err != nil {
		log.Warn(err.Error())
		return recipe, err
	}

	// insert recipe materials
	for _, recipeMaterial := range req.Materials {
		data := entity.RecipeMaterial{
			RecipeID:   recipe.ID,
			MaterialID: uint(recipeMaterial.MaterialID),
			Amount:     recipeMaterial.Amount,
		}
		recipeMaterialEntity, err := recipeMaterialRepo.Store(&data)
		if err != nil {
			log.Warn(err.Error())
			return recipe, err
		}
		recipe.RecipeMaterials = append(recipe.RecipeMaterials, recipeMaterialEntity)
	}

	// insert cook steps
	for _, cookStep := range req.CookingSteps {
		data := entity.CookStep{
			RecipeID:    recipe.ID,
			Description: cookStep.Description,
			Order:       cookStep.Order,
		}
		cookStepEntity, err := cookStepRepository.Store(&data)
		if err != nil {
			log.Warn(err.Error())
			return recipe, err
		}
		recipe.CookSteps = append(recipe.CookSteps, cookStepEntity)
	}

	return recipe, err
}
