package bot

import (
	"fmt"
	"strings"

	"Work.go/LPG-Bot/LPGBot/print"
	"github.com/bwmarrin/discordgo"
)

var newPollString = "Nouveau strawPoll de : "
var react = [6]string{"🇦", "🇧", "🇨", "🇩", "🇪", "🇫"}

//CreatePoll : function used for creqtion of new strawpoll on the server
func CreatePoll(s *discordgo.Session, m *discordgo.MessageCreate) error {
	print.DebugLog("[DEBUG] Start create Poll function - from poll command", "[SERVER]")

	// Print error if comand doesn't have arguments
	if m.Content == "!poll" {
		print.CheckError("[ERROR] No arguments provided for !poll command", "[SERVER]", nil)
		s.ChannelMessageSend(m.ChannelID, "Je ne peux pas créer de strawpoll avec si peu d'arguments "+m.Author.Username)
		return nil
	} else if !strings.Contains((m.Content), "\"") {
		print.CheckError("[ERROR] Double quotes are missing", "[SERVER]", nil)
		s.ChannelMessageSend(m.ChannelID, "La question et les choix doivent etre entre parenthese \" \" "+m.Author.Username)
		return nil
	}

	parts := strings.Split((m.Content), "\"")
	text := "> **" + parts[1] + "**"

	// Check if command has more than 1 argument
	if len(parts) < 2 {
		print.CheckError("[ERROR] Not enought arguments for poll commands", "[SERVER]", nil)
		s.ChannelMessageSend(m.ChannelID, "Je ne peux pas créer de strawpoll avec si peu d'arguments "+m.Author.Username)
		return nil
	} else if (len(parts) / 2) == 1 {
		print.DebugLog("[DEBUG] Single strawpoll created", "[SERVER]")
		s.ChannelMessageSend(m.ChannelID, newPollString+m.Author.Username)
		mBot, _ := s.ChannelMessageSend(m.ChannelID, text)
		s.MessageReactionAdd(m.ChannelID, mBot.ID, "👎")
		s.MessageReactionAdd(m.ChannelID, mBot.ID, "🤷")
		s.MessageReactionAdd(m.ChannelID, mBot.ID, "👍")
	} else if (len(parts) / 2) > len(react) {
		s.ChannelMessageSend(m.ChannelID, "Je ne peux pas créer de strawpoll avec autant d'arguments ¯\\_(ツ)_/¯")
	} else {
		print.DebugLog("[DEBUG] Complexe strawpoll created", "[SERVER]")
		s.ChannelMessageSend(m.ChannelID, newPollString+m.Author.Username)
		decal := 3
		for r := 0; r < ((len(parts) / 2) - 1); r++ {
			text = text + "\n" + react[r] + " " + parts[r+decal]
			decal++
			fmt.Println(text)
		}
		mBot, err := s.ChannelMessageSend(m.ChannelID, text)
		if err != nil {
			print.CheckError("[ERROR] Could not send message", "[SERVER]", err)
		}
		for r := 0; r < ((len(parts) / 2) - 1); r++ {
			err := s.MessageReactionAdd(m.ChannelID, mBot.ID, react[r])
			if err != nil {
				print.CheckError("[ERROR] Could not add reaction", "[SERVER]", err)
			}
		}
	}

	return nil
}
