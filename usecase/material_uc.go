package usecase

import (
	"context"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/middleware"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"

	"gorm.io/gorm"
)

type MaterialUC struct {
	Context context.Context
	DB      *gorm.DB
	JwtAuth *middleware.ConfigJWT
}

type MaterialUCInterface interface {
	Get() ([]entity.Material, error)
	Store(req *request.CreateMaterialRequest) (entity.Material, error)
}

func NewMaterialUC(ctx context.Context, db *gorm.DB, jwtAuth *middleware.ConfigJWT) MaterialUCInterface {
	return &MaterialUC{
		Context: ctx,
		DB:      db,
		JwtAuth: jwtAuth,
	}
}

func (uc *MaterialUC) Get() (res []entity.Material, err error) {
	materialRepo := repository.NewMaterialRepository(uc.Context, uc.DB)

	res, err = materialRepo.Get()
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc *MaterialUC) Store(req *request.CreateMaterialRequest) (res entity.Material, err error) {
	materialRepo := repository.NewMaterialRepository(uc.Context, uc.DB)

	material := entity.Material{
		Name:        req.Name,
		ImageFileID: req.ImageFileID,
	}

	err = materialRepo.Store(&material)
	if err != nil {
		return res, err
	}

	return material, err
}
