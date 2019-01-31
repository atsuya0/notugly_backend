package controller

import (
	"encoding/json"

	"github.com/tayusa/notugly_backend/app/domain"
	"github.com/tayusa/notugly_backend/app/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	Get(string) ([]byte, error)
	Create(string, []byte) error
	Update(string, []byte) error
}

func (u *userController) Get(uid string) ([]byte, error) {
	user, err := u.userService.Get(uid)
	if err != nil {
		return []byte{}, err
	}
	return user, nil
}

func (u *userController) Create(uid string, body []byte) error {
	user := domain.User{Id: uid}
	if err := json.Unmarshal(body, &user); err != nil {
		return err
	}

	if err := u.userService.Create(user); err != nil {
		return err
	}
	return nil
}

func (u *userController) Update(uid string, body []byte) error {
	user := domain.User{Id: uid}
	if err := json.Unmarshal(body, &user); err != nil {
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
