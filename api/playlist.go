package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sparrow-plus/hls"
)

func playlist(w http.ResponseWriter, r *http.Request) {
	videoName := r.PathValue("videoName")

	err := hls.CreateSegmentsAndPlaylist(fmt.Sprintf("%v%v", hls.RootDir, videoName))

	if err != nil {
		log.Fatalf("Failed to write playlist %v", err)
	}
	
	w.Header().Set("Content-type", hls.ContentType)
	playlist, filerr := os.ReadFile(fmt.Sprintf("%v%v", hls.PlaylistDir, hls.PlaylistFilename))

	if filerr != nil {
		log.Fatalf("Failed to load playlist file %v", err)
	}

	w.Write(playlist)
}