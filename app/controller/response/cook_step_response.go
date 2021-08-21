package response

import "go-resepee-api/entity"

type CookStepResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}

func CreateCookStepResponse(entity *entity.CookStep) CookStepResponse {
	return CookStepResponse{
		ID:          int(entity.ID),
		Description: entity.Description,
		Order:       entity.Order,
	}
}
