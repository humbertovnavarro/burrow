package youtubemuse

import (
	"io"
	"os"
	"path"

	"github.com/kkdai/youtube/v2"
)

var client youtube.Client

func GetBestAudioFormat(list youtube.FormatList) youtube.Format {
	list = list.WithAudioChannels()
	bestFormat := list[0]
	highest := 0
	for _, format := range list {
		score := 0

		switch format.AudioQuality {
		case "AUDIO_QUALITY_HIGH":
			score += 3
		case "AUDIO_QUALITY_MEDIUM":
			score += 2
		}

		switch format.AudioSampleRate {
		case "44100":
			score += 2
		case "48000":
			score += 1
		}

		if score > highest {
			highest = score
			bestFormat = format
		}
	}
	return bestFormat
}

func GetFileStream(videoId string) (file *os.File, err error) {
	fileName := path.Join("./cache", videoId)
	file, err = os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	video, err := client.GetVideo(videoId)
	if err != nil {
		return
	}
	format := GetBestAudioFormat(video.Formats)
	stream, _, err := client.GetStream(video, &format)
	if err != nil {
		return
	}
	_, err = io.Copy(file, stream)
	return
}
