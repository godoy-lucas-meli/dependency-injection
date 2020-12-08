package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	di "mercadolibre.com/di/practice"
	"mercadolibre.com/di/practice/entities"
)

func middleware(h func(io.Writer, *http.Request) (*entities.HandlerResult, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerResult, err := h(w, r)
		if err != nil {
			response := &entities.ErrorResponse{
				Success: false,
				Message: err.Error(),
			}
			sendResponse(w, handleErrors(err), response)
			return
		}

		sendResponse(w, handlerResult.Status, entities.SuccessResponse{
			Success: true,
			Data:    handlerResult.Body,
		})
	}
}

func sendResponse(w http.ResponseWriter, status int32, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	if status != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}

func handleErrors(err error) int32 {
	switch t := err.(type) {
	case di.CustomError:
		switch t.Cause {
		case di.ErrDependencyNotAvailable:
			return http.StatusFailedDependency
		case di.ErrBadRequest:
			return http.StatusBadRequest
		case di.ErrNotFound:
			return http.StatusNotFound
		case di.ErrNoWeatherInformationAsYet:
			return http.StatusNoContent
		case di.ErrForbiddenAccess:
			return http.StatusForbidden
		default:
			return http.StatusInternalServerError
		}
	default:
		return http.StatusInternalServerError
	}
}
