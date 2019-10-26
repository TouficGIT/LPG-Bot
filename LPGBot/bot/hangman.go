package bot

import (
	"fmt"
	"strconv"
	"strings"

	"Work.go/LPG-Bot/LPGBot/bot/hangman"
	"github.com/bwmarrin/discordgo"
)

var h *hangman.Game
var end int

// NewHangman : it used to launch hangman game
func NewHangman(s *discordgo.Session, g *discordgo.Guild, user *discordgo.User, m *discordgo.MessageCreate) *hangman.Game {
	fmt.Println("START : NewHangman function - from hangman command")
	end = 0
	err := hangman.Load("bot/hangman/words.txt")
	if err != nil {
		fmt.Printf("impossible de charger le dico: %v\n", err)
		return nil
	}
	h = hangman.New(8, hangman.PickWord())
	hangman.DrawWelcome(s, m)
	size := strconv.Itoa(len(h.Letters))
	s.ChannelMessageSend(m.ChannelID, "Mot de "+size+" lettres à trouver\nQuel est votre lettre ?\n\n **!h <lettre>** pour envoyer une lettre")
	return h
}

// GHangman : Function where Hangman compare try of the user with the actual word to guess
func GHangman(s *discordgo.Session, g *discordgo.Guild, user *discordgo.User, m *discordgo.MessageCreate, guess string) {
	fmt.Println("START : GHangman function - from hangman command")
	if end == 0 && h != nil {
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			s.ChannelMessageSend(m.ChannelID, "Lettre "+guess+" non valide")
			return
		}
		println("Lettre " + guess)
		h.MakeAGuess(guess)
		hangman.Draw(s, m, h, guess)
		switch h.State {
		case "won", "lost":
			end = 1
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Quel est votre lettre ? ")
	} else {
		s.ChannelMessageSend(m.ChannelID, "Veuillez relancer un nouveau jeu en tapant **!hangman** ou **!h**")
	}
}
