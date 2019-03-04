package handler

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/tayusa/notugly_backend/internal/infrastructure/api/property"
	"github.com/tayusa/notugly_backend/internal/interface/controller"
)

type favoriteHandler struct {
	FavoriteController controller.FavoriteController
}

type FavoriteHandler interface {
	PostFavorite(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteFavorite(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (f *favoriteHandler) PostFavorite(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if err := f.FavoriteController.Create(
		r.Context(), property.GetUserId(r.Context()), r.Body); err != nil {

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}

func (f *favoriteHandler) DeleteFavorite(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if err := f.FavoriteController.Delete(
		r.Context(), property.GetUserId(r.Context()), r.Body); err != nil {

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}

func NewFavoriteHandler(c controller.FavoriteController) FavoriteHandler {
	return &favoriteHandler{c}
}
