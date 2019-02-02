package repository

import (
	"github.com/tayusa/notugly_backend/app/domain"
)

type CoordinateRepository interface {
	FindById(int) (domain.Coordinate, error)
	GetAtRandom() (domain.Coordinate, error)
	FindByCoordinateIdAndUserId(int, string) (bool, error)
	FindByUserId(string) ([]domain.Coordinate, error)
	Store(domain.Coordinate) (int64, error)
	Delete(int) error
}
