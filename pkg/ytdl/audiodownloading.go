package ytdl

import (
	"io"
	"os"
	"os/exec"

	"github.com/jonas747/dca"
)

// Test that function later
func convertDca(filepath string) error {
	encodeSession, err := dca.EncodeFile(filepath, dca.StdEncodeOptions)
	if err != nil {
		return err
	}
	defer encodeSession.Cleanup()
	output, err := os.Create("music.dca")
	if err != nil {
		return err
	}
	io.Copy(output, encodeSession)
	return nil
}

// Downloads the video still but it is converted and replaced by an audio.
// Return an err later
func downloadAudio(title string) error {
	response, err := findVideos(title, 5)
	if err != nil {
		return err
	}
	videos := parseResponse(response)
	downloadUrl := formatDownloadUrl(videos.Items[0].Id.VideoId)
	cmd := exec.Command("youtube-dl", "--config-location", "./youtube-dl.conf", downloadUrl)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
