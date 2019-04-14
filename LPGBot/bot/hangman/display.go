package hangman

import (
	"github.com/bwmarrin/discordgo"
)

// DrawWelcome : it used to display Hangman when the game starting
func DrawWelcome(s *discordgo.Session, m *discordgo.MessageCreate) {
	var welc = "```" + `     
	 __ __                             
	/ // /__ ____   ___  ___ _ ___ ___
   / _  / _  / _ \/ _  /    \/ _  / _ \
  /_//_/\_,_/_//_/\_, /_/_/_/\_,_/_//_/
				 /___/                 

	` + "```"
	s.ChannelMessageSend(m.ChannelID, welc)
	return
}

func Draw(s *discordgo.Session, m *discordgo.MessageCreate, g *Game, guess string) {
	drawTurns(s, m, g.TurnsLeft)
	drawState(s, m, g, guess)
}

func drawTurns(s *discordgo.Session, m *discordgo.MessageCreate, l int) {
	var draw string
	switch l {
	case 0:
		draw = "```" + `
	 +---+
	 |   |
	 O   |
	/|\  |
	/ \  |
	   	 |
	=========
	  ` + "```"
	case 1:
		draw = "```" + `
	 +---+
	 |   |
	 O   |
	/|\  |
	     |
	   	 |
	=========
	  ` + "```"
	case 2:
		draw = "```" + `
	 +---+
	 |   |
	 O   |
	     |
	     |
	   	 |
	=========
	  ` + "```"
	case 3:
		draw = "```" + `
	 +---+
	 |   |
	     |
	     |
	     |
	   	 |
	=========
	  ` + "```"
	case 4:
		draw = "```" + `
	 +---+
	     |
	     |
	     |
	     |
	   	 |
	=========
	  ` + "```"
	case 5:
		draw = "```" + `
	
	     |
	     |
	     |
	     |
	   	 |
	=========
	  ` + "```"
	case 6:
		draw = "```" + `

	     |
	     |
	   	 |
	=========
	  ` + "```"
	case 7:
		draw = "```" + `
	=========
	  ` + "```"
	case 8:
		draw = "```" + `

	  ` + "```"
	}
	s.ChannelMessageSend(m.ChannelID, draw)
}

func drawState(s *discordgo.Session, m *discordgo.MessageCreate, g *Game, guess string) {
	s.ChannelMessageSend(m.ChannelID, "Essai: ")
	drawLetters(s, m, g.FoundLetters)

	s.ChannelMessageSend(m.ChannelID, "Utilisé:")
	drawLetters(s, m, g.UsedLetters)

	switch g.State {
	case "goodGuess":
		s.ChannelMessageSend(m.ChannelID, "Bien vu!")
	case "alreadyGuessed":
		s.ChannelMessageSend(m.ChannelID, "La lettre "+guess+" a déjà été utilisée")
	case "badGuess":
		s.ChannelMessageSend(m.ChannelID, "Dommage, "+guess+" n'est pas dans le mot")
	case "lost":
		s.ChannelMessageSend(m.ChannelID, "Perdu, le mot été: ")
		drawLetters(s, m, g.Letters)
	case "won":
		s.ChannelMessageSend(m.ChannelID, "Gagné! Le mot été: ")
		drawLetters(s, m, g.Letters)
	}
}

func drawLetters(s *discordgo.Session, m *discordgo.MessageCreate, l []string) {
	for _, c := range l {
		s.ChannelMessageSend(m.ChannelID, c)
	}
	s.ChannelMessageSend(m.ChannelID, "")
}
