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
