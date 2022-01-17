package bot

import (
	"math/rand"
	"strconv"

	"Work.go/LPG-Bot/LPGBot/print"
)

// Roll func : it used to roll a dice and return the result to the channel
func Roll() string {
	print.DebugLog("[DEBUG] Start Roll function", "[SERVER]")
	random := rand.Intn(6) + 1
	return "Dice shows **" + strconv.Itoa(random) + "**"
}
