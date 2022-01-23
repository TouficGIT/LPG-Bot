package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"Work.go/LPG-Bot/LPGBot/admin"
	"Work.go/LPG-Bot/LPGBot/bot"
	"Work.go/LPG-Bot/LPGBot/config"
	"Work.go/LPG-Bot/LPGBot/logs"
	"Work.go/LPG-Bot/LPGBot/print"
	"github.com/bwmarrin/discordgo"
)

var (
	// WaitNextCommand used to avoid peoples overusing sound cmd
	WaitNextCommand = 0

	// Current time : for logs
	cTime = time.Now().Format("01-02-2006 15:04:05")

	// var bot discord
	lpgBot        *discordgo.Session
	conf          *config.ConfigStruct
	lpgbotUser    *discordgo.User
	nsfwChannelID string

	//var generic
	PREFIX string
	err    error
)

func main() {
	fmt.Println("")
	print.InfoLog("[START] LPG Bot is starting ...", "[SERVER]")

	// -- CONFIG lpgbot --

	// Read the config file
	conf = config.ReadConfig()
	if conf == nil {
		print.CheckError("[ERROR] Could not read config file", "[SERVER]", err)
		return
	}

	// Load all lpg sounds into a buffer
	print.InfoLog("[INFO] Loading all lpg sounds into a buffer", "[SERVER]")
	bot.LPGSOUND.LoadAll()

	// Creation of lpgBot
	lpgBot, err = discordgo.New("Bot " + conf.Token)
	if err != nil {
		print.CheckError("[ERROR] Creation of lpgbot", "[SERVER]", err)
		return
	}

	// Create logs file if doesn't exist
	err = logs.CheckAndCreate()
	if err != nil {
		print.CheckError("[ERROR] Check or create Logs", "[SERVER]", err)
		return
	}

	//msg handler
	PREFIX = conf.BotPrefix
	lpgBot.AddHandler(ready)
	lpgBot.AddHandler(messageHandler)
	lpgBot.AddHandler(newUserHandler)

	// -- START lpgbot --

	// Identify is sent during initial handshake with the discord gateway.
	// It is now required to retrieve events from Discord servers
	lpgBot.Identify.Token = conf.Token
	lpgBot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	err = lpgBot.Open()
	if err != nil {
		print.CheckError("[ERROR] Opening connection", "[SERVER]", err)
		return
	}
	print.InfoLog("[CONNECTED] LPG Bot is connected !", "[SERVER]")

	// Set LPG Bot playing at !help
	lpgBot.UpdateGameStatus(0, "Need Help? !help")

	// Start listening commands from discord
	//make(chan struct{})
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

// Function to handle every message for the bot
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// -- GET INFOS --

	//Open log file
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		print.CheckError("[ERROR] Opening log file", "[SERVER]", err)
		return
	}
	defer f.Close()

	// Get Channel Id where message has been post
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		print.CheckError("[ERROR] Get Channel Id", "[SERVER]", err)
		return
	}

	// Get the guild (server)
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		print.CheckError("[ERROR] Get Guild Id", "[SERVER]", err)
		return
	}

	// Get the user (server)
	user, err := s.User(m.Author.ID)
	if err != nil {
		print.CheckError("[ERROR] Get User Id", "[SERVER]", err)
		return
	}

	// Check if user is admin
	var isAdmin bool = false
	for _, role := range m.Member.Roles {
		if role == "325384210133024768" {
			isAdmin = true
		}
	}

	// If user is a bot, we don't need to read the message
	lpgbotUser, err = lpgBot.User("@me")
	if user.ID == lpgbotUser.ID {
		return
	}

	// -- COMMANDS --

	// Read message from User, and check if command exists
	if strings.HasPrefix(m.Content, PREFIX) {
		if m.Author.Bot || m.Author.ID == s.State.User.ID || len(m.Content) <= 0 || m.Content[0] != '!' {
			return
		}

		// Get command from the content
		content := m.Content[len(PREFIX):]
		args := strings.Fields(strings.ToLower(content))
		command := args[0]

		// Write into the logs
		print.InfoLog("[INFO] Command: "+command+" From: "+m.Author.Username+"", "[SERVER]")
		_, err = f.WriteString("Time: " + cTime + " || Message: " + command + " || From: " + m.Author.Username + "\n")
		print.CheckError("[ERROR] During WriteString logs", "[SERVER]", err)

		// -- [ADMIN] --

		if isAdmin == true {
			switch command {

			// Check if welcome message works fine
			case "welcome":
				admin.WelcomeTest(s, m, g)
				return

			case "setnsfwchannel":
				nsfwChannelID = admin.SetNSFWChannel(args)
				_, _ = s.ChannelMessageSend(m.ChannelID, "[ADMIN] Channel for nsfw is set")
				return

			case "debug":
				print.SetDebug(args[1])
				_, _ = s.ChannelMessageSend(m.ChannelID, "[ADMIN] Debug mode is "+args[1])
				return

			default:
			}
		}

		// -- [NOT ADMIN] --

		switch command {

		// [FUN]

		case "hello", "salut", "hi":
			_, _ = s.ChannelMessageSend(m.ChannelID, "Hi "+m.Author.Username+" !")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤙")
		case "flip", "fp":
			coin, _ := bot.FlipCoin()
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🎰")
			_, _ = s.ChannelMessageSend(m.ChannelID, coin)
		case "roll":
			dice := bot.Roll()
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🎲")
			_, _ = s.ChannelMessageSend(m.ChannelID, dice)
		case "chuck":
			err = bot.ChuckFact(s, m)
			print.CheckError("[ERROR] Chuckfact function error", user.Username, err)
		case "meme":
			err = bot.Reddit(s, m, "memes", "")
			print.CheckError("[ERROR] Meme function", user.Username, err)

		// [PLAY A SOUND]

		case "say":
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					if WaitNextCommand == 0 {
						WaitNextCommand = 1
						err = bot.PlaySound(s, g.ID, vs.ChannelID, m.Content, args)
						WaitNextCommand = 0
					}

					if err != nil {
						print.CheckError("[ERROR] PlaySound function: ["+args[1]+" for "+m.Author.Username+"]", user.Username, err)
						_, _ = s.ChannelMessageSend(m.ChannelID, "I don't know this sound... 🤔")
					}
					return
				}
			}
			_, _ = s.ChannelMessageSend(m.ChannelID, "You are not connected to a Voice Channel!")

		// [ESPORT]

		case "rl":
			err = bot.GetRLStat(s, m, args)
			print.CheckError("[ERROR] RL function", user.Username, err)

		case "lol":
			err = bot.GetLolStat(s, m, args, conf.LolKey)
			print.CheckError("[ERROR] LOL function", user.Username, err)

		// [OTHER]

		case "poll":
			err = bot.CreatePoll(s, m)
			print.CheckError("[ERROR] Poll function", user.Username, err)

		case "nsfw":
			err = bot.Reddit(s, m, "nsfw", nsfwChannelID)
			print.CheckError("[ERROR] Nsfw function", user.Username, err)

		// [HELP]

		case "help":
			// create a private messaging channel between the bot and the user
			privChan, err := s.UserChannelCreate(m.Author.ID)
			print.CheckError("[ERROR] Could not create channel between bot and user: ", user.Username, err)
			if isAdmin {
				_, _ = s.ChannelMessageSendEmbed(privChan.ID, print.EmbedAdmin)
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, print.EmbedHelp)
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🧙")

		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "I don't know this command "+m.Author.Username+". Try `!help`")
			_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤔")
		}
	}
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "Need Help? !help")
}

