package hls

import (
	"fmt"
	"sparrow-plus/config"
)

const (
	ContentType = "application/vnd.apple.mpegurl"
	PlaylistFilename = "playlist.m3u8"

	hlsTime = 10.0
)

var RootDir = config.ReadConfig().RootDir
var PlaylistDir = fmt.Sprintf("%v%v", RootDir, "playlist/")
var FFMpeg = "ffmpeg"
var FFProbe = "ffprobe"