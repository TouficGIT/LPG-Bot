package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// User is a ricardo game user
type User struct {
	Username string `json:"username"`
	Points   int    `json:"points"`
	Badge    string `json:"badge"`
}

var rIntro = `
	 ___ ___ ___   _   ___ ___   ___     ___                  _ 
	| _ \_ _/ __| /_\ | _ \   \ / _ \   / __|__ _ _ __  ___  | |
	|   /| | (__ / _ \|   / |) | (_) | | (_ / _  | '  \/ -_) |_|
	|_|_\___\___/_/ \_\_|_\___/ \___/   \___\__,_|_|_|_\___| (_)`

var rRules = `
** Les règles du jeu:**
	- Postez des messages sur le serveur LPG pour gagner des points
	- Remportez de nouveaux badges ricardo !
	- Les badges vous permettent d'obtenir un nouveau rang

	- !ricardo stat: pour afficher votre score et votre badge

Have fun !
`

// Ricardo func : it used to register a player to ricardo game
// if the player is not already registered, else it return a msg
func Ricardo(msgUser string, msg string) (string, error) {
	fmt.Println("START : Ricardo function - from ricardo command")
	var u []User
	msgUser = strings.ToLower(msgUser)
	parts := strings.Split(strings.ToLower(msg), " ")

	// TODO: Refact the code below into a "Open file" function

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
			if len(parts) > 1 && parts[1] == "stat" {
				return "```" + rIntro + "```\n\nUsername: " + strings.ToUpper(msgUser) + "\nVotre score actuel: " + strconv.Itoa(u[i].Points) + "\nBadge: " + u[i].Badge, nil
			}
			return "Vous participez déjà au ricardo game " + strings.ToUpper(msgUser) + " ! 🎮\n```!ricardo stat: pour obtenir vos statistiques```\n", nil
		}
	}

	// TODO: Refact the code below in a "write file" function, and find a solution to lock the file
	// to avoid the case of 2 players trying to play in same time -> and their names are finally not
	// write into the json file.

	// Adding the new player
	data := append(u, User{Username: msgUser, Points: 0, Badge: ""})
	// Marshal the new user to the json file
	add, err := json.Marshal(data)
	if err != nil {
		println(err)
		return "", nil
	}
	err = ioutil.WriteFile("bot/ricardo/ricardoGame.json", add, 0644)
	if err != nil {
		println(err)
		return "", nil
	}
	return "**Bienvenue " + strings.ToUpper(msgUser) + " au :** \n\n```" + rIntro + "```\n```" + rRules + "```\n https://tenor.com/view/ricardo-milos-based-god-ricardo-milos-gif-13369215", nil
}

// RicardoGame func : its the main RicardoGame function
func RicardoGame(s *discordgo.Session, g *discordgo.Guild, user *discordgo.User) (string, error) {
	var u []User
	var newTag string
	uName := strings.ToLower(user.Username)

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
		if uName == u[i].Username {
			u[i].Points++

			switch u[i].Points {
			case 10:
				u[i].Badge = "https://tenor.com/view/ricardo-super-saiyan-smile-gif-13677081"
				newTag = "```" + rIntro + "```\n\nFélicitation " + strings.ToUpper(u[i].Username) + " !\nTon nouveau rôle: " + "\nTu obtiens le badge: \n" + u[i].Badge
			case 30:
				u[i].Badge = "https://tenor.com/view/ricardo-super-saiyan2-smile-gif-13677092"
				newTag = "```" + rIntro + "```\n\nFélicitation " + strings.ToUpper(u[i].Username) + " !\nTu obtiens le badge \n" + u[i].Badge
			case 60:
				u[i].Badge = "https://tenor.com/view/ricardo-super-saiyan3-flex-gif-13677095"
				newTag = "```" + rIntro + "```\n\nFélicitation " + strings.ToUpper(u[i].Username) + " !\nTu obtiens le badge \n" + u[i].Badge
			case 110:
				u[i].Badge = "https://tenor.com/view/ricardo-super-saiyan-god-gif-13677088"
				newTag = "```" + rIntro + "```\n\nFélicitation " + strings.ToUpper(u[i].Username) + " !\nTu obtiens le badge \n" + u[i].Badge
			case 200:
				u[i].Badge = "https://tenor.com/view/ricardo-super-saiyan-blue-naked-gif-13677086"
				newTag = "```" + rIntro + "```\n\nFélicitation " + strings.ToUpper(u[i].Username) + " !\nTu obtiens le badge \n" + u[i].Badge
			case 300:
				u[i].Badge = "https://tenor.com/view/ricardo-fused-super-saiyan-blue-gif-13677091"
				newTag = "```" + rIntro + "```\n\nFélicitation " + strings.ToUpper(u[i].Username) + " !\nTu obtiens le badge \n" + u[i].Badge
			case 600:
				u[i].Badge = "https://tenor.com/view/ricardo-ultra-instinct-sexy-dancing-gif-13677084"
				newTag = "```" + rIntro + "```\n\nFélicitation " + strings.ToUpper(u[i].Username) + " !\nTu obtiens le badge \n" + u[i].Badge
			}
		}
	}

	chg, err := json.Marshal(u)
	if err != nil {
		println(err)
		return "", nil
	}
	err = ioutil.WriteFile("bot/ricardo/ricardoGame.json", chg, 0644)
	if err != nil {
		println(err)
		return "", nil
	}

	return newTag, nil
}
