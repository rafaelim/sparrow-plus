package api

import (
	"net/http"
)

type APIServe struct {
	addr string
}

func NewAPIServe(addr string) *APIServe {
	return &APIServe{
		addr: addr,
	}
}

func (s *APIServe) Setup(router *http.ServeMux) {
	s.VideoStreamHandlers(router)
}
