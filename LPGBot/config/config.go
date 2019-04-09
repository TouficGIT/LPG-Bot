package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Global var for LPGBot
var (
	Token     string //Token of the bot (discord)
	BotPrefix string //For the exclamation mark before the cmd
	Webhook   string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	Webhook   string `json:"Webhook"`
}

// ReadConfig
func ReadConfig() error {
	fmt.Println("Reading from config file ...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	Webhook = config.Webhook

	return nil
}
