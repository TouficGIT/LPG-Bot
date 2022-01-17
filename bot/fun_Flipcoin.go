package bot

import (
	"math/rand"

	"Work.go/LPG-Bot/LPGBot/print"
)

// FlipCoin func: it's used to return the result of a coin flip
func FlipCoin() (string, error) {
	print.DebugLog("[DEBUG] Start FlipCoin function", "[SERVER]")
	random := rand.Intn(2)
	switch random {
	case 0:
		return "The coin falls on **Tails**", nil
	case 1:
		return "The coin falls on **Heads**", nil
	}
	return "", nil
}
