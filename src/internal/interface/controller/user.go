package controller

import (
	"context"
	"encoding/json"
	"io"

	"github.com/tayusa/notugly_backend/internal/domain"
	"github.com/tayusa/notugly_backend/internal/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	Get(context.Context, string) ([]byte, error)
	Create(context.Context, string, io.ReadCloser) error
	Update(context.Context, string, io.ReadCloser) error
}

func (u *userController) Get(ctx context.Context, uid string) ([]byte, error) {
	user, err := u.userService.Get(ctx, uid)
	if err != nil {
		return []byte{}, err
	}
	return user, nil
}

func (u *userController) Create(ctx context.Context, uid string, body io.ReadCloser) error {
	user := domain.User{Id: uid}
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		return err
	}

	if err := u.userService.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *userController) Update(ctx context.Context, uid string, body io.ReadCloser) error {
	user := domain.User{Id: uid}
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		return err
	}

	if err := u.userService.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func NewUserController(s service.UserService) UserController {
	return &userController{s}
}
