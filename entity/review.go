package entity

import "time"

type Review struct {
	ID          uint
	UserID      uint
	UserEntity  User
	Description string
	Rating      int
	RecipeID    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
