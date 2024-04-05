package movie

import (
	"fmt"
	"net/http"
	"sparrow-plus/types"
	"sparrow-plus/utils"
)

type Handler struct {
	store types.MovieStore
}

func NewHandler(store types.MovieStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/movies", h.handleGetMovies)
	router.HandleFunc("GET /api/movies/{movieId}", h.handleGetMovieById)

	router.HandleFunc("POST /api/movies", h.handleCreateMovie)
}

func (h *Handler) handleGetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.store.GetMovies()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, movies)
}

func (h *Handler) handleGetMovieById(w http.ResponseWriter, r *http.Request) {
	movieId := r.PathValue("movieId")
	movie, err := h.store.GetMovieById(movieId)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if movie == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no movie found with the id %v", movieId))
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, movie)
}

func (h *Handler) handleCreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie types.CreateMoviePayload
	if err := utils.ParseJSON(r, &movie); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateMovie(movie)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, movie)
}
