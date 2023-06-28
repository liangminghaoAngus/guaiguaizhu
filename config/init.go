package config

import (
	"encoding/json"
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"os"
	"path"
)

var config *Config
var systemFont font.Face

func GetConfig() *Config {
	return config
}

func GetSystemFont() font.Face {
	return systemFont
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

	fontPath := path.Join(dir, "assets/font/AlibabaPuHuiTi-3-55-Regular.ttf")
	raw, err = os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	tt, err := opentype.Parse(raw)
	if err != nil {
		panic(err)
	}
	systemFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
}
