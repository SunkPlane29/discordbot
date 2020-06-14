package youdl

import (
	"testing"
)

// Passed
func TestFormatWatchURL(t *testing.T) {
	URL, err := FormatWhatchURL("i don't know how choke")
	if err != nil {
		t.Error("Download URL not found.")
	}
	expected := "https://www.youtube.com/watch?v=mvJjmWTg7Qo"
	if URL != expected {
		t.Errorf("Error: got %s, expected %s", URL, expected)
	}
}
