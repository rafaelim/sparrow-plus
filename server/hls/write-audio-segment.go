package hls

import (
	"fmt"
	"io"
	"os/exec"
	customerrors "sparrow-plus/custom-errors"
)

func WriteAudioSegment(video string, segment, index int, w io.Writer) error {
	startTime := segment * hlsSegmentLenght
	args := []string{
		"-ss", fmt.Sprintf("%v.00", startTime),
		"-i", video,
		"-t", fmt.Sprintf("%v.00", hlsSegmentLenght),
		"-map", fmt.Sprintf("0:%d", index),
		"-c:a", "aac",
		"-ac", "2",
		"-b:a", "128k",
		"-muxdelay", "0",
		"-pix_fmt", "yuv420p",
		"-force_key_frames", "expr:gte(t,n_forced*5.000)",
		"-f", "ssegment",
		"-segment_time", fmt.Sprintf("%v.00", hlsSegmentLenght),
		"-initial_offset", fmt.Sprintf("%v.00", startTime),
		"pipe:out%03d.ts",
	}
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
