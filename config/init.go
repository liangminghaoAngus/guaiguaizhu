package config

import (
	"fmt"
	"os"
)

type Config struct {
	ScreenWidth  int `json:"screen_width"`
	ScreenHeight int `json:"screen_height"`
}

func Init(fileName string) {

	fmt.Println(os.Getwd())
	// os.OpenFile()
}
