package repository

import (
	"context"
	"fmt"
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
)

type RecipeMaterial struct {
	ID         uint `gorm:"primaryKey"`
	RecipeID   uint
	MaterialID uint
	Material   Material
	Amount     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type RecipeMaterialRepository struct {
	Context context.Context
	DB      *gorm.DB
}

type RecipeMaterialRepositoryInterface interface {
	FindByRecipeID(recipeID int) (res []entity.RecipeMaterial, err error)
	Store(recipeMaterial *entity.RecipeMaterial) (res entity.RecipeMaterial, err error)
}

func NewRecipeMaterialRepository(ctx context.Context, db *gorm.DB) RecipeMaterialRepositoryInterface {
	return &RecipeMaterialRepository{
		Context: ctx,
		DB:      db,
	}
}

func (repo *RecipeMaterialRepository) ToEntity(rec *RecipeMaterial) entity.RecipeMaterial {
	materialRepo := MaterialRepository{}
	return entity.RecipeMaterial{
		ID:             rec.ID,
		RecipeID:       rec.RecipeID,
		MaterialID:     rec.MaterialID,
		MaterialEntity: materialRepo.ToEntity(&rec.Material),
		Amount:         rec.Amount,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
	}
}

func (repo *RecipeMaterialRepository) ToRecord(entity *entity.RecipeMaterial) RecipeMaterial {
	return RecipeMaterial{
		ID:         entity.ID,
		RecipeID:   entity.RecipeID,
		MaterialID: entity.MaterialID,
		Amount:     entity.Amount,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}

func (repo *RecipeMaterialRepository) FindByRecipeID(recipeID int) (res []entity.RecipeMaterial, err error) {
	recs := []RecipeMaterial{}
	err = repo.DB.Preload("Material.ImageFile").Where("recipe_id = ?", recipeID).Find(&recs).Error
	if err != nil {
		return res, err
	}

	for _, rec := range recs {
		fmt.Println(rec)
		res = append(res, repo.ToEntity(&rec))
	}

	return res, err
}

func (repo *RecipeMaterialRepository) Store(recipeMaterial *entity.RecipeMaterial) (res entity.RecipeMaterial, err error) {
	rec := repo.ToRecord(recipeMaterial)
	err = repo.DB.Preload("Material.ImageFile").Create(&rec).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}
