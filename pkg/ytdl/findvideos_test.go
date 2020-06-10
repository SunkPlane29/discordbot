package ytdl

import (
	"fmt"
	"testing"
)

func TestFindVideosIDs(t *testing.T) {
	got := findVideosIDs("f society", 5)
	fmt.Println(got)
}
