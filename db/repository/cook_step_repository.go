package repository

import (
	"context"
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
)

type CookStep struct {
	ID          uint
	RecipeID    uint
	Description string
	Order       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CookStepRepository struct {
	Context context.Context
	DB      *gorm.DB
}

type CookStepRepositoryInterface interface {
	FindByRecipeID(recipeID int) (res []entity.CookStep, err error)
	Store(cookStep *entity.CookStep) (res entity.CookStep, err error)
}

func NewCookStepRepository(ctx context.Context, db *gorm.DB) CookStepRepositoryInterface {
	return &CookStepRepository{
		Context: ctx,
		DB:      db,
	}
}

func (repo *CookStepRepository) ToEntity(rec *CookStep) entity.CookStep {
	return entity.CookStep{
		ID:          rec.ID,
		RecipeID:    rec.RecipeID,
		Description: rec.Description,
		Order:       rec.Order,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func (repo *CookStepRepository) ToRecord(entity *entity.CookStep) CookStep {
	return CookStep{
		ID:          entity.ID,
		RecipeID:    entity.RecipeID,
		Description: entity.Description,
		Order:       entity.Order,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (repo *CookStepRepository) FindByRecipeID(recipeID int) (res []entity.CookStep, err error) {
	recs := []CookStep{}
	err = repo.DB.Where("recipe_id = ?", recipeID).Find(&recs).Error
	if err != nil {
		return res, err
	}

	for _, rec := range recs {
		res = append(res, repo.ToEntity(&rec))
	}

	return res, err
}

func (repo *CookStepRepository) Store(cookStep *entity.CookStep) (res entity.CookStep, err error) {
	rec := repo.ToRecord(cookStep)
	err = repo.DB.Create(&rec).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}
