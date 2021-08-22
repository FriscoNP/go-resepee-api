package entity

import "time"

type RecipeMaterial struct {
	ID             uint
	RecipeID       uint
	MaterialID     uint
	MaterialEntity Material
	Amount         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
