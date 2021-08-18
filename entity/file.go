package entity

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        int            `json:"id"`
	Type      string         `json:"type"`
	Path      string         `json:"path"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
