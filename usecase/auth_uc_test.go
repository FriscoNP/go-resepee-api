package usecase_test

import (
	"context"
	"errors"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/middleware"
	"go-resepee-api/db/repository/mocks"
	"go-resepee-api/entity"
	"go-resepee-api/helper/security"
	"go-resepee-api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRepository        mocks.UserRepositoryInterface
	abstractApiRepository mocks.AbstractApiRepositoryInterface

	mockRegisterRequest = request.RegisterRequest{
		Name:                 "user test",
		Email:                "user@test.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}

	authUC = usecase.NewAuthUC(context.Background(), &userRepository, &abstractApiRepository, &middleware.ConfigJWT{})
)

func TestRegister(t *testing.T) {
	t.Run("register happy case", func(t *testing.T) {
		user := entity.User{
			Name:     "user test",
			Email:    "user@test.com",
			Password: "password",
		}

		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, gorm.ErrRecordNotFound).Once()
		abstractApiRepository.On("ValidateEmail", mock.AnythingOfType("string")).Return(entity.AbstractEmailValidation{IsValidFormat: true, IsSMTPValid: true}, nil).Once()
		userRepository.On("Store", mock.AnythingOfType("*entity.User")).Return(user, nil).Once()

		res, err := authUC.Register(&mockRegisterRequest)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("register error find user by email", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, gorm.ErrInvalidDB).Once()

		res, err := authUC.Register(&mockRegisterRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("register error email registered", func(t *testing.T) {
		user := entity.User{
			Name:     "user test",
			Email:    "user@test.com",
			Password: "password",
		}
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(user, nil).Once()

		res, err := authUC.Register(&mockRegisterRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("register error abstract api", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, gorm.ErrRecordNotFound).Once()
		abstractApiRepository.On("ValidateEmail", mock.AnythingOfType("string")).Return(entity.AbstractEmailValidation{}, errors.New("abstract api error")).Once()

		res, err := authUC.Register(&mockRegisterRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("register error invalid email", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, gorm.ErrRecordNotFound).Once()
		abstractApiRepository.On("ValidateEmail", mock.AnythingOfType("string")).Return(entity.AbstractEmailValidation{IsValidFormat: false, IsSMTPValid: false}, nil).Once()

		res, err := authUC.Register(&mockRegisterRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("register error store user", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, gorm.ErrRecordNotFound).Once()
		abstractApiRepository.On("ValidateEmail", mock.AnythingOfType("string")).Return(entity.AbstractEmailValidation{IsValidFormat: true, IsSMTPValid: true}, nil).Once()
		userRepository.On("Store", mock.AnythingOfType("*entity.User")).Return(entity.User{}, errors.New("failed to store user")).Once()

		res, err := authUC.Register(&mockRegisterRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestLogin(t *testing.T) {
	t.Run("login happy case", func(t *testing.T) {
		hashedPassword := security.Hash("password")
		user := entity.User{
			ID:       1,
			Name:     "user test",
			Email:    "user@test.com",
			Password: hashedPassword,
		}
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(user, nil).Once()

		token, err := authUC.Login(user.Email, "password")
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("login error wrong password", func(t *testing.T) {
		hashedPassword := security.Hash("password")
		user := entity.User{
			ID:       1,
			Name:     "user test",
			Email:    "user@test.com",
			Password: hashedPassword,
		}
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(user, nil).Once()

		token, err := authUC.Login(user.Email, "password123")
		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("login error find email", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, errors.New("email not found")).Once()

		token, err := authUC.Login("user@mail.com", "password")
		assert.Error(t, err)
		assert.Empty(t, token)
	})
}
