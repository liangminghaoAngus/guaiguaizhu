package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene Scene
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

func NewGame() *Game {
	g := &Game{}
	return g
}
