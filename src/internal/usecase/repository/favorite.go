package repository

import (
	"context"

	"github.com/tayusa/notugly_backend/internal/domain"
)

type FavoriteRepository interface {
	Store(context.Context, domain.Favorite) error
	Delete(context.Context, domain.Favorite) error
}
