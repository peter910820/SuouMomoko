package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func SlashCommand(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "test",
			Description: "Test discord bot response",
		},
		{
			Name:        "ping",
			Description: "pong!",
		},
	}
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			fmt.Println("[ERROR] ", err)
		}
	}
}
