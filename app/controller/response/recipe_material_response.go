package response

import "go-resepee-api/entity"

type RecipeMaterialResponse struct {
	Material MaterialResponse `json:"material"`
	Amount   string           `json:"amount"`
}

func CreateRecipeMaterialResponse(entity *entity.RecipeMaterial) RecipeMaterialResponse {
	return RecipeMaterialResponse{
		Material: CreateMaterialResponse(&entity.MaterialEntity),
		Amount:   entity.Amount,
	}
}
