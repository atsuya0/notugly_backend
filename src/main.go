package main

import (
	"log"
	"net/http"

	config "github.com/tayusa/notugly_backend/configs"
	"github.com/tayusa/notugly_backend/internal/infrastructure/api/middleware"
	"github.com/tayusa/notugly_backend/internal/infrastructure/api/router"
	repository "github.com/tayusa/notugly_backend/internal/infrastructure/repository/mysql"
	"github.com/tayusa/notugly_backend/internal/registry"
)

const (
	configPath = "./configs/config.json"
)

func init() {
	log.SetFlags(log.Ltime | log.Llongfile)
}

func main() {
	config.LoadConfig(configPath)

	db := repository.NewDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	interactor := registry.NewInteractor(db, "images")

	handler := interactor.NewAppHandler()

	router := router.NewRouter(handler, middleware.Auth)

	log.Fatalln(http.ListenAndServe(":"+config.Data.Backend.Port, router))
}
