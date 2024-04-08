package hls

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	customerrors "sparrow-plus/custom-errors"
)

func WritePlaylist(urlTemplate string, file string, resolution int, w io.Writer) error {
	t := template.Must(template.New("urlTemplate").Parse(urlTemplate))

	vinfo, err := GetVideoInfo(file)

	if err != nil {
		return customerrors.NewHLSError(
			"Failed to load video info",
			"VIDEO_INFO_NOT_FOUND",
		)
	}

	duration := vinfo.Duration

	getUrl := func(segmentIndex int) string {
		buf := new(bytes.Buffer)
		t.Execute(buf, struct {
			Resolution int
			Segment    int
		}{
			resolution,
			segmentIndex,
		})
		return buf.String()
	}

	fmt.Fprint(w, "#EXTM3U\n")
	fmt.Fprint(w, "#EXT-X-VERSION:3\n")
	fmt.Fprint(w, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(w, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprint(w, "#EXT-X-TARGETDURATION:"+fmt.Sprintf("%.f", hlsSegmentLenght)+"\n")
	fmt.Fprint(w, "#EXT-X-PLAYLIST-TYPE:VOD\n")

	leftover := duration
	segmentIndex := 0

	for leftover > 0 {
		if leftover > hlsSegmentLenght {
			fmt.Fprintf(w, "#EXTINF: %f,\n", hlsSegmentLenght)
		} else {
			fmt.Fprintf(w, "#EXTINF: %f,\n", leftover)
		}
		fmt.Fprintf(w, getUrl(segmentIndex)+"\n")
		segmentIndex++
		leftover = leftover - hlsSegmentLenght
	}
	fmt.Fprint(w, "#EXT-X-ENDLIST\n")
	return nil
}
