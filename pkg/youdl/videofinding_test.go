package youdl

import (
	"fmt"
	"testing"
)

// Passed
func TestFindVideosIDs(t *testing.T) {
	response, err := findVideos("kero kero break", 5)
	if err != nil {
		t.Error(err)
	}
	items := parseResponse(response)
	fmt.Println("List of videos ids found.")
	for _, v := range items.Items {
		fmt.Println(v.Id.VideoId)
	}

}
