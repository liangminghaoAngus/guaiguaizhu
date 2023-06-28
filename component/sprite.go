package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type ColorOverride struct {
	R, G, B, A float64
}

type SpriteData struct {
	Image     *ebiten.Image
	ColorOver ColorOverride
	hidden    bool
	// The original rotation of the sprite
	// "Facing right" is considered 0 degrees
	OriginalRotation float64
}

func (s *SpriteData) Show() {
	s.hidden = false
}

func (s *SpriteData) Hidden() {
	s.hidden = true
}

var Sprite = donburi.NewComponentType[SpriteData](SpriteData{})
