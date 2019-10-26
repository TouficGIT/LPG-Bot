package bot

import (
	"fmt"
	"math/rand"
)

// FlipCoin func: it's used to return the result of a coin flip
func FlipCoin() (string, error) {
	fmt.Println("START : FlipCoin function")
	random := rand.Intn(2)
	switch random {
	case 0:
		return "La pièce tombe sur **Pile**", nil
	case 1:
		return "La pièce tombe sur **Face**", nil
	}
	return "", nil
}
