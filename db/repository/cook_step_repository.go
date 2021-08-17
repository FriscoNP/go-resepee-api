package repository

import (
	"context"
	"go-resepee-api/entity"

	"gorm.io/gorm"
)

type CookStepRepository struct {
	Context context.Context
	DB      *gorm.DB
}

type CookStepRepositoryInterface interface {
	FindByRecipeID(recipeID int) (res []entity.CookStep, err error)
	Store(cookStep *entity.CookStep) (err error)
}

func NewCookStepRepository(ctx context.Context, db *gorm.DB) CookStepRepositoryInterface {
	return &CookStepRepository{
		Context: ctx,
		DB:      db,
	}
}

func (repo *CookStepRepository) FindByRecipeID(recipeID int) (res []entity.CookStep, err error) {
	err = repo.DB.Where("recipe_id = ?", recipeID).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo *CookStepRepository) Store(cookStep *entity.CookStep) (err error) {
	err = repo.DB.Create(cookStep).Error
	if err != nil {
		return err
	}

	return err
}
