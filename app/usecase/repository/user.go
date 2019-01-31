package repository

import (
	"github.com/tayusa/notugly_backend/app/domain"
)

type UserRepository interface {
	FindById(string) (domain.User, error)
	Store(domain.User) error
	Update(domain.User) error
}
