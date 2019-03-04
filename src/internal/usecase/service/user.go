package service

import (
	"context"

	"github.com/tayusa/notugly_backend/internal/domain"
	"github.com/tayusa/notugly_backend/internal/usecase/presenter"
	"github.com/tayusa/notugly_backend/internal/usecase/repository"
)

type userService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserService interface {
	Get(context.Context, string) ([]byte, error)
	Create(context.Context, domain.User) error
	Update(context.Context, domain.User) error
}

func (u *userService) Get(ctx context.Context, uid string) ([]byte, error) {
	user, err := u.UserRepository.FindById(ctx, uid)
	if err != nil {
		return []byte{}, err
	}

	output, err := u.UserPresenter.ResponseUser(user)
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func (u *userService) Create(ctx context.Context, user domain.User) (err error) {
	err = u.UserRepository.Store(ctx, user)
	return
}

func (u *userService) Update(ctx context.Context, user domain.User) (err error) {
	err = u.UserRepository.Update(ctx, user)
	return
}

func NewUserService(
	r repository.UserRepository,
	p presenter.UserPresenter) UserService {

	return &userService{UserRepository: r, UserPresenter: p}
}
