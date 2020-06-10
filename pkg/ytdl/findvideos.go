package ytdl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Gets a video list through a GET request returning a json file, now need to fetch the videos ids
// in that json
func findVideosIDs(search string, maxResults int) string {

	APIKey := "AIzaSyA-NPrfWfSxI5v_fL1KI8HvD8Z8x9KOAnU"

	baseUrl := "https://www.googleapis.com/youtube/v3/search?"

	url := baseUrl + fmt.Sprintf("maxResults=%d", maxResults) + "&q=" + formatSearchUrl(search) + "&key=" + APIKey
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	respBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(respBytes)

}
