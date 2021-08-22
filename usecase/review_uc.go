package usecase

import (
	"context"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"

	log "github.com/sirupsen/logrus"
)

type ReviewUC struct {
	Context          context.Context
	ReviewRepository repository.ReviewRepositoryInterface
	RecipeRepository repository.RecipeRepositoryInterface
}

type ReviewUCInterface interface {
	Store(req *request.ReviewRequest, userID int) (res entity.Review, err error)
	FindByRecipeID(recipeID int) (res []entity.Review, err error)
}

func NewReviewUC(ctx context.Context, reviewRepo repository.ReviewRepositoryInterface, recipeRepo repository.RecipeRepositoryInterface) ReviewUCInterface {
	return &ReviewUC{
		Context:          ctx,
		ReviewRepository: reviewRepo,
		RecipeRepository: recipeRepo,
	}
}

func (uc *ReviewUC) Store(req *request.ReviewRequest, userID int) (res entity.Review, err error) {
	recipe, err := uc.RecipeRepository.FindByID(req.RecipeID)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}
	_, totalReview, err := uc.ReviewRepository.FindByRecipeID(req.RecipeID)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	// count new average rating
	totalRating := (float64(totalReview) * recipe.AverageRating) + float64(req.Rating)
	newAverageRating := totalRating / (float64(totalReview) + 1)

	err = uc.RecipeRepository.UpdateAverageRating(req.RecipeID, newAverageRating)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	review := entity.Review{
		RecipeID:    uint(req.RecipeID),
		Description: req.Description,
		Rating:      req.Rating,
		UserID:      uint(userID),
	}

	err = uc.ReviewRepository.Store(&review)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return review, err
}

func (uc *ReviewUC) FindByRecipeID(recipeID int) (res []entity.Review, err error) {
	res, _, err = uc.ReviewRepository.FindByRecipeID(recipeID)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}
