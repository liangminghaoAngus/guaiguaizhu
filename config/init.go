package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

var config *Config

func GetConfig() *Config {
	return config
}

type Config struct {
	ScreenWidth  int    `json:"screen_width"`
	ScreenHeight int    `json:"screen_height"`
	GameName     string `json:"game_name"`
}

func Init(fileName string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configJsonPath := path.Join(dir, "config", fmt.Sprintf("%s.json", fileName))
	raw, err := os.ReadFile(configJsonPath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(raw, &config); err != nil {
		panic(err)
	}
}
