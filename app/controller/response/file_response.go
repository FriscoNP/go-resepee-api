package response

import "go-resepee-api/entity"

type FileResponse struct {
	ID   int    `json:"id"`
	Type string `json:"file_type"`
	Path string `json:"file_path"`
}

func CreateFileResponse(entity *entity.File) FileResponse {
	return FileResponse{
		ID:   entity.ID,
		Type: entity.Type,
		Path: entity.Path,
	}
}
