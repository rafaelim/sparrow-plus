package show

import (
	"fmt"
	"net/http"
	"sparrow-plus/types"
	"sparrow-plus/utils"
	"strconv"
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

	router.HandleFunc("POST /api/shows", h.handleCreateShow)

	router.HandleFunc("GET /api/shows/episodes", h.handleGetEpisodes)
	router.HandleFunc("GET /api/shows/episodes/{episodeId}", h.handleGetEpisodeById)
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

func (h *Handler) handleCreateShow(w http.ResponseWriter, r *http.Request) {
	var show types.CreateShowPayload
	if err := utils.ParseJSON(r, &show); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateShow(show)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, show)
}

func (h *Handler) handleGetEpisodes(w http.ResponseWriter, r *http.Request) {
	showId := r.PathValue("showId")
	season, _ := strconv.Atoi(r.PathValue("season"))

	episodes, err := h.episodeStore.GetEpisodes(showId, season)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, episodes)
}

func (h *Handler) handleGetEpisodeById(w http.ResponseWriter, r *http.Request) {
	episodeId := r.PathValue("episodeId")

	episode, err := h.episodeStore.GetEpisodeById(episodeId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, episode)
}
