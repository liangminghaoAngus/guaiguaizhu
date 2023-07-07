package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type ControlData struct {
	TeleportKey ebiten.Key

	Left     ebiten.Key
	Right    ebiten.Key
	EnterKey ebiten.Key
	UeKey    ebiten.Key

	HpKey ebiten.Key
	MpKey ebiten.Key

	StoreKey ebiten.Key // 打开/关闭背包

	AbilityKeys map[ebiten.Key]int
}

var Control = donburi.NewComponentType[ControlData](ControlData{})

func NewPlayerControl() ControlData {
	return ControlData{
		TeleportKey: ebiten.KeyArrowUp,
		Left:        ebiten.KeyArrowLeft,
		Right:       ebiten.KeyArrowRight,
		EnterKey:    ebiten.KeyArrowUp,
		UeKey:       ebiten.KeyF,
		StoreKey:    ebiten.KeyB,
		HpKey:       ebiten.Key1,
		MpKey:       ebiten.Key2,
		AbilityKeys: make(map[ebiten.Key]int),
	}
}
