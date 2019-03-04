package repository

import (
	"context"

	"github.com/tayusa/notugly_backend/internal/domain"
)

type UserRepository interface {
	FindById(context.Context, string) (domain.User, error)
	Store(context.Context, domain.User) error
	Update(context.Context, domain.User) error
}
