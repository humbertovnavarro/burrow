package main

import (
	"github.com/humbertovnavarro/burrow/services/burrow-muse/pkg/youtubemuse"
)

func main() {
	videoID := "0mnXzCCpctU"
	f, err := youtubemuse.GetOpusFile(videoID)
	if err != nil {
		panic(err)
	}
	f.Close()
}
