package api

import (
	"log"
	"net/http"
)

type HandlerFuncWithError func(w http.ResponseWriter, r *http.Request) error

func withErrorHandling(handler HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			log.Print("Request failed", err)
		}
	}
}
