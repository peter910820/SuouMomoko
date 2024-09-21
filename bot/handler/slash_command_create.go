package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func BasicCommand(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "test",
			Description: "Test discord bot response",
		},
		{
			Name:        "ping",
			Description: "return bot heartbeatlatency",
		},
	}
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			fmt.Println("[ERROR] ", err)
		}
	}
}

func TestingCommand(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "voice_check",
			Description: "checking all VoiceStateUpdate event",
		},
	}
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			fmt.Println("[ERROR] ", err)
		}
	}
}

func MusicCommnad(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "join",
			Description: "add bot into voice channle",
		},
		{
			Name:        "leave",
			Description: "kick bot out to voice channle",
		},
		{
			Name:        "play",
			Description: "add bot into voice channle and play song or add song into playlist",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "url",
					Description: "enter your youtube url",
					Required:    true,
				},
			},
		},
		{
			Name:        "stop",
			Description: "stop song temporary",
		},
		{
			Name:        "resume",
			Description: "resume song",
		},
		{
			Name:        "insert",
			Description: "insert song to specify location",
		},
		{
			Name:        "now",
			Description: "show playing(current) song",
		},
		{
			Name:        "skip",
			Description: "skip playing(current) song",
		},
	}
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			fmt.Println("[ERROR] ", err)
		}
	}
}
