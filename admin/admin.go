package admin

import (
	"time"

	"Work.go/LPG-Bot/LPGBot/print"
	"github.com/bwmarrin/discordgo"
)

func WelcomeTest(s *discordgo.Session, m *discordgo.MessageCreate, g *discordgo.Guild) {
	print.InfoLog("[ADMIN] Start WelcomeTest function", "[SERVER]")
	// create a private messaging channel between the bot and the new guild member
	privChan, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		print.CheckError("[ERROR] Could not retrieve guild object from identifier", "[SERVER]", err)
		return
	}

	welcomeMsg := "**Welcome on server " + g.Name + " ! :video_game:**\n\n I'm LPGBot and before you start your adventure here, please have a look to our channel <#434444368418570243>.\n\n Voilà, that's all. Don't hesitate to come and say hi in <#373160766670503957> \n\n Enjoy ! :call_me:\n"
	embedWelc := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xFF8142, // Orange
		Description: welcomeMsg,
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:       "Hi " + m.Author.Username,
	}
	// send greet message to new guild member
	s.ChannelMessageSendEmbed(privChan.ID, embedWelc)
}

func SetNSFWChannel(args []string) (nsfwChannelID string) {
	print.InfoLog("[ADMIN] Start SetNSFWChannel function", "[SERVER]")
	if len(args) < 2 {
		nsfwChannelID = ""
		print.InfoLog("[ADMIN] Nsfw channeld id set to DEFAULT", "[SERVER]")
		return nsfwChannelID
	} else {
		nsfwChannelID = args[1]
		print.InfoLog("[ADMIN] New nsfw channel id = "+nsfwChannelID+"", "[SERVER]")
		return nsfwChannelID
	}
}
