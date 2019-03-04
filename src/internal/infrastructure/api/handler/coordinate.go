package handler

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/tayusa/notugly_backend/internal/infrastructure/api/property"
	"github.com/tayusa/notugly_backend/internal/interface/controller"
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

	coordinate, err := c.CoordinateController.Get(r.Context(), p.ByName("coordinateId"))
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

	coordinate, err := c.CoordinateController.GetAtRandom(
		r.Context(), property.GetUserId(r.Context()))
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

	coordinates, err := c.CoordinateController.GetByUserId(
		r.Context(), p.ByName("uid"))
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

	id, err := c.CoordinateController.Create(
		r.Context(), property.GetUserId(r.Context()), r.Body)
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

	if err := c.CoordinateController.Delete(r.Context(), r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func NewCoordinateHandler(c controller.CoordinateController) CoordinateHandler {
	return &coordinateHandler{c}
}
