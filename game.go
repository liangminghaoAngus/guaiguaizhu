package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func NewGame() *Game {
	g := &Game{}
	return g
}
