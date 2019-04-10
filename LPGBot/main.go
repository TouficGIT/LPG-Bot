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

	// Current time : for logs
	cTime := time.Now().Format("01-02-2006 15:04:05")
	fmt.Printf("Time: %v || Message: %+v || From: %s\n", cTime, m.Content, m.Author)
	// Open logs file
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	// Get Channel Id where message has been post
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err)
	}
	// Get the guild (server)
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		fmt.Println(err)
	}

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.Bot || m.Author.ID == s.State.User.ID {
			return
		}

		// Save messages into logs
		_, err = f.WriteString("Time: " + cTime + " || Message: " + m.Content + " || From: " + m.Author.Username + "\n")
		if err != nil {
			panic(err)
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
		case "!oui":
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					err = bot.PlaySound(s, g.ID, vs.ChannelID)
					if err != nil {
						fmt.Println("Error playing sound:", err)
						f.WriteString("Time: " + cTime + " || Error playing !oui for " + m.Author.Username)
					}
					return
				}
			}
		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Je n'ai pas compris ton message "+m.Author.Username+"  ¯\\_(ツ)_/¯")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤔")
		}

	}

}
