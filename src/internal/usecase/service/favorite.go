package service

import (
	"context"

	"github.com/tayusa/notugly_backend/internal/domain"
	"github.com/tayusa/notugly_backend/internal/usecase/repository"
)

type favoriteService struct {
	FavoriteRepository repository.FavoriteRepository
}

type FavoriteService interface {
	Create(context.Context, domain.Favorite) error
	Delete(context.Context, domain.Favorite) error
}

func (f *favoriteService) Create(ctx context.Context, favorite domain.Favorite) (err error) {
	err = f.FavoriteRepository.Store(ctx, favorite)
	return
}

func (f *favoriteService) Delete(ctx context.Context, favorite domain.Favorite) (err error) {
	err = f.FavoriteRepository.Delete(ctx, favorite)
	return
}

func NewFavoriteService(r repository.FavoriteRepository) FavoriteService {
	return &favoriteService{r}
}
