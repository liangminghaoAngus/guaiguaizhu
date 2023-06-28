package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/furex/v2"
	"image"
	"image/color"
	"liangminghaoangus/guaiguaizhu/config"
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
		ID:     "",
		//Text:   "start game",
		Handler: furex.NewHandler(furex.HandlerOpts{
			Update: nil,
			Draw: func(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
				t := "开始游戏"
				textBox := text.BoundString(f, t)
				x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)
				minW := textBox.Dx() / 2
				minH := textBox.Dy() / 2
				text.Draw(screen, t, f, int(x)-minW, int(y)+minH, color.White)
			},
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
		//Text:   "load game",
		Handler: furex.NewHandler(furex.HandlerOpts{
			Update: nil,
			Draw: func(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
				t := "加载存档"
				textBox := text.BoundString(f, t)
				x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)
				minW := textBox.Dx() / 2
				minH := textBox.Dy() / 2
				text.Draw(screen, t, f, int(x)-minW, int(y)+minH, color.White)
			},
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
		//Text:   "exit game",
		Handler: func() furex.Handler {
			//hover := false
			//buttonColor := color.NRGBA{R: 170, G: 170, B: 180, A: 255}
			//buttonHover := color.NRGBA{R: 130, G: 130, B: 150, A: 255}
			//buttonPressed := color.NRGBA{R: 100, G: 100, B: 120, A: 255}
			return furex.NewHandler(furex.HandlerOpts{
				Update: nil,
				Draw: func(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
					t := "退出游戏"
					textBox := text.BoundString(f, t)
					x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)
					minW := textBox.Dx() / 2
					minH := textBox.Dy() / 2
					text.Draw(screen, t, f, int(x)-minW, int(y)+minH, color.White)
				},
				HandlePress: func(x, y int, t ebiten.TouchID) {},
				HandleRelease: func(x, y int, isCancel bool) {
					if isCancel {
						return
					}
					os.Exit(0)
				},
			})
		}(),
	})
}
