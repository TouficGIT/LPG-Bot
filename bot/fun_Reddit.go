package bot

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"time"

	"Work.go/LPG-Bot/LPGBot/print"
	"github.com/bwmarrin/discordgo"
)

//MemeType : type Meme
type RedditType struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Meme : Fetch random meme from this API "https://meme-api.herokuapp.com/gimme"
func Reddit(s *discordgo.Session, m *discordgo.MessageCreate, subreddit string, definedChannelID string) (err error) {
	print.DebugLog("[DEBUG] Start Reddit function", "[SERVER]")
	resp, err := http.Get("https://meme-api.herokuapp.com/gimme/" + subreddit)
	if err != nil {
		print.CheckError("[ERROR] Could not contact URL", "[Server]", err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print.CheckError("[ERROR] Unknown response body", "[Server]", err)
		return err
	}

	var reddit RedditType
	var sendChannelID string
	print.DebugLog("[DEBUG] Unmarshal information", "[SERVER]")
	json.Unmarshal(body, &reddit)
	var title string
	var url string

	title = html.UnescapeString(reddit.Title)
	url = html.UnescapeString(reddit.URL)

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xFF9900, // Orange
		Description: "SubReddit: " + subreddit,
		Image: &discordgo.MessageEmbedImage{
			URL: url,
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     title,
	}

	print.DebugLog("[DEBUG] Retrieve informations to the discord channel", "[SERVER]")

	if definedChannelID != "" {
		sendChannelID = definedChannelID
	} else {
		sendChannelID = m.ChannelID
	}

	msg, err := s.ChannelMessageSendEmbed(sendChannelID, embed)
	if err != nil {
		fmt.Println(msg, err)
	}

	return err
}
