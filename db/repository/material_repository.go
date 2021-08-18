package repository

import (
	"context"
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
)

type Material struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	ImageFileID int
	ImageFile   File
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type MaterialRepositoryInterface interface {
	Get() (res []entity.Material, err error)
	Store(material *entity.Material) (res entity.Material, err error)
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

func (repo *MaterialRepository) ToEntity(rec *Material) entity.Material {
	return entity.Material{
		ID:          rec.ID,
		Name:        rec.Name,
		ImageFileID: int(rec.ImageFileID),
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
	}
}

func (repo *MaterialRepository) Get() (res []entity.Material, err error) {
	materials := []Material{}
	err = repo.DB.Find(&materials).Error
	if err != nil {
		return res, err
	}

	for _, material := range materials {
		res = append(res, repo.ToEntity(&material))
	}

	return res, err
}

func (repo *MaterialRepository) Store(material *entity.Material) (res entity.Material, err error) {
	rec := Material{
		Name:        material.Name,
		ImageFileID: material.ImageFileID,
	}
	err = repo.DB.Create(&rec).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}
