package bot

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"Work.go/LPG-Bot/LPGBot/print"
	"github.com/bwmarrin/discordgo"
)

type soundCollection struct {
	Sounds []*sound
}

// Sound represents a sound clip
type sound struct {
	Name   string
	buffer [][]byte
}

// LPGSOUND : Array of all the sounds availables
var LPGSOUND = &soundCollection{
	Sounds: []*sound{
		createSound("airhorn"),
		createSound("boi"),
		createSound("bruh"),
		createSound("coffin"),
		createSound("damage"),
		createSound("daniel"),
		createSound("deja"),
		createSound("fuck"),
		createSound("krabs"),
		createSound("mega"),
		createSound("mgs"),
		createSound("nani"),
		createSound("nice"),
		createSound("oof"),
		createSound("ooh"),
		createSound("oui"),
		createSound("ricardo"),
		createSound("spooky"),
		createSound("thug"),
		createSound("wow"),
	},
}

// PlaySound : function from github.com/bwmarrin/discordBot exemples
// it is used to play dca sound on discord channel
// retrieve an error if it can't play the sound
func PlaySound(s *discordgo.Session, guildID, channelID string, msg string, args []string) (err error) {
	print.DebugLog("[DEBUG] Start PlaySound function - from sound command", "[SERVER]")
	var sound *sound

	// If they passed a specific sound effect, find and select that (otherwise play nothing)
	if len(args) > 1 {
		for _, s := range LPGSOUND.Sounds {
			if args[1] == s.Name {
				print.DebugLog("[DEBUG] Request to play sound : "+s.Name+"", "[SERVER]")
				sound = s
			}
		}

		if sound == nil {
			print.CheckError("[ERROR] Sound doesn't exist", "[SERVER]", err)
			return
		}
	} else {
		print.CheckError("[ERROR] Not enought arguments", "[SERVER]", err)
		return err
	}

	// Play sound on a different channel that the one we are connected to
	if len(args) > 2 {
		var name string
		// Retrieve the channel name from the command
		for i := 2; i < len(args); i++ {
			name += args[i]
			if i+1 != len(args) {
				name += " "
			}
		}
		print.DebugLog("[DEBUG] Request to play sound in channel: "+name+"", "[SERVER]")

		// Check if channel exists
		channels, _ := s.GuildChannels(guildID)
		for _, c := range channels {
			// Check if channel is a guild voice channel and not a text or DM channel
			if strings.ToLower(c.Name) == name && c.Type == discordgo.ChannelTypeGuildVoice {
				channelID = c.ID
			}
		}
	}

	// Join the provided voice channel.
	print.DebugLog("[DEBUG] Bot is joining channel ID: "+channelID+"", "[SERVER]")

	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, false)

	if err != nil {
		print.CheckError("[ERROR] Bot not has not been able to join the channel properly", "[SERVER]", err)
		return err
	}

	print.DebugLog("[DEBUG] Bot has joined channel", "[SERVER]")
	// Sleep for a specified amount of time before playing the sound
	time.Sleep(10 * time.Millisecond)

	// Play the sound
	print.DebugLog("[DEBUG] Bot is playing Sound", "[SERVER]")
	//sound.Play(vc)

	if sound.buffer == nil {
		print.DebugLog("[DEBUG] Sound buffer was empty, reload sound ...", "[SERVER]")
		LPGSOUND.LoadAll()
	}

	vc.Speaking(true)
	for _, buff := range sound.buffer {
		vc.OpusSend <- buff
	}
	vc.Speaking(false)

	// Disconnect from the provided voice channel.
	print.DebugLog("[DEBUG] Bot is disconnecting from the channel", "[SERVER]")
	vc.Disconnect()

	return nil
}

// Plays this sound over the specified VoiceConnection
func (s *sound) Play(vc *discordgo.VoiceConnection) {
	vc.Speaking(true)
	defer vc.Speaking(false)

	for _, buff := range s.buffer {
		vc.OpusSend <- buff
	}
}

// Create a Sound struct
func createSound(Name string) *sound {
	return &sound{
		Name:   Name,
		buffer: make([][]byte, 0),
	}
}

// LoadAll : Is used to load all the sounds available
func (sc *soundCollection) LoadAll() {
	for _, sound := range sc.Sounds {
		sound.load(sc)
	}
}

// Load attempts to load an encoded sound file from disk
func (s *sound) load(c *soundCollection) error {
	path := fmt.Sprintf("bot/dca/%v.dca", s.Name)
	print.DebugLog("[DEBUG] Loading sound file: "+s.Name+"", "[SERVER]")
	file, err := os.Open(path)

	if err != nil {
		print.CheckError("[ERROR] While opening dca file", "[SERVER]", err)
		return err
	}

	var opuslen int16

	for {
		// read opus frame length from dca file
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil
		}

		if err != nil {
			print.CheckError("[ERROR] While reading from dca file", "[SERVER]", err)
			return err
		}

		// read encoded pcm from dca file
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			print.CheckError("[ERROR] While reading from dca file", "[SERVER]", err)
			return err
		}

		// append encoded pcm data to the buffer
		s.buffer = append(s.buffer, InBuf)
	}
}
