package usecase

import (
	"context"
	"errors"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/helper"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
	"github.com/sirupsen/logrus"
)

type authUsecase struct {
	userRepository model.UserRepository
}

func NewAuthUsecase(userRepository model.UserRepository) model.AuthUsecase {
	return &authUsecase{userRepository: userRepository}
}

func (u *authUsecase) Login(ctx context.Context, data model.Login) (token string, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"data": data,
	})

	// Find user by email
	user, err := u.userRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		logger.Error(err)
		return
	}

	// Check password
	if !checkPasswordHash(data.Password, user.Password) {
		err = errors.New("incorrect password")
		return
	}

	// Generate jwt token
	token, err = helper.GenerateToken(user.ID)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (u *authUsecase) Register(ctx context.Context, data model.Register) (token string, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"data": data,
	})

	passHashed, err := hashPassword(data.Password)
	if err != nil {
		logger.Error(err)
		return
	}
	newUser, err := u.userRepository.Create(ctx, model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: passHashed,
	})
	if err != nil {
		logger.Error(err)
		return
	}

	// Generate jwt token
	accessToken, err := helper.GenerateToken(newUser.ID)
	if err != nil {
		logger.Error(err)
		return
	}

	return accessToken, nil
}
