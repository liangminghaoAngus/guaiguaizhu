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
	return LevelData{
		LevelNum:     1,
		ExpNextLevel: nextLevelExp,
		Exp:          0,
	}
}
