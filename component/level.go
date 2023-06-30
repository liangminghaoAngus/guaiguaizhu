package component

import "github.com/yohamta/donburi"

type LevelData struct {
	LevelNum int
	Exp      int
}

var Level = donburi.NewComponentType[LevelData](LevelData{})

func NewLevelData() LevelData {
	return LevelData{
		LevelNum: 1,
		Exp:      0,
	}
}
