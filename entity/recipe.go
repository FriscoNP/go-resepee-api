package entity

import "time"

type RecipeEntity struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"decription"`
	ThumbnailFileID  int       `json:"thumbnail_file_id"`
	RecipeCategoryID int       `json:"recipe_category_id"`
	UserID           int       `json:"user_id"`
	AverageRating    float64   `json:"average_rating"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
}
