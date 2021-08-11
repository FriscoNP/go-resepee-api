package entity

import "time"

type ReviewEntity struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	RecipeID    int       `json:"recipe_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
