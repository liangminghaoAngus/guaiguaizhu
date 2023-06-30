package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/enums"
)

type RaceData struct {
	Race enums.Race
}

var Race = donburi.NewComponentType[RaceData](RaceData{})

func NewRaceData(data enums.Race) RaceData {
	return RaceData{Race: data}
}
