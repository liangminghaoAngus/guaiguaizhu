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

	closePopID := ""
	// start game button
	s.gameUI.AddChild(&furex.View{
		Width:  200,
		Height: 80,
		Text:   "开始游戏",
		Handler: &widgets.Button{FontFace: f, OnClick: func(attrs map[string]string) {
			//println("select player")
			if dialog, ok := s.gameUI.GetByID("new-game-pop"); ok {
				dialog.Display = furex.DisplayFlex
				dialog.SetHidden(false)
				closePopID = "new-game-pop"
			}
		}},
	})

	// new game select player window
	newGamePop := &furex.View{
		ID:       "new-game-pop",
		Width:    s.w,
		Height:   s.h / 2,
		Position: furex.PositionAbsolute,
		Left:     0,
		Top:      s.h / 4,
		Attrs:    nil,
		Hidden:   true,
		Display:  furex.DisplayNone,
	}

	// right-top close button
	rightPos := 4
	rightTopCloseButton := &furex.View{
		Width:    36,
		Height:   36,
		Position: furex.PositionAbsolute,
		Right:    &rightPos,
		Top:      4,
		Text:     "×",
		Handler: &widgets.Button{FontFace: f, OnClick: func(attrs map[string]string) {
			if dialog, ok := s.gameUI.GetByID(closePopID); ok {
				dialog.Display = furex.DisplayNone
				dialog.SetHidden(true)
			}
		}},
	}

	newGamePop.AddChild(rightTopCloseButton)
	// load game button
	s.gameUI.AddChild(&furex.View{
		Width:  200,
		Height: 80,
		Text:   "加载存档",
		Handler: &widgets.Button{FontFace: f, OnClick: func(attrs map[string]string) {
			//println("select load")
			if dialog, ok := s.gameUI.GetByID("load-game-pop"); ok {
				closePopID = "load-game-pop"
				dialog.Display = furex.DisplayFlex
				dialog.SetHidden(!dialog.Hidden)
			}
		}},
	})

	// load game select player window
	LoadGamePop := &furex.View{
		ID:       "load-game-pop",
		Width:    s.w,
		Height:   s.h / 2,
		Position: furex.PositionAbsolute,
		Left:     0,
		Top:      s.h / 4,
		Attrs:    nil,
		Hidden:   true,
		Display:  furex.DisplayNone,
	}

	LoadGamePop.AddChild(rightTopCloseButton)
	// exit game button
	s.gameUI.AddChild(&furex.View{
		Width:  200,
		Height: 80,
		Text:   "退出游戏",
		Handler: &widgets.Button{FontFace: f, OnClick: func(attrs map[string]string) {
			os.Exit(0)
		}},
	})

	// dialog window
	s.gameUI.AddChild(newGamePop)
	s.gameUI.AddChild(LoadGamePop)
}
