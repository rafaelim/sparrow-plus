package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"sparrow-plus/hls"
	"strconv"
)

func (s *APIServe) VideoStreamHandlers(router *http.ServeMux) {
	router.HandleFunc("GET /api/{videoName}", withErrorHandling(handlePlaylist))
	router.HandleFunc("GET /api/{videoName}/segments/{resolution}/{segment}/", withErrorHandling(handleSegment))
}

func handlePlaylist(w http.ResponseWriter, r *http.Request) error {
	videoName := r.PathValue("videoName")
	slog.Info("Init playlist handler", "video", videoName)

	template := fmt.Sprintf("%v://%v/api/%v/segments/{{.Resolution}}/{{.Segment}}/", "http", r.Host, videoName)
	return hls.WritePlaylist(template, fmt.Sprintf("%v%v", hls.RootDir, videoName), 1080, w)
}

func handleSegment(w http.ResponseWriter, r *http.Request) error {
	videoName := r.PathValue("videoName")
	segment, _ := strconv.ParseInt(r.PathValue("segment"), 0, 64)
	resolution, _ := strconv.ParseInt(r.PathValue("resolution"), 0, 64)
	slog.Info("Init segment handler", "video", videoName, "resolution", resolution, "segment", segment)

	return hls.WriteSegment(fmt.Sprintf("%v%v", hls.RootDir, videoName), segment, resolution, w)
}
