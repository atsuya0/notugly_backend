package service

import (
	"github.com/tayusa/notugly_backend/domain"
	"github.com/tayusa/notugly_backend/usecase/repository"
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

func NewFavoriteService(r repository.FavoriteRepository) FavoriteService {
	return &favoriteService{r}
}
