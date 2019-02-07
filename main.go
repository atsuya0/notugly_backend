package main

import (
	"log"
	"net/http"

	"github.com/tayusa/notugly_backend/app/config"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/router"
	"github.com/tayusa/notugly_backend/app/infrastructure/repository"
	"github.com/tayusa/notugly_backend/app/registry"
)

func init() {
	log.SetFlags(log.Ltime | log.Llongfile)
}

func main() {
	config.LoadConfig()

	db := repository.NewMySql()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	interactor := registry.NewInteractor(db)

	handler := interactor.NewAppHandler()

	router := router.NewRouter(handler)

	log.Fatalln(http.ListenAndServe(":"+config.Data.Backend.Port, router))
}
