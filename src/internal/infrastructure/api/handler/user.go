package handler

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/tayusa/notugly_backend/internal/infrastructure/api/property"
	"github.com/tayusa/notugly_backend/internal/interface/controller"
)

type userHandler struct {
	UserController controller.UserController
}

type UserHandler interface {
	GetUser(http.ResponseWriter, *http.Request, httprouter.Params)
	PostUser(http.ResponseWriter, *http.Request, httprouter.Params)
	PutUser(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (u *userHandler) GetUser(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	user, err := u.UserController.Get(p.ByName("uid"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(user)

	return
}

func (u *userHandler) PostUser(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if err := u.UserController.Create(
		property.GetUserId(r.Context()), r.Body); err != nil {

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func (u *userHandler) PutUser(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if err := u.UserController.Update(
		property.GetUserId(r.Context()), r.Body); err != nil {

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}

func NewUserHandler(c controller.UserController) UserHandler {
	return &userHandler{c}
}
