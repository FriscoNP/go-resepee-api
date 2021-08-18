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
	reviewRepository mocks.ReviewRepositoryInterface
	recipeRepository mocks.RecipeRepositoryInterface

	reviewUC = usecase.NewReviewUC(context.Background(), &reviewRepository, &recipeRepository)

	mockReviewRequest = request.ReviewRequest{
		RecipeID:    1,
		Description: "test",
		Rating:      4,
	}
)

func TestStoreReview(t *testing.T) {
	t.Run("store review happy case", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{ID: 1, AverageRating: 4}, nil).Once()
		reviewRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.Review{{ID: 1}, {ID: 2}}, 2, nil).Once()
		recipeRepository.On("UpdateAverageRating", mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(nil).Once()
		reviewRepository.On("Store", mock.AnythingOfType("*entity.Review")).Return(nil).Once()

		res, err := reviewUC.Store(&mockReviewRequest, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("store review error find recipe", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{}, errors.New("find recipe error")).Once()
		// reviewRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.Review{{ID: 1}, {ID: 2}}, 2, nil)
		// recipeRepository.On("UpdateAverageRating", mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(nil).Once()
		// reviewRepository.On("Store", mock.AnythingOfType("*entity.Review")).Return(nil).Once()

		res, err := reviewUC.Store(&mockReviewRequest, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("store review error find recipe reviews", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{ID: 1, AverageRating: 4}, nil).Once()
		reviewRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.Review{}, 0, errors.New("find recipe review error")).Once()
		// recipeRepository.On("UpdateAverageRating", mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(nil).Once()
		// reviewRepository.On("Store", mock.AnythingOfType("*entity.Review")).Return(nil).Once()

		res, err := reviewUC.Store(&mockReviewRequest, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("store review error update recipe average rating", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{ID: 1, AverageRating: 4}, nil).Once()
		reviewRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.Review{{ID: 1}, {ID: 2}}, 2, nil).Once()
		recipeRepository.On("UpdateAverageRating", mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(errors.New("update average rating error")).Once()
		// reviewRepository.On("Store", mock.AnythingOfType("*entity.Review")).Return(nil).Once()

		res, err := reviewUC.Store(&mockReviewRequest, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("store review error store data", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{ID: 1, AverageRating: 4}, nil).Once()
		reviewRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.Review{{ID: 1}, {ID: 2}}, 2, nil).Once()
		recipeRepository.On("UpdateAverageRating", mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(nil).Once()
		reviewRepository.On("Store", mock.AnythingOfType("*entity.Review")).Return(errors.New("store review error")).Once()

		res, err := reviewUC.Store(&mockReviewRequest, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestFindByRecipeID(t *testing.T) {
	t.Run("find review by recipe id happy case", func(t *testing.T) {
		reviewRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.Review{{ID: 1}, {ID: 2}}, 2, nil).Once()

		res, err := reviewUC.FindByRecipeID(1)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("find review by recipe id error", func(t *testing.T) {
		reviewRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.Review{}, 0, errors.New("find by recipe id error")).Once()

		res, err := reviewUC.FindByRecipeID(1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
