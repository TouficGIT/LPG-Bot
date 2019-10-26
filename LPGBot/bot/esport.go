package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type conEsport struct {
	Status  string `json:"success"`
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Time    string `json:"expires_in"`
}

type tournamentList struct {
	Count    int      `json:"count"`
	PageNext string   `json:"next"`
	Results  []result `json:"results"`
}

type result struct {
	EventName string `json:"full_name"`
	Title     title  `json:"title"`
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
}

type title struct {
	Slug string `json:"slug"`
	Name string `json:"abbreviation"`
}

// ConnectEsport : Function used for connection to the esport api
func ConnectEsport() (string, error) {

	param := strings.NewReader(`username=meffe.emilien@gmail.com&password=Miloudu47`)
	req, err := http.NewRequest("POST", "https://api.esportsdirectory.info/v1/auth/", param)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unknown response body")
		return "", err
	}

	var connect conEsport
	json.Unmarshal(body, &connect)

	return connect.Token, err
}

// GetTournament : Function used for retrieve list of esport tournament in coming.
func GetTournament(s *discordgo.Session, token string, msg *discordgo.MessageCreate) (err error) {

	req, err := http.NewRequest("GET", "https://api.esportsdirectory.info/v1/tournament/", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unknown response body")
		return err
	}

	var trnList tournamentList
	json.Unmarshal(body, &trnList)

	parts := strings.Split(strings.ToLower(msg.Content), " ")

	// Get the name of the game choosed by the user
	var gameExist bool
	if len(parts) > 1 {
		for _, g := range trnList.Results {
			if parts[1] == g.Title.Slug {
				gameExist = true
				fmt.Println("Request tournaments date events for : " + g.Title.Name)
				tournament := "Game : " + g.Title.Name + "\nTournois : " + g.EventName + "\nCommence le : " + g.DateStart + "\nFinis le : " + g.DateEnd + ""
				_, _ = s.ChannelMessageSend(msg.ChannelID, tournament)
			}
		}

		if gameExist == false {
			_, _ = s.ChannelMessageSend(msg.ChannelID, "Je ne connais pas ce jeu, dsl")
			fmt.Println("Error : game choosed doesn't exist")
			return err
		}
	}

	return nil
}
