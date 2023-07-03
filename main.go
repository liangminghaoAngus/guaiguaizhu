package main

import (
	assetsImage "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/data"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	config.Init("dev")
	data.Init()
	assetsImage.Init()
	gconf := config.GetConfig()
	ebiten.SetWindowSize(gconf.ScreenWidth, gconf.ScreenHeight)
	//ebiten.SetWindowIcon()
	ebiten.SetWindowTitle(gconf.GameName)
	rand.Seed(time.Now().UTC().UnixNano())

	err := ebiten.RunGame(NewGame(gconf.ScreenWidth, gconf.ScreenHeight))
	if err != nil {
		log.Fatal(err)
	}
}
