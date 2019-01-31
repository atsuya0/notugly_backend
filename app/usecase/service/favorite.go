package service

import (
	"github.com/tayusa/notugly_backend/app/domain"
	"github.com/tayusa/notugly_backend/app/usecase/repository"
)

type favoriteService struct {
	FavoriteRepository repository.FavoriteRepository
}

type FavoriteService interface {
	Create(domain.Favorite) error
	Delete(domain.Favorite) error
}

func (f *favoriteService) Create(favorite domain.Favorite) (err error) {
	err = f.FavoriteRepository.Store(favorite)
	return
}

func (f *favoriteService) Delete(favorite domain.Favorite) (err error) {
	err = f.FavoriteRepository.Delete(favorite)
	return
}

func NewFavoriteService(
	repository repository.FavoriteRepository) FavoriteService {
	return &favoriteService{repository}
}
