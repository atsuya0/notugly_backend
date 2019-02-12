package controller

import (
	"encoding/json"
	"io"

	"github.com/tayusa/notugly_backend/internal/domain"
	"github.com/tayusa/notugly_backend/internal/usecase/service"
)

type favoriteController struct {
	favoriteService service.FavoriteService
}

type FavoriteController interface {
	Create(string, io.ReadCloser) error
	Delete(string, io.ReadCloser) error
}

func (f *favoriteController) Create(uid string, body io.ReadCloser) error {
	favorite := domain.Favorite{UserId: uid}
	if err := json.NewDecoder(body).Decode(&favorite); err != nil {
		return err
	}

	if err := f.favoriteService.Create(favorite); err != nil {
		return err
	}
	return nil
}

func (f *favoriteController) Delete(uid string, body io.ReadCloser) error {
	favorite := domain.Favorite{UserId: uid}
	if err := json.NewDecoder(body).Decode(&favorite); err != nil {
		return err
	}

	if err := f.favoriteService.Delete(favorite); err != nil {
		return err
	}
	return nil
}

func NewFavoriteController(s service.FavoriteService) FavoriteController {
	return &favoriteController{s}
}
