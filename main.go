package main

import (
	"log"
	"net/http"

	"github.com/tayusa/notugly/app/config"
	"github.com/tayusa/notugly/app/infrastructure/api/router"
	"github.com/tayusa/notugly/app/infrastructure/repository"
	"github.com/tayusa/notugly/app/registry"
)

func main() {
	config.LoadConfig()

	conn := repository.NewMySql()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	registry := registry.NewInteractor(conn)

	handler := registry.NewAppHandler()

	router := router.NewRouter(handler)

	log.Fatalln(http.ListenAndServe(":"+config.Data.Backend.Port, router))
}
