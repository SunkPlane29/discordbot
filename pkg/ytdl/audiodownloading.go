package ytdl

import (
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/jonas747/dca"
)

// Download mp3 files to then convert them to dca (actually it downloads
// mp4 file first) to then, after 5 minutes deleting them from the computer
// (you don't want to have all musics stored on your pc)
func DownloadDca(title, filepath string) (string, error) {
	filename, err := downloadAudio(title)
	if err != nil {
		return "", err
	}
	err = convertDca(filepath, filename)
	if err != nil {
		return "", err
	}
	go func() {
		time.Sleep(time.Minute * 5)
		os.Remove(filepath + filename + ".dca")
	}()
	return filename + ".dca", nil
}

// Converts mp4 files given a filepath to dca files and deletes the original.
func convertDca(filepath, filename string) error {
	encodeSession, err := dca.EncodeFile(filepath+filename+".mp3", dca.StdEncodeOptions)
	if err != nil {
		return err
	}
	defer encodeSession.Cleanup()
	output, err := os.Create(filepath + filename + ".dca")
	if err != nil {
		return err
	}
	io.Copy(output, encodeSession)
	err = os.Remove(filepath + filename + ".mp3")
	if err != nil {
		return err
	}
	return nil
}

// Downloads the video still but it is converted and replaced by an audio.
// Return a the video id if succeeded
func downloadAudio(title string) (string, error) {
	response, err := findVideos(title, 5)
	if err != nil {
		return "", err
	}
	videos := parseResponse(response)
	downloadUrl := formatDownloadUrl(videos.Items[0].Id.VideoId)
	cmd := exec.Command("youtube-dl", "--config-location", "pkg/ytdl/youtube-dl.conf", downloadUrl)
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	return videos.Items[0].Id.VideoId, nil
}
