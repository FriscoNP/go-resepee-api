package response

type ReviewResponse struct {
	User        UserResponse `json:"user"`
	Description string       `json:"description"`
	Rating      float64      `json:"rating"`
	CreatedAt   string       `json:"created_at"`
}
