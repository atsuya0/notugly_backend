package presenter

import (
	"encoding/json"

	"github.com/tayusa/notugly_backend/app/domain"
)

type coordinatePresenter struct {
}

func (c *coordinatePresenter) ResponseCoordinate(
	coordinate domain.Coordinate) ([]byte, error) {

	output, err := json.Marshal(&coordinate)
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}

func (c *coordinatePresenter) ResponseCoordinates(
	coordinates []domain.Coordinate) ([]byte, error) {

	output, err := json.Marshal(&coordinates)
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}

func NewCoordinatePresenter() *coordinatePresenter {
	return &coordinatePresenter{}
}
