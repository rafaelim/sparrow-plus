package hls

import (
	"fmt"
	"io"
	"net/url"
)

func WriteMasterPlaylist(path string, values url.Values, w io.Writer) error {
	videoInfo, err := GetVideoInfo(path)

	if err != nil {
		return fmt.Errorf("error while getting video info %w", err)
	}

	fmt.Fprint(w, "#EXTM3U\n")
	for _, stream := range videoInfo.Streams {
		defaultLang := "NO"
		if stream.Language == "eng" {
			defaultLang = "YES"
		}
		if stream.Codec_type == "audio" {
			fmt.Fprintf(
				w,
				"#EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID=\"%v\",LANGUAGE=\"%v\",NAME=\"%v\",DEFAULT=%v,AUTOSELECT=YES,URI=\"%v\"\n",
				"stereo",
				stream.Language,
				stream.LanguageName,
				defaultLang,
				fmt.Sprintf("audio/%v/index.m3u8?%v", stream.Index, values.Encode()),
			)
		}
		if stream.Codec_type == "subtitle" {
			fmt.Fprintf(
				w,
				"#EXT-X-MEDIA:TYPE=SUBTITLES,GROUP-ID=\"%v\",LANGUAGE=\"%v\",NAME=\"%v\",DEFAULT=%v,AUTOSELECT=YES,URI=\"%v\"\n",
				"subs",
				stream.Language,
				stream.LanguageName,
				defaultLang,
				fmt.Sprintf("subtitles/%v/index.m3u8?%v", stream.Index, values.Encode()),
			)
		}
	}
	fmt.Fprint(w, "#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=258157,CODECS=\"avc1.4d400d,mp4a.40.2\",AUDIO=\"stereo\",RESOLUTION=422x180,SUBTITLES=\"subs\"\n")
	fmt.Fprintf(w, "index.m3u8?%v", values.Encode())

	return nil
}
