package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	config "github.com/tayusa/notugly_backend/configs"
	"github.com/tayusa/notugly_backend/internal/infrastructure/api/handler"
	"github.com/tayusa/notugly_backend/internal/infrastructure/api/middleware"
)

func handlePreFlight(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	origin := fmt.Sprintf(
		"http://%s:%s", config.Data.Frontend.Host, config.Data.Frontend.Port)
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Authorization")
}

func NewRouter(
	handler handler.AppHandler,
	auth func(next httprouter.Handle) httprouter.Handle) *httprouter.Router {

	router := httprouter.New()

	// user
	router.GET("/users/:uid",
		middleware.SetHeader(handler.GetUser))
	router.POST("/users/me",
		middleware.SetHeader(
			auth(handler.PostUser)))
	router.PUT("/users/me",
		middleware.SetHeader(
			auth(handler.PutUser)))

	// coordinate
	router.GET("/coordinates/:coordinateId",
		middleware.SetHeader(handler.GetCoordinate))
	router.GET("/coordinate",
		middleware.SetHeader(
			auth(handler.GetCoordinateAtRandom)))
	router.GET("/users/:uid/coordinates",
		middleware.SetHeader(handler.GetCoordinates))
	router.POST("/coordinates",
		middleware.SetHeader(
			auth(handler.PostCoordinate)))
	router.DELETE("/coordinates",
		middleware.SetHeader(
			auth(handler.DeleteCoordinate)))

	// favorite
	router.POST("/favorites",
		middleware.SetHeader(
			auth(handler.PostFavorite)))
	router.DELETE("/favorites",
		middleware.SetHeader(
			auth(handler.DeleteFavorite)))

	router.OPTIONS("/*path", handlePreFlight)
	router.ServeFiles("/images/*filepath", http.Dir("images"))

	return router
}
