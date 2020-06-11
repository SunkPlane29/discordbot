package ytdl

import (
	"fmt"
	"os/exec"
)

func DownloadVideo(title string) {
	response := findVideos(title, 5)
	videos := parseResponse(response)
	downloadUrl := formatDownloadUrl(videos.Items[0].Id.VideoId)
	cmd := exec.Command("youtube-dl", downloadUrl, "--no-playlist")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
