package print

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	cTime = time.Now().Format("01-02-2006 15:04:05")

	DEBUG = false

	HelpMsg = `
**FUN:**
` + "`!flip`" + ` : to play heads or tails
` + "`!roll`" + ` : to roll a dice
` + "`!chuck`" + ` : to get a fact on Chuck Norris
` + "`!meme`" + ` : to get a Random meme from Reddit/memes

**TO PLAY A SOUND:**
` + "`!say + <sound> + <channel's name>`" + ` : to play one of the following sounds (without the /)
	-> boi / bruh / coffin / damage / daniel / deja / fuck / krabs / mega
	-> mgs / nani / nice / ooh / oui / ricardo / spooky / thug / wow

**ESPORT (stats):**
` + "`!rl + <Epic username>`" + ` : to show your stats on Rocket League
` + "`!lol + <Riot username>`" + ` : to show your stats on League of Legends

**OTHER:**
` + "`!poll \"Question\" \"Choice1\" \"Choice2\" ... `" + ` : to create a Strawpoll

**NEED HELP?:**
` + "`!help`" + `: to show this message ^^ `

	EmbedHelp = &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xFF8142, // Orange
		Description: HelpMsg,
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:       "Hi, I'm LPGBot! I answer to the following commands:",
	}

	RLerrorMsg = ` 
How to use the !rl command:

You just need to type !rl, followed by your Epic username: 
**!rl** + __**<Epic username>**__
`

	EmbedErrorRL = &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x3D85C6,
		Description: RLerrorMsg,
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       "LPGBot: Command !rl",
	}

	AdminMsg = `
**ADMIN:**
` + "`!welcome`" + ` : to test the welcome msg
` + "`!setnsfwchannel <channel id>`" + ` : to set the nsfw channel
` + "`!debug <on/off>`" + ` : to set debug mode in the logs`

	EmbedAdmin = &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xFF8142, // Orange
		Description: AdminMsg,
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:       "Admin commands",
	}
)

func CheckError(msg string, user string, err error) {
	if err != nil {
		fmt.Println("Time: " + cTime + " || " + user + " || " + msg + " || " + err.Error())
		return
	}
}

func InfoLog(msg string, user string) {
	fmt.Println("Time: " + cTime + " || " + user + " || " + msg)
	return
}

func SetDebug(option string) {
	if option == "on" {
		DEBUG = true
	} else if option == "off" {
		DEBUG = false
	}
}

func DebugLog(msg string, user string) {
	if DEBUG == true {
		fmt.Println("Time: " + cTime + " || " + user + " || " + msg)
		return
	} else {
		return
	}

}
