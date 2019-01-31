package controller

import (
	"encoding/json"
	"strconv"

	"github.com/tayusa/notugly_backend/app/domain"
	"github.com/tayusa/notugly_backend/app/usecase/service"
)

type coordinateController struct {
	coordinateService service.CoordinateService
}

type CoordinateController interface {
	Get(string) ([]byte, error)
	GetByUserId(string) ([]byte, error)
	Create(string, []byte) ([]byte, error)
	Delete([]byte) error
}

func (c *coordinateController) Get(coordinateId string) ([]byte, error) {
	id, err := strconv.Atoi(coordinateId)
	if err != nil {
		return []byte{}, err
	}

	coordinate, err := c.coordinateService.Get(id)
	if err != nil {
		return []byte{}, err
	}
	return coordinate, nil
}

func (c *coordinateController) GetByUserId(uid string) ([]byte, error) {
	coordinates, err := c.coordinateService.GetByUserId(uid)
	if err != nil {
		return []byte{}, err
	}
	return coordinates, nil
}

func (c *coordinateController) Create(uid string, body []byte) ([]byte, error) {
	coordinate := domain.Coordinate{UserId: uid}
	if err := json.Unmarshal(body, &coordinate); err != nil {
		return []byte{}, err
	}

	coordinateId, err := c.coordinateService.Create(coordinate)
	if err != nil {
		return []byte{}, err
	}
	return coordinateId, nil
}

func (c *coordinateController) Delete(body []byte) error {
	var coordinate domain.Coordinate
	if err := json.Unmarshal(body, &coordinate); err != nil {
		return err
	}

	if err := c.coordinateService.Delete(coordinate.Id); err != nil {
		return err
	}
	return nil
}

func NewCoordinateController(service service.CoordinateService) CoordinateController {
	return &coordinateController{service}
}
