package stream

import (
	"fmt"
	"log/slog"
	"net/http"
	"sparrow-plus/hls"
	"sparrow-plus/types"
	"sparrow-plus/utils"
	"strconv"
)

type Handler struct {
	movieStore   types.MovieStore
	episodeStore types.EpisodeStore
}

func NewHandler(
	movieStore types.MovieStore,
	episodeStore types.EpisodeStore,
) *Handler {
	return &Handler{
		movieStore:   movieStore,
		episodeStore: episodeStore,
	}
}

func (h *Handler) withPath(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videoType := r.URL.Query().Get("watch")
		id := r.URL.Query().Get("id")
		values := r.URL.Query()

		fmt.Println(videoType, id)
		var path string
		switch videoType {
		case "movies":
			movie, _ := h.movieStore.GetMovieById(id)
			path = movie.Path
		case "episode":
			movie, _ := h.episodeStore.GetEpisodeById(id)
			path = movie.Path

		}

		if path == "" {
			fmt.Println(videoType, id)
			http.NotFound(w, r)
			return
		}

		values.Add("path", path)
		r.URL.RawQuery = values.Encode()

		next(w, r)
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/stream/master.m3u8", h.withPath(h.handleMasterPlaylist))
	router.HandleFunc("GET /api/stream/index.m3u8", h.withPath(h.handlePlaylist))
	router.HandleFunc("GET /api/stream/audio/{audioIndex}/index.m3u8", h.withPath(h.handleAudioPlaylist))
	router.HandleFunc("GET /api/stream/audio/{audioIndex}/segment/{segment}", h.withPath(h.handleAudioSegment))
	router.HandleFunc("GET /api/stream/subtitles/{subtitleIndex}/index.m3u8", h.withPath(h.handleSubtitlesPlaylist))
	router.HandleFunc("GET /api/stream/subtitles/{subtitleIndex}/segment/{segment}", h.withPath(h.handleSubtitleSegment))
	router.HandleFunc("GET /api/stream/segments/{resolution}/{segment}/", h.withPath(h.handleSegment))
	router.HandleFunc("GET /api/stream/info", h.withPath(h.handleVideoInfo))
	router.HandleFunc("GET /api/stream/subtitles/{subIndex}", h.withPath(h.handleSubtitles))
}

func (h *Handler) handleMasterPlaylist(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	values := r.URL.Query()
	values.Del("path")
	slog.Info("Init master playlist handler", "video", path, "values", values)
	w.Header().Add("Content-type", hls.ContentType)
	err := hls.WriteMasterPlaylist(path, values, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) handleAudioPlaylist(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	audioIndex := r.PathValue("audioIndex")
	values := r.URL.Query()
	values.Del("path")
	slog.Info("Init audio playlist handler", "path", path, "audioIndex", audioIndex)

	template := fmt.Sprintf("%v://%v/api/stream/audio/%v/segment/{{.Segment}}?%v", "http", r.Host, audioIndex, values.Encode())
	w.Header().Add("Content-type", hls.ContentType)
	err := hls.WritePlaylist(template, path, 1080, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) handleSubtitlesPlaylist(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	values := r.URL.Query()
	values.Del("path")
	subtitleIndex := r.PathValue("subtitleIndex")
	slog.Info("Init subtitle playlist handler", "path", path, "subtitleIndex", subtitleIndex)

	template := fmt.Sprintf("%v://%v/api/stream/subtitles/%v/segment/{{.Segment}}?%v", "http", r.Host, subtitleIndex, values.Encode())
	w.Header().Add("Content-type", hls.ContentType)
	err := hls.WriteSinglePlaylist(template, path, 1080, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) handlePlaylist(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	slog.Info("Init playlist handler", "path", path)
	values := r.URL.Query()
	values.Del("path")

	template := fmt.Sprintf("%v://%v/api/stream/segments/{{.Resolution}}/{{.Segment}}?%v", "http", r.Host, values.Encode())
	w.Header().Add("Content-type", hls.ContentType)
	err := hls.WritePlaylist(template, path, 1080, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) handleAudioSegment(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	segment, _ := strconv.Atoi(r.PathValue("segment"))
	audioIndex, _ := strconv.Atoi(r.PathValue("audioIndex"))
	slog.Info("Init segment handler", "path", path, "audioIndex", audioIndex, "segment", segment)

	err := hls.WriteAudioSegment(path, segment, audioIndex, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}
func (h *Handler) handleSubtitleSegment(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	segment, _ := strconv.Atoi(r.PathValue("segment"))
	subtitleIndex, _ := strconv.Atoi(r.PathValue("subtitleIndex"))
	slog.Info("Init segment handler", "path", path, "subtitleIndex", subtitleIndex, "segment", segment)

	err := hls.WriteSubtiteSegment(path, segment, subtitleIndex, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}
func (h *Handler) handleSegment(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	resolution, _ := strconv.ParseInt(r.PathValue("resolution"), 0, 64)
	segment, _ := strconv.ParseInt(r.PathValue("segment"), 0, 64)
	slog.Info("Init segment handler", "path", path, "resolution", resolution, "segment", segment)

	err := hls.WriteSegment(path, segment, resolution, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) handleVideoInfo(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	slog.Info("Init video info handler", "path", path)

	videoInfo, err := hls.GetVideoInfo(path)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, videoInfo)
}

func (h *Handler) handleSubtitles(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	subIndex, _ := strconv.Atoi(r.PathValue("subIndex"))
	slog.Info("Init subtitles handler", "path", path, "subIndex", subIndex)

	err := hls.WriteSubtitles(path, subIndex, w)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}
