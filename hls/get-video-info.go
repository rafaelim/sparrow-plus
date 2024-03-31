package hls

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	customerrors "sparrow-plus/custom-errors"
	"strconv"
)

type VideoInfo struct {
	Duration float64
}

func GetVideoInfo(path string) (*VideoInfo, error) {
	format, err := GetFFProbeJson(path)
	if err != nil {
		log.Fatal(err)
	}
	if _, ok := format["duration"]; !ok {
		return nil, customerrors.NewHLSError(
			"Duration was found on the video info",
			"FIELD_NOT_FOUND",
		)
	}
	durr, converr := strconv.ParseFloat(format["duration"].(string), 64)

	if converr != nil {
		return nil, customerrors.NewHLSError(
			"Failed to convert video duration to float",
			"CONVERSION_FAILED",
		)
	}
	var videoInfo = &VideoInfo{Duration: durr}

	return videoInfo, nil
}

func GetFFProbeJson(path string) (map[string]interface{}, error) {
	args := []string{
		"-v",
		"quiet",
		"-print_format",
		"json",
		"-show_format",
		path,
	}

	result, err := exec.Command(FFProbe, args...).Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run command to get video info; err: %v", err.Error())
	}

	var info map[string]interface{}
	jsonerr := json.Unmarshal([]byte(result), &info)
	if jsonerr != nil {
		return nil, fmt.Errorf("failed to convert video info to json; err: %v", jsonerr.Error())
	}

	if _, ok := info["format"]; !ok {
		return nil, fmt.Errorf("no format was found on the video info")
	}

	return info["format"].(map[string]interface{}), nil
}
