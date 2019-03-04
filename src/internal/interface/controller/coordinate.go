package controller

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"strconv"

	"github.com/tayusa/notugly_backend/internal/domain"
	"github.com/tayusa/notugly_backend/internal/usecase/service"
)

type coordinateController struct {
	coordinateService service.CoordinateService
}

type CoordinateController interface {
	Get(context.Context, string) ([]byte, error)
	GetAtRandom(context.Context, string) ([]byte, error)
	GetByUserId(context.Context, string) ([]byte, error)
	Create(context.Context, string, io.ReadCloser) ([]byte, error)
	Delete(context.Context, io.ReadCloser) error
}

func (c *coordinateController) Get(ctx context.Context, coordinateId string) ([]byte, error) {
	id, err := strconv.Atoi(coordinateId)
	if err != nil {
		return []byte{}, err
	}

	coordinate, err := c.coordinateService.Get(ctx, id)
	if err != nil {
		return []byte{}, err
	}
	return coordinate, nil
}

func (c *coordinateController) GetAtRandom(ctx context.Context, uid string) ([]byte, error) {
	coordinate, err := c.coordinateService.GetAtRandom(ctx, uid)
	if err != nil {
		return []byte{}, err
	}
	return coordinate, nil
}

func (c *coordinateController) GetByUserId(ctx context.Context, uid string) ([]byte, error) {
	coordinates, err := c.coordinateService.GetByUserId(ctx, uid)
	if err != nil {
		return []byte{}, err
	}
	return coordinates, nil
}

func (c *coordinateController) Create(
	ctx context.Context, uid string, body io.ReadCloser) ([]byte, error) {

	coordinate := domain.Coordinate{UserId: uid}
	if err := json.NewDecoder(body).Decode(&coordinate); err != nil {
		return []byte{}, err
	}

	image, err := base64.StdEncoding.DecodeString(coordinate.Image)
	if err != nil {
		return []byte{}, err
	}

	coordinateId, err := c.coordinateService.Create(ctx, coordinate, image)
	if err != nil {
		return []byte{}, err
	}
	return coordinateId, nil
}

func (c *coordinateController) Delete(ctx context.Context, body io.ReadCloser) error {
	var coordinate domain.Coordinate
	if err := json.NewDecoder(body).Decode(&coordinate); err != nil {
		return err
	}

	if err := c.coordinateService.Delete(ctx, coordinate.Id); err != nil {
		return err
	}
	return nil
}

func NewCoordinateController(s service.CoordinateService) CoordinateController {
	return &coordinateController{s}
}
