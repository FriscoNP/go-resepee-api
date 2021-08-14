package repository

import (
	"context"
	"go-resepee-api/entity"

	"gorm.io/gorm"
)

type RecipeCategoryRepository struct {
	Ctx context.Context
	DB  *gorm.DB
}

type RecipeCategoryRepositoryInterface interface {
	GetAll() (res []entity.RecipeCategory, err error)
	Store(category *entity.RecipeCategory) error
}

func NewRecipeCategoryRepository(ctx context.Context, db *gorm.DB) RecipeCategoryRepositoryInterface {
	return &RecipeCategoryRepository{
		Ctx: ctx,
		DB:  db,
	}
}

func (repo *RecipeCategoryRepository) GetAll() (res []entity.RecipeCategory, err error) {
	err = repo.DB.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo *RecipeCategoryRepository) Store(category *entity.RecipeCategory) error {
	err := repo.DB.Create(category).Error
	if err != nil {
		return err
	}

	return err
}
