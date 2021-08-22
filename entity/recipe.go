package entity

import (
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID                   uint
	Title                string
	Description          string
	ThumbnailFileID      int
	ThumbnailFileEntity  File
	RecipeCategoryID     int
	RecipeCategoryEntity RecipeCategory
	UserID               int
	UserEntity           User
	AverageRating        float64
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt
	RecipeMaterials      []RecipeMaterial
	CookSteps            []CookStep
}
