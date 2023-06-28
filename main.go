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
	ebiten.SetWindowSize(640, 320)
	rand.Seed(time.Now().UTC().UnixNano())

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
