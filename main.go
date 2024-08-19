package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"momoko-bot/bot/commands"
	"momoko-bot/bot/handler"
	"os"
)

var (
	botId string
	bot   *discordgo.Session
)

func main() {
	err := godotenv.Load(".env")
	token := os.Getenv("DISCORD_BOT_TOKEN")

	if err != nil {
		fmt.Println("Error loading .env file!!!")
		return
	}

	bot, err = discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := bot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botId = u.ID

	bot.AddHandler(ready) //註冊事件 建議換為指定事件
	bot.AddHandler(messageCreate)
	bot.AddHandler(onInteraction)

	err = bot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	<-make(chan struct{})
}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	fmt.Println("momoko is alreadyyyyyy!!!")
	s.UpdateGameStatus(0, "偶像大師")
	handler.SlashCommand(s)
	handler.MusicCommnad(s)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("Message: %s\n", m.Content)

	if m.Author.ID == botId { // avoid message loop
		return
	}

	switch m.Content {
	case "!ping":
		commands.PingCommand(s, m)
	case "!test":
		commands.TestCommand(s, m)

	}
}

func onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		cmdData, ok := i.Data.(discordgo.ApplicationCommandInteractionData)
		if !ok {
			fmt.Println("類型錯誤!")
			return
		}
		switch cmdData.Name {
		case "test":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "test!",
				},
			})
		case "ping":
			delay := bot.HeartbeatLatency()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("現在延遲為: %v", delay),
				},
			})
		case "join":
			a := i.ChannelID
			// member, err := s.State.Member(i.GuildID, i.Member.User.ID)
			// if err != nil {
			// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
			// 		Data: &discordgo.InteractionResponseData{
			// 			Content: fmt.Sprintf("[ERROR] %s", err),
			// 		},
			// 	})
			// 	return
			// }
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: a,
				},
			})
		}
	}
}
