package usecase

import (
	"context"
	"errors"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/middleware"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"
	"go-resepee-api/helper/security"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthUC struct {
	Context               context.Context
	JwtAuth               *middleware.ConfigJWT
	UserRepository        repository.UserRepositoryInterface
	AbstractApiRepository repository.AbstractApiRepositoryInterface
}

type AuthUCInterface interface {
	Login(email, password string) (res string, err error)
	Register(req *request.RegisterRequest) (err error)
}

func NewAuthUC(ctx context.Context, repo repository.UserRepositoryInterface, abstractApi repository.AbstractApiRepositoryInterface, jwtAuth *middleware.ConfigJWT) AuthUCInterface {
	return &AuthUC{
		Context:               ctx,
		UserRepository:        repo,
		AbstractApiRepository: abstractApi,
		JwtAuth:               jwtAuth,
	}
}

func (uc *AuthUC) Login(email, password string) (res string, err error) {
	user, err := uc.UserRepository.FindByEmail(email)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	valid := security.ValidateHash(user.Password, password)
	if !valid {
		return res, errors.New("wrong_password")
	}

	res = uc.JwtAuth.GenerateToken(int(user.ID))
	return res, err
}

func (uc *AuthUC) Register(req *request.RegisterRequest) (err error) {
	existedUser, err := uc.UserRepository.FindByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Warn(err.Error())
		return err
	}
	if existedUser.Email != "" {
		return errors.New("email_registered")
	}

	emailValidation, err := uc.AbstractApiRepository.ValidateEmail(req.Email)
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	if !emailValidation.IsValidFormat || !emailValidation.IsSMTPValid {
		return errors.New("invalid email format")
	}

	hashedPassword := security.Hash(req.Password)

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	err = uc.UserRepository.Store(&user)
	if err != nil {
		return err
	}

	return err
}
