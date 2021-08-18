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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRepository mocks.UserRepositoryInterface
	authUC         usecase.AuthUCInterface

	mockRegisterRequest = request.RegisterRequest{
		Name:                 "user test",
		Email:                "user@test.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
)

func setup() {
	authUC = usecase.NewAuthUC(context.Background(), &userRepository, &middleware.ConfigJWT{})
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	t.Run("register happy case", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, gorm.ErrRecordNotFound).Once()
		userRepository.On("Store", mock.AnythingOfType("*entity.User")).Return(nil).Once()

		err := authUC.Register(&mockRegisterRequest)
		assert.NoError(t, err)
	})

	t.Run("register error find user by email", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, gorm.ErrInvalidDB).Once()

		err := authUC.Register(&mockRegisterRequest)
		assert.Error(t, err)
	})

	t.Run("register error email registered", func(t *testing.T) {
		user := entity.User{
			Name:     "user test",
			Email:    "user@test.com",
			Password: "password",
		}
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(user, nil).Once()

		err := authUC.Register(&mockRegisterRequest)
		assert.Error(t, err)
	})

	t.Run("register error store user", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, gorm.ErrRecordNotFound).Once()
		userRepository.On("Store", mock.AnythingOfType("*entity.User")).Return(errors.New("failed to store user")).Once()

		err := authUC.Register(&mockRegisterRequest)
		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("login happy case", func(t *testing.T) {
		hashedPassword, _ := security.Hash("password")
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

	t.Run("login error find email", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(entity.User{}, errors.New("email not found")).Once()

		token, err := authUC.Login("user@mail.com", "password")
		assert.Error(t, err)
		assert.Empty(t, token)
	})
}
