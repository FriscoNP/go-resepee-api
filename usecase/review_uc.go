package usecase

import (
	"context"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReviewUC struct {
	Context context.Context
	DB      *gorm.DB
}

type ReviewUCInterface interface {
	Store(req *request.ReviewRequest) (res entity.Review, err error)
	FindByRecipeID(recipeID int) (res []entity.Review, err error)
}

func NewReviewUC(ctx context.Context, db *gorm.DB) ReviewUCInterface {
	return &ReviewUC{
		Context: ctx,
		DB:      db,
	}
}

func (uc *ReviewUC) Store(req *request.ReviewRequest) (res entity.Review, err error) {
	reviewRepo := repository.NewReviewRepository(uc.DB)
	recipeRepo := repository.NewRecipeRepository(uc.Context, uc.DB)

	recipe, err := recipeRepo.FindByID(req.RecipeID)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}
	_, totalReview, err := reviewRepo.FindByRecipeID(req.RecipeID)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	// count new average rating
	totalRating := (float64(totalReview) * recipe.AverageRating) + float64(req.Rating)
	newAverageRating := totalRating / (float64(totalReview) + 1)

	err = recipeRepo.UpdateAverageRating(req.RecipeID, newAverageRating)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	res.RecipeID = uint(req.RecipeID)
	res.Description = req.Description
	res.Rating = req.Rating

	err = reviewRepo.Store(&res)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}

func (uc *ReviewUC) FindByRecipeID(recipeID int) (res []entity.Review, err error) {
	reviewRepo := repository.NewReviewRepository(uc.DB)

	res, _, err = reviewRepo.FindByRecipeID(recipeID)
	if err != nil {
		return res, err
	}

	return res, err
}
