package entity

import "time"

type RecipeMaterialEntity struct {
	ID         int       `json:"id"`
	RecipeID   int       `json:"recipe_id"`
	MaterialID int       `json:"material_id"`
	Amount     string    `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
