package repository

import (
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *ReviewRepository) ToRecord(entity *entity.Review) Review {
	return Review{
		ID:          entity.ID,
		UserID:      entity.UserID,
		Description: entity.Description,
		Rating:      entity.Rating,
		RecipeID:    entity.RecipeID,
	}
}

func (repo *ReviewRepository) ToEntity(rec *Review) entity.Review {
	userRepo := UserRepository{}
	return entity.Review{
		ID:          rec.ID,
		UserID:      rec.UserID,
		UserEntity:  userRepo.ToEntity(&rec.User),
		Description: rec.Description,
		Rating:      rec.Rating,
		RecipeID:    rec.RecipeID,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func (repo *ReviewRepository) Store(review *entity.Review) error {
	rec := repo.ToRecord(review)
	err := repo.DB.Create(&rec).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *ReviewRepository) FindByRecipeID(recipeID int) (res []entity.Review, count int, err error) {
	recs := []Review{}
	err = repo.DB.Preload(clause.Associations).Where("recipe_id = ?", recipeID).Find(&recs).Error
	if err != nil {
		return res, count, err
	}

	for _, rec := range recs {
		res = append(res, repo.ToEntity(&rec))
	}
	return res, len(res), err
}
