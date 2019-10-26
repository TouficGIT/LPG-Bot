package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var newPollString = "Nouveau strawPoll de : "
var react = [6]string{"🇦", "🇧", "🇨", "🇩", "🇪", "🇫"}

//CreatePoll : function used for creqtion of new strawpoll on the server
func CreatePoll(s *discordgo.Session, m *discordgo.MessageCreate) error {
	fmt.Println("START : createPoll function - from poll command")

	if m.Content == "!poll" {
		fmt.Println("ERROR : No arguments provided for !poll command")
		s.ChannelMessageSend(m.ChannelID, "Je ne peux pas créer de strawpoll avec si peu d'arguments "+m.Author.Username)
		return nil
	}

	parts := strings.Split((m.Content), "\"")
	text := "> **" + parts[1] + "**"

	if len(parts) < 2 {
		fmt.Println("ERROR : Not enought arguments for poll commands")
		s.ChannelMessageSend(m.ChannelID, "Je ne peux pas créer de strawpoll avec si peu d'arguments "+m.Author.Username)
		return nil
	} else if (len(parts) / 2) == 1 {
		fmt.Println("INFO : Single strawpoll created")
		s.ChannelMessageSend(m.ChannelID, newPollString+m.Author.Username)
		mBot, _ := s.ChannelMessageSend(m.ChannelID, text)
		s.MessageReactionAdd(m.ChannelID, mBot.ID, "👎")
		s.MessageReactionAdd(m.ChannelID, mBot.ID, "🤷")
		s.MessageReactionAdd(m.ChannelID, mBot.ID, "👍")
	} else if (len(parts) / 2) > len(react) {
		fmt.Println("ERROR : Too much arguments for poll commands")
		s.ChannelMessageSend(m.ChannelID, "Je ne peux pas créer de strawpoll avec autant d'arguments ¯\\_(ツ)_/¯")
	} else {
		fmt.Println("INFO : Complexe strawpoll created")
		s.ChannelMessageSend(m.ChannelID, newPollString+m.Author.Username)
		decal := 3
		for r := 0; r < ((len(parts) / 2) - 1); r++ {
			text = text + "\n" + react[r] + " " + parts[r+decal]
			decal++
			fmt.Println(text)
		}
		mBot, err := s.ChannelMessageSend(m.ChannelID, text)
		if err != nil {
			fmt.Println(err)
		}
		for r := 0; r < ((len(parts) / 2) - 1); r++ {
			err := s.MessageReactionAdd(m.ChannelID, mBot.ID, react[r])
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}
