package main

import (
	"liangminghaoangus/guaiguaizhu/config"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	config.Init("dev")
	gconf := config.GetConfig()
	ebiten.SetWindowSize(gconf.ScreenWidth, gconf.ScreenHeight)
	//ebiten.SetWindowIcon()
	ebiten.SetWindowTitle(gconf.GameName)
	rand.Seed(time.Now().UTC().UnixNano())

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
