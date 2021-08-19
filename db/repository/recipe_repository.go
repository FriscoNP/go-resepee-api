package repository

import (
	"context"
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Recipe struct {
	ID               uint `gorm:"primaryKey"`
	Title            string
	Description      string
	ThumbnailFileID  int
	ThumbnailFile    File
	RecipeCategoryID int
	RecipeCategory   RecipeCategory
	UserID           int
	User             User
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
	Store(recipe *entity.Recipe) (res entity.Recipe, err error)
	UpdateAverageRating(recipeID int, averageRating float64) error
}

func NewRecipeRepository(ctx context.Context, db *gorm.DB) RecipeRepositoryInterface {
	return &RecipeRepository{
		Context: ctx,
		DB:      db,
	}
}

func (repo *RecipeRepository) ToEntity(rec *Recipe) entity.Recipe {
	fileRepo := FileRepository{}
	recipeCategoryRepo := RecipeCategoryRepository{}
	userRepo := UserRepository{}

	return entity.Recipe{
		ID:                   rec.ID,
		Title:                rec.Title,
		Description:          rec.Description,
		ThumbnailFileID:      int(rec.ThumbnailFileID),
		ThumbnailFileEntity:  fileRepo.ToEntity(&rec.ThumbnailFile),
		RecipeCategoryID:     int(rec.RecipeCategoryID),
		RecipeCategoryEntity: recipeCategoryRepo.ToEntity(&rec.RecipeCategory),
		UserID:               int(rec.UserID),
		UserEntity:           userRepo.ToEntity(&rec.User),
		AverageRating:        rec.AverageRating,
		CreatedAt:            rec.CreatedAt,
		UpdatedAt:            rec.UpdatedAt,
		DeletedAt:            rec.DeletedAt,
	}
}

func (repo *RecipeRepository) ToRecord(entity *entity.Recipe) Recipe {
	return Recipe{
		ID:               entity.ID,
		Title:            entity.Title,
		Description:      entity.Description,
		ThumbnailFileID:  entity.ThumbnailFileID,
		RecipeCategoryID: entity.RecipeCategoryID,
		UserID:           entity.UserID,
		AverageRating:    entity.AverageRating,
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
		DeletedAt:        entity.DeletedAt,
	}
}

func (repo *RecipeRepository) GetAll() (res []entity.Recipe, err error) {
	recs := []Recipe{}
	err = repo.DB.Preload(clause.Associations).Find(&recs).Error
	if err != nil {
		return res, err
	}

	for _, rec := range recs {
		res = append(res, repo.ToEntity(&rec))
	}

	return res, err
}

func (repo *RecipeRepository) FindByID(id int) (res entity.Recipe, err error) {
	rec := Recipe{}
	err = repo.DB.Preload(clause.Associations).Find(&rec, id).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}

func (repo *RecipeRepository) Store(recipe *entity.Recipe) (res entity.Recipe, err error) {
	rec := repo.ToRecord(recipe)
	err = repo.DB.Create(&rec).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}

func (repo *RecipeRepository) UpdateAverageRating(recipeID int, averageRating float64) error {
	rec := Recipe{}
	err := repo.DB.Model(&rec).Where("id = ?", recipeID).Update("average_rating", averageRating).Error
	if err != nil {
		return err
	}

	return nil
}
