package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sparrow-plus/hls"
)

type FileList struct {
	IsDir bool
	Filename string
}

func fileList(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(hls.RootDir)
	if err != nil {
		log.Fatalf("Failed to read dir %v; err: %v", hls.RootDir, err)
	}
	var response = []FileList{}
	for _, file := range files {
		response = append(response, FileList{IsDir: file.IsDir(), Filename: file.Name() })
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
} 
