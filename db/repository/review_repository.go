package repository

import (
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	User        User
	Description string
	Rating      int
	RecipeID    uint
	Recipe      Recipe
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ReviewRepository struct {
	DB *gorm.DB
}

type ReviewRepositoryInterface interface {
	Store(review *entity.Review) error
	FindByRecipeID(recipeID int) (res []entity.Review, count int, err error)
}

func NewReviewRepository(db *gorm.DB) ReviewRepositoryInterface {
	return &ReviewRepository{
		DB: db,
	}
}

func (repo *ReviewRepository) Store(review *entity.Review) error {
	err := repo.DB.Create(review).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *ReviewRepository) FindByRecipeID(recipeID int) (res []entity.Review, count int, err error) {
	recs := []Review{}
	err = repo.DB.Where("recipe_id = ?", recipeID).Find(&recs).Error
	if err != nil {
		return res, count, err
	}

	for _, rec := range recs {
		res = append(res, entity.Review{
			ID:          uint(rec.ID),
			UserID:      uint(rec.UserID),
			RecipeID:    uint(rec.RecipeID),
			Description: rec.Description,
			Rating:      rec.Rating,
			CreatedAt:   rec.CreatedAt,
			UpdatedAt:   rec.UpdatedAt,
		})
	}
	return res, len(res), err
}
