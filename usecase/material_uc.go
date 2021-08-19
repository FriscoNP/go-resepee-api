package usecase

import (
	"context"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"
)

type MaterialUC struct {
	Context            context.Context
	MaterialRepository repository.MaterialRepositoryInterface
}

type MaterialUCInterface interface {
	Get() ([]entity.Material, error)
	Store(req *request.CreateMaterialRequest) (entity.Material, error)
}

func NewMaterialUC(ctx context.Context, repo repository.MaterialRepositoryInterface) MaterialUCInterface {
	return &MaterialUC{
		Context:            ctx,
		MaterialRepository: repo,
	}
}

func (uc *MaterialUC) Get() (res []entity.Material, err error) {
	res, err = uc.MaterialRepository.Get()
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc *MaterialUC) Store(req *request.CreateMaterialRequest) (res entity.Material, err error) {
	material := entity.Material{
		Name:        req.Name,
		ImageFileID: req.ImageFileID,
	}

	res, err = uc.MaterialRepository.Store(&material)
	if err != nil {
		return res, err
	}

	return res, err
}
