package bot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"Work.go/LPG-Bot/LPGBot/print"
	"github.com/bwmarrin/discordgo"
)

type RLPlaylist struct {
	Data struct {
		Segments []struct {
			Type     string `json:"type"`
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			StatsPlaylist struct {
				Tier struct {
					Metadata struct {
						IconURL string `json:"iconUrl"`
						Name    string `json:"name"`
					} `json:"metadata"`
				} `json:"tier"`
				Division struct {
					Metadata struct {
						Name string `json:"name"`
					} `json:"metadata"`
				} `json:"division"`
			} `json:"stats,omitempty"`
		} `json:"segments"`
	} `json:"data"`
}

// GetRLStat : Function used for retrieve rocket stats of a player
func GetRLStat(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (err error) {
	print.DebugLog("[DEBUG] GetRLStats function - from rl command", "[SERVER]")

	var (
		url, username                                               string
		RLPlaylist                                                  RLPlaylist
		image2v2, rank1v1, rank2v2, rank3v3, div1v1, div2v2, div3v3 string
	)

	// Check if Rl command is used with more than 1 argument
	if len(args) < 2 {
		privChan, err := s.UserChannelCreate(m.Author.ID)
		print.CheckError("[ERROR] Could not create channel between bot and user", "[SERVER]", err)
		s.MessageReactionAdd(m.ChannelID, m.ID, "❓")
		s.ChannelMessageSendEmbed(privChan.ID, print.EmbedErrorRL)
		return nil
	}

	// Set up API url and Username passed by user's command
	username = args[1]
	url = "https://api.tracker.gg/api/v2/rocket-league/standard/profile/epic/" + username

	// Contact URL api
	print.DebugLog("[DEBUG] Contact URL api", "[SERVER]")
	resp, err := http.Get(url)
	if err != nil {
		print.CheckError("[ERROR] Could not contact RL api", "[SERVER]", err)
		return err
	}
	// Read the response from the URL
	print.DebugLog("[DEBUG] Read the response from the URL", "[SERVER]")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print.CheckError("[ERROR] Unable to read response from RL api", "[SERVER]", err)
		return err
	}

	// Unmarshal JSON response for PlaylistStats
	print.DebugLog("[DeBUG] Unmarshal RLPlaylist", "[SERVER]")
	json.Unmarshal(body, &RLPlaylist)

	// Set values to each playlist variables, from the JSON response
	print.DebugLog("[DEBUG] Set values to each playlist variables", "[SERVER]")
	for i := 0; i < len(RLPlaylist.Data.Segments); i++ {
		if RLPlaylist.Data.Segments[i].Type == "playlist" {
			if RLPlaylist.Data.Segments[i].Metadata.Name == "Ranked Doubles 2v2" {
				image2v2 = RLPlaylist.Data.Segments[i].StatsPlaylist.Tier.Metadata.IconURL
				rank2v2 = RLPlaylist.Data.Segments[i].StatsPlaylist.Tier.Metadata.Name
				div2v2 = RLPlaylist.Data.Segments[i].StatsPlaylist.Division.Metadata.Name
			} else if RLPlaylist.Data.Segments[i].Metadata.Name == "Ranked Standard 3v3" {
				rank3v3 = RLPlaylist.Data.Segments[i].StatsPlaylist.Tier.Metadata.Name
				div3v3 = RLPlaylist.Data.Segments[i].StatsPlaylist.Division.Metadata.Name
			} else if RLPlaylist.Data.Segments[i].Metadata.Name == "Ranked Duel 1v1" {
				rank1v1 = RLPlaylist.Data.Segments[i].StatsPlaylist.Tier.Metadata.Name
				div1v1 = RLPlaylist.Data.Segments[i].StatsPlaylist.Division.Metadata.Name
			}
		}
	}

	// Create EmbedMessage with the values collected from the response
	embedRLStat := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x3D85C6, // Green 00ff00
		Description: "**Doubles 2v2: ** " + rank2v2 + " " + div2v2,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Standard 3v3 rank",
				Value:  rank3v3 + " " + div3v3,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Duel 1v1 rank",
				Value:  rank1v1 + " " + div1v1,
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: image2v2,
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "Here's the **RocketLeague** stats of: **" + username + "**",
	}

	print.DebugLog("[DEBUG] Retrieve informations to the discord channel", "[SERVER]")
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embedRLStat)
	print.CheckError("[ERROR] While sending EmbedMsg", "[SERVER]", err)

	return err
}

/*

// Add overview stats? (wins, shots ...)

	// Set values to each stats variables, from the JSON response
	print.InfoLog("[INFO] Set values to each stats variables", "[SERVER]")
	for i := 0; i < len(RLStats.Segments); i++ {
		if RLStats.Segments[i].Type == "overview" {
			wins = RLStats.Segments[i].Stats.WinsStat.StatValue
			goals = RLStats.Segments[i].Stats.GoalsStat.StatValue
			saves = RLStats.Segments[i].Stats.SavesStat.StatValue
			assists = RLStats.Segments[i].Stats.AssistsStat.StatValue
			shots = RLStats.Segments[i].Stats.ShotsStat.StatValue


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
}
*/

/*
	StatsOverview struct {
		Wins struct {
			DisplayValue string `json:"displayValue"`
		} `json:"wins"`
		Goals struct {
			DisplayValue string `json:"displayValue"`
		} `json:"goals"`
		MVPs struct {
			DisplayValue string `json:"displayValue"`
		} `json:"mVPs"`
		Saves struct {
			DisplayValue string `json:"displayValue"`
		} `json:"saves"`
		Assists struct {
			DisplayValue string `json:"displayValue"`
		} `json:"assists"`
		Shots struct {
			DisplayValue string `json:"displayValue"`
		} `json:"shots"`
		Score struct {
			DisplayValue string `json:"displayValue"`
		} `json:"score"`
	} `json:"stats,omitempty"`
*/
