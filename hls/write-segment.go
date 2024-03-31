package hls

import (
	"fmt"
	"io"
	"os/exec"
	customerrors "sparrow-plus/custom-errors"
)

func WriteSegment(video string, segment, resolution int64, w io.Writer) error {
	startTime := segment * hlsSegmentLenght
	args := []string{
		"-timelimit", "45",
		"-ss", fmt.Sprintf("%v.00", startTime),
		"-i", video,
		"-t", fmt.Sprintf("%v.00", hlsSegmentLenght),
		"-vf", fmt.Sprintf("scale=-2:%v", resolution),
		"-vcodec", "libx264",
		"-preset", "veryfast",
		"-c:a", "aac",
		"-b:a", "128k",
		"-ac", "2",
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
			fmt.Sprintf("Failed to create playlist for the video %v", video),
			"COMMAND_FAILED",
		)
	}

	fmt.Fprint(w, string(output))
	return nil
}
