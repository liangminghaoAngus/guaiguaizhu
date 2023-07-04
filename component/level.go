package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/enums"
)

type LevelData struct {
	LevelNum     int
	Exp          int
	ExpNextLevel int
}

var Level = donburi.NewComponentType[LevelData](LevelData{})

func NewLevelData(level int) LevelData {
	nextLevelExp, _ := enums.GetLevelEXP(level)
	return LevelData{
		LevelNum: 1,
		Exp:      nextLevelExp,
	}
}
