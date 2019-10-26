package bot

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type soundCollection struct {
	Prefix   string
	Commands []string
	Sounds   []*sound
}

// Sound represents a sound clip
type sound struct {
	Name   string
	buffer [][]byte
}

// LPGSOUND : Array of all the sounds availables
var LPGSOUND = &soundCollection{
	Prefix: "lpgsound",
	Commands: []string{
		"!lpgsound",
	},
	Sounds: []*sound{
		createSound("airhorn"),
		createSound("boi"),
		createSound("bruh"),
		createSound("daniel"),
		createSound("deja"),
		createSound("fuck"),
		createSound("mgs"),
		createSound("nani"),
		createSound("nice"),
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
func PlaySound(s *discordgo.Session, guildID, channelID string, msg string) (err error) {
	fmt.Println("START : PlaySound function - from sound command")
	var sound *sound
	parts := strings.Split(strings.ToLower(msg), " ")

	// If they passed a specific sound effect, find and select that (otherwise play nothing)
	if len(parts) > 1 {
		for _, s := range LPGSOUND.Sounds {
			if parts[1] == s.Name {
				fmt.Println("Request to play sound : " + s.Name)
				sound = s
			}
		}

		if sound == nil {
			fmt.Println("ERROR : Sound doesn't exist")
			return
		}
	}

	// Join the provided voice channel.
	fmt.Println("Join channel ID : " + channelID)

	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, false)

	if err != nil {
		fmt.Println("ERROR : Bot not has not been able to join the channel properly")
		return err
	}

	fmt.Println("Chanel Joined")
	// Sleep for a specified amount of time before playing the sound
	time.Sleep(10 * time.Millisecond)

	// Play the sound
	fmt.Println("Play sound !")
	sound.Play(vc)

	// Disconnect from the provided voice channel.
	fmt.Println("Disconnect")
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
	fmt.Println("Loading sound file :", s.Name)
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("ERROR : opening dca file - ", err)
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
			fmt.Println("ERROR : reading from dca file - ", err)
			return err
		}

		// read encoded pcm from dca file
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("ERROR : reading from dca file - ", err)
			return err
		}

		// append encoded pcm data to the buffer
		s.buffer = append(s.buffer, InBuf)
	}
}