// Function to welcome new User on our Discord Server
func newUserHandler(s *discordgo.Session, event *discordgo.GuildMemberAdd) {

	print.InfoLog("[INFO] New user - User ID : "+event.User.Username+"", "[SERVER]")

	guild, err := s.Guild(event.GuildID)
	print.CheckError("[ERROR] Could not retrieve guild object from identifier", "[SERVER]", err)

	// create a private messaging channel between the bot and the new guild member
	privChan, err := s.UserChannelCreate(event.User.ID)
	print.CheckError("[ERROR] Could not create channel between bot and user", "[SERVER]", err)

	welcomeMsg := "**Welcome on server " + guild.Name + " ! :video_game:**\n\n I'm LPGBot and before you start your adventure here, please have a look to our channel <#434444368418570243>.\n\n Voilà, that's all. Don't hesitate to come and say hi in <#373160766670503957> \n\n Enjoy ! :call_me:\n"
	embedWelc := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xFF8142, // Orange
		Description: welcomeMsg,
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:       "Hi " + event.User.Username,
	}

	// send greet message to new guild member
	s.ChannelMessageSendEmbed(privChan.ID, embedWelc)
}

func Accum(s string) string {
	var result string
	var loop int

	for i, v := range s {
		result += strings.ToUpper(string(v))
		for i := 0; i < loop; i++ {
			result += strings.ToLower(string(v))
		}
		if i < len(s) {
			result += "-"
		}
		loop++
	}

	return result
}

func LongestConsec(strarr []string, k int) string {
	var biggest int
	var result string

	for i := 0; i < len(strarr)-(k-1); i++ {

		concat := make([]string, 0)

		for j := 0; j < k; j++ {
			concat[j] = strarr[i+j]
		}

		if len(concat) > biggest {
			for i := 0; i < len(concat); i++ {
				result = result + concat[i]
			}
		}
	}

	return result
}
