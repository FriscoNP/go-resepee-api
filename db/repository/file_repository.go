package repository

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        int `gorm:"primaryKey"`
	Type      string
	Path      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
