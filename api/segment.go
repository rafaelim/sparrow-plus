package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sparrow-plus/hls"
)

func segment(w http.ResponseWriter, r *http.Request) {
	segmentName := r.PathValue("segment")

	file, err := os.ReadFile(fmt.Sprintf("%v%v", hls.PlaylistDir, segmentName))

	if err != nil {
		log.Fatalf("Failed to load segment %v", err)
	}
	log.Printf("Returning video segment %v", segmentName)
	w.Write(file)
}