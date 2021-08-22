package repository

import (
	"context"
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UserRepositoryInterface interface {
	FindByID(id int) (res entity.User, err error)
	FindByEmail(email string) (res entity.User, err error)
	Store(user *entity.User) (res entity.User, err error)
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

func ToRecord(entity *entity.User) User {
	return User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

func (repo *UserRepository) ToEntity(rec *User) entity.User {
	return entity.User{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func (repo *UserRepository) FindByID(id int) (res entity.User, err error) {
	rec := User{}
	err = repo.DB.Find(&rec, id).Error
	if err != nil {
		return res, err
	}
	return repo.ToEntity(&rec), err
}

func (repo *UserRepository) FindByEmail(email string) (res entity.User, err error) {
	rec := User{}
	err = repo.DB.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}

func (repo *UserRepository) Store(user *entity.User) (res entity.User, err error) {
	rec := User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	err = repo.DB.Create(&rec).Error
	if err != nil {
		return res, err
	}
	return repo.ToEntity(&rec), err
}
