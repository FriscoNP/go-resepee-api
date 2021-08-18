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
	materialRepository mocks.MaterialRepositoryInterface

	mockMaterials = []entity.Material{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}

	mockMaterialRequest = request.CreateMaterialRequest{
		Name:        "Melon",
		ImageFileID: 1,
	}

	materialUC = usecase.NewMaterialUC(context.Background(), &materialRepository)
)

func TestGetMaterial(t *testing.T) {
	t.Run("get materials happy case", func(t *testing.T) {
		materialRepository.On("Get").Return(mockMaterials, nil).Once()

		res, err := materialUC.Get()
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("get materials error", func(t *testing.T) {
		materialRepository.On("Get").Return([]entity.Material{}, errors.New("get materials error")).Once()

		res, err := materialUC.Get()
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestStoreMaterial(t *testing.T) {
	t.Run("store material happy case", func(t *testing.T) {
		materialRepository.On("Store", mock.AnythingOfType("*entity.Material")).Return(entity.Material{ID: 1}, nil).Once()

		res, err := materialUC.Store(&mockMaterialRequest)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("store material error", func(t *testing.T) {
		materialRepository.On("Store", mock.AnythingOfType("*entity.Material")).Return(entity.Material{}, errors.New("store material error")).Once()

		res, err := materialUC.Store(&mockMaterialRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
