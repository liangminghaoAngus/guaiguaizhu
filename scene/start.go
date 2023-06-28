package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// StartScene game start scene
type StartScene struct {
	w int
	h int

	newGameCallBack  func()
	loadGameCallBack func()
}

func NewStart(w, h int, newGame, loadGame func()) *StartScene {
	return &StartScene{
		w:                w,
		h:                h,
		newGameCallBack:  newGame,
		loadGameCallBack: loadGame,
	}
}

func (s *StartScene) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeySpace) {
		s.newGameCallBack()
		return
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {

	//text.Draw(screen, "m110's Airplanes", assets.NarrowFont, t.screenWidth/4, 100, color.White)
	//text.Draw(screen, "Player 1: WASD + Space", assets.NarrowFont, t.screenWidth/6, 250, color.White)
	//text.Draw(screen, "Player 2: Arrows + Enter", assets.NarrowFont, t.screenWidth/6, 350, color.White)
	//text.Draw(screen, "Press space to start", assets.NarrowFont, t.screenWidth/5, 500, color.White)
}
