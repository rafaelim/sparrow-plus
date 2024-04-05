package api

import (
	"database/sql"
	"net/http"
	"sparrow-plus/services/category"
	"sparrow-plus/services/movie"
)

type APIServe struct {
	addr string
	db   *sql.DB
}

func NewAPIServe(addr string, db *sql.DB) *APIServe {
	return &APIServe{
		addr: addr,
		db:   db,
	}
}

func (s *APIServe) Setup(router *http.ServeMux) {
	s.VideoStreamHandlers(router)

	categoryStore := category.NewStore(s.db)
	categoryHandler := category.NewHandler(categoryStore)
	categoryHandler.RegisterRoutes(router)

	movieStore := movie.NewStore(s.db)
	movieHandler := movie.NewHandler(movieStore)
	movieHandler.RegisterRoutes(router)
}
