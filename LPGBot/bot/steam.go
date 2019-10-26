package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// SteamID structure for getstats commands
type SteamID struct {
	Username string `json:"username"`
	ID       string `json:"steamID"`
}

// AddSteamID : func used to add the steam ID of a user into a json file
func AddSteamID(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	fmt.Println("START : addSteamID function - from steam command")
	var steamID []SteamID
	parts := strings.Split(strings.ToLower(m.Content), " ")

	if len(parts) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Commande **!steam + <ton ID>** pour ajouter ton ID steam "+m.Author.Username)
		return nil
	}

	// Open our jsonFile
	stFile, err := os.Open("bot/steam/steamIDs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer stFile.Close()
	// read content of the file
	ct, err := ioutil.ReadAll(stFile)
	if err != nil {
		fmt.Println("Unknown response body")
	}

	json.Unmarshal(ct, &steamID)
	for i := 0; i < len(steamID); i++ {
		if m.Author.Username == steamID[i].Username {
			fmt.Println("INFO : " + m.Author.Username + " have already added is SteamID")
			s.ChannelMessageSend(m.ChannelID, "Tu as déjà fournis ton steamID "+strings.ToUpper(m.Author.Username))
			return nil
		}
	}

	// Adding the steamID of the user
	data := append(steamID, SteamID{Username: m.Author.Username, ID: parts[1]})
	// Marshal the new user to the json file
	add, err := json.Marshal(data)
	if err != nil {
		println(err)
		return nil
	}
	err = ioutil.WriteFile("bot/steam/steamIDs.json", add, 0644)
	if err != nil {
		println(err)
		return nil
	}
	fmt.Println("INFO : Steam ID of : " + m.Author.Username + " added")
	s.ChannelMessageSend(m.ChannelID, "Ton **SteamID** a été ajouté "+m.Author.Username)
	return err
}
