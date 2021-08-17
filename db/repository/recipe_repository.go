package repository

import (
	"context"
	"go-resepee-api/entity"

	"gorm.io/gorm"
)

type RecipeRepository struct {
	Context context.Context
	DB      *gorm.DB
}

type RecipeRepositoryInterface interface {
	GetAll() (res []entity.Recipe, err error)
	FindByID(id int) (res entity.Recipe, err error)
	Store(recipe *entity.Recipe) (err error)
}

func NewRecipeRepository(ctx context.Context, db *gorm.DB) RecipeRepositoryInterface {
	return &RecipeRepository{
		Context: ctx,
		DB:      db,
	}
}

func (repo *RecipeRepository) GetAll() (res []entity.Recipe, err error) {
	err = repo.DB.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo *RecipeRepository) FindByID(id int) (res entity.Recipe, err error) {
	err = repo.DB.Find(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo *RecipeRepository) Store(recipe *entity.Recipe) (err error) {
	err = repo.DB.Create(recipe).Error
	if err != nil {
		return err
	}

	return err
}
