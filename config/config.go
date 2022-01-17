package config

import (
	"encoding/json"
	"io/ioutil"

	"Work.go/LPG-Bot/LPGBot/print"
)

type ConfigStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	Webhook   string `json:"Webhook"`
}

// ReadConfig
func ReadConfig() *ConfigStruct {
	print.InfoLog("[INFO] Reading from config file ...", "[SERVER]")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		print.CheckError("[ERROR] Reading config.json file", "[Server]", err)
		return nil
	}

	var config ConfigStruct

	print.DebugLog("[DEBUG] "+string(file)+"", "[SERVER]")
	err = json.Unmarshal(file, &config)

	if err != nil {
		print.CheckError("[ERROR] Unmarshal config.json data", "[Server]", err)
		return nil
	}

	return &config
}
