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
	botId      string
	bot        *discordgo.Session
	voiceState map[string]string
)

func main() {
	voiceState = make(map[string]string)

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
	bot.AddHandler(onInteractionTesting)
	bot.AddHandler(voiceStateUpdate)

	err = bot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	<-make(chan struct{})
}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	fmt.Println("momoko is alreadyyyyyy!!!")
	s.UpdateGameStatus(0, "青夏軌跡")
	handler.BasicCommand(s)
	handler.TestingCommand(s)
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
func voiceStateUpdate(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	user, err := s.User(vs.UserID)
	if err != nil {
		fmt.Printf("無法獲取用戶資訊: %v", err)
		return
	}
	if vs.ChannelID == "" {
		delete(voiceState, user.Username)
		fmt.Printf("%s 離開了頻道\n", user.Username)
	} else {
		voiceState[user.Username] = vs.ChannelID
		fmt.Printf("%s 加入了頻道 %s\n", user.Username, vs.ChannelID)
	}
}

func onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		cmdData, ok := i.Data.(discordgo.ApplicationCommandInteractionData)
		if !ok {
			fmt.Println("Type ERROR!")
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
			go commands.JoinRoom(s, i, &voiceState)

		}
	}
}

func onInteractionTesting(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		cmdData, ok := i.Data.(discordgo.ApplicationCommandInteractionData)
		if !ok {
			fmt.Println("類型錯誤!")
			return
		}
		switch cmdData.Name {
		case "voice_check":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("**This is develop mode**\nAll VoiceStateUpdate event: %v", voiceState),
				},
			})
		}
	}
}
