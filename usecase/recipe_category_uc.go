package usecase

import (
	"context"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RecipeCategoryUC struct {
	Context context.Context
	DB      *gorm.DB
}

type RecipeCategoryUCInterface interface {
	GetAll() (res []entity.RecipeCategory, err error)
	Store(req *request.RecipeCategoryRequest) (res entity.RecipeCategory, err error)
}

func NewRecipeCategoryUC(ctx context.Context, db *gorm.DB) RecipeCategoryUCInterface {
	return &RecipeCategoryUC{
		Context: ctx,
		DB:      db,
	}
}

func (uc *RecipeCategoryUC) GetAll() (res []entity.RecipeCategory, err error) {
	recipeCategoryRepo := repository.NewRecipeCategoryRepository(uc.Context, uc.DB)

	res, err = recipeCategoryRepo.GetAll()
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}

func (uc *RecipeCategoryUC) Store(req *request.RecipeCategoryRequest) (res entity.RecipeCategory, err error) {
	recipeCategoryRepo := repository.NewRecipeCategoryRepository(uc.Context, uc.DB)

	res.Name = req.Name
	err = recipeCategoryRepo.Store(&res)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}
