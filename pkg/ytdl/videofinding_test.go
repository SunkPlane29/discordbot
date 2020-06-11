package ytdl

import (
	"fmt"
	"testing"
)

func TestFindVideosIDs(t *testing.T) {
	response := findVideos("kero kero break", 5)
	items := parseResponse(response)
	fmt.Println(items.Items[0].Id.VideoId)
	for _, v := range items.Items {
		fmt.Println(v.Id.VideoId)
	}

}


