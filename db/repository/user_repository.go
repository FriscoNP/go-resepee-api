package repository

import (
	"context"
	"go-resepee-api/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindByID(user *entity.UserEntity, id int) error
	FindByEmail(user *entity.UserEntity, email string) error
	Store(user *entity.UserEntity) error
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

func (repo *UserRepository) FindByID(user *entity.UserEntity, id int) error {
	err := repo.DB.Find(user, id).Error
	if err != nil {
		logrus.Warn(err.Error())
		return err
	}
	return err
}

func (repo *UserRepository) FindByEmail(user *entity.UserEntity, email string) error {
	err := repo.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		logrus.Warn(err.Error())
		return err
	}
	return err
}

func (repo *UserRepository) Store(user *entity.UserEntity) error {
	err := repo.DB.Create(user).Error
	if err != nil {
		logrus.Warn(err.Error())
		return err
	}
	return err
}
