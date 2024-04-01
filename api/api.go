package api

import (
	"net/http"
	"sparrow-plus/db"
)

type APIServe struct {
	addr string
	d *db.DB
}

func NewAPIServe(addr string, d *db.DB) *APIServe {
	return &APIServe{
		addr: addr,
		d: d,
	}
}

func (s *APIServe) Setup(router *http.ServeMux) {
	s.VideoStreamHandlers(router)
	s.CategoriesHandlers(router)
}
