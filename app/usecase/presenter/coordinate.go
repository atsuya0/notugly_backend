package presenter

import (
	"github.com/tayusa/notugly_backend/app/domain"
)

type CoordinatePresenter interface {
	ResponseCoordinate(domain.Coordinate) ([]byte, error)
	ResponseCoordinates([]domain.Coordinate) ([]byte, error)
}
