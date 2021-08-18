package entity

import (
	"time"

	"gorm.io/gorm"
)

type Material struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name"`
	ImageFileID     int            `json:"image_file_id"`
	ImageFileEntity File           `json:"image_file"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
