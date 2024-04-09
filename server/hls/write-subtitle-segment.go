package hls

import (
	"fmt"
	"io"
	"os/exec"
	customerrors "sparrow-plus/custom-errors"
)

func WriteSubtiteSegment(video string, segment, index int, w io.Writer) error {
	args := []string{
		"-i", video,
		"-vn",
		"-an",
		"-map", fmt.Sprintf("0:%v", index),
		"-f", "webvtt",
		"-muxdelay", "0",
		"pipe:",
	}
	fmt.Println(args)
	output, err := exec.Command(FFMpeg, args...).Output()

	if err != nil {
		return customerrors.NewHLSError(
			fmt.Sprintf("Failed to create segment for the video %v", video),
			"COMMAND_FAILED",
		)
	}

	fmt.Fprint(w, string(output))
	return nil
}
