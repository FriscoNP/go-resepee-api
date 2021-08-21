package response

import "go-resepee-api/entity"

type MaterialResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}

func CreateMaterialResponse(entity *entity.Material) MaterialResponse {
	return MaterialResponse{
		ID:        int(entity.ID),
		Name:      entity.Name,
		ImagePath: entity.ImageFileEntity.Path,
	}
}
