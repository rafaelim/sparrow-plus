package hls

import (
	"fmt"
	"io"
)

func WriteMasterPlaylist(w io.Writer) error {
	fmt.Fprint(w, "#EXTM3U\n")
	fmt.Fprint(w, "#EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID=\"stereo\",LANGUAGE=\"en\",NAME=\"English\",DEFAULT=YES,AUTOSELECT=YES,URI=\"audio/1/index.m3u8\"\n")
	fmt.Fprint(w, "#EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID=\"stereo\",LANGUAGE=\"jpn\",NAME=\"Japonese\",DEFAULT=YES,AUTOSELECT=YES,URI=\"audio/0/index.m3u8\"\n")
	fmt.Fprint(w, "#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=258157,CODECS=\"avc1.4d400d,mp4a.40.2\",AUDIO=\"stereo\",RESOLUTION=422x180,SUBTITLES=\"subs\"\n")
	fmt.Fprint(w, "http://localhost:3000/api/stream/naruto.mkv\n")

	return nil
}
