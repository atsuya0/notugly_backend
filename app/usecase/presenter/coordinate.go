package presenter

import (
	"github.com/tayusa/notugly_backend/domain"
)

type CoordinatePresenter interface {
	ResponseCoordinate(domain.Coordinate) ([]byte, error)
	ResponseCoordinates([]domain.Coordinate) ([]byte, error)
}
