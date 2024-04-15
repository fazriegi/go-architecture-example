package usecase

import (
	"errors"
	"net/http"

	"github.com/fazriegi/go-architecture-example/libs/helper"
	"github.com/fazriegi/go-architecture-example/model"
	"github.com/fazriegi/go-architecture-example/repository"
	"github.com/gofiber/fiber/v2/log"
)

type UserUsecase struct {
	repository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{repository: userRepository}
}

func (u *UserUsecase) Create(props *model.User) (int, error) {
	var (
		err  error
		user model.User
	)

	if user, err = u.repository.FindEmail(props.Email); err != nil {
		log.Errorf("error finding email: %s", err.Error())
		return http.StatusInternalServerError, errors.New("failed to create user")
	}

	if user.Email != "" {
		return http.StatusBadRequest, errors.New("email already exist")
	}

	if props.Password, err = helper.HashPassword(props.Password); err != nil {
		log.Errorf("error hashing password: %s", err.Error())
		return http.StatusInternalServerError, errors.New("failed to create user")
	}

	if err := u.repository.Create(props); err != nil {
		log.Errorf("error creating user: %s", err.Error())
		return http.StatusInternalServerError, errors.New("failed to create user")
	}

	return http.StatusCreated, nil
}
