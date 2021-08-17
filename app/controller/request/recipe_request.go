package request

type RecipeRequest struct {
	Title            string                     `json:"title"`
	Description      string                     `json:"description"`
	ThumbnailFileID  int                        `json:"thumbnail_file_id"`
	RecipeCategoryID int                        `json:"recipe_category_id"`
	Materials        []RecipeMaterialRequest    `json:"materials"`
	CookingSteps     []RecipeCookingStepRequest `json:"cooking_steps"`
}

type RecipeMaterialRequest struct {
	MaterialID int    `json:"material_id"`
	Amount     string `json:"amount"`
}

type RecipeCookingStepRequest struct {
	Description string `json:"description"`
	Order       int    `json:"order"`
}
