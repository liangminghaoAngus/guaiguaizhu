package component

import "github.com/yohamta/donburi"

type MovementData struct {
	speed float64
}

var Movement = donburi.NewComponentType[MovementData](MovementData{})
