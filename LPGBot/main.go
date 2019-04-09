package main

import (
	"fmt"
	"strings"

	"Work.go/LPG-Bot/LPGBot/config"
	"github.com/bwmarrin/discordgo"
)

func main() {

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lpgBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	lpgBot.AddHandler(messageHandler)
	err = lpgBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running !")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.Bot {
			return
		}

		switch m.Content {
		case "!ping":
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Je ne comprend pas ¯\\_(ツ)_/¯")

		}

	}

}
