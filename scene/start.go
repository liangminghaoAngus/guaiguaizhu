package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"image/color"
	"os"
)

// StartScene game start scene
type StartScene struct {
	w int
	h int

	gameUI *furex.View

	newGameCallBack  func()
	loadGameCallBack func()
}

func NewStart(w, h int, newGame, loadGame func()) *StartScene {
	s := &StartScene{
		w:                w,
		h:                h,
		newGameCallBack:  newGame,
		loadGameCallBack: loadGame,
	}
	s.setupUI()
	return s
}

func (s *StartScene) Update() {
	s.gameUI.Update()

	//if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeySpace) {
	//	s.newGameCallBack()
	//	return
	//}
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x3d, 0x55, 0x0c, 0xff})
	s.gameUI.Draw(screen)

}

func (s *StartScene) setupUI() {
	furex.Debug = true

	s.gameUI = &furex.View{
		Width:        s.w,
		Height:       s.h,
		Direction:    furex.Column,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
		AlignContent: furex.AlignContentCenter,
		Wrap:         furex.Wrap,
	}
	// start game button
	s.gameUI.AddChild(&furex.View{
		Width:  200,
		Height: 80,
		ID:     "",
		Text:   "start game",
		Handler: furex.NewHandler(furex.HandlerOpts{
			Update:      nil,
			Draw:        nil,
			HandlePress: func(x, y int, t ebiten.TouchID) {},
			HandleRelease: func(x, y int, isCancel bool) {
				if isCancel {
					return
				}
				// select player
				println("select player")
			},
		}),
	})

	// load game button
	s.gameUI.AddChild(&furex.View{
		Width:  200,
		Height: 80,
		ID:     "",
		Text:   "load game",
		Handler: furex.NewHandler(furex.HandlerOpts{
			Update:      nil,
			Draw:        nil,
			HandlePress: nil,
			HandleRelease: func(x, y int, isCancel bool) {
				if isCancel {
					return
				}
				// select load
				println("select load")
			},
		}),
	})
	// exit game button
	s.gameUI.AddChild(&furex.View{
		Width:  200,
		Height: 80,
		ID:     "",
		Text:   "exit game",
		Handler: furex.NewHandler(furex.HandlerOpts{
			Update:      nil,
			Draw:        nil,
			HandlePress: nil,
			HandleRelease: func(x, y int, isCancel bool) {
				if isCancel {
					return
				}
				os.Exit(0)
			},
		}),
	})
}
