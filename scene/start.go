package scene

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"image/color"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/enums"
	"liangminghaoangus/guaiguaizhu/scene/widgets"
	"math"
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

	// new game select player window
	newGamePop := &furex.View{
		ID:         "new-game-pop",
		Width:      s.w,
		Height:     s.h / 2,
		Position:   furex.PositionAbsolute,
		Left:       0,
		Top:        s.h / 4,
		Attrs:      nil,
		Hidden:     true,
		Display:    furex.DisplayNone,
		AlignItems: furex.AlignItemCenter,
		Justify:    furex.JustifyCenter,
	}

	// new game select view and text
	newGamePop.AddChild(s.newGamePopViews()...)

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

type raceInfoSelect struct {
	race  enums.Race
	name  string
	intro string
}

func (s *StartScene) newGamePopViews() []*furex.View {
	boxCenter := &furex.View{
		WidthInPct:  80,
		HeightInPct: 100,
		Justify:     furex.JustifySpaceBetween,
		AlignItems:  furex.AlignItemStart,
		Display:     furex.DisplayFlex,
		Hidden:      false,
	}

	q := make([]*furex.View, 0)
	cardWidth := math.Ceil(float64(s.w)*0.66) / 3
	// three race card select
	l := []raceInfoSelect{
		{
			race:  enums.RaceGod,
			name:  enums.GetRaceName(enums.RaceGod),
			intro: "ff",
		},
		{
			race:  enums.RaceHuman,
			name:  enums.GetRaceName(enums.RaceHuman),
			intro: "fff",
		},
		{
			race:  enums.RaceDevil,
			name:  enums.GetRaceName(enums.RaceDevil),
			intro: "ffff",
		},
	}
	for _, val := range l {
		item := val
		boxCenter.AddChild(&furex.View{
			Width:       int(math.Ceil(cardWidth)),
			HeightInPct: 80,
			ID:          fmt.Sprintf("race_%d", item.race),
			Text:        item.name,
			Handler: &widgets.Button{FontFace: config.GetSystemFont(), OnClick: func(attr map[string]string) {
				println(item.name)
			}},
		})
	}
	q = append(q, boxCenter)
	return q
}
