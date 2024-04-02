package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	customerrors "sparrow-plus/custom-errors"
	"sparrow-plus/hls"
	"strconv"
)

func (s *APIServe) VideoStreamHandlers(router *http.ServeMux) {
	router.HandleFunc("GET /api/stream/{videoName}", withErrorHandling(handlePlaylist))
	router.HandleFunc("GET /api/stream/info/{videoName}", withErrorHandling(handleVideoInfo))
	router.HandleFunc("GET /api/stream/subtitles/{videoName}/{subIndex}", withErrorHandling(handleSubtitles))
	router.HandleFunc("GET /api/stream/{videoName}/segments/{resolution}/{segment}/", withErrorHandling(handleSegment))
}

func handlePlaylist(w http.ResponseWriter, r *http.Request) error {
	videoName := r.PathValue("videoName")
	slog.Info("Init playlist handler", "video", videoName)

	template := fmt.Sprintf("%v://%v/api/stream/%v/segments/{{.Resolution}}/{{.Segment}}/", "http", r.Host, videoName)

	err := hls.WritePlaylist(template, filepath.Join(hls.RootDir, videoName), 1080, w)

	if err != nil {
		return customerrors.NewRequestError(
			http.StatusBadRequest,
			err.Error(),
			"FAILED_TO_CREATE_PLAYLIST",
		)
	}

	return nil
}

func handleSegment(w http.ResponseWriter, r *http.Request) error {
	videoName := r.PathValue("videoName")
	segment, _ := strconv.ParseInt(r.PathValue("segment"), 0, 64)
	resolution, _ := strconv.ParseInt(r.PathValue("resolution"), 0, 64)
	slog.Info("Init segment handler", "video", videoName, "resolution", resolution, "segment", segment)

	err := hls.WriteSegment(filepath.Join(hls.RootDir, videoName), segment, resolution, w)

	if err != nil {
		return customerrors.NewRequestError(
			http.StatusBadRequest,
			err.Error(),
			"FAILED_TO_CREATE_SEGMENT",
		)
	}

	return nil
}

func handleVideoInfo(w http.ResponseWriter, r *http.Request) error {
	videoName := r.PathValue("videoName")
	slog.Info("Init video info handler", "video", videoName)

	videoInfo, err := hls.GetVideoInfo(filepath.Join(hls.RootDir, videoName))

	if err != nil {
		fmt.Println(err)
		return customerrors.NewRequestError(
			http.StatusBadRequest,
			err.Error(),
			"FAILED_TO_GET_VIDEO_INFO",
		)
	}

	json.NewEncoder(w).Encode(videoInfo)

	return nil
}
func handleSubtitles(w http.ResponseWriter, r *http.Request) error {
	videoName := r.PathValue("videoName")
	subIndex, _ := strconv.Atoi(r.PathValue("subIndex"))
	slog.Info("Init subtitles handler", "video", videoName, "subIndex", subIndex)

	err := hls.WriteSubtitles(filepath.Join(hls.RootDir, videoName), subIndex, w)

	if err != nil {
		fmt.Println(err)
		return customerrors.NewRequestError(
			http.StatusBadRequest,
			err.Error(),
			"FAILED_TO_CREATE_SUBTITLES",
		)
	}

	return nil
}
