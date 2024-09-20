package test

import (
	"fmt"

	"github.com/kkdai/youtube/v2"
)

func Yt() {
	videoID := "url"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(video.Author)
}
