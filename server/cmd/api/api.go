package api

import (
	"database/sql"
	"log"
	"net/http"
	"sparrow-plus/services/episode"
	"sparrow-plus/services/movie"
	"sparrow-plus/services/show"
	"sparrow-plus/services/stream"

	"github.com/rs/cors"
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

func (s *APIServe) Run() error {
	router := http.NewServeMux()
	handler := cors.AllowAll().Handler(router)

	streamHandler := stream.NewHandler()
	streamHandler.RegisterRoutes(router)

	movieStore := movie.NewStore(s.db)
	movieHandler := movie.NewHandler(movieStore)
	movieHandler.RegisterRoutes(router)

	showStore := show.NewStore(s.db)
	episodeStore := episode.NewStore(s.db)
	showHandler := show.NewHandler(showStore, episodeStore)
	showHandler.RegisterRoutes(router)

	log.Printf("Listening on %v", s.addr)

	return http.ListenAndServe(s.addr, handler)
}
