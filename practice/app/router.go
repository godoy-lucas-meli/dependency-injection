package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"mercadolibre.com/di/practice/entities"
)

type beerForecastFetcher interface {
	Do(w io.Writer, r *http.Request) (*entities.HandlerResult, error)
}

func NewRouter(bff beerForecastFetcher) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/beer-forecast", middleware(bff.Do)).Methods(http.MethodGet)

	return router
}
