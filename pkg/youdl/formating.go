package youdl

import (
	"strings"
)

func formatSearchUrl(search string) string {

	searchContents := strings.Split(search, " ")
	searchUrl := strings.Join(searchContents, "%20")
	return searchUrl
}

func formatDownloadUrl(videoId string) string {
	baseUrl := "https://www.youtube.com/watch?v="
	return baseUrl + videoId
}

// Formats a watch URL given the title
func FormatWhatchURL(title string) (string, error) {

	response, err := findVideos(title, 5)
	if err != nil {
		return "", err
	}
	videos := parseResponse(response)
	downloadURL := formatDownloadUrl(videos.Items[0].Id.VideoId)
	return downloadURL, nil
}
