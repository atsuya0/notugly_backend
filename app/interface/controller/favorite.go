package controller

import (
	"encoding/json"

	"github.com/tayusa/notugly_backend/app/domain"
	"github.com/tayusa/notugly_backend/app/usecase/service"
)

type favoriteController struct {
	favoriteService service.FavoriteService
}

type FavoriteController interface {
	Create(string, []byte) error
	Delete(string, []byte) error
}

func (f *favoriteController) Create(uid string, body []byte) error {
	favorite := domain.Favorite{UserId: uid}
	if err := json.Unmarshal(body, &favorite); err != nil {
		return err
	}

	if err := f.favoriteService.Create(favorite); err != nil {
		return err
	}
	return nil
}

func (f *favoriteController) Delete(uid string, body []byte) error {
	favorite := domain.Favorite{UserId: uid}
	if err := json.Unmarshal(body, &favorite); err != nil {
		return err
	}

	if err := f.favoriteService.Delete(favorite); err != nil {
		return err
	}
	return nil
}

func NewFavoriteController(service service.FavoriteService) FavoriteController {
	return &favoriteController{service}
}
