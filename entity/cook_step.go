package entity

import "time"

type CookStep struct {
	ID          uint
	RecipeID    uint
	Description string
	Order       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
