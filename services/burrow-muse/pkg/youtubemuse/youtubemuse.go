package youtubemuse

import (
	"io"
	"os"
	"os/exec"

	"github.com/kkdai/youtube/v2"
)

var client youtube.Client

func GetBestAudioFormat(list youtube.FormatList) youtube.Format {
	list = list.WithAudioChannels()
	bestFormat := list[0]
	highest := 0
	for _, format := range list {
		score := 0
		// self explanatory
		switch format.AudioQuality {
		case "AUDIO_QUALITY_HIGH":
			score += 3
		case "AUDIO_QUALITY_MEDIUM":
			score += 2
		case "AUDIO_QUALITY_LOW":
			score -= 1
		}

		switch format.AudioSampleRate {
		// same sample rate as discord voice = good
		case "48000":
			score += 1
		case "44100":
			score -= 1
		default:
			// strange sample rate = bad
			score -= 2
		}
		// easier to encode
		if format.FPS <= 30 {
			score += 1
		}
		// harder to encode
		if format.FPS <= 60 {
			score -= 1
		}

		if score > highest {
			highest = score
			bestFormat = format
		}
	}
	return bestFormat
}

func GetVideoStream(videoId string) (reader io.ReadCloser, err error) {
	video, err := client.GetVideo(videoId)
	if err != nil {
		return
	}
	format := GetBestAudioFormat(video.Formats)
	stream, _, err := client.GetStream(video, &format)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

/*
GetOpusFile downloads the best audio stream from a youtube video,
encodes it to opus and pipes it to a file. Returns file pointer that must be closed.
*/
func GetOpusFile(videoId string) (file *os.File, err error) {
	stream, err := GetVideoStream(videoId)
	if err != nil {
		return nil, err
	}
	file, err = os.Create(videoId + ".opus")
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("ffmpeg", "-i", "-", "-vn", "-f", "opus", "-")
	cmd.Stdin = stream
	cmd.Stdout = file
	cmdErr := cmd.Start()
	if cmdErr != nil {
		// clean up on error to not leave a file behind
		err = file.Close()
		if err != nil {
			panic(err)
		}
		err = os.Remove(file.Name())
		if err != nil {
			panic(err)
		}
		return nil, cmdErr
	}
	return file, nil
}
