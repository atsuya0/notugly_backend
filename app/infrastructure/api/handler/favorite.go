package handler

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/firebase"
	"github.com/tayusa/notugly_backend/app/interface/controller"
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

	token, err := firebase.FetchToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := f.FavoriteController.Create(token.UID, r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}

func (f *favoriteHandler) DeleteFavorite(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	token, err := firebase.FetchToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := f.FavoriteController.Delete(token.UID, r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}

func NewFavoriteHandler(
	controller controller.FavoriteController) FavoriteHandler {
	return &favoriteHandler{controller}
}
