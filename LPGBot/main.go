package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"Work.go/LPG-Bot/LPGBot/bot"
	"Work.go/LPG-Bot/LPGBot/config"
	"github.com/bwmarrin/discordgo"
)

var helpMsg = `Salut, je suis **LPG Bot** ! Je répond aux commandes suivantes :

- **!hello** ou **!hi** : pour me dire bonjour et je te répondrai
- **!poll "Question" "Choix"** : pour créer un strawpoll
- **!steam <SteamID>** : pour renseigner ton steam ID
- **!rl** : pour afficher ton rank sur Rocket League (en 2v2)

- **!sd** ou **!dit** + **<son>** : pour jouer l'un des sons suivants
	-> boi / bruh / daniel / deja / fuck / mgs / nani / nice / ooh ...
	-> oui / ricardo / spooky / thug et wow

- **!ricardo** : pour participer au RicardoGame !
- **!hangman** ou **!h** : pour jouer au pendu avec LPG Bot
- **!chuck** : pour balancer une fact sur chuck norris
- **!météo** ou **!mt** + **<ville>** : pour obtenir la météo sur cette ville

- **!flip** : pour jouer à pile ou face
- **!roll** : pour lancer un dé

- **!help** ou **!lpg**: pour afficher ce message d'aide ^^ `

func main() {

	// Read the config file
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Load all lpg sounds into a buffer
	fmt.Println("")
	fmt.Println("INFO : -- Loading all lpg sounds into a buffer --")
	fmt.Println("")
	bot.LPGSOUND.LoadAll()

	// creation of lpgBot
	lpgBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lpgBot.AddHandler(messageHandler)
	lpgBot.AddHandler(newUserHandler)
	err = lpgBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("")
	fmt.Println("LPG Bot is connected !")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Current time : for logs
	cTime := time.Now().Format("01-02-2006 15:04:05")

	// Set LPG Bot playing at !help
	s.UpdateStatus(0, "Besoins d'aide ? !help")

	// Open logs file
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	// Get Channel Id where message has been post
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get the guild (server)
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get the user (server)
	user, err := s.User(m.Author.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Update points of the player + display his new rank if he get a new one
	rank, _ := bot.RicardoGame(s, g, user)
	if len(rank) != 0 {
		_, _ = s.ChannelMessageSend(m.ChannelID, rank)
	}

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.Bot || m.Author.ID == s.State.User.ID || len(m.Content) <= 0 || m.Content[0] != '!' {
			return
		}

		// Write into the logs
		_, err = f.WriteString("Time: " + cTime + " || Message: " + m.Content + " || From: " + m.Author.Username + "\n")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Time: %v || Message: %+v || From: %s\n", cTime, m.Content, m.Author)

		// Split the message content
		parts := strings.Split(strings.ToLower(m.Content), " ")

		// Parsing of the content message in order to launch the proper command
		switch parts[0] {
		case "!hello", "!salut", "!hi":
			_, _ = s.ChannelMessageSend(m.ChannelID, "Salut "+m.Author.Username+" !")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤙")
		case "!chuck":
			fact, err := bot.ChuckFact()
			if err != nil {
				fmt.Println("Time: " + cTime + " || ERROR : chuckfact function error")
				fmt.Println(err)
			}
			_, _ = s.ChannelMessageSend(m.ChannelID, fact)
		case "!ricardo":
			ric, _ := bot.Ricardo(m.Author.Username, m.Content)
			_, _ = s.ChannelMessageSend(m.ChannelID, ric)
		case "!hangman", "!h":
			if len(parts) < 2 {
				_ = bot.NewHangman(s, g, user, m)
			} else {
				bot.GHangman(s, g, user, m, parts[1])
			}
		case "!poll":
			err = bot.CreatePoll(s, m)
			if err != nil {
				fmt.Println("Time: " + cTime + " || ERROR : during poll function process")
				fmt.Println(err)
			}
		case "!steam":
			err = bot.AddSteamID(s, m)
			if err != nil {
				fmt.Println("Time: " + cTime + " || ERROR : during RL Getstat function process")
				fmt.Println(err)
			}
		case "!rl":
			err = bot.GetStats(s, m)
			if err != nil {
				fmt.Println("Time: " + cTime + " || ERROR : during RL Getstat function process")
				fmt.Println(err)
			}
		case "!sd", "!sound", "!dit":
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					fmt.Println("Time: " + cTime + " || User :" + user.Username + " ||  Playsound command")
					err = bot.PlaySound(s, g.ID, vs.ChannelID, m.Content)
					if err != nil {
						fmt.Println("Error playing sound:", err)
						f.WriteString("Time: " + cTime + " || Error playing " + parts[1] + " for " + m.Author.Username)
					}
					return
				}
			}
		case "!prank":

		case "!wt", "!weather", "!meteo", "!météo", "!mt":
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🌞")
			weather, _ := bot.Weather(parts[1])
			_, _ = s.ChannelMessageSend(m.ChannelID, weather)
		case "!flip", "!fp":
			coin, _ := bot.FlipCoin()
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🎰")
			_, _ = s.ChannelMessageSend(m.ChannelID, coin)
		case "!roll":
			dice := bot.Roll()
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🎲")
			_, _ = s.ChannelMessageSend(m.ChannelID, dice)
		case "!help", "!lpg":
			// create a private messaging channel between the bot and the user
			privChan, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				println("ERROR : -" + parts[0] + "- Could not create channel between bot and user. ")
				return
			}
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🧙")
			_, _ = s.ChannelMessageSend(privChan.ID, helpMsg)
		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Je n'ai pas compris ton message "+m.Author.Username+"  ¯\\_(ツ)_/¯")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤔")
		}
		fmt.Println("")
	}
}

func newUserHandler(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	cTime := time.Now().Format("01-02-2006 15:04:05")
	fmt.Println("INFO : New user - guild member add event triggered")
	fmt.Println("Joined at : " + cTime + "\n User ID : " + event.User.Username)

	guild, err := s.Guild(event.GuildID)
	if err != nil {
		println("ERROR : Could not retrieve guild object from identifier. ")
		return
	}

	// create a private messaging channel between the bot and the new guild member
	privChan, err := s.UserChannelCreate(event.User.ID)
	if err != nil {
		println("ERROR : Could not create channel between bot and user. ")
		return
	}

	// send greet message to new guild member
	s.ChannelMessageSend(privChan.ID, "Salut "+event.User.Username+" ! \n\n**Bienvenu sur le serveur "+guild.Name+" ! :video_game:**\n\n Je suis le bot de ce serveur et avant de démarrer ton aventure ici, je vais juste te demander de jeter un petit coup d'oeil au channel des <#434444368418570243>.\n\n Voilà, c'est tout. N'hésites pas à venir nous saluer sur <#373160766670503957> \n\n Enjoy ! :call_me:\n Toufic & Saya")
}

// Commands deleted - copy/past them into 'switch' to bring them back to life
/*
   case "!esport":
   token, err := bot.ConnectEsport()
   if err != nil {
   	fmt.Println("Time: " + cTime + " || ERROR: Command esport")
   	fmt.Println(err)
   }
   err = bot.GetTournament(s, token, m)
   if err != nil {
   	fmt.Println("Time: " + cTime + " || ERROR: Command esport - GetTournament")
   	fmt.Println(err)
   }
*/
