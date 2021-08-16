package entity

import "time"

type CookStep struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	RecipeID    uint      `json:"recipe_id"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
