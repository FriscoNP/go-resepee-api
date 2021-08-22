package response

import "go-resepee-api/entity"

type ReviewResponse struct {
	User        UserResponse `json:"user"`
	Description string       `json:"description"`
	Rating      float64      `json:"rating"`
	CreatedAt   string       `json:"created_at"`
}

func CreateReviewResponse(entity *entity.Review) ReviewResponse {
	joinedAt := entity.CreatedAt.Format("02 January 2006")
	return ReviewResponse{
		User:        CreateUserResponse(&entity.UserEntity),
		Description: entity.Description,
		Rating:      float64(entity.Rating),
		CreatedAt:   joinedAt,
	}
}
