package bot

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"

	"Work.go/LPG-Bot/LPGBot/print"
	"github.com/bwmarrin/discordgo"
)

type chuckJoke struct {
	Categorie string `json:"categories"`
	CreatedAt string `json:"created_at"`
	Icon      string `json:"icon_url"`
	ID        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
	Fact      string `json:"value"`
}

// ChuckFact : Fetch Chuck Norris Joke
func ChuckFact(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	print.DebugLog("[DEBUG] Start ChuckFact function - from chuck command", "[SERVER]")
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		print.CheckError("[ERROR] Could not fetch joke", "[Server]", err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print.CheckError("[ERROR] Unknown response body", "[Server]", err)
		return err
	}

	var joke chuckJoke
	var fact string

	print.DebugLog("[DEBUG] Unmarshal the chuck norris quote", "[SERVER]")
	json.Unmarshal(body, &joke)

	//UnescapeString used for accented characters (like "é")
	fact = html.UnescapeString(joke.Fact)

	msg, err := s.ChannelMessageSend(m.ChannelID, fact)
	if err != nil {
		fmt.Println(msg, err)
	}

	return err
}
