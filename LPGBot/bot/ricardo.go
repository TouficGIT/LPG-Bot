package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// User is a ricardo game user
type User struct {
	Username string `json:"username"`
	Points   int    `json:"points"`
	Badge    string `json:"badge"`
}

// Ricardo func : it used to register a player to ricardo game
// if the player is not already registered
func Ricardo(msgUser string, msg string) (string, error) {
	var u []User
	parts := strings.Split(strings.ToLower(msg), " ")

	msgUser = strings.ToLower(msgUser)
	// Open our jsonFile
	rgFile, err := os.Open("bot/ricardo/ricardoGame.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer rgFile.Close()
	// read content of the file
	ct, err := ioutil.ReadAll(rgFile)
	if err != nil {
		fmt.Println("Unknown response body")
	}
	json.Unmarshal(ct, &u)
	for i := 0; i < len(u); i++ {
		if msgUser == u[i].Username {
			if parts[1] == "stat" {
				return "**Ricardo Game:**\n\n" + msgUser + "Votre score actuel -> " + strconv.Itoa(u[i].Points) + "\n" + u[i].Badge, nil
			}
			return "Vous participez déjà au ricardo game " + msgUser, nil
		}
	}
	// Adding the new player
	data := append(u, User{Username: msgUser, Points: 0, Badge: ""})
	// Marshal the new user to the json file
	add, _ := json.Marshal(data)
	_ = ioutil.WriteFile("bot/ricardo/ricardoGame.json", add, 0644)
	return "Bienvenue au **Ricardo Game** " + msgUser + " !\n https://github.com/TouficGIT/LPG-bot/tree/master/LPGBot/bot/ricardo/dance.gif", nil
}

// RicardoGame func : its the main RicardoGame function
func RicardoGame(user string) {
	var u []User
	user = strings.ToLower(user)
	// Open our jsonFile
	rgFile, err := os.Open("bot/ricardo/ricardoGame.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer rgFile.Close()
	// read content of the file
	ct, err := ioutil.ReadAll(rgFile)
	if err != nil {
		fmt.Println("Unknown response body")
	}
	json.Unmarshal(ct, &u)
	for i := 0; i < len(u); i++ {
		if user == u[i].Username {
			u[i].Points++
			switch u[i].Points {
			case 10:
				u[i].Badge = "https://github.com/TouficGIT/LPG-bot/tree/master/LPGBot/bot/ricardo/saiyan.gif"
			case 30:
				u[i].Badge = "https://github.com/TouficGIT/LPG-bot/tree/master/LPGBot/bot/ricardo/saiyan2.gif"
			case 60:
				u[i].Badge = "https://github.com/TouficGIT/LPG-bot/tree/master/LPGBot/bot/ricardo/saiyan3.gif"
			case 110:
				u[i].Badge = "https://github.com/TouficGIT/LPG-bot/tree/master/LPGBot/bot/ricardo/saiyanGod.gif"
			case 200:
				u[i].Badge = "https://github.com/TouficGIT/LPG-bot/tree/master/LPGBot/bot/ricardo/saiyanBlue.gif"
			case 300:
				u[i].Badge = "https://github.com/TouficGIT/LPG-bot/tree/master/LPGBot/bot/ricardo/saiyanBlueFused.gif"
			case 600:
				u[i].Badge = "https://github.com/TouficGIT/LPG-bot/tree/master/LPGBot/bot/ricardo/saiyanUltra.gif"
			}
		}
	}
	chg, _ := json.Marshal(u)
	_ = ioutil.WriteFile("bot/ricardo/ricardoGame.json", chg, 0644)
}
