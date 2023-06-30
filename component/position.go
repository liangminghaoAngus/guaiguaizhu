package component

import "github.com/yohamta/donburi"

type PositionData struct {
	X, Y float64
}

var Position = donburi.NewComponentType[PositionData](PositionData{})

func NewPositionData() PositionData { // todo 需要修改起始点
	return PositionData{
		X: 0,
		Y: 0,
	}
}
