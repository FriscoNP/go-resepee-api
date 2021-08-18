package repository

import (
	"context"
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
)

type RecipeCategory struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type RecipeCategoryRepository struct {
	Ctx context.Context
	DB  *gorm.DB
}

type RecipeCategoryRepositoryInterface interface {
	GetAll() (res []entity.RecipeCategory, err error)
	Store(category *entity.RecipeCategory) (res entity.RecipeCategory, err error)
}

func NewRecipeCategoryRepository(ctx context.Context, db *gorm.DB) RecipeCategoryRepositoryInterface {
	return &RecipeCategoryRepository{
		Ctx: ctx,
		DB:  db,
	}
}

func (repo *RecipeCategoryRepository) ToEntity(rec *RecipeCategory) entity.RecipeCategory {
	return entity.RecipeCategory{
		ID:        int(rec.ID),
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func (repo *RecipeCategoryRepository) GetAll() (res []entity.RecipeCategory, err error) {
	recs := []RecipeCategory{}
	err = repo.DB.Find(&recs).Error
	if err != nil {
		return res, err
	}

	for _, rec := range recs {
		res = append(res, repo.ToEntity(&rec))
	}

	return res, err
}

func (repo *RecipeCategoryRepository) Store(category *entity.RecipeCategory) (res entity.RecipeCategory, err error) {
	rec := RecipeCategory{
		Name: category.Name,
	}
	err = repo.DB.Create(&rec).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}
