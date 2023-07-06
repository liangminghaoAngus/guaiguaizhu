package component

import (
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/yohamta/donburi"
)

type LevelData struct {
	LevelNum     int
	Exp          int
	ExpNextLevel int
}

var Level = donburi.NewComponentType[LevelData](LevelData{})

func NewLevelData(level int) LevelData {
	nextLevelExp, _ := enums.GetLevelEXP(level)
	prevLevelExp, _ := enums.GetLevelEXP(level - 1)
	return LevelData{
		LevelNum:     level,
		ExpNextLevel: nextLevelExp,
		Exp:          prevLevelExp,
	}
}
