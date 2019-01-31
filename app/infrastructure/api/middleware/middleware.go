package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tayusa/notugly_backend/app/config"
	"github.com/tayusa/notugly_backend/app/infrastructure/api/firebase"
)

func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		_, err := firebase.FetchToken(r)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying ID token\n"))
			return
		}
		next(w, r, p)
	}
}

func SetHeader(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		origin := fmt.Sprintf(
			"http://%s:%s", config.Data.Frontend.Host, config.Data.Frontend.Port)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		next(w, r, p)
	}
}
