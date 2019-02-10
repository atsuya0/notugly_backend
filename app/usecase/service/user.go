package service

import (
	"github.com/tayusa/notugly_backend/domain"
	"github.com/tayusa/notugly_backend/usecase/presenter"
	"github.com/tayusa/notugly_backend/usecase/repository"
)

type userService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserService interface {
	Get(string) ([]byte, error)
	Create(domain.User) error
	Update(domain.User) error
}

func (u *userService) Get(uid string) ([]byte, error) {
	user, err := u.UserRepository.FindById(uid)
	if err != nil {
		return []byte{}, err
	}

	output, err := u.UserPresenter.ResponseUser(user)
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func (u *userService) Create(user domain.User) (err error) {
	err = u.UserRepository.Store(user)
	return
}

func (u *userService) Update(user domain.User) (err error) {
	err = u.UserRepository.Update(user)
	return
}

func NewUserService(
	repository repository.UserRepository,
	presenter presenter.UserPresenter) UserService {

	return &userService{UserRepository: repository, UserPresenter: presenter}
}
