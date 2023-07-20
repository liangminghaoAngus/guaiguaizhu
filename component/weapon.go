package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type WeaponData struct {
	Level     int
	AttackNum int

	Image *ebiten.Image

	Angle         float64
	Width, Height float64
}
