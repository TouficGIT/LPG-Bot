package bot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"Work.go/LPG-Bot/LPGBot/print"
	"github.com/bwmarrin/discordgo"
)

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

type Ranked []struct {
	LeagueID     string `json:"leagueId"`
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	SummonerID   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

type Mastery []struct {
	ChampionID                   int    `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int    `json:"tokensEarned"`
	SummonerID                   string `json:"summonerId"`
}

type Champion struct {
	Champions []struct {
		ChampionID int    `json:"championId"`
		Name       string `json:"name"`
		Full       string `json:"full"`
	} `json:"champions"`
}

// GetLolStat: Function used for retrieve lol stats of a player
func GetLolStat(s *discordgo.Session, m *discordgo.MessageCreate, args []string, apiKey string) (err error) {
	print.DebugLog("[DEBUG] GetLolStat function - from lol command", "[SERVER]")

	var (
		championID, championPoints  int
		url, username, summonerID   string
		championName                string
		tierUser, rankUser, rankImg string
		ranked                      Ranked
		summoner                    Summoner
		mastery                     Mastery
		champion                    Champion
	)

	// Check numbers of arguments from Lol command
	if len(args) < 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Not enough arguments. I need your Username")
		return nil
	}

	if len(args) >= 3 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Too much arguments. I just need your username")
		return nil
	}

	// Get summoner username, from args
	username = args[1]

	// Get summoner ID from summoners API
	url = "https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + username + "?api_key=" + apiKey
	body, err := GetByte(url)

	// Unmarshal JSON response for Summoner info
	print.DebugLog("[DEBUG] Unmarshal Summoner info", "[SERVER]")
	if err != nil {
		print.DebugLog("[DEBUG] Error, body response is empty", "[SERVER]")
		return
	} else {
		json.Unmarshal(body, &summoner)
	}

	summonerID = summoner.ID

	if summonerID == "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "No information found for summoner: "+username)
		return nil
	}

	// Get Rank info with summoner ID
	url = "https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/" + summonerID + "?api_key=" + apiKey
	body, err = GetByte(url)

	if err != nil {
		print.DebugLog("[DEBUG] Error, body response is empty", "[SERVER]")
		return
	} else {
		json.Unmarshal(body, &ranked)
	}

	if ranked.IsStructureEmpty() {
		rankUser = "You have no rank atm."
		rankImg = "https://github.com/TouficGIT/LPG-Bot/blob/main/_files/lol_rank/emblem_norank.png"
	} else {
		rankUser = ranked[0].Rank
		tierUser = ranked[0].Tier
		rankImg = "https://raw.githubusercontent.com/TouficGIT/LPG-Bot/main/_files/lol_rank/emblem_" + strings.ToLower(tierUser) + ".png"
	}

	// Get champion mastery with champion-masteries API
	url = "https://euw1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/" + summonerID + "?api_key=" + apiKey
	body, err = GetByte(url)

	if err != nil {
		print.DebugLog("[DEBUG] Error, body response is empty", "[SERVER]")
		return
	} else {
		json.Unmarshal(body, &mastery)
	}

	championID = mastery[0].ChampionID
	championPoints = mastery[0].ChampionPoints

	// Translate Champion ID to name with our JSON file
	jsonFile, err := os.Open("_files/lol_champions/championsIDs.json")
	if err != nil {
		print.CheckError("[ERROR] Could not open championsIDs.json", "[SERVER]", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &champion)

	for _, v := range champion.Champions {
		if v.ChampionID == championID {
			championName = v.Name
			//championImg = v.Full
		}
	}

	// Create EmbedMessage with the values collected from the responses
	embedRLStat := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xFDDC5C, // Green 00ff00
		Description: "**Rank 5v5: ** " + tierUser + " " + rankUser,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Summoner Level",
				Value:  strconv.Itoa(summoner.SummonerLevel),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Champion Mastery",
				Value:  "Name: " + championName + " / " + "Points: " + strconv.Itoa(championPoints),
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: rankImg,
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "Here's the **League of Legend** stats of: **" + username + "**",
	}

	print.DebugLog("[DEBUG] Retrieve informations to the discord channel", "[SERVER]")
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embedRLStat)
	print.CheckError("[ERROR] While sending EmbedMsg", "[SERVER]", err)

	return err
}

func (ranked Ranked) IsStructureEmpty() bool {
	return reflect.DeepEqual(ranked, Ranked{})
}

func GetByte(url string) (body []byte, err error) {
	// Contact URL api
	print.DebugLog("[DEBUG] Contact URL api", "[SERVER]")
	resp, err := http.Get(url)
	if err != nil {
		print.CheckError("[ERROR] Could not contact LOL api: "+url, "[SERVER]", err)
		return nil, err
	}
	// Read the response from the URL
	print.DebugLog("[DEBUG] Read the response from the URL", "[SERVER]")
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		print.CheckError("[ERROR] Unable to read response from LOL api", "[SERVER]", err)
		return nil, err
	}

	return body, nil
}
