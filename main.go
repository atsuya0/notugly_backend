package main

import (
	"log"
	"net/http"

	"github.com/tayusa/notugly_backend/app/config"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/middleware"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/router"
	"github.com/tayusa/notugly_backend/app/infrastructure/repository/mysql"
	"github.com/tayusa/notugly_backend/app/registry"
)

const (
	configPath = "./app/config/config.json"
)

func init() {
	log.SetFlags(log.Ltime | log.Llongfile)
}

func main() {
	config.LoadConfig(configPath)

	db := mysql.NewDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	interactor := registry.NewInteractor(db)

	handler := interactor.NewAppHandler()

	router := router.NewRouter(handler, middleware.Auth)

	log.Fatalln(http.ListenAndServe(":"+config.Data.Backend.Port, router))
}
