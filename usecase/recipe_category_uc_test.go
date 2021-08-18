package usecase_test

import (
	"context"
	"errors"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/db/repository/mocks"
	"go-resepee-api/entity"
	"go-resepee-api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	recipeCategoryRepository mocks.RecipeCategoryRepositoryInterface

	recipeCategoryUC          = usecase.NewRecipeCategoryUC(context.Background(), &recipeCategoryRepository)
	mockRecipeCategoryRequest = request.RecipeCategoryRequest{Name: "Sarapan"}
)

func TestGetAllRecipeCategory(t *testing.T) {
	t.Run("get all recipe category happy case", func(t *testing.T) {
		recipeCategoryRepository.On("GetAll").Return([]entity.RecipeCategory{{Name: "Sarapan"}}, nil).Once()

		res, err := recipeCategoryUC.GetAll()
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("get all recipe category error", func(t *testing.T) {
		recipeCategoryRepository.On("GetAll").Return([]entity.RecipeCategory{}, errors.New("get all recipe category error")).Once()

		res, err := recipeCategoryUC.GetAll()
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestStoreRecipeCategory(t *testing.T) {
	t.Run("store recipe category happy case", func(t *testing.T) {
		recipeCategoryRepository.On("Store", mock.AnythingOfType("*entity.RecipeCategory")).Return(entity.RecipeCategory{ID: 1}, nil).Once()

		res, err := recipeCategoryUC.Store(&mockRecipeCategoryRequest)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("store recipe category error", func(t *testing.T) {
		recipeCategoryRepository.On("Store", mock.AnythingOfType("*entity.RecipeCategory")).Return(entity.RecipeCategory{}, errors.New("store recipe category error")).Once()

		res, err := recipeCategoryUC.Store(&mockRecipeCategoryRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
