package usecase

import (
	"context"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"

	log "github.com/sirupsen/logrus"
)

type RecipeUC struct {
	Context                  context.Context
	RecipeRepository         repository.RecipeRepositoryInterface
	RecipeMaterialRepository repository.RecipeMaterialRepositoryInterface
	CookStepRepository       repository.CookStepRepositoryInterface
}

type RecipeUCInterface interface {
	GetAll() (res []entity.Recipe, err error)
	FindByID(id int) (recipe entity.Recipe, err error)
	Store(req *request.RecipeRequest, userID int) (recipe entity.Recipe, err error)
}

func NewRecipeUC(
	ctx context.Context,
	recipeRepo repository.RecipeRepositoryInterface,
	recipeMaterialRepo repository.RecipeMaterialRepositoryInterface,
	cookStepRepo repository.CookStepRepositoryInterface) RecipeUCInterface {
	return &RecipeUC{
		Context:                  ctx,
		RecipeRepository:         recipeRepo,
		RecipeMaterialRepository: recipeMaterialRepo,
		CookStepRepository:       cookStepRepo,
	}
}

func (uc *RecipeUC) GetAll() (res []entity.Recipe, err error) {
	res, err = uc.RecipeRepository.GetAll()
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}

func (uc *RecipeUC) FindByID(id int) (res entity.Recipe, err error) {
	recipe, err := uc.RecipeRepository.FindByID(id)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	recipe.RecipeMaterials, err = uc.RecipeMaterialRepository.FindByRecipeID(id)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	recipe.CookSteps, err = uc.CookStepRepository.FindByRecipeID(id)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return recipe, err
}

func (uc *RecipeUC) Store(req *request.RecipeRequest, userID int) (res entity.Recipe, err error) {
	// insert recipe
	newRecipe := entity.Recipe{
		Title:            req.Title,
		Description:      req.Description,
		ThumbnailFileID:  req.ThumbnailFileID,
		RecipeCategoryID: req.RecipeCategoryID,
		UserID:           userID,
	}
	recipe, err := uc.RecipeRepository.Store(&newRecipe)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	// insert recipe materials
	for _, recipeMaterial := range req.Materials {
		data := entity.RecipeMaterial{
			RecipeID:   recipe.ID,
			MaterialID: uint(recipeMaterial.MaterialID),
			Amount:     recipeMaterial.Amount,
		}
		recipeMaterialEntity, err := uc.RecipeMaterialRepository.Store(&data)
		if err != nil {
			log.Warn(err.Error())
			return res, err
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
		cookStepEntity, err := uc.CookStepRepository.Store(&data)
		if err != nil {
			log.Warn(err.Error())
			return res, err
		}
		recipe.CookSteps = append(recipe.CookSteps, cookStepEntity)
	}

	return recipe, err
}
