package repository

import (
	"context"
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID               uint
	Title            string
	Description      string
	ThumbnailFileID  uint
	RecipeCategoryID uint
	UserID           uint
	AverageRating    float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

type RecipeRepository struct {
	Context context.Context
	DB      *gorm.DB
}

type RecipeRepositoryInterface interface {
	GetAll() (res []entity.Recipe, err error)
	FindByID(id int) (res entity.Recipe, err error)
	Store(recipe *entity.Recipe) (err error)
	UpdateAverageRating(recipeID int, averageRating float64) error
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

func (repo *RecipeRepository) UpdateAverageRating(recipeID int, averageRating float64) error {
	rec := Recipe{}
	err := repo.DB.Model(&rec).Where("id = ?", recipeID).Update("average_rating", averageRating).Error
	if err != nil {
		return err
	}

	return nil
}
