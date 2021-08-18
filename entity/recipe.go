package entity

import (
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID                   uint             `json:"id" gorm:"primaryKey"`
	Title                string           `json:"title"`
	Description          string           `json:"decription"`
	ThumbnailFileID      int              `json:"thumbnail_file_id"`
	ThumbnailFileEntity  File             `json:"thumbnail_file"`
	RecipeCategoryID     int              `json:"recipe_category_id"`
	RecipeCategoryEntity RecipeCategory   `json:"recipe_category"`
	UserID               int              `json:"user_id"`
	UserEntity           User             `json:"user"`
	AverageRating        float64          `json:"average_rating"`
	CreatedAt            time.Time        `json:"created_at"`
	UpdatedAt            time.Time        `json:"updated_at"`
	DeletedAt            gorm.DeletedAt   `json:"deleted_at"`
	RecipeMaterials      []RecipeMaterial `json:"recipe_materials"`
	CookSteps            []CookStep       `json:"cook_steps"`
}
