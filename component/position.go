package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/enums"
)

type PositionData struct {
	X, Y float64
	Map  enums.Map
}

var Position = donburi.NewComponentType[PositionData](PositionData{})

func NewPlayerPositionData() PositionData { // todo 需要修改起始点
	return PositionData{
		X:   20,
		Y:   0,
		Map: enums.MapRookie,
	}
}
