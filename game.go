package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"liangminghaoangus/guaiguaizhu/enums"
	"liangminghaoangus/guaiguaizhu/scene"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene Scene
	w     int
	h     int
}

func (g *Game) switch2Start() {
	newGameFunc := func(raceInt enums.Race) {
		g.switch2Game(raceInt)
	}
	loadGameFunc := func() {
		// todo
	}

	g.scene = scene.NewStart(g.w, g.h, newGameFunc, loadGameFunc)
}

func (g *Game) switch2Game(raceInt enums.Race) {
	g.scene = scene.NewGame(raceInt)
}

func (g *Game) Update() error {
	if g.scene != nil {
		g.scene.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene != nil {
		g.scene.Draw(screen)
	}
	debugStr := fmt.Sprintf("FPS:%0.2f \nTPS:%0.2f", ebiten.ActualFPS(), ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, debugStr)
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func NewGame(w, h int) *Game {
	g := &Game{
		w: w,
		h: h,
	}
	g.switch2Start()
	return g
}
