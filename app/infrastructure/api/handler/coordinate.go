package handler

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/firebase"
	"github.com/tayusa/notugly_backend/app/interface/controller"
)

type coordinateHandler struct {
	CoordinateController controller.CoordinateController
}

type CoordinateHandler interface {
	GetCoordinate(http.ResponseWriter, *http.Request, httprouter.Params)
	GetCoordinateAtRandom(http.ResponseWriter, *http.Request, httprouter.Params)
	GetCoordinates(http.ResponseWriter, *http.Request, httprouter.Params)
	PostCoordinate(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteCoordinate(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (c *coordinateHandler) GetCoordinate(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	coordinate, err := c.CoordinateController.Get(p.ByName("coordinateId"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(coordinate)

	return
}

func (c *coordinateHandler) GetCoordinateAtRandom(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	token, err := firebase.FetchToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	coordinate, err := c.CoordinateController.GetAtRandom(token.UID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(coordinate)

	return
}

func (c *coordinateHandler) GetCoordinates(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	coordinates, err := c.CoordinateController.GetByUserId(p.ByName("uid"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(coordinates)

	return
}

func (c *coordinateHandler) PostCoordinate(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	token, err := firebase.FetchToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	id, err := c.CoordinateController.Create(token.UID, body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(id)

	return
}

func (c *coordinateHandler) DeleteCoordinate(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	if err := c.CoordinateController.Delete(body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func NewCoordinateHandler(
	controller controller.CoordinateController) CoordinateHandler {
	return &coordinateHandler{controller}
}
