package entity

import "time"

type Material struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	ImageFileID int       `json:"image_file_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
