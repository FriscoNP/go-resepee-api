package request

type ReviewRequest struct {
	RecipeID    int    `json:"recipe_id"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
