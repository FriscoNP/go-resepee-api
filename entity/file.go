package entity

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        int
	Type      string
	Path      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
