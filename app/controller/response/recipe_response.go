package response

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

type RecipeMaterialResponse struct {
	Material MaterialResponse `json:"material"`
	Amount   string           `json:"amount"`
}
