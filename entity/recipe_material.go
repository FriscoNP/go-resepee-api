package entity

import "time"

type RecipeMaterial struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	RecipeID   uint      `json:"recipe_id"`
	MaterialID uint      `json:"material_id"`
	Amount     string    `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
