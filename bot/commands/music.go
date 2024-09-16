package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func JoinRoom(s *discordgo.Session, i *discordgo.InteractionCreate, voiceState *map[string]string) {
	username := i.Member.User.Username
	voiceChannel, exists := (*voiceState)[username]
	if exists {
		go func() {
			_, err := s.ChannelVoiceJoin(i.GuildID, voiceChannel, false, false)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "已加入頻道",
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "未加入頻道",
			},
		})
	}
}
