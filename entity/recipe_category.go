package entity

import (
	"time"

	"gorm.io/gorm"
)

type RecipeCategory struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
