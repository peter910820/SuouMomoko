package commands

import (
	// "fmt"

	"io"
	"log"
	"os"
	"os/exec"

	"github.com/bwmarrin/discordgo"
	"github.com/kkdai/youtube/v2"
)

func sendResponse(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	interaction := i.Interaction
	err := s.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
	if err != nil {
		log.Printf("Failed to send response: %v", err)
	}
}

// func joinRoom(s *discordgo.Session, i *discordgo.InteractionCreate, voiceState *map[string]string) (bool) {
// 	username := i.Member.User.Username
// 	voiceChannel, exists := (*voiceState)[username]
// 	if exists {
// 		voiceConn, _ := s.ChannelVoiceJoin(i.GuildID, voiceChannel, false, false) // ignore error
// 		return true
// 	} else {
// 		return false
// 	}
// }

func Play(s *discordgo.Session, i *discordgo.InteractionCreate, voiceState *map[string]string, url string) {
	// download song
	var video *youtube.Video
	var err error
	client := youtube.Client{}
	video, err = client.GetVideo(url)
	if err != nil {
		sendResponse(s, i, err.Error())
	}
	file, err := os.Create("tmp/" + video.Title + ".mp3")
	if err != nil {
		sendResponse(s, i, err.Error())
	}

	defer file.Close()
	format := video.Formats.Quality("tiny")
	response, _, err := client.GetStream(video, &format[0])
	if err != nil {
		sendResponse(s, i, err.Error())
	}
	_, err = io.Copy(file, response)
	if err != nil {
		sendResponse(s, i, err.Error())
	}
	username := i.Member.User.Username
	voiceChannel, exists := (*voiceState)[username]
	if exists {
		playSound(s, i, voiceChannel, video) // ignore error
	} else {
		sendResponse(s, i, "加入失敗")
	}

}

func playSound(s *discordgo.Session, i *discordgo.InteractionCreate, voiceChannel string, video *youtube.Video) {
	// start FFmpeg
	cmd := exec.Command("ffmpeg", "-i", "tmp/"+video.Title+".mp3", "-f", "opus", "-ar", "48000", "-ac", "2", "-b:a", "192k", "pipe:1")
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return
	}
	if err := cmd.Start(); err != nil {
		log.Fatalln("Error starting command:", err)
		return
	}
	vc, err := s.ChannelVoiceJoin(i.GuildID, voiceChannel, false, true)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		defer cmd.Wait()
		buf := make([]byte, 1024)
		for {
			n, err := pipe.Read(buf)
			if err != nil {
				break
			}
			if n > 0 {
				vc.OpusSend <- buf[:n]
			}
		}
	}()

}
