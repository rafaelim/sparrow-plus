package watchstatus

import (
	"net/http"
	"sparrow-plus/types"
	"sparrow-plus/utils"
)

type Handler struct {
	store types.WatchStatusStore
}

func NewHandler(
	store types.WatchStatusStore,
) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/watchStatus", h.handleGetWatchStatus)
	router.HandleFunc("POST /api/watchStatus", h.handleCreateWatchStatus)
}

func (h *Handler) handleCreateWatchStatus(w http.ResponseWriter, r *http.Request) {
	var payload types.UpsertWatchStatusPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	watchStatus, err := h.store.UpsertWatchStatus(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, watchStatus)
}

func (h *Handler) handleGetWatchStatus(w http.ResponseWriter, r *http.Request) {
	watchStatus, err := h.store.GetWatchStatus()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, watchStatus)
}
