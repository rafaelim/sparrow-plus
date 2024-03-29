package hls

import (
	"fmt"
	"log"
	"os/exec"
)

func CreateSegmentsAndPlaylist(video string) error {
	log.Printf("Creating playlist for the video %v", video)
	args := []string{
		"-i",
		video,
		"-codec:",
		"copy",
		"-start_number",
		"0",
		"-hls_time",
		fmt.Sprintf("%v", hlsTime),
		"-hls_list_size",
		"0",
		"-f",
		"hls",
		fmt.Sprintf("%v%v", PlaylistDir, PlaylistFilename),
	}

	_, err := exec.Command(FFMpeg, args...).Output()

	if err != nil {
		log.Fatalf("Failed to create playlist for the video %v; err: %v", video, err)
		return err
	}

	return nil
}
