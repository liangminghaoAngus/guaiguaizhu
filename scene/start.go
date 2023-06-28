package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"image/color"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/scene/widgets"
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

	f := config.GetSystemFont()

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
		Text:   "开始游戏",
		Handler: &widgets.Button{FontFace: f, OnClick: func(attrs map[string]string) {
			println("select player")
		}},
	})

	// load game button
	s.gameUI.AddChild(&furex.View{
		Width:  200,
		Height: 80,
		Text:   "加载存档",
		Handler: &widgets.Button{FontFace: f, OnClick: func(attrs map[string]string) {
			println("select load")
		}},
	})
	// exit game button
	s.gameUI.AddChild(&furex.View{
		Width:  200,
		Height: 80,
		Text:   "退出游戏",
		Handler: &widgets.Button{FontFace: f, OnClick: func(attrs map[string]string) {
			os.Exit(0)
		}},
	})
}
