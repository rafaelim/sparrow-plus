package cmd

import (
	"log"
	"os/exec"
	"sparrow-plus/hls"
)

func SetupEnv() {
	ffmpeg, ffmpegerr := exec.LookPath("ffmpeg")
	if ffmpegerr != nil {
		log.Fatalf("ffmpeg was not found %v", ffmpegerr)
	}

	ffprobe, ffprobeerr := exec.LookPath("ffprobe")
	if ffprobeerr != nil {
		log.Fatalf("ffprobe was not found %v", ffprobeerr)
	}

	hls.FFMpeg = ffmpeg
	hls.FFProbe = ffprobe
	hls.RootDir = "/home/rafael/personal/sparrow-plus/videos"
}
