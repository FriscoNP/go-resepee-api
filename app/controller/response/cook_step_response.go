package response

type CookStepResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}
