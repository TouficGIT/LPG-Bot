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
		createSound("fuck"),
		createSound("mgs"),
		createSound("nice"),
		createSound("ooh"),
		createSound("oui"),
		createSound("thug"),
		createSound("wow"),
	},
}

// PlaySound : function from github.com/bwmarrin/discordBot exemples
// it is used to play dca sound on discord channel
// retrieve an error if it can't play the sound
func PlaySound(s *discordgo.Session, guildID, channelID string, msg string) (err error) {

	var sound *sound
	parts := strings.Split(strings.ToLower(msg), " ")

	// If they passed a specific sound effect, find and select that (otherwise play nothing)
	if len(parts) > 1 {
		for _, s := range LPGSOUND.Sounds {
			if parts[1] == s.Name {
				sound = s
			}
		}

		if sound == nil {
			return
		}
	}

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, false)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(10 * time.Millisecond)

	// Play the sound
	sound.Play(vc)

	// Disconnect from the provided voice channel.
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

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("error opening dca file :", err)
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
			fmt.Println("error reading from dca file :", err)
			return err
		}

		// read encoded pcm from dca file
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("error reading from dca file :", err)
			return err
		}

		// append encoded pcm data to the buffer
		s.buffer = append(s.buffer, InBuf)
	}
}
