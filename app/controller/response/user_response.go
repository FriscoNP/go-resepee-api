package response

import (
	"go-resepee-api/entity"
)

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	JoinedAt string `json:"joined_at"`
}

func CreateUserResponse(entity *entity.User) UserResponse {
	joinedAt := entity.CreatedAt.Format("02 January 2006")
	return UserResponse{
		ID:       int(entity.ID),
		Name:     entity.Name,
		Email:    entity.Email,
		JoinedAt: joinedAt,
	}
}
