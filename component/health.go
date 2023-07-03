package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"liangminghaoangus/guaiguaizhu/engine"
	"time"
)

type HealthData struct {
	HP int
	MP int

	HPuiPos math.Vec2     // ui的位置
	MPuiPos math.Vec2     // ui的位置
	HPui    *ebiten.Image // 判断是否需要定制 HP 界面
	MPui    *ebiten.Image

	JustDamage           bool
	DamageIndicatorTimer *engine.Timer
	//DamageIndicator      *SpriteData
}

var Health = donburi.NewComponentType[HealthData](HealthData{
	DamageIndicatorTimer: engine.NewTimer(time.Millisecond * 100),
})

func NewPlayerHealthData() HealthData {
	return HealthData{
		HP:                   100,
		MP:                   100,
		JustDamage:           false,
		DamageIndicatorTimer: engine.NewTimer(time.Millisecond * 100),
	}
}
