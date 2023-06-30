package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type ControlData struct {
	Left     ebiten.Key
	Right    ebiten.Key
	EnterKey ebiten.Key
	UeKey    ebiten.Key
}

var Control = donburi.NewComponentType[ControlData](ControlData{})

func NewPlayerControl() ControlData {
	return ControlData{
		Left:     ebiten.KeyArrowLeft,
		Right:    ebiten.KeyArrowRight,
		EnterKey: ebiten.KeyArrowUp,
		UeKey:    ebiten.KeyF,
	}
}
