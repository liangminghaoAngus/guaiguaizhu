package component

import (
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/yohamta/donburi"
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
			Width:  40,
			Height: 60,
		}
	case enums.RaceHuman:
		r = BoxData{
			Width:  40,
			Height: 60,
		}
	case enums.RaceDevil:
		r = BoxData{
			Width:  40,
			Height: 60,
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
