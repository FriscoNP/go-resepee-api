package repository

import (
	"context"
	"go-resepee-api/entity"

	"gorm.io/gorm"
)

type RecipeMaterialRepository struct {
	Context context.Context
	DB      *gorm.DB
}

type RecipeMaterialRepositoryInterface interface {
	FindByRecipeID(recipeID int) (res []entity.RecipeMaterial, err error)
	Store(recipeMaterial *entity.RecipeMaterial) (err error)
}

func NewRecipeMaterialRepository(ctx context.Context, db *gorm.DB) RecipeMaterialRepositoryInterface {
	return &RecipeMaterialRepository{
		Context: ctx,
		DB:      db,
	}
}

func (repo *RecipeMaterialRepository) FindByRecipeID(recipeID int) (res []entity.RecipeMaterial, err error) {
	err = repo.DB.Where("recipe_id = ?", recipeID).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo *RecipeMaterialRepository) Store(recipeMaterial *entity.RecipeMaterial) (err error) {
	err = repo.DB.Create(recipeMaterial).Error
	if err != nil {
		return err
	}

	return err
}
