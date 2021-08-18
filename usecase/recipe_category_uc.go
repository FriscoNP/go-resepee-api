package usecase

import (
	"context"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"

	log "github.com/sirupsen/logrus"
)

type RecipeCategoryUC struct {
	Context    context.Context
	Repository repository.RecipeCategoryRepositoryInterface
}

type RecipeCategoryUCInterface interface {
	GetAll() (res []entity.RecipeCategory, err error)
	Store(req *request.RecipeCategoryRequest) (res entity.RecipeCategory, err error)
}

func NewRecipeCategoryUC(ctx context.Context, repo repository.RecipeCategoryRepositoryInterface) RecipeCategoryUCInterface {
	return &RecipeCategoryUC{
		Context:    ctx,
		Repository: repo,
	}
}

func (uc *RecipeCategoryUC) GetAll() (res []entity.RecipeCategory, err error) {
	res, err = uc.Repository.GetAll()
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}

func (uc *RecipeCategoryUC) Store(req *request.RecipeCategoryRequest) (res entity.RecipeCategory, err error) {
	res.Name = req.Name
	res, err = uc.Repository.Store(&res)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}
