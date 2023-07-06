package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/enums"
)

type PositionData struct {
	X, Y          float64
	Map           enums.Map
	FaceDirection int
}

func (d *PositionData) ChangeFaceDirection() {
	if d.FaceDirection == 0 {
		d.FaceDirection = 1
		return
	} else {
		d.FaceDirection = 0
		return
	}
}

func (d *PositionData) IsFaceLeft() bool {
	return d.FaceDirection == 1
}

var Position = donburi.NewComponentType[PositionData](PositionData{})

func NewPlayerPositionData() PositionData {
	return PositionData{
		X:             20,
		Y:             0,
		Map:           enums.MapRookie,
		FaceDirection: 0,
	}
}
