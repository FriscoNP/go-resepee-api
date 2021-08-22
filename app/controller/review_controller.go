package controller

import (
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/controller/response"
	"go-resepee-api/app/middleware"
	"go-resepee-api/db/repository"
	"go-resepee-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReviewController struct {
	DB *gorm.DB
}

func NewReviewController(db *gorm.DB) *ReviewController {
	return &ReviewController{
		DB: db,
	}
}

func (controller *ReviewController) FindByRecipeID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.QueryParam("recipe_id")
	recipeID, err := strconv.Atoi(id)
	if err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusBadRequest, err)
	}

	reviewRepository := repository.NewReviewRepository(controller.DB)
	recipeRepository := repository.NewRecipeRepository(ctx, controller.DB)
	reviewUC := usecase.NewReviewUC(ctx, reviewRepository, recipeRepository)

	reviews, err := reviewUC.FindByRecipeID(recipeID)
	if err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusBadRequest, err)
	}

	res := []response.ReviewResponse{}
	for _, review := range reviews {
		res = append(res, response.CreateReviewResponse(&review))
	}

	return SendSuccess(c, res, "get_recipe_reviews")
}

func (controller *ReviewController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	jwtUser := middleware.GetJWTUser(c)

	req := request.ReviewRequest{}
	if err := c.Bind(&req); err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusBadRequest, err)
	}

	tx := controller.DB.Begin()
	reviewRepository := repository.NewReviewRepository(tx)
	recipeRepository := repository.NewRecipeRepository(ctx, tx)
	reviewUC := usecase.NewReviewUC(ctx, reviewRepository, recipeRepository)
	review, err := reviewUC.Store(&req, jwtUser.ID)
	if err != nil {
		tx.Rollback()
		log.Warn(err.Error())
		return SendError(c, http.StatusBadRequest, err)
	}

	tx.Commit()
	return SendSuccess(c, response.CreateReviewResponse(&review), "review_created")
}
