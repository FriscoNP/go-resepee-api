package controller

import (
	"go-resepee-api/app/controller/request"
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
		return SendError(c, http.StatusInternalServerError, err)
	}

	reviewUC := usecase.ReviewUC{
		Context: ctx,
		DB:      controller.DB,
	}

	reviews, err := reviewUC.FindByRecipeID(recipeID)
	if err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusInternalServerError, err)
	}

	return SendSuccess(c, reviews, "get_recipe_reviews")
}

func (controller *ReviewController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.ReviewRequest{}
	if err := c.Bind(&req); err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusBadRequest, err)
	}

	tx := controller.DB.Begin()
	reviewUC := usecase.NewReviewUC(ctx, tx)
	review, err := reviewUC.Store(&req)
	if err != nil {
		tx.Rollback()
		log.Warn(err.Error())
		return SendError(c, http.StatusInternalServerError, err)
	}

	tx.Commit()
	return SendSuccess(c, review, "review_created")
}
