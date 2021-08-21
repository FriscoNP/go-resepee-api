package entity

import (
	"time"

	"gorm.io/gorm"
)

type Material struct {
	ID              uint
	Name            string
	ImageFileID     int
	ImageFileEntity File
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
