package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpritePivot int

const (
	SpritePivotCenter SpritePivot = iota
	SpritePivotTopLeft
)

type ColorOverride struct {
	R, G, B, A float64
}

type SpriteData struct {
	Image     *ebiten.Image
	ColorOver ColorOverride
	HiddenED  bool
	Layer     int
	Pivot     SpritePivot
	// The original rotation of the sprite
	// "Facing right" is considered 0 degrees
	OriginalRotation float64
}

func (s *SpriteData) Show() {
	s.HiddenED = false
}

func (s *SpriteData) Hidden() {
	s.HiddenED = true
}

var Sprite = donburi.NewComponentType[SpriteData](SpriteData{})
