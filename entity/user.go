package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
