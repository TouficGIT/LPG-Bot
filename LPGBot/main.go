package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"Work.go/LPG-Bot/LPGBot/bot"
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

	fmt.Printf("Time: %v || Message: %+v || From: %s\n", time.UnixDate, m.Content, m.Author)
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString("Time: " + time.UnixDate + " || Message: " + m.Content + " || From: " + m.Author.Username + "\n")
	if err != nil {
		panic(err)
	}

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.Bot {
			return
		}

		switch m.Content {
		case "!ping":
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		case "!hello", "!salut", "!hi":
			_, _ = s.ChannelMessageSend(m.ChannelID, "Salut "+m.Author.Username+" !")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤙")
		case "!chuck":
			fact, _ := bot.ChuckFact()
			_, _ = s.ChannelMessageSend(m.ChannelID, fact)
		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Je n'ai pas compris ton message "+m.Author.Username+"  ¯\\_(ツ)_/¯")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤔")
		}

	}

}
