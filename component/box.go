package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/enums"
)

type BoxData struct {
	Width, Height int
}

var Box = donburi.NewComponentType[BoxData](BoxData{})

func NewPlayerBox(raceInt enums.Race) BoxData {
	r := BoxData{}
	switch raceInt {
	case enums.RaceGod:
		r = BoxData{
			Width:  50,
			Height: 80,
		}
	case enums.RaceHuman:
		r = BoxData{
			Width:  50,
			Height: 80,
		}
	case enums.RaceDevil:
		r = BoxData{
			Width:  50,
			Height: 80,
		}
	}
	return r
}

func NewTeleportBox() BoxData {
	return BoxData{
		Width:  80,
		Height: 100,
	}
}
