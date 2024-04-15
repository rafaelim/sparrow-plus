package episode

import (
	"fmt"
	"net/http"
	"sparrow-plus/types"
	"sparrow-plus/utils"
)

type Handler struct {
	store     types.EpisodeStore
	showStore types.ShowStore
}

func NewHandler(store types.EpisodeStore, showStore types.ShowStore) *Handler {
	return &Handler{store: store, showStore: showStore}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /api/episodes", h.handleCreateShow)

	router.HandleFunc("GET /api/episodes/{episodeId}", h.handleGetEpisodeById)
}

func (h *Handler) handleCreateShow(w http.ResponseWriter, r *http.Request) {
	var episode types.CreateEpisodePayload
	if err := utils.ParseJSON(r, &episode); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	show, errf := h.showStore.GetShowByName(episode.ShowName)
	if errf != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to find show by name %w", errf))
		return
	}

	if show != nil && show.ShowId == "" {
		show, errf = h.showStore.CreateShow(types.CreateShowPayload{Name: episode.ShowName})

		if errf != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to create show with name %v %w", episode.ShowName, errf))
			return
		}

	}
	episode.ShowId = show.ShowId
	err := h.store.CreateEpisode(episode)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, episode)
}

func (h *Handler) handleGetEpisodeById(w http.ResponseWriter, r *http.Request) {
	episodeId := r.PathValue("episodeId")

	episode, err := h.store.GetEpisodeById(episodeId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, episode)
}
