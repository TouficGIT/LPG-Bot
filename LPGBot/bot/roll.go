package bot

import (
	"math/rand"
	"strconv"
)

// Roll func : it used to roll a dice and return the result to the channel
func Roll() string {
	random := rand.Intn(6) + 1
	return "Le dé affiche " + strconv.Itoa(random)
}
