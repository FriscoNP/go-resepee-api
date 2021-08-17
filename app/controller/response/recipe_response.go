package response

import "time"

type RecipeResponse struct {
	ID               int                      `json:"id"`
	Title            string                   `json:"title"`
	Description      string                   `json:"description"`
	ThumbnailFileID  int                      `json:"thumbnail_file_id"`
	ThumbnailFile    FileResponse             `json:"thumbnail_file"`
	RecipeCategoryID int                      `json:"recipe_category_id"`
	RecipeCategory   RecipeCategoryResponse   `json:"recipe_category"`
	UserID           int                      `json:"user_id"`
	User             RecipeUserResponse       `json:"user"`
	AverageRating    float64                  `json:"average_rating"`
	Materials        []RecipeMaterialResponse `json:"materials"`
	CookSteps        []CookStepResponse       `json:"cook_steps"`
}

type RecipeCategoryResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RecipeUserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RecipeMaterialResponse struct {
	Material MaterialResponse `json:"material"`
	Amount   string           `json:"amount"`
}
