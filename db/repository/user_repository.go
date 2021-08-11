package repository

import (
	"context"
	"go-resepee-api/entity"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindByID(user *entity.User, id int) error
	FindByEmail(user *entity.User, email string) error
	Store(user *entity.User) error
}

type UserRepository struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewUserRepository(ctx context.Context, db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		ctx: ctx,
		DB:  db,
	}
}

func (repo *UserRepository) FindByID(user *entity.User, id int) error {
	err := repo.DB.Find(user, id).Error
	if err != nil {
		return err
	}
	return err
}

func (repo *UserRepository) FindByEmail(user *entity.User, email string) error {
	err := repo.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return err
	}
	return err
}

func (repo *UserRepository) Store(user *entity.User) error {
	err := repo.DB.Create(user).Error
	if err != nil {
		return err
	}
	return err
}
