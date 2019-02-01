package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tayusa/notugly_backend/app/config"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/handler"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/middleware"
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

func NewRouter(handler handler.AppHandler) *httprouter.Router {
	router := httprouter.New()

	// user
	router.GET("/users/:uid",
		middleware.SetHeader(middleware.Auth(handler.GetUser)))
	router.POST("/users/me", middleware.SetHeader(handler.PostUser))
	router.PUT("/users/me", middleware.SetHeader(handler.PutUser))

	// coordinate
	router.GET("/coordinates/:coordinateId",
		middleware.SetHeader(middleware.Auth(handler.GetCoordinate)))
	router.GET("/coordinates/:uid",
		middleware.SetHeader(middleware.Auth(handler.GetCoordinates)))
	router.POST("/coordinates",
		middleware.SetHeader(handler.PostCoordinate))
	router.DELETE("/coordinates",
		middleware.SetHeader(middleware.Auth(handler.DeleteCoordinate)))

	// favorite
	router.POST("/favorites", middleware.SetHeader(handler.PostFavorite))
	router.DELETE("/favorites", middleware.SetHeader(handler.DeleteFavorite))

	router.OPTIONS("/*path", handlePreFlight)
	router.ServeFiles("/images/*filepath", http.Dir("images"))

	return router
}
