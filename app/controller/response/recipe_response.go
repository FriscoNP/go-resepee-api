package response

import "go-resepee-api/entity"

type RecipeResponse struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	ThumbnailPath string  `json:"thumbnail_path"`
	Category      string  `json:"category"`
	CreatedBy     string  `json:"created_by"`
	AverageRating float64 `json:"average_rating"`
}

type RecipeDetailResponse struct {
	ID            int                      `json:"id"`
	Title         string                   `json:"title"`
	Description   string                   `json:"description"`
	ThumbnailPath string                   `json:"thumbnail_path"`
	Category      string                   `json:"category"`
	CreatedBy     string                   `json:"created_by"`
	AverageRating float64                  `json:"average_rating"`
	Materials     []RecipeMaterialResponse `json:"materials"`
	CookSteps     []CookStepResponse       `json:"cook_steps"`
}

func CreateRecipeResponse(entity *entity.Recipe) RecipeResponse {
	return RecipeResponse{
		ID:            int(entity.ID),
		Title:         entity.Title,
		Description:   entity.Description,
		ThumbnailPath: entity.ThumbnailFileEntity.Path,
		Category:      entity.RecipeCategoryEntity.Name,
		CreatedBy:     entity.UserEntity.Name,
		AverageRating: entity.AverageRating,
	}
}

func CreateRecipeDetailResponse(entity *entity.Recipe) RecipeDetailResponse {
	materials := []RecipeMaterialResponse{}
	for _, material := range entity.RecipeMaterials {
		materials = append(materials, CreateRecipeMaterialResponse(&material))
	}

	cookSteps := []CookStepResponse{}
	for _, cook := range entity.CookSteps {
		cookSteps = append(cookSteps, CreateCookStepResponse(&cook))
	}

	return RecipeDetailResponse{
		ID:            int(entity.ID),
		Title:         entity.Title,
		Description:   entity.Description,
		ThumbnailPath: entity.ThumbnailFileEntity.Path,
		Category:      entity.RecipeCategoryEntity.Name,
		CreatedBy:     entity.UserEntity.Name,
		AverageRating: entity.AverageRating,
		Materials:     materials,
		CookSteps:     cookSteps,
	}
}
