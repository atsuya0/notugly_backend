package service

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/tayusa/notugly_backend/internal/domain"
	"github.com/tayusa/notugly_backend/internal/usecase/presenter"
	"github.com/tayusa/notugly_backend/internal/usecase/repository"
	"github.com/tayusa/notugly_backend/pkg/random"
)

type coordinateService struct {
	imagePath            string
	CoordinateRepository repository.CoordinateRepository
	CoordinatePresenter  presenter.CoordinatePresenter
}

type CoordinateService interface {
	Get(context.Context, int) ([]byte, error)
	GetAtRandom(context.Context, string) ([]byte, error)
	isFavorite(context.Context, int, string) (bool, error)
	GetByUserId(context.Context, string) ([]byte, error)
	saveImage(string, []byte) error
	Create(context.Context, domain.Coordinate, []byte) ([]byte, error)
	Delete(context.Context, int) error
}

func (c *coordinateService) Get(
	ctx context.Context, coordinateId int) ([]byte, error) {

	coordinate, err := c.CoordinateRepository.FindById(ctx, coordinateId)
	if err != nil {
		return []byte{}, err
	}

	output, err := c.CoordinatePresenter.ResponseCoordinate(coordinate)
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func (c *coordinateService) GetAtRandom(
	ctx context.Context, uid string) ([]byte, error) {

	coordinate, err := c.CoordinateRepository.GetAtRandom(ctx)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return []byte{}, nil
		default:
			return []byte{}, err
		}
	}

	isFavorited, err := c.isFavorite(ctx, coordinate.Id, uid)
	if err != nil {
		return []byte{}, err
	}
	coordinate.IsFavorited = isFavorited

	output, err := c.CoordinatePresenter.ResponseCoordinate(coordinate)
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func (c *coordinateService) isFavorite(
	ctx context.Context, coordinateId int, uid string) (bool, error) {

	_, err := c.CoordinateRepository.
		FindFavoriteByCoordinateIdAndUserId(ctx, coordinateId, uid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func (c *coordinateService) GetByUserId(
	ctx context.Context, uid string) ([]byte, error) {

	coordinates, err := c.CoordinateRepository.FindByUserId(ctx, uid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return []byte{}, nil
		default:
			return []byte{}, err
		}
	}

	output, err := c.CoordinatePresenter.ResponseCoordinates(coordinates)
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func (c *coordinateService) saveImage(fileName string, image []byte) error {
	file, err := os.Create(filepath.Join(c.imagePath, fileName))
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return err
	}

	_, err = file.Write(image)
	if err != nil {
		return err
	}

	return nil
}

func (c *coordinateService) Create(
	ctx context.Context, coordinate domain.Coordinate, image []byte) (
	[]byte, error) {

	fileName := random.RandomString(16)
	err := c.saveImage(fileName, image)
	if err != nil {
		return []byte{}, err
	}
	coordinate.ImageName = fileName

	id, err := c.CoordinateRepository.Store(ctx, coordinate)
	if err != nil {
		return []byte{}, err
	}

	output, err := c.CoordinatePresenter.ResponseCoordinate(
		domain.Coordinate{Id: int(id), ImageName: fileName})
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func (c *coordinateService) Delete(
	ctx context.Context, coordinateId int) (err error) {

	err = c.CoordinateRepository.Delete(ctx, coordinateId)
	return
}

func NewCoordinateService(
	imagePath string,
	r repository.CoordinateRepository,
	p presenter.CoordinatePresenter) CoordinateService {

	return &coordinateService{
		imagePath:            imagePath,
		CoordinateRepository: r,
		CoordinatePresenter:  p}
}
