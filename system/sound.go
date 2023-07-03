package system

import (
	"liangminghaoangus/guaiguaizhu/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Sound struct {
	query *query.Query
}

func NewSound() *Sound {
	return &Sound{
		query: query.NewQuery(filter.Contains(component.Sound)),
	}
}

func (s *Sound) Update(w donburi.World) {
	s.query.Each(w, func(e *donburi.Entry) {
		sound := component.Sound.Get(e)

		if sound.Paused {
			sound.AudioPlayer.Pause()
			return
		}
		if sound.AudioPlayer.Current() == sound.Total && sound.Loop {
			_ = sound.AudioPlayer.Seek(0)
		}
		sound.AudioPlayer.SetVolume(float64(sound.Volume))
		sound.AudioPlayer.Play()
	})
}
