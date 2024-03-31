package api

import (
	"encoding/json"
	"net/http"
	customerrors "sparrow-plus/custom-errors"
)

type HandlerFuncWithError func(w http.ResponseWriter, r *http.Request) error

func withErrorHandling(handler HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		if err, ok := err.(*customerrors.RequestError); ok {
			status, headers := err.ResponseHeaders()

			for header, value := range headers {
				w.Header().Set(header, value)
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(err)
		}

	}
}
