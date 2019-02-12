package presenter

import (
	"github.com/tayusa/notugly_backend/internal/domain"
)

type CoordinatePresenter interface {
	ResponseCoordinate(domain.Coordinate) ([]byte, error)
	ResponseCoordinates([]domain.Coordinate) ([]byte, error)
}
