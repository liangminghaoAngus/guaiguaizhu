package component

import "github.com/yohamta/donburi"

type MovementData struct {
	Speed float64
}

var Movement = donburi.NewComponentType[MovementData](MovementData{})

func NewMovementData() MovementData {
	return MovementData{
		Speed: 0.8,
	}
}
