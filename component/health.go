package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/engine"
	"time"
)

type HealthData struct {
	HP                   int
	MP                   int
	JustDamage           bool
	DamageIndicatorTimer *engine.Timer
	//DamageIndicator      *SpriteData
}

var Health = donburi.NewComponentType[HealthData](HealthData{
	DamageIndicatorTimer: engine.NewTimer(time.Millisecond * 100),
})

func NewPlayerHealthData() HealthData {
	return HealthData{
		100,
		100,
		false,
		engine.NewTimer(time.Millisecond * 100),
	}
}
