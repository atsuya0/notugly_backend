package handler

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/firebase"
	"github.com/tayusa/notugly_backend/app/interface/controller"
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

	token, err := firebase.FetchToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	if err := u.UserController.Create(token.UID, body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func (u *userHandler) PutUser(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	token, err := firebase.FetchToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	if err := u.UserController.Update(token.UID, body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}

func NewUserHandler(controller controller.UserController) UserHandler {
	return &userHandler{controller}
}
