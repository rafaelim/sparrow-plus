package stream

import (
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	"sparrow-plus/hls"
	"sparrow-plus/utils"
	"strconv"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/stream/{videoName}", h.handlePlaylist)
	router.HandleFunc("GET /api/stream/{videoName}/segments/{resolution}/{segment}/", h.handleSegment)
	router.HandleFunc("GET /api/stream/{videoName}/info", h.handleVideoInfo)
	router.HandleFunc("GET /api/stream/{videoName}/subtitles/{subIndex}", h.handleSubtitles)
}

func (h *Handler) handlePlaylist(w http.ResponseWriter, r *http.Request) {
	videoName := r.PathValue("videoName")
	slog.Info("Init playlist handler", "video", videoName)

	template := fmt.Sprintf("%v://%v/api/stream/%v/segments/{{.Resolution}}/{{.Segment}}/", "http", r.Host, videoName)

	err := hls.WritePlaylist(template, filepath.Join(hls.RootDir, videoName), 1080, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) handleSegment(w http.ResponseWriter, r *http.Request) {
	videoName := r.PathValue("videoName")
	segment, _ := strconv.ParseInt(r.PathValue("segment"), 0, 64)
	resolution, _ := strconv.ParseInt(r.PathValue("resolution"), 0, 64)
	slog.Info("Init segment handler", "video", videoName, "resolution", resolution, "segment", segment)

	err := hls.WriteSegment(filepath.Join(hls.RootDir, videoName), segment, resolution, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) handleVideoInfo(w http.ResponseWriter, r *http.Request) {
	videoName := r.PathValue("videoName")
	slog.Info("Init video info handler", "video", videoName)

	videoInfo, err := hls.GetVideoInfo(filepath.Join(hls.RootDir, videoName))

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, videoInfo)
}

func (h *Handler) handleSubtitles(w http.ResponseWriter, r *http.Request) {
	videoName := r.PathValue("videoName")
	subIndex, _ := strconv.Atoi(r.PathValue("subIndex"))
	slog.Info("Init subtitles handler", "video", videoName, "subIndex", subIndex)

	err := hls.WriteSubtitles(filepath.Join(hls.RootDir, videoName), subIndex, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}
