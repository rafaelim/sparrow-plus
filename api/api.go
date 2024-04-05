package api

import (
	"database/sql"
	"net/http"
	"sparrow-plus/services/movie"
	"sparrow-plus/services/show"
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

	movieStore := movie.NewStore(s.db)
	movieHandler := movie.NewHandler(movieStore)
	movieHandler.RegisterRoutes(router)

	showStore := show.NewStore(s.db)
	showHandler := show.NewHandler(showStore)
	showHandler.RegisterRoutes(router)
}
