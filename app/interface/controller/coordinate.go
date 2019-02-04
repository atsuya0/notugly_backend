package controller

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"strconv"

	"github.com/tayusa/notugly_backend/app/domain"
	"github.com/tayusa/notugly_backend/app/usecase/service"
)

type coordinateController struct {
	coordinateService service.CoordinateService
}

type CoordinateController interface {
	Get(string) ([]byte, error)
	GetAtRandom(string) ([]byte, error)
	GetByUserId(string) ([]byte, error)
	Create(string, io.ReadCloser) ([]byte, error)
	Delete(io.ReadCloser) error
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

func (c *coordinateController) GetAtRandom(uid string) ([]byte, error) {
	coordinate, err := c.coordinateService.GetAtRandom(uid)
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

func (c *coordinateController) Create(uid string, body io.ReadCloser) ([]byte, error) {
	coordinate := domain.Coordinate{UserId: uid}
	if err := json.NewDecoder(body).Decode(&coordinate); err != nil {
		return []byte{}, err
	}

	image, err := base64.StdEncoding.DecodeString(coordinate.Image)
	if err != nil {
		return []byte{}, err
	}

	coordinateId, err := c.coordinateService.Create(coordinate, image)
	if err != nil {
		return []byte{}, err
	}
	return coordinateId, nil
}

func (c *coordinateController) Delete(body io.ReadCloser) error {
	var coordinate domain.Coordinate
	if err := json.NewDecoder(body).Decode(&coordinate); err != nil {
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
