package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"mercadolibre.com/di/practice/entities"
)

type bpfHandler interface {
	Do(w io.Writer, r *http.Request) (*entities.HandlerResult, error)
}

func NewRouter(handler bpfHandler) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/beer-forecast", middleware(handler.Do)).Methods(http.MethodGet)

	return router
}
