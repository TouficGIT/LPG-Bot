package bot

import (
	"fmt"
	"strings"

	"Work.go/LPG-Bot/LPGBot/bot/hangman"
	"github.com/bwmarrin/discordgo"
)

var h *hangman.Game

// Hangman : it used to launch hangman game
func NewHangman(s *discordgo.Session, g *discordgo.Guild, user *discordgo.User, m *discordgo.MessageCreate) *hangman.Game {

	err := hangman.Load("bot/hangman/words.txt")
	if err != nil {
		fmt.Printf("impossible de charger le dico: %v\n", err)
		return nil
	}
	h := hangman.New(8, hangman.PickWord())
	hangman.DrawWelcome(s, m)
	s.ChannelMessageSend(m.ChannelID, "Quel est votre lettre ? ")
	return h
}

func GHangman(s *discordgo.Session, g *discordgo.Guild, user *discordgo.User, m *discordgo.MessageCreate, guess string) {
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
		return
	}
	s.ChannelMessageSend(m.ChannelID, "Quel est votre lettre ? ")
}
