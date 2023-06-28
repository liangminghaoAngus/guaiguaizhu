package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type Game struct {
}

func (g *Game) Update() {

}

func (s *Game) Draw(screen *ebiten.Image) {

	//text.Draw(screen, "m110's Airplanes", assets.NarrowFont, t.screenWidth/4, 100, color.White)
	//text.Draw(screen, "Player 1: WASD + Space", assets.NarrowFont, t.screenWidth/6, 250, color.White)
	//text.Draw(screen, "Player 2: Arrows + Enter", assets.NarrowFont, t.screenWidth/6, 350, color.White)
	//text.Draw(screen, "Press space to start", assets.NarrowFont, t.screenWidth/5, 500, color.White)
}
