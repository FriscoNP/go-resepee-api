package repository

import (
	"context"
	"go-resepee-api/entity"

	"gorm.io/gorm"
)

type MaterialRepositoryInterface interface {
	Get() (res []entity.Material, err error)
	Store(material *entity.Material) error
}

type MaterialRepository struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewMaterialRepository(ctx context.Context, db *gorm.DB) MaterialRepositoryInterface {
	return &MaterialRepository{
		ctx: ctx,
		DB:  db,
	}
}

func (repo *MaterialRepository) Get() (res []entity.Material, err error) {
	err = repo.DB.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo *MaterialRepository) Store(material *entity.Material) error {
	err := repo.DB.Create(material).Error
	if err != nil {
		return err
	}

	return err
}
