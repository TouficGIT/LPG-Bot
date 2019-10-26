package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
)

// GetStats : Function used for retrieve rocket stats of a player
func GetStats(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	fmt.Println("START : GetStats function - from rocketleague command")
	var steamid string
	var steamID []SteamID

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

	// If the user didn't give his ID, then ask him to send it via !steam command
	// Otherwise, set his ID into SteamID variable
	for i := 0; i < len(steamID); i++ {
		if m.Author.Username == steamID[i].Username {
			steamid = steamID[i].ID
		} else {
			fmt.Println("ERROR : SteamID not provided")
			s.ChannelMessageSend(m.ChannelID, "Il me faut ton **SteamID** @"+m.Author.Username+" pour que je trouve tes stats")
			s.ChannelMessageSend(m.ChannelID, "Commande **!steam + <ton ID>** pour l'ajouter")
			return nil
		}
	}

	if len(steamID) == 0 {
		fmt.Println("INFO : Steam ID json file is empty")
		s.ChannelMessageSend(m.ChannelID, "Il me faut ton **SteamID** @"+m.Author.Username+" pour que je trouve tes stats")
		s.ChannelMessageSend(m.ChannelID, "Commande **!steam + <ton ID>** pour l'ajouter")
		return nil
	}

	url := "https://rocketleague.tracker.network/profile/steam/" + steamid

	fmt.Println("INFO : Getting html content from : " + url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("ERROR : code d'erreur HTML ")
		fmt.Print(res.StatusCode)
	}

	// Load the HTML document
	fmt.Println("INFO : Load the HTML content into the scraper")
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	statsArray := doc.Find(".stat")
	secondStat := statsArray.Eq(2)
	thirdStat := statsArray.Eq(2)
	fourthStat := statsArray.Eq(3)
	fifthStat := statsArray.Eq(4)
	sixthStat := statsArray.Eq(5)
	eigthStat := statsArray.Eq(7)

	ratioGoalShots := strings.TrimSpace(secondStat.Find(".value").Text())
	wins := strings.TrimSpace(thirdStat.Find(".value").Text())
	goals := strings.TrimSpace(fourthStat.Find(".value").Text())
	saves := strings.TrimSpace(fifthStat.Find(".value").Text())
	shots := strings.TrimSpace(sixthStat.Find(".value").Text())
	assists := strings.TrimSpace(eigthStat.Find(".value").Text())

	rankSeasonImgGb := doc.Find("img")
	rankSeasonImg := rankSeasonImgGb.Eq(14).AttrOr(`src`, ``)
	rankSeasonTitleTd := doc.Find("td")
	rankSeasonTitle := rankSeasonTitleTd.Eq(16).Text()

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00, // Green
		Description: rankSeasonTitle,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Victoires",
				Value:  wins,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Buts",
				Value:  goals,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Tirs",
				Value:  shots,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Saves",
				Value:  saves,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Assistes",
				Value:  assists,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Ratio Buts/Tirs",
				Value:  ratioGoalShots,
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://rocketleague.tracker.network/" + rankSeasonImg,
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "Voici tes stats **RocketLeague** " + m.Author.Username,
	}

	fmt.Println("INFO : Retrieve informations to the discord channel")
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		fmt.Println(msg, err)
	}

	return err
}
