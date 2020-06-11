package ytdl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Items struct {
	Items []Item
}

type Item struct {
	Id Id
}

type Id struct {
	Kind    string
	VideoId string
}

// Gets a video list through a GET request returning a json file, now need to fetch the videos ids
// in that json
func findVideos(search string, maxResults int) []byte {

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

	return respBytes

}

// Parses the json youtube response to the struct Items
func parseResponse(response []byte) Items {
	var items Items
	json.Unmarshal(response, &items)
	return items

}
