package ytdl

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
