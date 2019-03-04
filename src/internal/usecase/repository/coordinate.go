package repository

import (
	"context"

	"github.com/tayusa/notugly_backend/internal/domain"
)

type CoordinateRepository interface {
	FindById(context.Context, int) (domain.Coordinate, error)
	GetAtRandom(context.Context) (domain.Coordinate, error)
	FindFavoriteByCoordinateIdAndUserId(
		context.Context, int, string) (domain.Favorite, error)
	FindByUserId(context.Context, string) ([]domain.Coordinate, error)
	Store(context.Context, domain.Coordinate) (int64, error)
	Delete(context.Context, int) error
}
