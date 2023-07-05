package component

import (
	"image"
	"liangminghaoangus/guaiguaizhu/engine"
	"liangminghaoangus/guaiguaizhu/enums"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type AnimateHealthItem int

const (
	AnimateHp AnimateHealthItem = iota + 1
	AnimateMp
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

	AnimationTime     time.Duration
	LastAnimationTime map[AnimateHealthItem]time.Time
	LastAnimationItem map[AnimateHealthItem]int

	JustDamage           bool
	DamageIndicatorTimer *engine.Timer
	//DamageIndicator      *SpriteData
}

var Health = donburi.NewComponentType[HealthData](HealthData{
	DamageIndicatorTimer: engine.NewTimer(time.Millisecond * 100),
	AnimationTime:        time.Second * 1,
	LastAnimationTime:    make(map[AnimateHealthItem]time.Time),
	LastAnimationItem:    make(map[AnimateHealthItem]int),
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
		AnimationTime:        time.Second * 1,
		DamageIndicatorTimer: engine.NewTimer(time.Millisecond * 100),
		LastAnimationTime:    make(map[AnimateHealthItem]time.Time),
		LastAnimationItem:    make(map[AnimateHealthItem]int),
	}
}

func (d *HealthData) ChangeHP(targetHP int, nowTime time.Time) {
	d.LastAnimationItem[AnimateHp] = d.HP
	//d.LastAnimationTime[AnimateHp] = nowTime
	d.HP = targetHP
}

func (d *HealthData) ChangeMP(targetMP int, nowTime time.Time) {
	d.LastAnimationItem[AnimateMp] = d.MP
	//d.LastAnimationTime[AnimateMp] = nowTime
	d.MP = targetMP
}

func (d *HealthData) DrawPlayerHPImage(screen, hpUI *ebiten.Image, x, y int, hp int, a float32) {
	hpImage := hpUI
	percent := float64(hp) / float64(d.HPMax)
	x0 := 0
	y0 := hpImage.Bounds().Dy()
	x1 := hpImage.Bounds().Dx()
	y1 := float64(hpImage.Bounds().Dy()) * (float64(1) - percent)

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.SetA(a)
	op.GeoM.Translate(float64(x+10), float64(y+12+int(y1)))
	screen.DrawImage(hpImage.SubImage(image.Rect(x0, y0, x1, int(y1))).(*ebiten.Image), op)
}

func (d *HealthData) DrawPlayerMPImage(screen, mpUI *ebiten.Image, x, y int, mp int, a float32) {
	mpImage := mpUI
	percent := float64(mp) / float64(d.MPMax)
	x0 := 0
	y0 := mpImage.Bounds().Dy()
	x1 := mpImage.Bounds().Dx()
	y1 := float64(mpImage.Bounds().Dy()) * (float64(1) - percent)

	transx := x - mpImage.Bounds().Dx()

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.SetA(a)
	op.GeoM.Translate(float64(transx-16), float64(y+12+int(y1)))
	screen.DrawImage(mpImage.SubImage(image.Rect(x0, y0, x1, int(y1))).(*ebiten.Image), op)
}

type HealData struct {
	level int

	itemNums map[AnimateHealthItem]int
	itemHeal map[AnimateHealthItem]int
}

func (d *HealData) AddItem(hp, mp int) {
	d.itemNums[AnimateHp] += hp
	d.itemNums[AnimateMp] += mp
}

func (d *HealData) LevelUp() {
	d.level++
	for item := range d.itemHeal {
		d.itemHeal[item] = enums.HealItemLevel[d.level]
	}
}

func (d *HealData) UseHP() int {
	if count, ok := d.itemNums[AnimateHp]; ok && count > 0 {
		if value, ok := d.itemHeal[AnimateHp]; ok {
			return value
		}
	}
	return 0
}

func (d *HealData) UseMP() int {
	if count, ok := d.itemNums[AnimateMp]; ok && count > 0 {
		if value, ok := d.itemHeal[AnimateMp]; ok {
			return value
		}
	}
	return 0
}

var defaultHealData = HealData{
	level: 1,
	itemNums: map[AnimateHealthItem]int{
		AnimateHp: 1,
		AnimateMp: 1,
	},
	itemHeal: map[AnimateHealthItem]int{
		AnimateHp: enums.HealItemLevel[1],
		AnimateMp: enums.HealItemLevel[1],
	},
}

var Heal = donburi.NewComponentType[HealData](defaultHealData)
