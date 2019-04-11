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

	// Load all lpg sounds into a buffer
	bot.LPGSOUND.LoadAll()

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
		if m.Author.Bot || m.Author.ID == s.State.User.ID || len(m.Content) <= 0 || m.Content[0] != '!' {
			return
		}

		// Write into the logs
		_, err = f.WriteString("Time: " + cTime + " || Message: " + m.Content + " || From: " + m.Author.Username + "\n")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Time: %v || Message: %+v || From: %s\n", cTime, m.Content, m.Author)

		// Split the message content
		parts := strings.Split(strings.ToLower(m.Content), " ")

		switch parts[0] {
		case "!hello", "!salut", "!hi":
			_, _ = s.ChannelMessageSend(m.ChannelID, "Salut "+m.Author.Username+" !")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤙")
		case "!chuck":
			fact, _ := bot.ChuckFact()
			_, _ = s.ChannelMessageSend(m.ChannelID, fact)
		case "!sd", "!sound", "!dit":
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					err = bot.PlaySound(s, g.ID, vs.ChannelID, m.Content)
					if err != nil {
						fmt.Println("Error playing sound:", err)
						f.WriteString("Time: " + cTime + " || Error playing " + parts[1] + " for " + m.Author.Username)
					}
					return
				}
			}
		case "!help", "!lpg":
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤙")
			_, _ = s.ChannelMessageSend(m.ChannelID, `Salut, je suis **LPG Bot** ! Je répond aux commandes suivantes :

- **!hello** ou **!hi** : pour me dire bonjour et je te répondrai
- **!chuck** : pour balancer une fact sur chuck norris
- **!sd** ou **!dit <son>** : pour jouer l'un des sons suivants
	-> boi / bruh / fuck / mgs / nice / ooh / oui / thug et wow 
- **!help** ou **!lpg**: pour afficher ce message d'aide ^^ `)
		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Je n'ai pas compris ton message "+m.Author.Username+"  ¯\\_(ツ)_/¯")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤔")
		}

	}

}
