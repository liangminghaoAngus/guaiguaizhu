package component

import "github.com/yohamta/donburi"

type MovementData struct {
	VelocityX     float64
	AccelerationX float64
	MaxSpeed      float64
}

var Movement = donburi.NewComponentType[MovementData](NewMovementData())

func NewMovementData() MovementData {
	return MovementData{
		VelocityX:     0,
		AccelerationX: 0.68,
		MaxSpeed:      2.4,
	}
}
