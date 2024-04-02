package hls

import (
	"fmt"
	"io"
	"os/exec"
	customerrors "sparrow-plus/custom-errors"
)

func WriteSubtitles(video string, index int, w io.Writer) error {
	args := []string{
		"-i", video,
		"-map", fmt.Sprintf("0:%v", index),
		"-f", "webvtt",
		"pipe:",
	}
	output, err := exec.Command(FFMpeg, args...).Output()

	if err != nil {
		return customerrors.NewHLSError(
			fmt.Sprintf("Failed to extract subtitles and audiotracks from the video %v", video),
			"COMMAND_FAILED",
		)
	}

	fmt.Fprint(w, string(output))
	return nil
}
