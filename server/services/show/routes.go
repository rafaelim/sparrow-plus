package show

import (
	"fmt"
	"net/http"
	"sparrow-plus/types"
	"sparrow-plus/utils"
)

type Handler struct {
	store        types.ShowStore
	episodeStore types.EpisodeStore
}

func NewHandler(store types.ShowStore, episodeStore types.EpisodeStore) *Handler {
	return &Handler{store: store, episodeStore: episodeStore}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/shows", h.handleGetShows)
	router.HandleFunc("GET /api/shows/{showId}", h.handleGetShowById)
}

func (h *Handler) handleGetShows(w http.ResponseWriter, r *http.Request) {
	shows, err := h.store.GetShows()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, shows)
}

func (h *Handler) handleGetShowById(w http.ResponseWriter, r *http.Request) {
	showId := r.PathValue("showId")
	show, err := h.store.GetShowById(showId)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if show == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no show found with the id %v", showId))
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, show)
}
