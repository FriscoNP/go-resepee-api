package response

import "time"

type MaterialResponse struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	ImageFileID int          `json:"image_file_id"`
	ImageFile   FileResponse `json:"image_file"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
