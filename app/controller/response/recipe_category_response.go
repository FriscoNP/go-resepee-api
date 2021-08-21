package response

import "go-resepee-api/entity"

type RecipeCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateRecipeCategoryResponse(entity *entity.RecipeCategory) RecipeCategoryResponse {
	return RecipeCategoryResponse{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
