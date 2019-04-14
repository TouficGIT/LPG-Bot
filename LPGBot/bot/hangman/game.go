package hangman

import "strings"

// Game : structure de notre jeu
type Game struct {
	State        string   // Etat du jeu: perdu ou gagné
	Letters      []string // Les lettres a trouver dans le mot
	FoundLetters []string // Bonnes propositions de lettres
	UsedLetters  []string // Les lettres utilisés
	TurnsLeft    int      // Nombres d'essais restants
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "") // Split: permet de séparer lettres - renvoi un slice
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.BadTry(guess)
		if hasLost(g.TurnsLeft) {
			g.State = "lost"
		}
	}
}

func (g *Game) BadTry(guess string) {
	g.State = "badGuess"
	g.UsedLetters = append(g.UsedLetters, guess)
	g.TurnsLeft--
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func hasLost(turnsLeft int) bool {
	if turnsLeft <= 0 {
		return true
	}
	return false
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
