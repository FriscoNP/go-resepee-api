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
	recipeRepository         mocks.RecipeRepositoryInterface
	recipeMaterialRepository mocks.RecipeMaterialRepositoryInterface
	cookStepRepository       mocks.CookStepRepositoryInterface

	mockRecipeRequest = request.RecipeRequest{
		Title: "test recipe",
		Materials: []request.RecipeMaterialRequest{
			{
				MaterialID: 1,
				Amount:     "1gr",
			},
		},
		CookingSteps: []request.RecipeCookingStepRequest{
			{
				Description: "test cook step",
				Order:       1,
			},
		},
	}

	recipeUC = usecase.NewRecipeUC(context.Background(), &recipeRepository, &recipeMaterialRepository, &cookStepRepository)
)

func TestRecipeGetAll(t *testing.T) {
	t.Run("get all recipe happy case", func(t *testing.T) {
		recipeRepository.On("GetAll").Return([]entity.Recipe{{ID: 1}}, nil).Once()

		res, err := recipeUC.GetAll()
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("get all recipe error", func(t *testing.T) {
		recipeRepository.On("GetAll").Return([]entity.Recipe{}, errors.New("get all recipe error")).Once()

		res, err := recipeUC.GetAll()
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestRecipeFindByID(t *testing.T) {
	t.Run("recipe find by id happy case", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{ID: 1}, nil).Once()
		recipeMaterialRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.RecipeMaterial{{ID: 1}}, nil).Once()
		cookStepRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.CookStep{{ID: 1}}, nil).Once()

		res, err := recipeUC.FindByID(1)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
	t.Run("recipe find by id error find recipe", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{}, errors.New("find recipe error")).Once()

		res, err := recipeUC.FindByID(1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("recipe find by id error get recipe material", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{ID: 1}, nil).Once()
		recipeMaterialRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.RecipeMaterial{}, errors.New("get recipe material error")).Once()

		res, err := recipeUC.FindByID(1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("recipe find by id error get cook step", func(t *testing.T) {
		recipeRepository.On("FindByID", mock.AnythingOfType("int")).Return(entity.Recipe{ID: 1}, nil).Once()
		recipeMaterialRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.RecipeMaterial{{ID: 1}}, nil).Once()
		cookStepRepository.On("FindByRecipeID", mock.AnythingOfType("int")).Return([]entity.CookStep{}, errors.New("get cook step error")).Once()

		res, err := recipeUC.FindByID(1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestRecipeStore(t *testing.T) {
	t.Run("recipe store happy case", func(t *testing.T) {
		recipeRepository.On("Store", mock.AnythingOfType("*entity.Recipe")).Return(entity.Recipe{ID: 1}, nil).Once()
		recipeMaterialRepository.On("Store", mock.AnythingOfType("*entity.RecipeMaterial")).Return(entity.RecipeMaterial{ID: 1}, nil).Once()
		cookStepRepository.On("Store", mock.AnythingOfType("*entity.CookStep")).Return(entity.CookStep{ID: 1}, nil).Once()

		res, err := recipeUC.Store(&mockRecipeRequest, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("recipe store error store recipe data", func(t *testing.T) {
		recipeRepository.On("Store", mock.AnythingOfType("*entity.Recipe")).Return(entity.Recipe{}, errors.New("store recipe error")).Once()

		res, err := recipeUC.Store(&mockRecipeRequest, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("recipe store error store recipe material", func(t *testing.T) {
		recipeRepository.On("Store", mock.AnythingOfType("*entity.Recipe")).Return(entity.Recipe{ID: 1}, nil).Once()
		recipeMaterialRepository.On("Store", mock.AnythingOfType("*entity.RecipeMaterial")).Return(entity.RecipeMaterial{}, errors.New("store recipe material error")).Once()

		res, err := recipeUC.Store(&mockRecipeRequest, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("recipe store error store cook step", func(t *testing.T) {
		recipeRepository.On("Store", mock.AnythingOfType("*entity.Recipe")).Return(entity.Recipe{ID: 1}, nil).Once()
		recipeMaterialRepository.On("Store", mock.AnythingOfType("*entity.RecipeMaterial")).Return(entity.RecipeMaterial{ID: 1}, nil).Once()
		cookStepRepository.On("Store", mock.AnythingOfType("*entity.CookStep")).Return(entity.CookStep{}, errors.New("store cook step error")).Once()

		res, err := recipeUC.Store(&mockRecipeRequest, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
