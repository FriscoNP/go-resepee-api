package entity

import "time"

type Review struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	UserEntity  User      `json:"user"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	RecipeID    uint      `json:"recipe_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
