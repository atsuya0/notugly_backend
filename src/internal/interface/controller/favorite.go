package controller

import (
	"context"
	"encoding/json"
	"io"

	"github.com/tayusa/notugly_backend/internal/domain"
	"github.com/tayusa/notugly_backend/internal/usecase/service"
)

type favoriteController struct {
	favoriteService service.FavoriteService
}

type FavoriteController interface {
	Create(context.Context, string, io.ReadCloser) error
	Delete(context.Context, string, io.ReadCloser) error
}

func (f *favoriteController) Create(ctx context.Context, uid string, body io.ReadCloser) error {
	favorite := domain.Favorite{UserId: uid}
	if err := json.NewDecoder(body).Decode(&favorite); err != nil {
		return err
	}

	if err := f.favoriteService.Create(ctx, favorite); err != nil {
		return err
	}
	return nil
}

func (f *favoriteController) Delete(ctx context.Context, uid string, body io.ReadCloser) error {
	favorite := domain.Favorite{UserId: uid}
	if err := json.NewDecoder(body).Decode(&favorite); err != nil {
		return err
	}

	if err := f.favoriteService.Delete(ctx, favorite); err != nil {
		return err
	}
	return nil
}

func NewFavoriteController(s service.FavoriteService) FavoriteController {
	return &favoriteController{s}
}
