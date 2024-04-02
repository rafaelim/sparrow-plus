package hls

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	customerrors "sparrow-plus/custom-errors"
	"strconv"
)

type VideoInfo struct {
	Duration float64      `json:"duration"`
	Streams  []StreamInfo `json:"streams"`
}

type StreamInfo struct {
	Index      float64 `json:"streamIndex"`
	Codec_type string  `json:"codecType"`
	Language   string  `json:"language"`
}

func GetVideoInfo(path string) (*VideoInfo, error) {
	info, err := GetFFProbeJson(path)
	if err != nil {
		return nil, err
	}
	format, forerr := GetFormat(info)
	if forerr != nil {
		return nil, forerr
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
	streams := GetStreams(info)
	var videoInfo = &VideoInfo{Duration: durr, Streams: streams}

	return videoInfo, nil
}

func GetFormat(info map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := info["format"]; !ok {
		return nil, customerrors.NewHLSError(
			"Format was found on the video info",
			"FIELD_NOT_FOUND",
		)
	}
	return info["format"].(map[string]interface{}), nil
}

func GetStreams(info map[string]interface{}) []StreamInfo {
	if _, ok := info["streams"]; !ok {
		slog.Info("No streams data was found for this video")
		return nil
	}
	var streamsInfo []StreamInfo
	for _, v := range info["streams"].([]interface{}) {
		streamInfo := v.(map[string]interface{})

		if streamInfo["codec_name"] == "hdmv_pgs_subtitle" {
			continue
		}

		language, ok := streamInfo["tags"].(map[string]interface{})["language"].(string)
		if !ok {
			continue
		}

		streamsInfo = append(streamsInfo, StreamInfo{
			Index:      streamInfo["index"].(float64),
			Codec_type: streamInfo["codec_type"].(string),
			Language:   language,
		})
	}
	return streamsInfo
}

func GetFFProbeJson(path string) (map[string]interface{}, error) {
	args := []string{
		"-i", path,
		"-v",
		"quiet",
		"-show_streams",
		"-show_entries", "stream=index,codec_type,codec_name:stream_tags=language",
		"-show_format",
		"-print_format", "json",
	}

	result, err := exec.Command(FFProbe, args...).Output()
	if err != nil {
		return nil, customerrors.NewHLSError(
			fmt.Sprintf("Failed to get info from the video %v", path),
			"COMMAND_FAILED",
		)
	}
	var info map[string]interface{}
	jsonerr := json.Unmarshal([]byte(result), &info)
	if jsonerr != nil {
		return nil, customerrors.NewHLSError(
			fmt.Sprintf("Failed to convert info from the video o json %v; err: %v", path, jsonerr.Error()),
			"COMMAND_FAILED",
		)
	}

	return info, nil
}
