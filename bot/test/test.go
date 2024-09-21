package test

import (
	"fmt"

	"github.com/kkdai/youtube/v2"
)

func Yt() {
	videoID := "https://music.youtube.com/watch?v=fWVhcojPqQI&list=PLM0jRdHj2C1nzyw0HWlFDXOf_F3QOFCf4"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(video.Author)
}
