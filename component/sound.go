package component

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type SoundData struct {
	Paused       bool
	Loop         bool
	Mp3Byte      []byte
	AudioContext *audio.Context
	AudioPlayer  *audio.Player
	Volume       int
}

var Sound = donburi.NewComponentType[SoundData](SoundData{})

var BgSound = donburi.NewTag()

func FindBgSound(w donburi.World) *SoundData {
	bg, ok := query.NewQuery(filter.Contains(Sound, BgSound)).First(w)
	if !ok {
		panic("bg not found")
	}
	return Sound.Get(bg)
}
