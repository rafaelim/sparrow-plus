package hls

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
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
		var err = fmt.Errorf("no duration was found on the video info")
		log.Fatal(err)
		return nil, err
	}
	durr, converr := strconv.ParseFloat(format["duration"].(string), 64)

	if converr != nil {
		log.Fatal(converr)
		return nil, converr
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
		log.Fatal(err)
		return nil, err
	}
	
	var info map[string]interface{}
	jsonerr := json.Unmarshal([]byte(result), &info)

	if jsonerr != nil {
		log.Fatalf("Failed to convert video info to json %v",err)
		return nil, err
	}
	if _, ok := info["format"]; !ok {
		return nil, fmt.Errorf("no format was found on the video info")
	}
	return info["format"].(map[string]interface{}), nil
}