package usecase

import (
	"context"
	"errors"
	"go-resepee-api/app/controller/request"
	"go-resepee-api/app/middleware"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"
	"go-resepee-api/helper/security"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthUC struct {
	Context context.Context
	DB      *gorm.DB
	JwtAuth *middleware.ConfigJWT
}

type AuthUCInterface interface {
	Login(email, password string) (res string, err error)
	Register(req *request.RegisterRequest) (res entity.User, err error)
}

func NewAuthUC(ctx context.Context, db *gorm.DB, jwtAuth *middleware.ConfigJWT) AuthUCInterface {
	return &AuthUC{
		Context: ctx,
		DB:      db,
		JwtAuth: jwtAuth,
	}
}

func (uc *AuthUC) Login(email, password string) (res string, err error) {
	userRepo := repository.NewUserRepository(uc.Context, uc.DB)

	user := entity.User{}
	err = userRepo.FindByEmail(&user, email)
	if err != nil {
		logrus.Warn(err.Error())
		return res, err
	}

	valid := security.ValidateHash(user.Password, password)
	if !valid {
		return res, errors.New("wrong_password")
	}

	res = uc.JwtAuth.GenerateToken(int(user.ID))
	return res, err
}

func (uc *AuthUC) Register(req *request.RegisterRequest) (res entity.User, err error) {
	user := entity.User{}
	userRepo := repository.NewUserRepository(uc.Context, uc.DB)

	err = userRepo.FindByEmail(&user, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return res, err
	}

	hashedPassword, err := security.Hash(req.Password)
	if err != nil {
		return res, err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Password = hashedPassword
	err = userRepo.Store(&user)
	if err != nil {
		return res, err
	}

	return user, err
}
