package request

type CreateMaterialRequest struct {
	Name        string `json:"name"`
	ImageFileID int    `json:"image_file_id"`
}
