package ytdl

import (
	"strings"
)

func formatSearchUrl(search string) string {

	searchContents := strings.Split(search, " ")
	searchUrl := strings.Join(searchContents, "%20")
	return searchUrl
}
