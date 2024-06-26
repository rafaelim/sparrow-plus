package api

import (
	"database/sql"
	"log"
	"net/http"
	"sparrow-plus/services/episode"
	"sparrow-plus/services/movie"
	"sparrow-plus/services/show"
	"sparrow-plus/services/stream"
	watchstatus "sparrow-plus/services/watch-status"

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

	movieStore := movie.NewStore(s.db)
	movieHandler := movie.NewHandler(movieStore)
	movieHandler.RegisterRoutes(router)

	episodeStore := episode.NewStore(s.db)
	showStore := show.NewStore(s.db)

	showHandler := show.NewHandler(showStore, episodeStore)
	showHandler.RegisterRoutes(router)

	episodeHandler := episode.NewHandler(episodeStore, showStore)
	episodeHandler.RegisterRoutes(router)

	watchStatusStore := watchstatus.NewStore(s.db)
	watchStatusHandler := watchstatus.NewHandler(watchStatusStore)
	watchStatusHandler.RegisterRoutes(router)

	streamHandler := stream.NewHandler(movieStore, episodeStore)
	streamHandler.RegisterRoutes(router)

	log.Printf("Listening on %v", s.addr)

	return http.ListenAndServe(s.addr, handler)
}
