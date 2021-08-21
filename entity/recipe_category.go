package entity

import (
	"time"

	"gorm.io/gorm"
)

type RecipeCategory struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
