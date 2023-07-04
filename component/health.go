package component

import (
	"liangminghaoangus/guaiguaizhu/engine"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type HealthData struct {
	HP    int
	HPMax int
	MP    int
	MPMax int

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

func NewPlayerHealthData(hp, mp *ebiten.Image) HealthData {
	return HealthData{
		HP:                   80,
		HPMax:                100,
		HPui:                 hp,
		MP:                   50,
		MPMax:                100,
		MPui:                 mp,
		JustDamage:           false,
		DamageIndicatorTimer: engine.NewTimer(time.Millisecond * 100),
	}
}
