package repository

import (
	"go-resepee-api/entity"
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        int `gorm:"primaryKey"`
	Type      string
	Path      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type FileRepository struct {
	DB *gorm.DB
}

type FileRepositoryInterface interface {
	FindByID(id int) (res entity.File, err error)
	Store(file *entity.File) (res entity.File, err error)
}

func NewFileRepository(db *gorm.DB) FileRepositoryInterface {
	return &FileRepository{
		DB: db,
	}
}

func (repo *FileRepository) ToEntity(rec *File) entity.File {
	return entity.File{
		ID:        rec.ID,
		Type:      rec.Type,
		Path:      rec.Path,
		CreatedAt: rec.CreatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func (repo *FileRepository) ToRecord(entity *entity.File) File {
	return File{
		ID:        entity.ID,
		Type:      entity.Type,
		Path:      entity.Path,
		CreatedAt: entity.CreatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

func (repo *FileRepository) FindByID(id int) (res entity.File, err error) {
	rec := File{}
	err = repo.DB.Find(&rec, id).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}

func (repo *FileRepository) Store(file *entity.File) (res entity.File, err error) {
	rec := repo.ToRecord(file)
	err = repo.DB.Create(&rec).Error
	if err != nil {
		return res, err
	}

	return repo.ToEntity(&rec), err
}
