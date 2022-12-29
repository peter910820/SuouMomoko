package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly/v2"
)

var (
	Token  string = "MTA1Mzk2MTgxODI2NTEwODUzMA.GSNE99.4UDATCG8h7RJB5cmnCchfA_ztT6cbDaDsLGSDo"
	Prefix string = "$"
	botId  string
)

func main() {
	bot, err := discordgo.New("Bot " + Token)

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

	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)

	err = bot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("啟動成功!")

	<-make(chan struct{})
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)

	if m.Author.ID == botId {
		return
	}

	if m.Content == Prefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	} else if m.Content == Prefix+"IT-q" {
		crawler, crawlerTime := itCrawler("https://ithelp.ithome.com.tw/questions")

		result := "```\n"
		count := 0
		for _, cw := range crawler {
			result = result[:] + cw + " on " + crawlerTime[count]
			result = result[:] + "\n"
			count++
		}
		result = result[:] + "```"
		fmt.Println(crawler)
		_, _ = s.ChannelMessageSend(m.ChannelID, result)
	} else if m.Content == Prefix+"IT-a" {
		crawler, crawlerTime := itCrawler("https://ithelp.ithome.com.tw/articles?tab=tech")

		result := "```\n"
		count := 0
		for _, cw := range crawler {
			result = result[:] + cw + " on " + crawlerTime[count]
			result = result[:] + "\n"
			count++
		}
		result = result[:] + "```"
		fmt.Println(crawler)
		_, _ = s.ChannelMessageSend(m.ChannelID, result)
	}

}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateGameStatus(0, "偶像大師")
}

func itCrawler(url string) ([]string, []string) {

	countLink, countTime := 0, 0
	var (
		data     []string
		dataTime []string
	)

	c := colly.NewCollector()

	c.OnHTML(".qa-list__title-link", func(title *colly.HTMLElement) {
		data = append(data, title.Text)
		// fmt.Println(data)
		countLink++
	})

	c.OnHTML(".qa-list__info-time", func(title *colly.HTMLElement) {
		dataTime = append(dataTime, title.Text)
		// fmt.Println(data)
		countTime++
	})

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println(string(r.Body))
	// })

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.46")
	})

	c.Visit(url)

	return data, dataTime
}
