package component

import "github.com/yohamta/donburi"

type MovementData struct {
	speed float64
}

var Movement = donburi.NewComponentType[MovementData](MovementData{})

func NewMovementData() MovementData {
	return MovementData{
		speed: 0.8,
	}
}
