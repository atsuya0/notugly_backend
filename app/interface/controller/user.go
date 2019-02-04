package controller

import (
	"encoding/json"
	"io"

	"github.com/tayusa/notugly_backend/app/domain"
	"github.com/tayusa/notugly_backend/app/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	Get(string) ([]byte, error)
	Create(string, io.ReadCloser) error
	Update(string, io.ReadCloser) error
}

func (u *userController) Get(uid string) ([]byte, error) {
	user, err := u.userService.Get(uid)
	if err != nil {
		return []byte{}, err
	}
	return user, nil
}

func (u *userController) Create(uid string, body io.ReadCloser) error {
	user := domain.User{Id: uid}
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		return err
	}

	if err := u.userService.Create(user); err != nil {
		return err
	}
	return nil
}

func (u *userController) Update(uid string, body io.ReadCloser) error {
	user := domain.User{Id: uid}
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		return err
	}

	if err := u.userService.Update(user); err != nil {
		return err
	}
	return nil
}

func NewUserController(service service.UserService) UserController {
	return &userController{service}
}
