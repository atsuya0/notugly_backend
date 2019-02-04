package service

import (
	"database/sql"
	"log"
	"os"

	"github.com/tayusa/notugly_backend/app/domain"
	"github.com/tayusa/notugly_backend/app/usecase/presenter"
	"github.com/tayusa/notugly_backend/app/usecase/repository"
	"github.com/tayusa/notugly_backend/app/utils/random"
)

type coordinateService struct {
	CoordinateRepository repository.CoordinateRepository
	CoordinatePresenter  presenter.CoordinatePresenter
}

type CoordinateService interface {
	Get(int) ([]byte, error)
	GetAtRandom(string) ([]byte, error)
	IsFavorite(int, string) (bool, error)
	GetByUserId(string) ([]byte, error)
	Create(domain.Coordinate, []byte) ([]byte, error)
	Delete(int) error
}

func (c *coordinateService) Get(coordinateId int) ([]byte, error) {
	coordinate, err := c.CoordinateRepository.FindById(coordinateId)
	if err != nil {
		return []byte{}, err
	}

	output, err := c.CoordinatePresenter.ResponseCoordinate(coordinate)
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func (c *coordinateService) GetAtRandom(uid string) ([]byte, error) {
	coordinate, err := c.CoordinateRepository.GetAtRandom()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return []byte{}, nil
		default:
			return []byte{}, err
		}
	}

	isFavorited, err := c.IsFavorite(coordinate.Id, uid)
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

func (c *coordinateService) IsFavorite(
	coordinateId int, uid string) (bool, error) {

	_, err := c.CoordinateRepository.
		FindFavoriteByCoordinateIdAndUserId(coordinateId, uid)
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

func (c *coordinateService) GetByUserId(uid string) ([]byte, error) {
	coordinates, err := c.CoordinateRepository.FindByUserId(uid)
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


func (c *coordinateService) Create(
	coordinate domain.Coordinate, image []byte) ([]byte, error) {

	fileName := random.RandomString(16)
	err := c.SaveImage(fileName, image)
	if err != nil {
		return []byte{}, err
	}
	coordinate.ImageName = fileName

	id, err := c.CoordinateRepository.Store(coordinate)
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

func (c *coordinateService) Delete(coordinateId int) (err error) {
	err = c.CoordinateRepository.Delete(coordinateId)
	return
}

func NewCoordinateService(
	repository repository.CoordinateRepository,
	presenter presenter.CoordinatePresenter) CoordinateService {

	return &coordinateService{CoordinateRepository: repository, CoordinatePresenter: presenter}
}
